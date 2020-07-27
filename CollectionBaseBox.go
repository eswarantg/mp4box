package mp4box

import (
	"bytes"
)

//CollectionBaseBox - Collection of boxes
type CollectionBaseBox struct {
	BaseBox
	childBoxes map[string]Box
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *CollectionBaseBox) getLeafBox() Box {
	return b
}

//GetCollectionBaseBox - Implement Box method for this object
func (b *CollectionBaseBox) GetCollectionBaseBox() (*CollectionBaseBox, error) {
	return b, nil
}

//GetChildByName - Get child by name
func (b *CollectionBaseBox) GetChildByName(boxType string) (Box, error) {
	_, ok := b.childBoxes[boxType]
	if ok {
		return b.childBoxes[boxType], nil
	}
	for _, childBox := range b.childBoxes {
		if childBox.isCollection() {
			box, err := childBox.GetChildByName(boxType)
			if err == nil {
				return box, nil
			}
		}
	}
	return nil, ErrBoxNotFound
}

//isCollection -
func (b *CollectionBaseBox) isCollection() bool {
	return true
}

//Interface methods Impl - End

//NewBaseBox - Create a new base box
func (b *CollectionBaseBox) initData(boxSize int64, boxType string, payload *[]byte, parent Box) error {
	b.BaseBox.initData(boxSize, boxType, nil, parent)
	b.childBoxes = make(map[string]Box)
	b.populateChildBoxes(payload)
	return nil
}

func (b *CollectionBaseBox) String() string {
	var ret string
	ret += b.BaseBox.String()
	for _, child := range b.childBoxes {
		ret += child.String()
	}
	return ret
}

func (b *CollectionBaseBox) populateChildBoxes(payload *[]byte) error {
	if payload != nil {
		r := bytes.NewReader(*payload)
		decoder := newBoxReader(r, b)
		for {
			box, err := decoder.NextBox()
			if err != nil {
				return err
			}
			b.childBoxes[box.Boxtype()] = box
		}
	}
	return nil
}
