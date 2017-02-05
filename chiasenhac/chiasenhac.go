package chiasenhac

import (
	"errors"
	"io/ioutil"
	"net/http"
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

	res, err := http.Get(linkDownload)
	if err != nil {
		return nil, errors.New("Invalid Link")
	}
	defer res.Body.Close()
	buffer, _ := ioutil.ReadAll(res.Body)
	parseString := string(buffer)

	_LinkTmp := strings.Split(parseString, "document.write")
	_LinkTmp2 := strings.Split(_LinkTmp[2], "href=\"")
	_LinkTmp3 := strings.Split(_LinkTmp2[1], "\" onmouseover")
	_LinkTmp4 := strings.Split(_LinkTmp3[0], " ")
	directLink := _LinkTmp4[0] + _LinkTmp4[1] + _LinkTmp4[2]

	listStream = append(listStream, directLink)
	return listStream, nil
}
