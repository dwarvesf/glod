package vimeo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var vimeo Vimeo

func TestLinkEmptyString(t *testing.T) {
	_, err := vimeo.GetDirectLink("")
	assert.Equal(t, "Empty Link", err.Error())
}

func TestDownloadSuccess(t *testing.T) {
	_, err := vimeo.GetDirectLink("https://vimeo.com/150178241")
	assert.Nil(t, err, "We are expecting nil error here")
}
