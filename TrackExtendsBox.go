package mp4box

import (
	"encoding/binary"
	"fmt"
)

//TrackExtendsBox -
/*
aligned(8) class TrackExtendsBox extends FullBox(‘trex’, 0, 0){
	unsigned int(32) track_ID;
	unsigned int(32) default_sample_description_index;
	unsigned int(32) default_sample_duration;
	unsigned int(32) default_sample_size;
	unsigned int(32) default_sample_flags
}
*/
type TrackExtendsBox struct {
	FullBox
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *TrackExtendsBox) getLeafBox() AccessBoxType {
	return b
}

//GetMTrackHeaderBox - Implement AccessBoxType method for this object
func (b *TrackExtendsBox) GetMTrackHeaderBox() (*TrackExtendsBox, error) {
	return b, nil
}

//Interface methods Impl - End

//TrackID - Returns TrackID
func (b *TrackExtendsBox) TrackID() uint32 {
	p := b.FullBox.getPayload()
	return binary.BigEndian.Uint32(p[0:4])
}

//String - Display
func (b *TrackExtendsBox) String() string {
	var ret string
	ret += b.FullBox.String()
	ret += fmt.Sprintf("\n%v TrackID:%v ", b.leadString(), b.TrackID())
	return ret
}