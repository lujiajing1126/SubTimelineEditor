package sub

import (
	"os"
	"bufio"
	"strconv"
	"log"
	"bytes"
)

// check the type at compile time
var _ bufio.SplitFunc = scanSub

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}

func scanSub(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data,[]byte("\n\n")); i >= 0 {
		return i + 2,dropCR(data[0:i+1]), nil
	}

	if atEOF {
		return len(data), dropCR(data), nil
	}
	return 0, nil, nil
}

type Sub struct {
	subFile	*os.File
	entryies []*Entry
}

func NewSub(filepath string) (*Sub, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil,err
	}
	return &Sub{subFile: file},nil
}


func (S *Sub) Parse() error {
	defer S.subFile.Close()
	fileScanner := bufio.NewScanner(S.subFile)
	fileScanner.Split(scanSub)
	for fileScanner.Scan() {
		rawStr := []byte(fileScanner.Text())
		noIndex := bytes.IndexByte(rawStr,'\n')
		no, err := strconv.Atoi(string(rawStr[0:noIndex]))
		if err != nil {
			log.Fatal(err)
		}
		durationIndex := bytes.IndexByte(rawStr[noIndex+1:],'\n')
		rawDuration := rawStr[noIndex + 1:noIndex + durationIndex + 1]
		S.entryies = append(S.entryies,NewEntry(no,rawDuration,rawStr[noIndex+1+durationIndex+1:]))
	}
	for _, v := range S.entryies {
		log.Printf("%d: %d --> %d / %s",v.No,v.StartTime,v.EndTime, v.Content)
	}
	return nil
}