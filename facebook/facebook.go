package facebook

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Facebook struct {
}

func (fb *Facebook) GetDirectLink(link string) ([]string, error) {
	if link == "" {
		return nil, errors.New("Empty Link")
	}
	var listStream []string

	response, err := http.Get(link)
	if err != nil {
		fmt.Println("Link is invalid")
		return nil, err
	}
	defer response.Body.Close()
	buffer, _ := ioutil.ReadAll(response.Body)

	parseString := string(buffer)
	if !strings.Contains(parseString, "sd_src") {
		return nil, errors.New("This is private video or invalid link")
	}
	splitSrcSdSource := strings.Split(parseString, "sd_src")

	splitLinkNotSanitize := strings.Split(splitSrcSdSource[len(splitSrcSdSource)-1], "hd_tag")

	linkShort := splitLinkNotSanitize[0]

	linkShort = linkShort[3 : len(linkShort)-3]

	linkSanitized := strings.Replace(linkShort, "\\/", "/", -1)

	listStream = append(listStream, linkSanitized)

	return listStream, nil

}
