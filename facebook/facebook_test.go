package facebook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var fb Facebook

func TestLinkEmptyString(t *testing.T) {
	_, err := fb.GetDirectLink("")
	assert.Equal(t, "Empty Link", err.Error())
}

func TestDownload(t *testing.T) {
	_, err := fb.GetDirectLink("https://www.facebook.com/pokerorganization/videos/508013846071164/")
	assert.Nil(t, err, "We are expecting nil error here")
}
