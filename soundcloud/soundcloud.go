package soundcloud

import "github.com/dwarvesf/glod"

type Response struct {
	Artist    string
	StreamURL string
	Title     string
}

// Soundcloud const
const (
	ApiLink     string = "http://api.soundcloud.com/resolve.json?"
	MediaLink   string = "http://media.soundcloud.com/stream/"
	CLIENT_ID   string = "70faad763a2546deb1bdabc3a6bfa722"
	urlSong     int    = 5
	urlPlaylist int    = 6
)

type ReponseSoundCloud struct {
	StreamURL string `json:"stream_url"`
	TrackList []struct {
		StreamURL string `json:"stream_url"`
		Title     string `json:"title"`
	} `json:"tracks"`
	Title  string `json:"title"`
	UserID int    `json:"id"`
}

type SoundCloud struct {
}

func (s *SoundCloud) GetDirectLink(link string) ([]glod.Response, error) {
	var listSong []glod.Response
	return listSong, nil
}
