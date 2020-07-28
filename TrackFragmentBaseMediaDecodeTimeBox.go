package mp4box

import (
	"encoding/binary"
	"fmt"
)

//TrackFragmentBaseMediaDecodeTimeBox -
/*
aligned(8) class TrackFragmentBaseMediaDecodeTimeBox extends FullBox("tfdt", version, 0) {
	if (version==1) {
		unsigned int(64) baseMediaDecodeTime;
	} else { // version==0
		unsigned int(32) baseMediaDecodeTime; }
	}
}
*/
type TrackFragmentBaseMediaDecodeTimeBox struct {
	FullBox
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *TrackFragmentBaseMediaDecodeTimeBox) getLeafBox() Box {
	return b
}

//GetTrackFragmentBaseMediaDecodeTimeBox - Implement Box method for this object
func (b *TrackFragmentBaseMediaDecodeTimeBox) GetTrackFragmentBaseMediaDecodeTimeBox() (*TrackFragmentBaseMediaDecodeTimeBox, error) {
	return b, nil
}

//Interface methods Impl - End

//BaseMediaDecodeTime - Duration of the media
func (b *TrackFragmentBaseMediaDecodeTimeBox) BaseMediaDecodeTime() uint64 {
	var ret uint64
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		if len(p) >= 4 {
			return uint64(binary.BigEndian.Uint32(p[0:4]))
		}
	case 1:
		if len(p) >= 8 {
			return binary.BigEndian.Uint64(p[0:8])
		}
	}
	//Improper Box
	return ret
}

//String - Display
func (b *TrackFragmentBaseMediaDecodeTimeBox) String() string {
	var ret string
	ret += b.FullBox.String()
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf(" BaseMediaDecodeTime:%v ", b.BaseMediaDecodeTime())
	return ret
}
