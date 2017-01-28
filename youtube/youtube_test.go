package youtube

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var youtube Youtube

func TestLinkEmptyString(t *testing.T) {
	_, err := youtube.GetDirectLink("")
	assert.Equal(t, "Empty Link", err.Error())
}

func TestFetchMetaWithWrongVideoId(t *testing.T) {
	_, err := fetchMeta("wrongvideo")
	assert.Equal(t, "Invalid Video Id", err.Error())
}
