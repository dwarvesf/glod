package chiasenhac

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var csn ChiaSeNhac

func TestLinkEmptyString(t *testing.T) {
	_, err := csn.GetDirectLink("")
	assert.Equal(t, "Empty Link", err.Error())
}

func TestDownloadSucces(t *testing.T) {
	_, err := csn.GetDirectLink("http://m1.chiasenhac.vn/mp3/vietnam/v-pop/ai-ra-xu-hue~quang-le-le-minh-trung~ts3vtm07q2f8nt.html")
	assert.Nil(t, err, "We are expecting nil error here")
}
