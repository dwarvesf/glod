package vimeo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Vimeo struct {
}

func (vimeo *Vimeo) GetDirectLink(link string) ([]string, error) {
	if link == "" {
		return nil, errors.New("Empty Link")
	}

	var listStream []string
	splitVideoIds := strings.Split(link, "/")
	videoId := splitVideoIds[len(splitVideoIds)-1]
	var linkDownload = "https://player.vimeo.com/video/" + videoId

	response, err := http.Get(linkDownload)
	if err != nil {
		fmt.Println("Link is invalid")
		return nil, err
	}
	defer response.Body.Close()

	buffer, _ := ioutil.ReadAll(response.Body)
	parseString := string(buffer)

	splitStringSplitInitial := strings.Split(parseString, "video/mp4")

	splitStringUrlStart := strings.Split(splitStringSplitInitial[1], "url")
	splitStringUrlStop := strings.Split(splitStringUrlStart[1], "cdn")
	linkStream := splitStringUrlStop[0][3:] + "cdn" + splitStringUrlStop[1][:len(splitStringUrlStop[1])-3]

	splitStringTitleStart := strings.Split(parseString, "<title>")
	splitstringTitleStop := strings.Split(splitStringTitleStart[1], "</title>")

	listStream = append(listStream, linkStream+"~"+strings.Replace(splitstringTitleStop[0], "//", "", -1))
	return listStream, nil

}
