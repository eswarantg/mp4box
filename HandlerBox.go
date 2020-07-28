package mp4box

import (
	"encoding/binary"
	"fmt"
)

//HandlerBox -
/*
aligned(8) class HandlerBox extends FullBox(‘hdlr’, version = 0, 0) {
	unsigned int(32) pre_defined = 0;
	unsigned int(32) handler_type;
	const unsigned int(32)[3] reserved = 0;
   	string   name;
}
*/
type HandlerBox struct {
	FullBox
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *HandlerBox) getLeafBox() Box {
	return b
}

//GetHandlerBox - Implement Box method for this object
func (b *HandlerBox) GetHandlerBox() (*HandlerBox, error) {
	return b, nil
}

//HandlerType - Returns Type of the Handler
func (b *HandlerBox) HandlerType() uint32 {
	var ret uint32
	p := b.FullBox.getPayload()
	if p != nil {
		//4 skip bytes
		if len(p) >= 8 {
			return binary.BigEndian.Uint32(p[4:8])
		}
	}
	return ret
}

//Name - Returns Handler Name
func (b *HandlerBox) Name() string {
	var ret string
	p := b.FullBox.getPayload()
	if p != nil {
		if len(p) >= 21 {
			ret, _ = getNullTermString(p[20:])
		}
	}
	return ret
}

//Interface methods Impl - End

//String - Display
func (b *HandlerBox) String() string {
	var ret string
	ret += b.FullBox.String()
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf(" HandlerType:%v Name:%v ", b.HandlerType(), b.Name())
	return ret
}
