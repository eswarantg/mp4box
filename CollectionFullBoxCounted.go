package mp4box

import (
	"bytes"
	"encoding/binary"
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
	childBoxes map[string]AccessBoxType
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *CollectionFullBoxCounted) getLeafBox() AccessBoxType {
	return b
}

//GetCollectionFullBoxCounted - Implement AccessBoxType method for this object
func (b *CollectionFullBoxCounted) GetCollectionFullBoxCounted() (*CollectionFullBoxCounted, error) {
	return b, nil
}

//isCollection -
func (b *CollectionFullBoxCounted) isCollection() bool {
	return true
}

//Interface methods Impl - End
//NewBaseBox - Create a new base box
func (b *CollectionFullBoxCounted) initData(boxSize int64, boxType string, payload *[]byte, parent AccessBoxType) error {
	b.FullBox.initData(boxSize, boxType, nil, parent)
	b.childBoxes = make(map[string]AccessBoxType)
	b.populateChildBoxes(payload)
	return nil
}

func (b *CollectionFullBoxCounted) String() string {
	var ret string
	ret += b.FullBox.String()
	for _, child := range b.childBoxes {
		ret += child.String()
	}
	return ret
}

func (b *CollectionFullBoxCounted) populateChildBoxes(payload *[]byte) error {
	var err error
	if payload != nil {
		counter := binary.BigEndian.Uint32(*payload)
		r := bytes.NewReader(*payload)
		decoder := newBoxDecoder(r, b)
		for {
			var box AccessBoxType
			box, err = decoder.NextBox()
			if err != nil {
				break
			}
			b.childBoxes[box.Boxtype()] = box
		}
		if int(counter) != len(b.childBoxes) {
			return err
		}
	}
	return nil
}
