package youtube

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/dwarvesf/glod"
)

// YoutubePrefix ...
const (
	YoutubePrefix = "youtube"
)

const (
	linkDownload        = "http://www.youtube.com/get_video_info?&video_id="
	album        string = "playlist"
	isFail       string = "status=ok"
)

// Format ...
type Format struct {
	Itag                    int
	VideoType, Quality, URL string
}

// Youtube ...
type Youtube struct {
}

// DownloadSingleVideo is a function that download single video
func DownloadSingleVideo(videoID string) (glod.Response, error) {
	var song glod.Response

	queryString, err := fetchMeta(videoID)
	if err != nil {
		return song, errors.New("Cannot fetchMeta")
	}

	song, err = parseMeta(videoID, queryString)
	if err != nil {
		return song, errors.New("Cannot parseMeta")
	}

	return song, nil
}

// fetchMeta function readall body of request and return string body
func fetchMeta(videoID string) (string, error) {
	resp, err := http.Get(linkDownload + videoID)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	queryString, _ := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(queryString), isFail) != true {
		return "", errors.New("Invalid Video Id")
	}
	return string(queryString), nil
}

// parseMeta is a function parse string to Video struct
func parseMeta(videoID, queryString string) (glod.Response, error) {
	var song glod.Response

	u, _ := url.Parse("?" + queryString)
	query := u.Query()

	if query.Get("errorcode") != "" || query.Get("status") == "fail" {
		return song, errors.New(query.Get("reason"))
	}

	song.Title = query.Get("title")
	song.Artist = query.Get("author")

	formatParam := strings.Split(query.Get("url_encoded_fmt_stream_map"), ",")
	var formats []Format
	for _, f := range formatParam {
		furl, _ := url.Parse("?" + f)
		fquery := furl.Query()

		itag, _ := strconv.Atoi(fquery.Get("itag"))

		formats = append(
			formats,
			Format{
				Itag:      itag,
				VideoType: fquery.Get("type"),
				Quality:   fquery.Get("quality"),
				URL:       fquery.Get("url"),
			})
	}
	if strings.Contains(formats[0].URL, "signature=") {
		song.StreamURL = formats[0].URL
	}
	return song, nil
}

// GetTitleAndArtist returns Title and Artist in case we can't crawl URL
func GetTitleAndArtist(link string, index int) glod.Response {
	var song glod.Response

	doc, err := goquery.NewDocument(link)
	if err != nil {
		return song
	}

	if strings.Contains(link, album) {
		doc.Find(".pl-video").EachWithBreak(func(i int, s *goquery.Selection) bool {
			if i == index {
				dataTitle, _ := s.Attr("data-title")
				song.Artist = dataTitle

				if strings.Contains(dataTitle, "-") {
					temp := strings.Split(dataTitle, "-")
					song.Title = temp[1]
					song.Artist = temp[0]
				}
				return false
			}
			return true
		})
	} else {
		doc.Find("#eow-title").Each(func(i int, s *goquery.Selection) {
			dataTitle, _ := s.Attr("title")
			song.Artist = dataTitle

			if strings.Contains(dataTitle, "-") {
				temp := strings.Split(dataTitle, "-")
				song.Title = temp[1]
				song.Artist = temp[0]
			}
		})
	}
	return song
}

// ToPlayList ...
func ToPlayList(link string) string {
	temp := strings.Split(link, "&list=")
	link = temp[1]
	return link
}

// GetDirectLink is a function that receive input is a link and output doesnt matter(but it override GetDirectLink of Glod interface)
func (youtube *Youtube) GetDirectLink(link string) ([]glod.Response, error) {
	var listSong []glod.Response

	var song glod.Response

	if link == "" {
		return nil, errors.New("Empty Link")
	}

	if strings.Contains(link, "&list=") {
		link = ToPlayList(link)
	}

	if strings.Contains(link, album) {
		var listVideoID []string
		doc, err := goquery.NewDocument(link)
		if err != nil {
			return nil, err
		}

		doc.Find(".pl-video").Each(func(i int, s *goquery.Selection) {
			videoID, _ := s.Attr("data-video-id")
			listVideoID = append(listVideoID, videoID)
		})
		for i := 0; i < len(listVideoID); i++ {
			song, err := DownloadSingleVideo(listVideoID[i])
			if err != nil {
				song = GetTitleAndArtist(link, i)
			}
			listSong = append(listSong, song)
		}
		for i := range listSong {
			fmt.Println("Artist : " + listSong[i].Artist)
			fmt.Println("Title : " + listSong[i].Title)
			fmt.Println("URL : " + listSong[i].StreamURL)
		}
		return listSong, nil
	}

	urlList := strings.Split(link, "/")
	if len(urlList) < 4 {
		return nil, errors.New("Invalid link")
	}

	_videoID := urlList[3]
	videoID := _videoID[8:len(_videoID)]
	song, err := DownloadSingleVideo(videoID)
	if err != nil {

		// 1 means single video
		song = GetTitleAndArtist(link, 1)

		fmt.Println("Artist : " + song.Artist)
		fmt.Println("Title : " + song.Title)
		fmt.Println("URL : " + song.StreamURL)

	}
	listSong = append(listSong, song)

	return listSong, nil
}
