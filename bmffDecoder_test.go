package mp4box

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"testing"
)

//HELPER - Begin
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
	decoder := readMp4File("test/bbb_30fps_1024x576_2500k_0.m4v")
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
}
