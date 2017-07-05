package nhaccuatui

type NhacCuaTui struct {
}

type Response struct {
	Artist    string
	StreamURL string
	Title     string
}

const (
	song  string = "http://www.nhaccuatui.com/bai-hat/"
	album string = "http://www.nhaccuatui.com/playlist/"

	isDownloadSong     string = "http://www.nhaccuatui.com/flash/xml?html5=true&key1"
	isDownloadPlaylist string = "http://www.nhaccuatui.com/flash/xml?html5=true&key2"
)

// TODO: code nhaccuatui
// function that input is a link then return an slice of url that permantly download file and error(if it has)
func (nct *NhacCuaTui) GetDirectLink(link string) ([]Response, error) {
	var listSong []Response
	return listSong, nil
}
