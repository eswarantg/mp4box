package mp4box

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

//BaseBox - Base box holding the bytes
type BaseBox struct {
	naAccessBoxTypeImpl               //get all methods of ErrBoxNotFound responses other than overridden
	boxSize             int64         //Box Size
	boxType             string        //Box Type
	payload             *[]byte       //Payload of the Box
	level               int           //Level of the box in heirachy
	parent              AccessBoxType //Parent Box access
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *BaseBox) getLeafBox() AccessBoxType {
	return b
}

//GetBaseBox - Implement AccessBoxType method for this object
func (b *BaseBox) GetBaseBox() (*BaseBox, error) {
	return b, nil
}

//GetParentByName - Get parent by name
func (b *BaseBox) GetParentByName(boxType string) (AccessBoxType, error) {
	if b.parent != nil {
		if b.parent.Boxtype() == boxType {
			return b.getLeafBox(), nil
		}
		return b.parent.GetParentByName(boxType)
	}
	return nil, ErrBoxNotFound
}

//GetChildByName - Get child by name
func (b *BaseBox) GetChildByName(boxType string) (AccessBoxType, error) {
	return nil, ErrBoxNotFound
}

//getLevel - Get Object
func (b *BaseBox) getLevel() int {
	return b.level
}

//isCollection -
func (b *BaseBox) isCollection() bool {
	return true
}

//Interface methods Impl - End

//NewBaseBox - Create a new base box
func (b *BaseBox) initData(boxSize int64, boxType string, payload *[]byte, parent AccessBoxType) error {
	b.boxSize = boxSize
	b.boxType = boxType
	b.payload = payload
	if parent != nil {
		b.parent = parent
		b.level = b.parent.getLevel() + 1
	} else {
		b.level = 0
	}
	return nil
}

//Boxtype - Returns BoxType of the Box
func (b *BaseBox) Boxtype() string {
	return b.boxType
}

//Size - Returns Size of the Box
func (b *BaseBox) Size() int64 {
	return b.boxSize
}

//getPayload - Returns the payload excluding headers
func (b *BaseBox) getPayload() []byte {
	return *b.payload
}

func (b *BaseBox) leadString() string {
	var lead string
	for i := 0; i < b.level; i++ {
		lead += "\t"
	}
	return lead
}

//String - Returns User Readable description of content
func (b *BaseBox) String() string {
	var ret string
	if b.payload != nil {
		ret += fmt.Sprintf("\n%d%vType: %v, Size: %v, Payload:%v", b.level, b.leadString(), b.Boxtype(), b.Size(), len(*b.payload))
	} else {
		ret += fmt.Sprintf("\n%d%vType: %v, Size: %v, Payload:<nil>", b.level, b.leadString(), b.Boxtype(), b.Size())
	}
	return ret
}

//Payload - Returns the payload excluding headers
func (b *BaseBox) writeBytes(w io.Writer) error {
	if b.boxSize == 0 {
		tmpInt32 := int32(0)
		binary.Write(w, binary.BigEndian, tmpInt32)
	}
	if b.boxSize < math.MaxInt32 {
		tmpInt32 := int32(b.boxSize)
		binary.Write(w, binary.BigEndian, tmpInt32)
		binary.Write(w, binary.BigEndian, []byte(b.boxType[0:4]))
	} else {
		tmpInt32 := int32(1)
		binary.Write(w, binary.BigEndian, tmpInt32)
		binary.Write(w, binary.BigEndian, []byte(b.boxType[0:4]))
		binary.Write(w, binary.BigEndian, b.boxSize)
	}
	if b.payload != nil {
		if len(*b.payload) > 0 {
			binary.Write(w, binary.BigEndian, b.payload)
		}
	}
	return nil
}

//FullBox - Base box holding the bytes
/*
aligned(8) class FullBox(unsigned int(32) boxtype, unsigned int(8) v, bit(24) f) extends Box(boxtype) {
unsigned int(8) version = v;
bit(24) flags = f;
}
*/
type FullBox struct {
	BaseBox
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *FullBox) getLeafBox() AccessBoxType {
	return b
}

//GetFullBox - Implement AccessBoxType method for this object
func (b *FullBox) GetFullBox() (*FullBox, error) {
	return b, nil
}

//Interface methods Impl - End

//getPayload - Returns the payload excluding headers
func (b *FullBox) getPayload() []byte {
	return (*b.payload)[4:]
}

//Version - returns the version of the box
func (b *FullBox) Version() int8 {
	var ret int8
	p := b.BaseBox.getPayload()
	ret = int8(p[0])
	return ret
}

//Flags - Returns flags
func (b *FullBox) Flags() []uint8 {
	p := b.BaseBox.getPayload()
	return p[1:4]
}

//String - Returns User Readable description of content
func (b *FullBox) String() string {
	var ret string
	ret += b.BaseBox.String()
	ret += fmt.Sprintf("\n%v ", b.leadString())
	ret += fmt.Sprintf(" Version: %v, Flags: %v", b.Version(), b.Flags())
	return ret
}
