package sub

import (
	"time"
	"bytes"
	"log"
	"fmt"
	"errors"
	"strconv"
)

var ParamsParseError = errors.New("duration params parsed error")

type Entry struct {
	No int
	StartTime time.Duration
	EndTime time.Duration
	Content string
}

func getDuration(formattedStr string) (time.Duration,error) {
	// fmt.Println("------")
	var hours,minutes,seconds,mseconds int
	// fmt.Println(formattedStr)
	paramsParsed, err := fmt.Sscanf(formattedStr,"%d:%d:%d,%d",&hours,&minutes,&seconds,&mseconds)
	if err != nil {
		return 0, err
	}
	if paramsParsed < 4 {
		return 0, ParamsParseError
	}
	durationStr := strconv.Itoa(hours) + "h" + strconv.Itoa(minutes) + "m" + strconv.Itoa(seconds) + "s" + strconv.Itoa(mseconds) + "ms"
	// fmt.Println(durationStr)
	return time.ParseDuration(durationStr)
}

func NewEntry(no int,rawDuration []byte,content []byte) *Entry {
	index := bytes.Index(rawDuration,[]byte(" --> "))
	if index < 0 {
		log.Fatal("invalid duration")
	}
	start,err := getDuration(string(rawDuration[:index]))
	if err != nil {
		log.Fatal(err)
	}
	end,err := getDuration(string(rawDuration[index+5:]))
	if err != nil {
		log.Fatal(err)
	}
	return &Entry{ No: no,StartTime: start,EndTime: end,Content: string(content) }
}