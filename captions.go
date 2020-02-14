package y2bcaptions

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"regexp"
)

type Captions struct {
	VideoId      string
	BaseUrl      string `json:"baseUrl"`
	CaptionTrack []*CaptionTrack
}
type CaptionTrack struct {
	BaseUrl        string `json:"baseUrl"`
	IsTranslatable bool   `json:"isTranslatable"`
	Kind           string `json:"kind"`
	LanguageCode   string `json:"languageCode"`
	Name           string `json:name.simpleText`
	VssId          string `json:"vssId"`
}

func GetSubtitleInfoFromUrl(videoUrl string) (caption *Captions, err error) {
	u, err := url.Parse(videoUrl)
	if err != nil {
		return nil, err
	}
	videoId := extractVideoID(u)
	return GetSubtitleInfoFromId(videoId)
}

func GetSubtitleInfoFromId(videoId string) (caption *Captions, err error) {
	url := fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoId)
	content, err := getFromUrl(url)
	if err != nil {
		return nil, err
	}
	return getCaptionInfoFromHtml(videoId, content)
}

var ErrParseHtmlFailed = errors.New("parse html failed")
var ErrCaptionNotFound = errors.New("no available captions")
var ytInitialPlayerResponse = regexp.MustCompile(`\["ytInitialPlayerResponse"\] = (.+);`)

func getCaptionInfoFromHtml(videoId string, html []byte) (captions *Captions, err error) {
	matches := ytInitialPlayerResponse.FindSubmatch(html)
	if len(matches) < 1 {
		return nil, ErrParseHtmlFailed
	}
	responseContext := new(ResponseContext)
	err = json.Unmarshal(matches[1], responseContext)
	if err != nil {
		return nil, err
	}
	if len(responseContext.Captions.PlayerCaptionsTracklistRenderer.CaptionTracks) == 0 {
		return nil, ErrCaptionNotFound
	}
	captions = &Captions{
		VideoId:      videoId,
		BaseUrl:      responseContext.Captions.PlayerCaptionsRenderer.BaseUrl,
		CaptionTrack: nil,
	}
	for _, c := range responseContext.Captions.PlayerCaptionsTracklistRenderer.CaptionTracks {
		captionTrack := &CaptionTrack{
			BaseUrl:        c.BaseUrl,
			IsTranslatable: c.IsTranslatable,
			Kind:           c.Kind,
			LanguageCode:   c.LanguageCode,
			Name:           c.Name.SimpleText,
			VssId:          c.VssId,
		}
		captions.CaptionTrack = append(captions.CaptionTrack, captionTrack)
	}
	return captions, nil

}

func (c *Captions) DownloadToFile(vssId string, tlang string, filePath string) (err error) {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	return c.Download(vssId, tlang, f)
}

var ErrWrongVssIdParam = errors.New("vssId not exists")

func (c *Captions) Download(vssId string, tlang string, writer io.Writer) (err error) {
	var cap *CaptionTrack
	for i := range c.CaptionTrack {
		if c.CaptionTrack[i].VssId == vssId {
			cap = c.CaptionTrack[i]
		}
	}
	if cap == nil {
		return ErrWrongVssIdParam
	}
	var url = fmt.Sprintf("%s&lang=%s&fmt=json3", c.BaseUrl, cap.LanguageCode)
	if cap.Kind == "asr" {
		url += "&kind=asr"
	}
	if tlang != "" {
		url += "&tlang=" + tlang
	}
	json3Bytes, err := getFromUrl(url)
	if err != nil {
		return err
	}
	json3 := new(Json3)
	err = json.Unmarshal(json3Bytes, json3)
	if err != nil {
		return err
	}
	_, err = writer.Write(json3.ToSrt())
	return
}
