package main

import (
	"flag"
	"fmt"
	"sub"
	"log"
)

const (
	defaultDuraion = iota
)

type EditParams struct {
	advance bool
	hour float64
	minute float64
	second float64
}

func main() {
	params := new(EditParams)
	var advance = flag.Bool("advance",false,"advance or back the timeline")
	var duration = flag.Duration("duration",defaultDuraion,"duration which you want to change")
	flag.Parse()
	if *advance == true {
		params.advance = true
	} else {
		params.advance = false
	}
	params.hour = duration.Hours()
	params.minute = duration.Minutes()
	params.second = duration.Seconds()
	subtitle,err := sub.NewSub("./test.srt")
	//subtitle,err := sub.NewSub("/Users/megrez/Downloads/6d3e011105dd071752cba160c5f18fc8/The.Shannara.Chronicles.S01E02.720p.WEB-DL.AAC2.0.H.264-VietHD/The.Shannara.Chronicles.S01E02.720p.WEB-DL.AAC2.0.H.264-VietHD.简体&英文.srt")
	if err != nil {
		log.Fatal(err)
	}
	subtitle.Parse()
	fmt.Println(duration)
}