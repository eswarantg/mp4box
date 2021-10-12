package mp4box

import (
	"encoding/binary"
	"fmt"
)

//FileBox - Base box holding the bytes
/*
aligned(8) class FileTypeBox
extends Box(‘ftyp’) {
unsigned int(32) major_brand; unsigned int(32) minor_version; unsigned int(32) compatible_brands[];
}
*/
type FileBox struct {
	BaseBox
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *FileBox) getLeafBox() Box {
	return b
}

//GetFileBox - Implement Box method for this object
func (b *FileBox) GetFileBox() (*FileBox, error) {
	return b, nil
}

//Interface methods Impl - End

//MajorBrand - returns the major brand of the file
func (b *FileBox) MajorBrand() string {
	var ret string
	p := b.BaseBox.getPayload()
	if len(p) >= 4 {
		ret = string(p[0:4])
	}
	//Improper box
	return ret
}

//MinorVersion - returns the minor version of the file
func (b *FileBox) MinorVersion() uint32 {
	var ret uint32
	p := b.BaseBox.getPayload()
	if len(p) >= 8 {
		ret = binary.BigEndian.Uint32(p[4:8])
		return ret
	}
	//Improper box
	return 0
}

//CompatibleBrands - returns the Compatible brands
func (b *FileBox) CompatibleBrands() []string {
	var ret []string
	p := b.BaseBox.getPayload()
	if len(p) >= 8 {
		nEntries := (len(p) - 8) / 4
		ret = make([]string, nEntries)
		bytesRead := 8
		for i := 0; i < nEntries; i++ {
			if len(p) >= bytesRead+4 {
				ret[i] = string(p[bytesRead : bytesRead+4])
				bytesRead += 4
			} else {
				//Improper box
				break
			}
		}
	}
	//Improper box
	return ret
}

//String - Returns User Readable description of content
func (b *FileBox) String() string {
	var ret string
	ret += b.BaseBox.String()
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf(" MajorBrand: %v, MinorVersion: %v, CompatibleBrands:%v ", b.MajorBrand(), b.MinorVersion(), b.CompatibleBrands())
	return ret
}

//SegmentBox - 'styp' is same as 'ftyp'
type SegmentBox struct {
	FileBox
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *SegmentBox) getLeafBox() Box {
	return b
}

//GetSegmentBox - Implement Box method for this object
func (b *SegmentBox) GetSegmentBox() (*SegmentBox, error) {
	return b, nil
}

//Interface methods Impl - End
