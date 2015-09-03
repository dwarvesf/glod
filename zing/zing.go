package zing

import (
	"fmt"
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

func (z *Zing) GetDirectLink(link string) ([]string, error) {
	if link == "" {
		return nil, nil
	}
	var listStream []string
	if strings.Contains(link, song) {
		urlList := strings.Split(link, "/")

		//req := httplib.Get(linkDownloadSong + urlList[3])

		linkDownload := linkDownloadSong + urlList[5]
		// cut .html
		substring := linkDownload[0 : len(linkDownload)-5]
		fmt.Println(substring)
		listStream = append(listStream, substring)
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
	}
	return listStream, nil
}
