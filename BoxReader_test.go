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

var fragScale *uint32

func buildFragDuration(t *testing.T, box Box) *time.Duration {
	var ret *time.Duration
	var err error
	var sidxbox *SegmentIndexBox
	var moofbox *MovieFragmentBox
	switch box.Boxtype() {
	case "styp":
		fragScale = nil
		t.Logf("Fragment Start")
	case "sidx":
		if box != nil {
			sidxbox, err = box.GetSegmentIndexBox()
			if err == nil {
				fragScale = new(uint32)
				*fragScale = sidxbox.TimeScale()
				t.Logf("Fragment Scale: %v", *fragScale)
			}
		}
	case "moof":
		if box != nil {
			moofbox, err = box.GetMovieFragmentBox()
			if err == nil {
				dur := moofbox.TotalDuration()
				if fragScale != nil {
					secs := float64(dur) / float64(*fragScale)
					ret = new(time.Duration)
					*ret = time.Duration(int64(secs*1000000)) * time.Microsecond
					t.Logf("Fragment Duration: (%v/%v) = %v", dur, *fragScale, *ret)
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
			dur := buildFragDuration(t, box)
			if dur != nil {
				t.Logf("Fragment Duration : %v", dur)
			}
		}
		t.Logf("*************** %v %v **********************", testno, name)
	}
}
