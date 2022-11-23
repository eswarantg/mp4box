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
var sequence *uint32
var baseMediaDecodetime *uint64

func buildChunkDuration(box mp4box.Box) *time.Duration {
	var ret *time.Duration
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
				if timeScale == nil {
					timeScale = new(uint32)
				}
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
				if ts != nil {
					if timeScale == nil {
						timeScale = new(uint32)
					}
					*timeScale = *ts
					fmt.Fprintf(os.Stdout, "\nChunk Time Scale: %v", *timeScale)
				}
				if tid != nil {
					if trackID == nil {
						trackID = new(uint32)
					}
					*trackID = *tid
				}
				if duration != nil {
					if *duration > 0 && *duration != 0xFFFFFFFFFFFFFFFF {
						if timeScale != nil {
							dur := time.Duration(uint64(float64(*duration)*1000000/float64(*timeScale))) * time.Microsecond
							fmt.Fprintf(os.Stdout, "\nChunk Time Scale: %v %v/%v %v", *trackID, duration, *timeScale, dur)
							return &dur
						}
					}
				}
			}
		}
	case "moof":
		var lastSequence uint32
		var lastBaseMediaDecodetime uint64
		if box != nil {
			var moofbox *mp4box.MovieFragmentBox
			moofbox, err = box.GetMovieFragmentBox()
			if err == nil {
				s, bmdt, tid, ts := moofbox.Summary()
				if ts != nil {
					if timeScale == nil {
						timeScale = new(uint32)
					}
					*timeScale = *ts
					fmt.Fprintf(os.Stdout, "\nChunk Time Scale: %v", *timeScale)
				}
				if tid != nil {
					if trackID == nil {
						trackID = new(uint32)
					}
					*trackID = *tid
				}
				if s != nil {
					if sequence == nil {
						sequence = new(uint32)
					}
					lastSequence = *sequence
					*sequence = *s
				}
				if bmdt != nil {
					if baseMediaDecodetime == nil {
						baseMediaDecodetime = new(uint64)
					}
					lastBaseMediaDecodetime = *baseMediaDecodetime
					*baseMediaDecodetime = *bmdt
				}
			}
			if timeScale != nil && sequence != nil && baseMediaDecodetime != nil {
				if *sequence == lastSequence+1 {
					duration := *baseMediaDecodetime - lastBaseMediaDecodetime
					if duration > 0 {
						dur := time.Duration(uint64(float64(duration)*1000000/float64(*timeScale))) * time.Microsecond
						fmt.Fprintf(os.Stdout, "\nChunk Time Scale: %v %v/%v %v", *trackID, duration, *timeScale, dur)
						return &dur
					}

				}
			}
		}
	}
	return ret
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
		// fmt.Println("boxType: ", box.Boxtype())
		fmt.Fprintf(os.Stdout, "\n%v %v", boxType, box.String())
		// fmt.Println("\nBox String: ", box.String())
		if box.Boxtype() == "emsg" {
			fmt.Println("\nEMSG BOX INFO: ")
			eMsgBox, err := box.GetEmsgBox()
			if err != nil {
				fmt.Println("err: ", err)
				break
			}
			err = eMsgBox.ParseAllData()
			if err != nil {
				fmt.Println("err: ", err)
				break
			}
			fmt.Println("\neMsgBox Payload: ", eMsgBox.GetPayload())
			msgData := eMsgBox.GetMsgData()
			fmt.Println("eMsgBox MsgData: ", msgData)
			schemeIdUri, schemeIdVal := eMsgBox.GetSchemeInfo()
			fmt.Println("eMsgBox schemeInfo: ", schemeIdUri, " : ", schemeIdVal)
			version := eMsgBox.FullBox.Version()
			fmt.Println("eMsgBox version: ", version)
		}
		//Get TimeScale from sidx box
		dur := buildChunkDuration(box)
		if dur != nil {
			fmt.Fprintf(os.Stdout, "\nChunk Duration : %v", dur)
		}
	}
	fmt.Fprintf(os.Stdout, "\n*************** %v %v **********************", testno, name)
}
