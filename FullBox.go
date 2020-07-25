package mp4box

import "fmt"

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
func (b *FullBox) getLeafBox() Box {
	return b
}

//GetFullBox - Implement Box method for this object
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
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf(" Version: %v, Flags: %v", b.Version(), b.Flags())
	return ret
}
