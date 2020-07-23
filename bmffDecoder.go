package mp4box

import (
	"encoding/binary"
	"errors"
	"io"
	"strings"
)

//BoxDecoder - Decode Box from Stream
type BoxDecoder struct {
	r      io.Reader
	parent AccessBoxType
}

//NewBoxDecoder - Create a decoder for stream
func NewBoxDecoder(r io.Reader) *BoxDecoder {
	return newBoxDecoder(r, nil)
}
func newBoxDecoder(r io.Reader, parent AccessBoxType) *BoxDecoder {
	ret := &BoxDecoder{}
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

//NextBox - Returns the Box read
func (d *BoxDecoder) NextBox() (AccessBoxType, error) {
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
		return nil, err
	}
	bytesRead += 4
	if boxHdr.Size == 0 {
		//Last box
		return nil, ErrLastBox
	}
	err = binary.Read(d.r, binary.BigEndian, &boxHdr.BoxType)
	if err != nil {
		return nil, err
	}
	bytesRead += 4
	if boxHdr.Size == 1 {
		err = binary.Read(d.r, binary.BigEndian, &boxSize)
		if err != nil {
			return nil, err
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
			return nil, err
		}
		bytesRead += len(extendedType)
	}
	if boxSize-int64(bytesRead) > 0 {
		payload = make([]byte, boxSize-int64(bytesRead))
		err = binary.Read(d.r, binary.BigEndian, payload)
		if err != nil {
			return nil, err
		}
	}
	return d.makeBox(boxSize, boxType, &payload, extendedType)
}

//makeBox - Make Box based on boxType
func (d *BoxDecoder) makeBox(boxSize int64, boxType string, payload *[]byte, extendedType []int8) (AccessBoxType, error) {
	var ret AccessBoxType
	var err error
	ret = makeEmptyBoxObject(boxType)
	if ret == nil {
		err = ErrUnknownBox
	} else {
		err = ret.initData(boxSize, boxType, payload, d.parent)
	}
	return ret, err
}
