package youtube

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var youtube Youtube

// func TestDownloadSingleVideoSuccess(t *testing.T) {
// 	t.Log("Try to download valid video from youtube... expecting error is nil")
// 	_, err := DownloadSingleVideo("glhPVLwZZzA")
// 	if err != nil {
// 		t.Errorf("Expected error of nil, but it was %s instead", err.Error())
// 	}
// }
func TestLinkEmptyString(t *testing.T) {
	_, err := youtube.GetDirectLink("")
	assert.Equal(t, "Empty Link", err.Error())
}
