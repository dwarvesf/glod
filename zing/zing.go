package zing

import (
	"strings"

	"github.com/ivkean/MDwn/engine"
	"github.com/k0kubun/pp"
)

const (
	song  string = "http://v3.mp3.zing.vn/download/vip/song/"
	album string = "http://mp3.zing.vn/album/"
)

func GetDirectLink(songID string) (string, error) {
	// get direct link
	url := "http://v3.mp3.zing.vn/download/vip/song/" + songID

	DirectLink, err := engine.Request(url, "GET")
	if err != nil {
		return "", err
	}

	pp.Println(DirectLink)
	url = strings.Replace(DirectLink.HttpResponse.Request.URL.String(), " ", "%20", -1)
	// res, err := engine.Request(url, "GET")
	// d1 := []byte(res.Body)
	// err = ioutil.WriteFile(file, d1, 0644)
	// if err != nil {
	// 	return err
	// }

	return url, nil
}
