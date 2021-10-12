package mp4box

import (
	"encoding/binary"
	"errors"
	"io"
	"strings"
)

//BoxReader - Decode Box from Stream
type BoxReader struct {
	r      io.Reader
	parent Box
}

//NewBoxReader - Create a reader for stream
func NewBoxReader(r io.Reader) *BoxReader {
	return newBoxReader(r, nil)
}
func newBoxReader(r io.Reader, parent Box) *BoxReader {
	ret := &BoxReader{}
	ret.r = r
	if parent != nil {
		ret.parent = parent
	}
	return ret
}

//ErrLastBox - Last box of the block
var ErrLastBox = errors.New("LAST BOX")

//ErrUnknownBox - Box is not known. couldn't build
var ErrUnknownBox = errors.New("UNKNOWN BOX")

/*
aligned(8) class Box (unsigned int(32) boxtype, optional unsigned int(8)[16] extended_type) {
   unsigned int(32) size;
   unsigned int(32) type = boxtype;
   if (size==1) {
      unsigned int(64) largesize;
   } else if (size==0) {
      // box extends to end of file
   }
if (boxtype==‘uuid’) {
unsigned int(8)[16] usertype = extended_type;
} }
*/

func (d *BoxReader) readBoxHdr() (int64, string, []int8, []byte, error) {
	var boxHdr struct {
		Size    int32
		BoxType [4]byte
	}
	var boxSize int64
	var boxType string
	var extendedType []int8
	var payload []byte
	var err error
	var bytesRead int
	err = binary.Read(d.r, binary.BigEndian, &boxHdr.Size)
	if err != nil {
		return 0, "", nil, nil, err
	}
	bytesRead += 4
	if boxHdr.Size == 0 {
		//Last box
		return 0, "", nil, nil, ErrLastBox
	}
	err = binary.Read(d.r, binary.BigEndian, &boxHdr.BoxType)
	if err != nil {
		return 0, "", nil, nil, err
	}
	bytesRead += 4
	if boxHdr.Size == 1 {
		err = binary.Read(d.r, binary.BigEndian, &boxSize)
		if err != nil {
			return 0, "", nil, nil, err
		}
		bytesRead += 8
	} else {
		boxSize = int64(boxHdr.Size)
	}
	boxType = string(boxHdr.BoxType[:])
	if strings.Compare(boxType, "uuid") == 0 {
		extendedType = make([]int8, 16)
		err = binary.Read(d.r, binary.BigEndian, &extendedType)
		if err != nil {
			return boxSize, "", nil, nil, err
		}
		bytesRead += len(extendedType)
	}
	if boxSize-int64(bytesRead) > 0 {
		payload = make([]byte, boxSize-int64(bytesRead))
		err = binary.Read(d.r, binary.BigEndian, payload)
		if err != nil {
			return boxSize, boxType, extendedType, nil, err
		}
	}
	return boxSize, boxType, extendedType, payload, err
}

//NextBox - Returns the Box read
func (d *BoxReader) NextBox() (Box, error) {
	boxSize, boxType, extendedType, payload, err := d.readBoxHdr()
	if err != nil {
		return nil, err
	}
	return d.makeBox(boxSize, boxType, &payload, extendedType)
}

//makeBox - Make Box based on boxType
func (d *BoxReader) makeBox(boxSize int64, boxType string, payload *[]byte, extendedType []int8) (Box, error) {
	var ret Box
	var err error
	factory := BoxFactory{}
	ret = factory.MakeEmptyBoxObject(boxType)
	if ret == nil {
		err = ErrUnknownBox
	} else {
		err = ret.initData(boxSize, boxType, payload, d.parent)
	}
	return ret, err
}
