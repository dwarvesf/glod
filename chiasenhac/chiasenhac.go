package chiasenhac

import "github.com/dwarvesf/glod"

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

// TODO: test chiesenhac
func (csn *ChiaSeNhac) GetDirectLink(link string) ([]glod.Response, error) {
	var listSong []glod.Response
	return listSong, nil
}
