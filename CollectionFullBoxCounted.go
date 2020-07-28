package mp4box

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//CollectionFullBoxCounted - Collection of boxes
/*
aligned(8) class XXXXX
extends FullBox(xxxx, version = 0, 0) {
	unsigned int(32) entry_count;
	for (i=1; i • entry_count; i++) {
		BOX()
	}
}
*/
type CollectionFullBoxCounted struct {
	FullBox
	counter     uint32
	childBoxMap map[string][]Box
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *CollectionFullBoxCounted) getLeafBox() Box {
	return b
}

//GetCollectionFullBoxCounted - Implement Box method for this object
func (b *CollectionFullBoxCounted) GetCollectionFullBoxCounted() (*CollectionFullBoxCounted, error) {
	return b, nil
}

//isCollection -
func (b *CollectionFullBoxCounted) isCollection() bool {
	return true
}

//GetChildrenByName - Get child by name
func (b *CollectionFullBoxCounted) GetChildrenByName(boxType string) ([]Box, error) {
	return getChildBoxHelper(b.childBoxMap, boxType)
}

//Interface methods Impl - End
//NewBaseBox - Create a new base box
func (b *CollectionFullBoxCounted) initData(boxSize int64, boxType string, payload *[]byte, parent Box) error {
	var err error
	b.childBoxMap = make(map[string][]Box)
	if payload != nil && len(*payload) >= 4 {
		fullboxpayload := (*payload)[0:4]
		err = b.FullBox.initData(boxSize, boxType, &fullboxpayload, parent)
	}
	if err == nil && payload != nil && len(*payload) >= 8 {
		childpayload := (*payload)[4:]
		return b.populateChildBoxes(&childpayload)
	}
	return err

}

func (b *CollectionFullBoxCounted) String() string {
	var ret string
	ret += b.FullBox.String()
	ret += b.detailString()
	ret += getChildBoxString(b.childBoxMap)
	return ret
}
func (b *CollectionFullBoxCounted) detailString() string {
	var ret string
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf(" Count:%v ", b.counter)
	return ret
}

func (b *CollectionFullBoxCounted) populateChildBoxes(payload *[]byte) error {
	if payload != nil {
		b.counter = binary.BigEndian.Uint32(*payload)
		r := bytes.NewReader((*payload)[4:])
		reader := newBoxReader(r, b)
		return populateChildBoxesHelper(b.childBoxMap, reader)
	}
	return nil
}
