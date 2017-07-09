package soundcloud

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/dwarvesf/glod"
)

// Soundcloud const
const (
	APILink     string = "http://api.soundcloud.com/resolve.json?"
	MediaLink   string = "http://media.soundcloud.com/stream/"
	ClientID    string = "70faad763a2546deb1bdabc3a6bfa722"
	urlSong     int    = 5
	urlPlaylist int    = 6
)

// ResponseSoundCloud is a struct that will be returned by soundcloud
type ResponseSoundCloud struct {
	StreamURL string `json:"stream_url"`
	TrackList []struct {
		StreamURL string `json:"stream_url"`
		Title     string `json:"title"`
	} `json:"tracks"`
	Title  string `json:"title"`
	UserID int    `json:"id"`
}

// SoundCloud ...
type SoundCloud struct {
}

// GetDirectLink input is a link then return an object contains download url, song title and artist
func (s *SoundCloud) GetDirectLink(link string) ([]glod.Response, error) {
	var listSong []glod.Response
	var song glod.Response
	var linkRequest = APILink + "url=" + link + "&client_id=" + ClientID

	var res ResponseSoundCloud
	response, err := http.Get(linkRequest)
	if err != nil {
		return nil, errors.New("Error download this song")
	}
	defer response.Body.Close()

	// server forbidden, get song title only
	if response.StatusCode == 403 {
		song.Title = crawlTitleFromURL(link)
		listSong = append(listSong, song)
	} else {

		buffer, _ := ioutil.ReadAll(response.Body)

		err = json.Unmarshal(buffer, &res)
		if err != nil {
			return nil, errors.New("Error parsing")
		}

		if len(strings.Split(link, "/")) == urlSong {
			if res.StreamURL == "" {
				return nil, errors.New("This song is not streamable")
			}

			url := res.StreamURL + "?client_id=" + ClientID

			song.StreamURL = url
			song.Title = res.Title
			listSong = append(listSong, song)

		} else if len(strings.Split(link, "/")) == urlPlaylist {
			for i := range res.TrackList {
				if res.TrackList[i].StreamURL != "" {
					url := res.TrackList[i].StreamURL + "?client_id=" + ClientID

					song.StreamURL = url
					song.Title = res.TrackList[i].Title
					listSong = append(listSong, song)
					if strings.Contains(song.Title, "-") {
						tmp := strings.Split(song.Title, "-")

						song.Title = tmp[0]
						song.Artist = tmp[1]
					}
				}
			}
		} else {
			return nil, errors.New("Wrong Format Link")
		}
	}
	return listSong, nil
}

// crawlTitleFromURL return a song title from given url
func crawlTitleFromURL(link string) string {
	_title := strings.Split(link, "/")
	title := strings.Split(_title[4], "-ft-")

	title[0] = strings.Replace(title[0], "-", " ", -1)

	return title[0]
}
