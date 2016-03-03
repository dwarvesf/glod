package zing

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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

		doc, err := goquery.NewDocument(link)
		if err != nil {
			return nil, err
		}
		doc.Find(".zm-player-song").Each(func(i int, s *goquery.Selection) {
			a, _ := s.Attr("data-xml")

			response, err := http.Get(a)
			if err != nil {
				fmt.Println("Error while downloading", a, "-", err)
				return
			}
			defer response.Body.Close()
			buffer, _ := ioutil.ReadAll(response.Body)

			parseString := string(buffer)

			splitStringSourceStart := strings.Split(parseString, "<source>")
			splitStringSourceEnd := strings.Split(splitStringSourceStart[1], "</source>")
			_s := splitStringSourceEnd[0]
			listStream = append(listStream, _s[9:len(_s)-3])

		})
		return listStream, nil
	}

	if strings.Contains(link, album) {
		doc, err := goquery.NewDocument(link)
		if err != nil {
			return nil, err
		}
		doc.Find(".zm-player-song").Each(func(i int, s *goquery.Selection) {
			a, _ := s.Attr("data-xml")

			response, err := http.Get(a)
			if err != nil {
				fmt.Println("Error while downloading", a, "-", err)
				return
			}
			defer response.Body.Close()
			buffer, _ := ioutil.ReadAll(response.Body)

			parseString := string(buffer)

			splitStringSourceStart := strings.Split(parseString, "<source>")
			for i, v := range splitStringSourceStart {
				if i != 0 && i != len(splitStringSourceStart) {
					splitStringSourceEnd := strings.Split(v, "</source>")
					_s := splitStringSourceEnd[0]
					listStream = append(listStream, _s[9:len(_s)-3])
				}
			}

		})

		return listStream, nil
	}
	return listStream, errors.New("Unable to dowload this link")
}
