package nhaccuatui

import (
	"errors"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/astaxie/beego/httplib"
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
		return nil, nil
	}

	// implement
	var listStream []string
	if strings.Contains(link, song) {
		urlList := strings.Split(link, ".")
		if len(urlList) < 4 {
			return nil, errors.New("Invalid link")
		}
		req := httplib.Get(linkDownloadSong + urlList[3])

		var res ResponseNhacCuaTui
		err := req.ToJson(&res)
		if err != nil {
			return nil, err
		}
		listStream = append(listStream, res.Data.StreamUrl)
	}

	if strings.Contains(link, album) {
		doc, err := goquery.NewDocument(link)
		if err != nil {
			return nil, err
		}
		// read all html, find class "item_content" then file tag a, get href infomation
		doc.Find(".item_content").Each(func(i int, s *goquery.Selection) {
			a, _ := s.Find("a").Attr("href")
			urlList := strings.Split(a, ".")
			if len(urlList) < 4 {
				return nil, errors.New("Invalid link")
			}
			req := httplib.Get(linkDownloadSong + urlList[3])

			var res ResponseNhacCuaTui
			req.ToJson(&res)
			listStream = append(listStream, res.Data.StreamUrl)
		})
	}
	return listStream, nil

}
