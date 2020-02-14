package main

import (
	"flag"
	"fmt"
	"github.com/treant5612/y2bcaptions"
	"log"
	"strings"
)

var (
	vid = flag.String("vid","","list captions of specific videoId")
	fname = flag.String("file","subtitle","download file name")
	)

func main(){
	flag.Parse()
	if *vid!=""{
		captions,err:= y2bcaptions.GetSubtitleInfoFromId(*vid)
		if err!=nil{
			log.Fatal("get caption info failed:",err)
		}
		if len(captions.CaptionTrack)==0{
			log.Fatal("no caption for specific video")
		}
		fmt.Println(strings.Repeat("-",48))
		for i,c := range captions.CaptionTrack{
			fmt.Printf("%d:\tlang=%s[%s]\tkind=%s\n",i,c.LanguageCode,c.Name,c.Kind)
		}
		fmt.Printf("input a number and translation option to download \n")
		fmt.Printf("for example : 1,zh\n")
		var n int
		var tlang string
		fmt.Scanf("%d,%s",&n,&tlang)
		err =captions.DownloadToFile(captions.CaptionTrack[n].VssId,tlang,*fname+".srt")
		if err!=nil{
			log.Fatal(err)
		}
	}
}
