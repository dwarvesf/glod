package chiasenhac

const (
	CsnPrefix         = "chiasenhac"
	csnDownloadSuffix = "download.html"
	csnAblum          = "chiasenhac.vn/nghe-album/"
	csnSong1          = "chiasenhac.vn/mp3"
	csnSong2          = "chiasenhac.vn/nhac-hot"

	csnStreamURL = "/128/"
)

type ChiaSeNhac struct {
}

type Response struct {
	Artist    string
	StreamURL string
	Title     string
}

// TODO: test chiesenhac
func (csn *ChiaSeNhac) GetDirectLink(link string) ([]Response, error) {
	var listSong []Response
	return listSong, nil
}
