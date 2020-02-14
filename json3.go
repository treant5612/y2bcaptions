package y2bcaptions

import (
	"bytes"
	"fmt"
	"time"
)

type Json3 struct {
	WireMagic      string  `json:"wireMagic"`
	Pens           JsonObj `json:"pens"`
	WsWinStyles    JsonObj `json:"wsWinStyles"`
	WpWinPositions JsonObj `json:"wpWinPositions"`
	Events         []Event `json:"events"`
}

type Event struct {
	TStartMs    int64 `json:"tStartMs"`
	DDurationMs int64 `json:"dDurationMs"`
	Segs        []Seg `json:"segs"`
}

type Seg struct {
	Utf8 string `json:"utf8"`
}

type JsonObj []map[string]interface{}

func (j *Json3) ToSrt() (b []byte) {
	buf := new(bytes.Buffer)
	n := 0
	for _, event := range j.Events {
		start, end := formatDuration(time.Duration(event.TStartMs)*time.Millisecond), formatDuration(time.Duration(event.TStartMs+event.DDurationMs)*time.Millisecond)
		seg := ""
		for _, s := range event.Segs {
			seg += s.Utf8
		}
		buf.WriteString(fmt.Sprintf("%d\n%s --> %s\n%s\n\n", n, start, end, seg))
		n++
	}
	return buf.Bytes()
}
