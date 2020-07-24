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
func (b *FileBox) getLeafBox() AccessBoxType {
	return b
}

//GetFileBox - Implement AccessBoxType method for this object
func (b *FileBox) GetFileBox() (*FileBox, error) {
	return b, nil
}

//Interface methods Impl - End

//MajorBrand - returns the major brand of the file
func (b *FileBox) MajorBrand() string {
	p := b.BaseBox.getPayload()
	if p != nil && len(p) >= 4 {
		ret := string(p[0:4])
		return ret
	}
	return ""
}

//MinorVersion - returns the minor version of the file
func (b *FileBox) MinorVersion() uint32 {
	p := b.BaseBox.getPayload()
	if p != nil && len(p) >= 8 {
		ret := binary.BigEndian.Uint32(p[4:8])
		return ret
	}
	return 0
}

//CompatibleBrands - returns the Compatible brands
func (b *FileBox) CompatibleBrands() []string {
	var ret []string
	p := b.BaseBox.getPayload()
	if p != nil && len(p) >= 12 {
		nEntries := (len(p) - 12) / 4
		ret = make([]string, nEntries)
		bytesRead := 8
		for i := 0; i < nEntries; i++ {
			ret[i] = string(p[bytesRead : bytesRead+4])
			bytesRead += 4
		}
	}
	return ret
}

func (b *FileBox) String() string {
	var ret string
	ret += b.BaseBox.String()
	ret += fmt.Sprintf("\n%v MajorBrand: %v, MinorVersion: %v, CompatibleBrands:%v ", b.leadString(), b.MajorBrand(), b.MinorVersion(), b.CompatibleBrands())
	return ret
}

//SegmentBox - 'styp' is same as 'ftyp'
type SegmentBox struct {
	FileBox
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *SegmentBox) getLeafBox() AccessBoxType {
	return b
}

//GetSegmentBox - Implement AccessBoxType method for this object
func (b *SegmentBox) GetSegmentBox() (*SegmentBox, error) {
	return b, nil
}

//Interface methods Impl - End
