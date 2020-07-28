package mp4box

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
)

//HELPER - Begin
var urlsToTest []string

func loadURLs(t *testing.T) {
	urllistfilename := "test/urls.txt"
	file, err := os.Open(urllistfilename) // For read access.
	if err != nil {
		t.Errorf("Error opening file %v: %v", urllistfilename, err.Error())
		return
	}
	c := make(chan string)
	go func(file *os.File) {
		defer file.Close()
		defer close(c)
		rdr := bufio.NewReader(file)
		url := ""
		for {
			line, isprefix, err := rdr.ReadLine()
			if err != nil {
				if err != io.EOF {
					t.Errorf("Error during read file %v: %v", urllistfilename, err.Error())
				}
				break
			}
			if isprefix == true {
				url += string(line)
			} else {
				url += string(line)
				if len(url) > 0 {
					if url[0] != '#' {
						c <- string(url)
					}
				}
				url = ""
			}
		}
	}(file)
	for urlStr := range c {
		u, err := url.Parse(urlStr)
		if err != nil {
			t.Errorf("Error Parsing url%v: %v", urlStr, err.Error())
			continue
		}
		filename := "test/" + strings.ReplaceAll(u.Hostname(), ".", "_") + strings.ReplaceAll(u.Path, "/", "_")
		_, err = os.Stat(filename)
		if err == nil {
			t.Logf("File exists. %s", filename)
			urlsToTest = append(urlsToTest, filename)
			continue
		} else {
			if !os.IsNotExist(err) {
				t.Errorf("File stat error. %s. %v", filename, err.Error())
				continue
			}
		}
		resp, err := http.Get(u.String())
		if err != nil {
			t.Errorf("Error fetching url%v: %v", u.String(), err.Error())
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Error fetching url%v : Resp:%v", urlStr, resp.Status)
			continue
		}
		fileW, err := os.Create(filename)
		if err != nil {
			t.Errorf("Error creating file %v: %v", filename, err.Error())
			return
		}
		defer fileW.Close()
		_, err = io.Copy(fileW, resp.Body)
		if err != nil {
			t.Errorf("Error copying filecontent to %v: %v", filename, err.Error())
		}
		t.Logf("Created %v from %v", filename, urlStr)
		urlsToTest = append(urlsToTest, filename)
	}
}

func readMp4File(filename string) *BoxReader {
	file, err := os.Open(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	rdr := bufio.NewReader(file)
	return NewBoxReader(rdr)
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

func buildChunkDuration(t *testing.T, box Box) *time.Duration {
	var ret *time.Duration
	var err error
	switch box.Boxtype() {
	case "ftyp":
		timeScale = nil
		t.Logf("Chunk Restart")
	case "sidx":
		if box != nil {
			var sidxbox *SegmentIndexBox
			sidxbox, err = box.GetSegmentIndexBox()
			if err == nil {
				timeScale = new(uint32)
				*timeScale = sidxbox.TimeScale()
				t.Logf("Chunk Time Scale: %v", *timeScale)
			}
		}
	case "moov":
		if box != nil {
			var moovbox *MovieBox
			moovbox, err = box.GetMovieBox()
			if err == nil {
				duration, ts, tid := moovbox.Summary()
				if ts > 0 {
					timeScale = new(uint32)
					*timeScale = ts
					t.Logf("Chunk Time Scale: %v", *timeScale)
				}
				trackID = new(uint32)
				*trackID = tid
				if duration > 0 {
					if timeScale != nil {
						dur := time.Duration(uint64(float64(duration)*1000000/float64(*timeScale))) * time.Microsecond
						t.Logf("Chunk Time Scale: %v %v/%v %v", *trackID, duration, *timeScale, dur)
						return &dur
					}
				}
			}
		}
	case "moof":
		if box != nil {
			var moofbox *MovieFragmentBox
			moofbox, err = box.GetMovieFragmentBox()
			if err == nil {
				duration, ts, tid, _ := moofbox.Summary()
				if ts > 0 {
					timeScale = new(uint32)
					*timeScale = ts
				}
				trackID = new(uint32)
				*trackID = tid
				if duration > 0 {
					if timeScale != nil {
						dur := time.Duration(uint64(float64(duration)*1000000/float64(*timeScale))) * time.Microsecond
						t.Logf("Chunk Time Scale: %v %v/%v %v", *trackID, duration, *timeScale, dur)
						return &dur
					}
				}
			}
		}
	}
	return ret
}

//HELPER - End
func TestReadBox1(t *testing.T) {
	loadURLs(t)
	testno := 0
	for _, name := range urlsToTest {
		testno++
		t.Logf("*************** %v %v **********************", testno, name)
		reader := readMp4File(name)
		i := 0
		for {
			i++
			box, err := reader.NextBox()
			if err != nil {
				if err != io.EOF {
					t.Errorf("Error : %s", err)
				}
				break
			}
			boxType := getType(box)
			t.Logf("%v %v", boxType, box)
			//Get TimeScale from sidx box
			dur := buildChunkDuration(t, box)
			if dur != nil {
				t.Logf("Chunk Duration : %v", dur)
			}
		}
		t.Logf("*************** %v %v **********************", testno, name)
	}
}
