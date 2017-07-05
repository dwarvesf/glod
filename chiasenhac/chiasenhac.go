package chiasenhac

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/goquery"
)

const (
	CsnPrefix         = "chiasenhac"
	csnDownloadSuffix = "download.html"
	csnAblum          = "chiasenhac.vn/nghe-album/"
	csnSong1          = "chiasenhac.vn/mp3"
	csnSong2          = "chiasenhac.vn/nhac-hot"

	csnStreamURL = "/128/"
)

type ChiaSeNhac struct {
}

type Response struct {
	Artist    string
	StreamURL string
	Title     string
}

// TODO: test chiesenhac
func (csn *ChiaSeNhac) GetDirectLink(link string) ([]Response, error) {
	if link == "" {
		return nil, errors.New("Empty Link")
	}

	var listLink []string

	if strings.Contains(link, csnSong1) || strings.Contains(link, csnSong2) {
		if !strings.Contains(link, csnDownloadSuffix) {
			url := strings.Replace(link, ".html", "_download.html", -1)
			listLink = append(listLink, url)
		} else {
			listLink = append(listLink, link)
		}
	} else if strings.Contains(link, csnAblum) {
		doc, err := goquery.NewDocument(link)
		if err != nil {
			return nil, err
		}

		var url string
		for i := 1; ; i++ {
			url = ""
			doc.Find("#playlist-" + strconv.Itoa(i)).Find("td").Find("a").EachWithBreak(func(i int, s *goquery.Selection) bool {
				url, _ = s.Attr("href")
				if strings.Contains(url, csnDownloadSuffix) {
					return false
				}
				return true
			})
			if strings.Contains(url, csnDownloadSuffix) {
				listLink = append(listLink, url)
			} else {
				break
			}
		}
	}

	listSong := getSongs(listLink)

	return listSong, nil
}

func getSongs(listLink []string) []Response {
	var listSong []Response
	var song Response
	for i := range listLink {
		res, err := http.Get(listLink[i])
		if err != nil {
			return nil
		}

		defer res.Body.Close()

		buffer, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil
		}

		parseString := string(buffer)

		tmp := strings.Split(parseString, "document.write")
		urlString := strings.Split(tmp[1], "href=\"")
		url := strings.Split(urlString[1], "\" onmouseover")

		streamURL := url[0]

		title, artist := getTitleAndArtist(streamURL)

		song.StreamURL = streamURL
		song.Title = title
		song.Artist = artist
		listSong = append(listSong, song)
	}
	return listSong
}

func getTitleAndArtist(link string) (string, string) {
	_title := strings.Split(link, "/")
	if strings.Contains(_title[8], "%20-%20") {
		title := strings.Split(_title[8], "%20-%20")
		title[0] = strings.Replace(title[0], "%20", " ", -1)
		_artist := strings.Split(title[1], "[MP3")
		artist := strings.Replace(_artist[0], "%20", " ", -1)
		return title[0], artist
	}
	return "", ""
}
