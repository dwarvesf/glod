package soundcloud

import (
	"errors"
	"os"
	"strings"

	"github.com/astaxie/beego/httplib"
)

const (
	ApiLink   string = "http://api.soundcloud.com/resolve.json?"
	MediaLink string = "http://media.soundcloud.com/stream/"
)

type ReponseSoundCloud struct {
	WaveFormUrl string `json:"waveform_url"`
}

type SoundCloud struct {
}

func (s *SoundCloud) GetDirectLink(link string) ([]string, error) {
	if link == "" {
		return nil, errors.New("Empty Link")
	}
	var listStream []string
	var linkRequest = ApiLink + "url=" + link + "&client_id=" + os.Getenv("SOUNDCLOUD_CLIENT_ID")

	var res ReponseSoundCloud
	req := httplib.Get(linkRequest)
	req.ToJson(&res)

	if res.WaveFormUrl == "" {
		return nil, errors.New("This song is not streamable")
	}

	urlList := strings.Split(res.WaveFormUrl, "/")
	if len(urlList) < 4 {
		return nil, errors.New("Wrong Format Link")
	}

	longString := urlList[3]
	substring := longString[0 : len(longString)-6]

	listStream = append(listStream, MediaLink+substring)
	return listStream, nil
}