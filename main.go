package main

import (
	"strings"

	"github.com/ivkean/MDwn/zing"
)

func main() {
	url := "http://mp3.zing.vn/bai-hat/Em-Cua-Qua-Khu-Nguyen-Dinh-Vu/ZW7009IZ.html"

	surl := strings.Split(url, ".")
	if len(surl) < 2 {
		panic("Wrong Url!")
	}

	domain := surl[1]

	switch domain {
	case "zing":
		// check to get songID
		if len(surl) < 3 {
			panic("Wrong Url!")
		}
		pathSongID := surl[2]

		// get songID
		sPathSongID := strings.Split(pathSongID, "/")
		if len(sPathSongID) < 4 {
			panic("Wrong Url!")
		}
		songID := sPathSongID[3]

		zing.GetDirectLink(songID)
	}
}
