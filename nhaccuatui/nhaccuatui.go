package nhaccuatui

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type NhacCuaTui struct {
}

const (
	song             string = "http://www.nhaccuatui.com/bai-hat/"
	album            string = "http://www.nhaccuatui.com/playlist/"
	linkDownloadSong string = "http://www.nhaccuatui.com/download/song/"
)

type ResponseNhacCuaTui struct {
	Data           ResponseData `json:"data"`
	ErrorMessage   string       `json:"error_message"`
	ErrorCode      int          `json:"error_code"`
	StatusReadMode bool         `json:"STATUS_READ_MODE"`
}

type ResponseData struct {
	StreamUrl string `json:"stream_url"`
	IsCharge  string `json:"is_charge"`
}

// function that input is a link then return an slice of url that permantly download file and error(if it has)
func (nct *NhacCuaTui) GetDirectLink(link string) ([]string, error) {
	if link == "" {
		return nil, errors.New("Empty Link")
	}

	// implement
	var listStream []string
	if strings.Contains(link, song) {
		urlList := strings.Split(link, ".")
		if len(urlList) < 4 {
			return nil, errors.New("Wrong Format Link")
		}

		var res ResponseNhacCuaTui
		response, err := http.Get(linkDownloadSong + urlList[3])
		defer response.Body.Close()

		buffer, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(buffer, &res)
		if err != nil {
			return nil, errors.New("Error parsing")
		}

		if res.Data.StreamUrl == "" {
			return nil, errors.New("Invalid Link")
		}
		listStream = append(listStream, res.Data.StreamUrl)
		return listStream, nil
	}

	if strings.Contains(link, album) {
		doc, err := goquery.NewDocument(link)
		if err != nil {
			return nil, err
		}
		// read all html, find class "item_content" then file tag a, get href infomation
		doc.Find("#idScrllSongInAlbum").Find(".item_content").Each(func(i int, s *goquery.Selection) {
			a, _ := s.Find("a").Attr("href")
			urlList := strings.Split(a, ".")
			var res ResponseNhacCuaTui
			response, err := http.Get(linkDownloadSong + urlList[3])
			defer response.Body.Close()

			buffer, _ := ioutil.ReadAll(response.Body)
			err = json.Unmarshal(buffer, &res)
			if err != nil {
				return
			}

			if res.Data.StreamUrl == "" {
				return
			}
			listStream = append(listStream, res.Data.StreamUrl)
		})

		if len(listStream) == 0 {
			return nil, errors.New("Invalid Link")
		}
		return listStream, nil
	}
	return listStream, errors.New("Unable to dowload this link")

}
