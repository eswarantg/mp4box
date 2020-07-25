package mp4box

import (
	"encoding/binary"
	"fmt"
)

//MovieFragmentHeaderBox -
/*
aligned(8) class MovieFragmentHeaderBox extends FullBox(‘mfhd’, 0, 0){
	unsigned int(32)  sequence_number;
 }
*/
type MovieFragmentHeaderBox struct {
	FullBox
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *MovieFragmentHeaderBox) getLeafBox() Box {
	return b
}

//GetMovieFragmentHeaderBox - Implement Box method for this object
func (b *MovieFragmentHeaderBox) GetMovieFragmentHeaderBox() (*MovieFragmentHeaderBox, error) {
	return b, nil
}

//Interface methods Impl - End

//SequenceNumber - Segment sequence number
func (b *MovieFragmentHeaderBox) SequenceNumber() uint32 {
	p := b.FullBox.getPayload()
	return binary.BigEndian.Uint32(p[0:4])
}

//String - Display
func (b *MovieFragmentHeaderBox) String() string {
	var ret string
	ret += b.FullBox.String()
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf(" SequenceNumber:%v ", b.SequenceNumber())
	return ret
}
