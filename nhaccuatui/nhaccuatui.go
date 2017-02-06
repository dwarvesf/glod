package nhaccuatui

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type NhacCuaTui struct {
}

const (
	song  string = "http://www.nhaccuatui.com/bai-hat/"
	album string = "http://www.nhaccuatui.com/playlist/"

	isDownloadSong     string = "http://www.nhaccuatui.com/flash/xml?html5=true&key1"
	isDownloadPlaylist string = "http://www.nhaccuatui.com/flash/xml?html5=true&key2"
)

// function that input is a link then return an slice of url that permantly download file and error(if it has)
func (nct *NhacCuaTui) GetDirectLink(link string) ([]string, error) {
	if link == "" {
		return nil, errors.New("Empty Link")
	}

	urlList := strings.Split(link, ".")
	if len(urlList) < 4 {
		return nil, errors.New("Wrong Format Link")
	}

	if len(urlList[3]) != 12 {
		return nil, errors.New("Invalid Link")
	}

	var listStream []string
	res, _ := http.Get(link)
	defer res.Body.Close()
	buffer, _ := ioutil.ReadAll(res.Body)

	parseString := string(buffer)

	xmlURL := strings.Split(parseString, "player.peConfig.xmlURL = \"")
	_xmlURL := strings.Split(xmlURL[1], "\";")

	if (strings.Contains(_xmlURL[0], isDownloadSong) != true) && (strings.Contains(_xmlURL[0], isDownloadPlaylist) != true) {
		return nil, errors.New("Invalid Link")
	}

	res, _ = http.Get(_xmlURL[0])
	buffer, _ = ioutil.ReadAll(res.Body)

	parseString = string(buffer)

	if strings.Contains(link, song) {
		directLink := strings.Split(parseString, "http://")
		_directLink := strings.Split(directLink[1], "]]>")
		_directLink[0] = "http://" + _directLink[0]

		listStream = append(listStream, _directLink[0])
	}

	if strings.Contains(link, album) {
		locationLink := strings.Split(parseString, "<location>")
		for index := range locationLink {
			if index > 1 {
				directLink := strings.Split(locationLink[index], "http://")
				_directLink := strings.Split(directLink[1], "]]>")
				_directLink[0] = "http://" + _directLink[0]

				listStream = append(listStream, _directLink[0])
			}
		}
	}
	if len(listStream) == 0 {
		return nil, errors.New("Invalid Link")
	}
	return listStream, nil

}
