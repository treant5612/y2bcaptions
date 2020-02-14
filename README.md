# youtube-caption-dl
### download youtube captions.
inspired by github.com/rylio/ytdl

### how to 
for inline use
```go
    
    captions,_ := y2bcaptions.GetSubtitleInfoFromId( ${videoId} )
    // or captions ,_ := y2bcaptions.GetSubtitleInfoFromUrl( ${url} )
    
    
    vssId := captinos.CaptionTrack[0].VssId
    tlang := "zh-Hans"
    file  := "subtitle.srt"

    captions.DownloadToFile(vssId,tlang,file)
   
```

for cmd
```
    ./y2bcap -vid ${vid} -file ${filename}
    // -> you will get a list of captions,type the no and translate option to download
    
```