package nhaccuatui

import (
"fmt"
"net/http"
"strings"

"github.com/PuerkitoBio/goquery"
"github.com/astaxie/beego/httplib"
)
const (
	song string ="http://www.nhaccuatui.com/bai-hat/"
	album string ="http://www.nhaccuatui.com/playlist/"
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

func GetDirectLink(link string)(string[],error) {
	if link == "" {
		panic("error")
		return
	}

	// implement
	var listStream []string
	if strings.Contains(link, song) {
		urlList := strings.Split(link, ".")

		req := httplib.Get(linkDownloadSong + urlList[3])

		var res ResponseNhacCuaTui
		err := req.ToJson(&res)
		if err != nil {
			panic(err)
			return
		}
		listStream = append(listStream,res.Data.StreamUrl)
		return listStream,nill
		// c.Data["json"] = res.Data.StreamUrl
		// c.ServeJson()
		// return
	}
	if strings.Contains(link, album) {
		// urlList := strings.Split(link, ".")
		
		doc, err := goquery.NewDocument(link)
		if err != nil {
			panic(err)
		}
		doc.Find(".item_content").Each(func(i int, s *goquery.Selection) {
			a, _ := s.Find("a").Attr("href")
			urlList := strings.Split(a, ".")

			req := httplib.Get(linkDownloadSong + urlList[3])

			var res ResponseNhacCuaTui
			err := req.ToJson(&res)
			if err != nil {
				panic(err)
				return
			}
			listStream = append(listStream, res.Data.StreamUrl)
			})
		return listStream,nill
	}
}
