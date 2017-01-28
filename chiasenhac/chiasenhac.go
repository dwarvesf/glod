package chiasenhac

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type ChiaSeNhac struct {
}

func (csn *ChiaSeNhac) GetDirectLink(link string) ([]string, error) {
	if link == "" {
		return nil, errors.New("Empty Link")
	}

	var listStream []string

	linkDownloadTmp := strings.Split(link, ".html")
	linkDownload := linkDownloadTmp[0] + "_download.html"

	doc, err := goquery.NewDocument(linkDownload)
	if err != nil {
		return nil, err
	}

	doc.Find("#downloadlink").Find("a[href]").Each(func(i int, s *goquery.Selection) {
		fmt.Println("a")
		a, _ := s.Attr("href")
		fmt.Println(a)
		if strings.Contains(a, "/320/") {
			listStream = append(listStream, a)
		}

	})

	return listStream, nil
}
