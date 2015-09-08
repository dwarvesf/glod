package zing

import (
	"errors"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	song             string = "http://mp3.zing.vn/bai-hat/"
	album            string = "http://mp3.zing.vn/album/"
	linkDownloadSong string = "http://v3.mp3.zing.vn/download/vip/song/"
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
		urlList := strings.Split(link, "/")
		if len(urlList) < 6 {
			return nil, errors.New("Wrong Format link")
		}
		linkDownload := linkDownloadSong + urlList[5]
		// cut .html
		substring := linkDownload[0 : len(linkDownload)-5]
		listStream = append(listStream, substring)
		return listStream, nil
	}

	if strings.Contains(link, album) {
		doc, err := goquery.NewDocument(link)
		if err != nil {
			return nil, err
		}

		doc.Find(".fn-playlist-item").Each(func(i int, s *goquery.Selection) {
			a, _ := s.Attr("data-id")
			linkDownload := linkDownloadSong + a
			listStream = append(listStream, linkDownload)
		})
		if len(listStream) == 0 {
			return nil, errors.New("Invalid Link")
		}
		return listStream, nil
	}
	return listStream, errors.New("Unable to dowload this link")
}
