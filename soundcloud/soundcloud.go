package soundcloud

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	ApiLink   string = "http://api.soundcloud.com/resolve.json?"
	MediaLink string = "http://media.soundcloud.com/stream/"
	CLIENT_ID        = "70faad763a2546deb1bdabc3a6bfa722"
)

type ReponseSoundCloud struct {
	WaveFormUrl string `json:"waveform_url"`
	StreamURL   string `json:"stream_url"`
}

type SoundCloud struct {
}

func (s *SoundCloud) GetDirectLink(link string) ([]string, error) {
	if link == "" {
		return nil, errors.New("Empty Link")
	}
	var listStream []string
	var linkRequest = ApiLink + "url=" + link + "&client_id=" + CLIENT_ID

	var res ReponseSoundCloud
	response, err := http.Get(linkRequest)
	if err != nil {
		return nil, errors.New("Error while downloading")
	}

	defer response.Body.Close()

	buffer, _ := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(buffer, &res)
	if err != nil {
		return nil, errors.New("Error parsing")
	}

	if res.StreamURL == "" {
		return nil, errors.New("This song is not streamable")
	}

	// urlList := strings.Split(res.WaveFormUrl, "/")
	// if len(urlList) < 4 {
	// return nil, errors.New("Wrong Format Link")
	// }

	// longString := urlList[3]
	// substring := longString[0 : len(longString)-6]

	// listStream = append(listStream, MediaLink+substring)

	url := res.StreamURL + "?client_id=" + CLIENT_ID
	listStream = append(listStream, url)

	return listStream, nil
}
