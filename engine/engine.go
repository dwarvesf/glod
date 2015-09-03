package engine

import (
	"log"
	"time"

	"github.com/nareix/curl"
)

func Request(url, method string) (res curl.Response, err error) {
	req := curl.New(url)

	req.Method(method) // can be "PUT"/"POST"/"DELETE" ...

	req.DialTimeout(time.Second * 10) // TCP Connection Timeout
	req.Timeout(time.Second * 30)     // Download Timeout

	// Print progress status per one second
	req.Progress(func(p curl.ProgressStatus) {
		log.Println(
			"Stat", p.Stat, // one of curl.Connecting / curl.Downloading / curl.Closed
			"speed", curl.PrettySpeedString(p.Speed),
			"len", curl.PrettySizeString(p.ContentLength),
			"got", curl.PrettySizeString(p.Size),
			"percent", p.Percent,
			"paused", p.Paused,
		)
	}, time.Second)
	/*
		2015/05/20 15:34:15 Stat 2 speed 0.0B/s len 78.5M got 0.0B percent 0 paused true
		2015/05/20 15:34:16 Stat 2 speed 0.0B/s len 78.5M got 0.0B percent 0 paused true
		2015/05/20 15:34:16 Stat 2 speed 394.1K/s len 78.5M got 197.5K percent 0.0024564497 paused false
		2015/05/20 15:34:17 Stat 2 speed 87.8K/s len 78.5M got 241.5K percent 0.0030038392 paused false
		2015/05/20 15:34:17 Stat 2 speed 79.8K/s len 78.5M got 281.5K percent 0.003501466 paused false
		2015/05/20 15:34:18 Stat 2 speed 63.9K/s len 78.5M got 313.5K percent 0.0038995675 paused false
	*/

	return req.Do()
}
