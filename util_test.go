package y2bcaptions

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestConvertDuration(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<rand.Intn(1000) ;i++{
		var h, m, s, ms = rand.Intn(24), rand.Intn(60), rand.Intn(60), rand.Intn(1000)
		d := time.Hour*time.Duration(h) + time.Minute*time.Duration(m) + time.Second*time.Duration(s) + time.Millisecond*time.Duration(ms)
		if r := formatDuration(d); r != fmt.Sprintf("%02d:%02d:%02d,%03d", h, m, s, ms) {
			t.Logf("h=%d,m=%d,s=%d,ms=%d,wrong format result:%v", h, m, s, ms, r)
			t.Fail()
		}
	}
}
