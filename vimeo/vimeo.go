package vimeo

type Vimeo struct {
}

// Response
type Response struct {
	Artist    string
	StreamURL string
	Title     string
}

// TODO: code vimeo
func (vimeo *Vimeo) GetDirectLink(link string) ([]Response, error) {
	var listSong []Response
	return listSong, nil
}
