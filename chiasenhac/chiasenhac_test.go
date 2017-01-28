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
	_, err := csn.GetDirectLink("http://chiasenhac.com/nhac-hot-2/we-dont-talk-anymore~charlie-puth-selena-gomez~1621445.html")
	assert.Nil(t, err, "We are expecting nil error here")
}
