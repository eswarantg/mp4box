package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/eswarantg/mp4box"
)

func main() {
	for i, name := range os.Args[1:] {
		parseFile(i, name)
	}
}

func readMp4File(filename string) *mp4box.BoxReader {
	file, err := os.Open(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	rdr := bufio.NewReader(file)
	return mp4box.NewBoxReader(rdr)
}

func getType(myvar interface{}) string {
	var t reflect.Type
	t = reflect.TypeOf(myvar)
	if t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	}
	return t.Name()
}

var timeScale *uint32
var trackID *uint32

func buildChunkDuration(box mp4box.Box) *time.Duration {
	var err error
	switch box.Boxtype() {
	case "ftyp":
		timeScale = nil
		fmt.Fprintf(os.Stdout, "\nChunk Restart")
	case "sidx":
		if box != nil {
			var sidxbox *mp4box.SegmentIndexBox
			sidxbox, err = box.GetSegmentIndexBox()
			if err == nil {
				timeScale = new(uint32)
				*timeScale = sidxbox.TimeScale()
				fmt.Fprintf(os.Stdout, "\nChunk Time Scale: %v", *timeScale)
			}
		}
	case "moov":
		if box != nil {
			var moovbox *mp4box.MovieBox
			moovbox, err = box.GetMovieBox()
			if err == nil {
				duration, ts, tid := moovbox.Summary()
				if ts > 0 {
					timeScale = new(uint32)
					*timeScale = ts
					fmt.Fprintf(os.Stdout, "\nChunk Time Scale: %v", *timeScale)
				}
				trackID = new(uint32)
				*trackID = tid
				if duration > 0 && duration != 0xFFFFFFFFFFFFFFFF {
					if timeScale != nil {
						dur := time.Duration(uint64(float64(duration)*1000000/float64(*timeScale))) * time.Microsecond
						fmt.Fprintf(os.Stdout, "\nChunk Time Scale: %v %v/%v %v", *trackID, duration, *timeScale, dur)
						return &dur
					}
				}
			}
		}
	case "moof":
		if box != nil {
			var moofbox *mp4box.MovieFragmentBox
			moofbox, err = box.GetMovieFragmentBox()
			if err == nil {
				duration, ts, tid, _ := moofbox.Summary()
				if ts > 0 {
					timeScale = new(uint32)
					*timeScale = ts
				}
				trackID = new(uint32)
				*trackID = tid
				if duration > 0 && duration != 0xFFFFFFFFFFFFFFFF {
					if timeScale != nil {
						dur := time.Duration(uint64(float64(duration)*1000000/float64(*timeScale))) * time.Microsecond
						fmt.Fprintf(os.Stdout, "\nChunk Time Scale: %v %v/%v %v", *trackID, duration, *timeScale, dur)
						return &dur
					}
				}
			}
		}
	}
	return nil
}

func parseFile(testno int, name string) {
	testno++
	fmt.Fprintf(os.Stdout, "\n*************** %v %v **********************", testno, name)
	reader := readMp4File(name)
	i := 0
	for {
		i++
		box, err := reader.NextBox()
		if err != nil {
			if err != io.EOF {
				fmt.Fprintf(os.Stderr, "\nError : %s", err)
			}
			break
		}
		boxType := getType(box)
		fmt.Fprintf(os.Stdout, "\n%v %v", boxType, box.String())
		//Get TimeScale from sidx box
		dur := buildChunkDuration(box)
		if dur != nil {
			fmt.Fprintf(os.Stdout, "\nChunk Duration : %v", dur)
		}
	}
	fmt.Fprintf(os.Stdout, "\n*************** %v %v **********************", testno, name)
}
