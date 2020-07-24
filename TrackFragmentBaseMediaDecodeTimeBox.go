package mp4box

import (
	"encoding/binary"
	"fmt"
	"time"
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
func (b *TrackFragmentBaseMediaDecodeTimeBox) getLeafBox() AccessBoxType {
	return b
}

//GetTrackFragmentBaseMediaDecodeTimeBox - Implement AccessBoxType method for this object
func (b *TrackFragmentBaseMediaDecodeTimeBox) GetTrackFragmentBaseMediaDecodeTimeBox() (*TrackFragmentBaseMediaDecodeTimeBox, error) {
	return b, nil
}

//Interface methods Impl - End

//BaseMediaDecodeTime - Duration of the media
func (b *TrackFragmentBaseMediaDecodeTimeBox) BaseMediaDecodeTime(timescale uint32) time.Duration {
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		ticks := binary.BigEndian.Uint32(p[0:4])
		secs := ticks / timescale
		return time.Duration(secs)
	case 1:
		ticks := binary.BigEndian.Uint64(p[0:8])
		secs := ticks / uint64(timescale)
		return time.Duration(secs)
	}
	return 0
}

//String - Display
func (b *TrackFragmentBaseMediaDecodeTimeBox) String() string {
	var ret string
	ret += b.FullBox.String()
	ret += fmt.Sprintf("\n%v BaseMediaDecodeTime:%v ", b.leadString(), b.BaseMediaDecodeTime(90000))
	return ret
}
