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

func readMp4File(filename string) *BoxDecoder {
	file, err := os.Open(filename) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	rdr := bufio.NewReader(file)
	return NewBoxDecoder(rdr)
}

func getType(myvar interface{}) string {
	var t reflect.Type
	t = reflect.TypeOf(myvar)
	if t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	}
	return t.Name()
}

//HELPER - End
func TestReadBox1(t *testing.T) {
	loadURLs(t)
	for _, name := range urlsToTest {
		t.Logf("*************************************")
		decoder := readMp4File(name)
		i := 0
		for {
			i++
			box, err := decoder.NextBox()
			if err != nil {
				t.Logf("Error : %s", err.Error())
				break
			}
			boxType := getType(box)
			t.Logf("%v %v", boxType, box.String())
		}
		t.Logf("*************************************")
	}
}
