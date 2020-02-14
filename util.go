package y2bcaptions

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func formatDuration(d time.Duration) (duration string) {
	return fmt.Sprintf("%02d:%02d:%02d,%03d", d/time.Hour, d%time.Hour/time.Minute, d%time.Minute/time.Second, d%time.Second/time.Millisecond)
}

var ErrGetUrlFailed = errors.New("get from url failed")

func getFromUrl(url string) (data []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,zh-TW;q=0.8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, ErrGetUrlFailed
	}
	return ioutil.ReadAll(resp.Body)
}

// this function is borrowed form github.com/rylio/ytdl/video_info.go
// GetVideoInfoFromShortURL fetches video info from a short youtube url
func extractVideoID(u *url.URL) string {
	switch u.Host {
	case "www.youtube.com", "youtube.com", "m.youtube.com":
		if u.Path == "/watch" {
			return u.Query().Get("v")
		}
		if strings.HasPrefix(u.Path, "/embed/") {
			return u.Path[7:]
		}
	case "youtu.be":
		if len(u.Path) > 1 {
			return u.Path[1:]
		}
	}
	return ""
}
