package facebook

type Response struct {
	Artist    string
	StreamURL string
	Title     string
}

type Facebook struct {
}

// TODO : code facebook
func (fb *Facebook) GetDirectLink(link string) ([]Response, error) {
	var listSong []Response
	return listSong, nil
}
