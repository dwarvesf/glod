package zing

import "testing"

var zing Zing

func TestLinkInputIsEmpty(t *testing.T) {
	t.Log("Try to empty link, expect error output")
	_, err := zing.GetDirectLink("")
	if err == nil {
		t.Error("Expected error of not nil but it was nil instead")
	}
}

func TestSplitStringOutOfRange(t *testing.T) {
	t.Log("Try to input invalid link, expect error output")
	_, err := zing.GetDirectLink("http://mp3.zing.vn/bai-hat/Lam-Vo-Anh-Nhe-Chi-Dan")
	if err == nil {
		t.Error("Expected error of not nil but it was nil instead")
	}
}

func TestLinkAlbumIsInvalid(t *testing.T) {
	t.Log("Try to input invalid link, expect empty listStream return")
	listStream, _ := zing.GetDirectLink("http://mp3.zing.vn/album/Noi-Em-Cho-Anh-Dan-Kim/ZDOA.html")
	if len(listStream) > 0 {
		t.Errorf("Expected length of listStream equals 0 but it is %d instead", len(listStream))
	}
}

func TestDownloadSongSuccess(t *testing.T) {
	t.Log("Try to input valid song link, expect length of listStream return bigger than 0")
	listStream, _ := zing.GetDirectLink("http://mp3.zing.vn/bai-hat/Lam-Vo-Anh-Nhe-Chi-Dan/ZW7IU0UC.html")
	if listStream == nil || len(listStream) == 0 {
		t.Errorf("Expected length of listStream is bigger than 0 but it is %d instead", len(listStream))
	}
}

func TestDownloadAlbumSuccess(t *testing.T) {
	t.Log("Try to input valid album link, expect length of listStream return bigger than 0")
	listStream, _ := zing.GetDirectLink("http://mp3.zing.vn/album/Noi-Em-Cho-Anh-Dan-Kim/ZWZC0DOA.html")
	if listStream == nil || len(listStream) == 0 {
		t.Errorf("Expected length of listStream is bigger than 0 but it is %d instead", len(listStream))
	}
}
