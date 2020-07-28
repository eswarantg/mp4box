package mp4box

import (
	"bytes"
)

//CollectionFullBox - Collection of boxes
type CollectionFullBox struct {
	FullBox
	childBoxMap map[string][]Box
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *CollectionFullBox) getLeafBox() Box {
	return b
}

//GetCollectionFullBox - Implement Box method for this object
func (b *CollectionFullBox) GetCollectionFullBox() (*CollectionFullBox, error) {
	return b, nil
}

//isCollection -
func (b *CollectionFullBox) isCollection() bool {
	return true
}

//GetChildrenByName - Get child by name
func (b *CollectionFullBox) GetChildrenByName(boxType string) ([]Box, error) {
	return getChildBoxHelper(b.childBoxMap, boxType)
}

//Interface methods Impl - End
//NewBaseBox - Create a new base box
func (b *CollectionFullBox) initData(boxSize int64, boxType string, payload *[]byte, parent Box) error {
	var err error
	b.childBoxMap = make(map[string][]Box)
	if payload != nil && len(*payload) >= 5 {
		fullboxpayload := (*payload)[0:4]
		err = b.FullBox.initData(boxSize, boxType, &fullboxpayload, parent)
	}
	if err == nil && payload != nil && len(*payload) >= 6 {
		childpayload := (*payload)[4:]
		return b.populateChildBoxes(&childpayload)
	}
	return err
}

//String - Returns User Readable description of content
func (b *CollectionFullBox) String() string {
	var ret string
	ret += b.FullBox.String()
	ret += b.detailString()
	ret += getChildBoxString(b.childBoxMap)
	return ret
}

func (b *CollectionFullBox) detailString() string {
	var ret string
	return ret
}

func (b *CollectionFullBox) populateChildBoxes(payload *[]byte) error {
	if payload != nil {
		r := bytes.NewReader(*payload)
		reader := newBoxReader(r, b)
		return populateChildBoxesHelper(b.childBoxMap, reader)
	}
	return nil
}
