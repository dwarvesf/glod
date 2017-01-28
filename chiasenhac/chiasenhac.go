package chiasenhac

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
)

type ChiaSeNhac struct {
}

func (csn *ChiaSeNhac) GetDirectLink(link string) ([]string, error) {
	if link == "" {
		return nil, errors.New("Empty Link")
	}

	var listStream []string

	linkDownloadSong := link[:len(link)-5] + "_download" + link[len(link)-5:]

	doc, err := goquery.NewDocument(linkDownloadSong)
	if err != nil {
		return nil, err
	}

	doc.Find("#downloadlink").Find("a").Each(func(i int, s *goquery.Selection) {
		if i == 1 {
			a, _ := s.Attr("href")
			listStream = append(listStream, a)
		}
	})
	return listStream, nil
}
