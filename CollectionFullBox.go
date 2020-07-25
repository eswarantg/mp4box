package mp4box

import (
	"bytes"
)

//CollectionFullBox - Collection of boxes
type CollectionFullBox struct {
	FullBox
	childBoxes map[string]Box
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

//Interface methods Impl - End
//NewBaseBox - Create a new base box
func (b *CollectionFullBox) initData(boxSize int64, boxType string, payload *[]byte, parent Box) error {
	b.FullBox.initData(boxSize, boxType, nil, parent)
	b.childBoxes = make(map[string]Box)
	b.populateChildBoxes(payload)
	return nil
}

func (b *CollectionFullBox) String() string {
	var ret string
	ret += b.FullBox.String()
	for _, child := range b.childBoxes {
		ret += child.String()
	}
	return ret
}

func (b *CollectionFullBox) populateChildBoxes(payload *[]byte) error {
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
