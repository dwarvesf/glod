package youtube

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Youtube struct {
}
type Video struct {
	Id, Title, Author, Keywords, Thumbnail_url string
	Avg_rating                                 float32
	View_count, Length_seconds                 int
	Formats                                    []Format
}

const (
	linkDownload        = "http://www.youtube.com/get_video_info?&video_id="
	album        string = "playlist"
)

var FORMATS []string = []string{"3gp", "mp4", "flv", "webm", "avi"}

type Format struct {
	Itag                     int
	Video_type, Quality, Url string
}

// function that download single video
func DownloadSingleVideo(video_id string) ([]string, error) {
	query_string, err := fetchMeta(video_id)
	if err != nil {
		return nil, err
	}
	video, err := parseMeta(video_id, query_string)

	if err != nil {
		return nil, err
	}
	//generate file name
	filename := video.Title + "." + video.GetExtension(0)
	out, err := os.Create(filename)
	defer out.Close()

	if err != nil {
		return nil, errors.New("Unable to write to file " + filename)
	}
	resp, err := http.Get(video.Formats[0].Url)
	defer resp.Body.Close()

	if err != nil {
		return nil, errors.New("Unable to download video content from Yotutube")
	}
	io.Copy(out, resp.Body)
	return nil, nil
}

// function that receive input is a link and output doesnt matter(but it override GetDirectLink of Glod interface)
func (youtube *Youtube) GetDirectLink(link string) ([]string, error) {
	if link == "" {
		return nil, nil
	}

	if strings.Contains(link, album) {
		var listLink []string
		doc, err := goquery.NewDocument(link)
		if err != nil {
			return nil, err
		}

		doc.Find(".pl-video").Each(func(i int, s *goquery.Selection) {
			a, _ := s.Attr("data-video-id")
			listLink = append(listLink, a)
		})
		for i := 0; i < len(listLink); i++ {
			go DownloadSingleVideo(listLink[i])
		}
		return nil, nil
	}

	urlList := strings.Split(link, "/")
	if len(urlList) < 4 {
		return nil, errors.New("Invalid link")
	}
	_videoId := urlList[3]
	video_id := _videoId[8:len(_videoId)]
	DownloadSingleVideo(video_id)

	return nil, nil
}

// return extension of video
func (video *Video) GetExtension(index int) string {
	for i := 0; i < len(FORMATS); i++ {
		if strings.Contains(video.Formats[index].Video_type, FORMATS[i]) {
			return FORMATS[i]
		}
	}

	return "avi"
}

// function readall body of request and return string body
func fetchMeta(video_id string) (string, error) {
	resp, err := http.Get(linkDownload + video_id)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	query_string, _ := ioutil.ReadAll(resp.Body)

	return string(query_string), nil
}

//function parse string to Video struct
func parseMeta(video_id, query_string string) (Video, error) {
	u, _ := url.Parse("?" + query_string)

	query := u.Query()

	if query.Get("errorcode") != "" || query.Get("status") == "fail" {
		return Video{}, errors.New(query.Get("reason"))
	}

	video := Video{
		Id:            video_id,
		Title:         query.Get("title"),
		Author:        query.Get("author"),
		Keywords:      query.Get("keywords"),
		Thumbnail_url: query.Get("thumbnail_url"),
	}

	v, _ := strconv.Atoi(query.Get("view_count"))
	video.View_count = v

	r, _ := strconv.ParseFloat(query.Get("avg_rating"), 32)
	video.Avg_rating = float32(r)

	l, _ := strconv.Atoi(query.Get("length_seconds"))
	video.Length_seconds = l

	format_params := strings.Split(query.Get("url_encoded_fmt_stream_map"), ",")

	for _, f := range format_params {
		furl, _ := url.Parse("?" + f)
		fquery := furl.Query()

		itag, _ := strconv.Atoi(fquery.Get("itag"))

		video.Formats = append(
			video.Formats,
			Format{
				Itag:       itag,
				Video_type: fquery.Get("type"),
				Quality:    fquery.Get("quality"),
				Url:        fquery.Get("url") + "&signature=" + fquery.Get("sig"),
			})
	}

	return video, nil
}
