package zing

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	InitMp3          string = "http://www.mp3.zing.vn"
	song             string = "http://mp3.zing.vn/bai-hat/"
	album            string = "http://mp3.zing.vn/album/"
	linkDownloadSong string = "http://www.mp3.zing.vn/json/song/get-download?code="
)

type Zing struct {
}

// function that input is a link then return an slice of url that permantly download file and error(if it has)
func (z *Zing) GetDirectLink(link string) ([]string, error) {
	if link == "" {
		return nil, errors.New("Empty Link")

	}
	var listStream []string
	if strings.Contains(link, song) {
		if len(strings.Split(link, "/")) < 6 {
			return nil, errors.New("Invalid link")

		}

		doc, err := goquery.NewDocument(link)
		if err != nil {
			return nil, err

		}

		doc.Find("#tabService").Each(func(i int, s *goquery.Selection) {

			dataCode, _ := s.Attr("data-code")

			linkDownload := linkDownloadSong + dataCode
			response, _ := http.Get(linkDownload)

			defer response.Body.Close()
			buffer, _ := ioutil.ReadAll(response.Body)

			parseString := string(buffer)

			stringSource := strings.Split(parseString, "\"link\":\"")
			_stringSource := strings.Split(stringSource[1], dataCode)
			listStreamTmp := strings.Split(_stringSource[0], "\",\"size")

			stringTitle := strings.Split(listStreamTmp[0], "/song/")
			_Title := strings.Split(stringTitle[1], "/")

			listStream = append(listStream, InitMp3+listStreamTmp[0]+"~"+_Title[0])

		})
		return listStream, nil

	}

	if strings.Contains(link, album) {
		doc, err := goquery.NewDocument(link)
		if err != nil {
			return nil, err

		}

		doc.Find(".fn-playlist-item").Each(func(i int, s *goquery.Selection) {
			dataCode, _ := s.Attr("data-code")

			linkDownload := linkDownloadSong + dataCode
			response, _ := http.Get(linkDownload)
			defer response.Body.Close()
			buffer, _ := ioutil.ReadAll(response.Body)

			parseString := string(buffer)

			stringSource := strings.Split(parseString, "\"link\":\"")
			_stringSource := strings.Split(stringSource[1], dataCode)
			listStreamTmp := strings.Split(_stringSource[0], "\",\"size")

			stringTitle := strings.Split(listStreamTmp[0], "/song/")
			_Title := strings.Split(stringTitle[1], "/")

			listStream = append(listStream, InitMp3+listStreamTmp[0]+"~"+_Title[0])

		})
		return listStream, nil

	}

	return listStream, errors.New("Unable to dowload this link")

}
