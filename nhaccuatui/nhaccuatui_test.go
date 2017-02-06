package nhaccuatui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var nct NhacCuaTui

//Unit Test
func TestLinkEmptyString(t *testing.T) {
	_, err := nct.GetDirectLink("")
	assert.Equal(t, "Empty Link", err.Error())
}

func TestSplitStringOutOfRange(t *testing.T) {
	_, err := nct.GetDirectLink("http://www.nhaccuatui.com/bai-hat/")
	assert.Equal(t, "Wrong Format Link", err.Error())
}

func TestLinkSongIsInvalid(t *testing.T) {
	_, err := nct.GetDirectLink("http://www.nhaccuatui.com/bai-hat/boi-vi-em-het-yeu-anh-chi-dan.Z4Y057sG.html")
	assert.Equal(t, "Invalid Link", err.Error())
}

func TestLinkAlbumIsInvalid(t *testing.T) {
	_, err := nct.GetDirectLink("http://www.nhaccuatui.com/playlist/love-me-right-repackage-album-exo.AIG6kI0.html")
	assert.Equal(t, "Invalid Link", err.Error())
}

func TestDownloadSongSuccess(t *testing.T) {
	_, err := nct.GetDirectLink("http://www.nhaccuatui.com/bai-hat/cho-em-gan-anh-them-chut-nua-huong-tram.zcHcXTHdZsSD.html")
	assert.Nil(t, err, "We are expecting nil error here")
}

func TestDownloadAlbumSuccess(t *testing.T) {
	_, err := nct.GetDirectLink("http://www.nhaccuatui.com/playlist/anh-yeu-em-single-pham-truong.G5cC3EbnxuPs.html")
	assert.Nil(t, err, "We are expecting nil error here")
}
