package mp4box

import (
	"bytes"
	"io"
)

func getChildBoxHelper(childBoxMap map[string][]Box, boxType string) ([]Box, error) {
	_, ok := childBoxMap[boxType]
	if ok {
		return childBoxMap[boxType], nil
	}
	for _, childBoxes := range childBoxMap {
		for _, childBox := range childBoxes {
			if childBox.isCollection() {
				box, err := childBox.GetChildrenByName(boxType)
				if err == nil {
					return box, nil
				}
			}
		}
	}
	return nil, ErrBoxNotFound
}

func getChildBoxString(childBoxMap map[string][]Box) string {
	var ret string
	for _, childBoxes := range childBoxMap {
		for _, childBox := range childBoxes {
			ret += childBox.String()
		}
	}
	return ret
}

func populateChildBoxesHelper(childBoxMap map[string][]Box, reader *BoxReader) error {
	for {
		box, err := reader.NextBox()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if _, ok := childBoxMap[box.Boxtype()]; ok {
			childBoxMap[box.Boxtype()] = append(childBoxMap[box.Boxtype()], box)
		} else {
			childBoxMap[box.Boxtype()] = []Box{box}
		}
	}
}

//CollectionBaseBox - Collection of boxes
type CollectionBaseBox struct {
	BaseBox
	childBoxMap map[string][]Box
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

//GetChildrenByName - Get child by name
func (b *CollectionBaseBox) GetChildrenByName(boxType string) ([]Box, error) {
	return getChildBoxHelper(b.childBoxMap, boxType)
}

//isCollection -
func (b *CollectionBaseBox) isCollection() bool {
	return true
}

//Interface methods Impl - End

//NewBaseBox - Create a new base box
func (b *CollectionBaseBox) initData(boxSize int64, boxType string, payload *[]byte, parent Box) error {
	b.BaseBox.initData(boxSize, boxType, nil, parent)
	b.childBoxMap = make(map[string][]Box)
	b.populateChildBoxes(payload)
	return nil
}

//String - Returns User Readable description of content
func (b *CollectionBaseBox) String() string {
	var ret string
	ret += b.BaseBox.String()
	ret += b.detailString()
	ret += getChildBoxString(b.childBoxMap)
	return ret
}

func (b *CollectionBaseBox) detailString() string {
	var ret string
	return ret
}

func (b *CollectionBaseBox) populateChildBoxes(payload *[]byte) error {
	if payload != nil {
		r := bytes.NewReader(*payload)
		reader := newBoxReader(r, b)
		return populateChildBoxesHelper(b.childBoxMap, reader)
	}
	return nil
}
