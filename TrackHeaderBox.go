package mp4box

import (
	"encoding/binary"
	"fmt"
	"time"
)

//TrackHeaderBox -
/*
aligned(8)
class TrackHeaderBox extends FullBox(‘tkhd’, version, flags)
{ if (version==1) {
      unsigned int(64)  creation_time;
      unsigned int(64)  modification_time;
      unsigned int(32)  track_ID;
      const unsigned int(32)  reserved = 0;
      unsigned int(64)  duration;
   } else { // version==0
      unsigned int(32)  creation_time;
      unsigned int(32)  modification_time;
      unsigned int(32)  track_ID;
      const unsigned int(32)  reserved = 0;
      unsigned int(32)  duration;
}
const unsigned int(32)[2] reserved = 0;
template int(16) layer = 0;
template int(16) alternate_group = 0;
template int(16) volume = {if track_is_audio 0x0100 else 0}; const unsigned int(16) reserved = 0;
template int(32)[9] matrix=
{ 0x00010000,0,0,0,0x00010000,0,0,0,0x40000000 };
      // unity matrix
   unsigned int(32) width;
   unsigned int(32) height;
}
*/
type TrackHeaderBox struct {
	FullBox
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *TrackHeaderBox) getLeafBox() AccessBoxType {
	return b
}

//GetMTrackHeaderBox - Implement AccessBoxType method for this object
func (b *TrackHeaderBox) GetMTrackHeaderBox() (*TrackHeaderBox, error) {
	return b, nil
}

//Interface methods Impl - End

//CreationTime - CreationTime of the content
func (b *TrackHeaderBox) CreationTime() time.Time {
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		secs := binary.BigEndian.Uint32(p[0:4])
		t := epochTimeMp4
		return t.Add(time.Duration(secs))
	case 1:
		secs := binary.BigEndian.Uint64(p[0:8])
		t := epochTimeMp4
		return t.Add(time.Duration(secs))
	}
	return time.Time{}
}

//ModificationTime - ModificationTime of the content
func (b *TrackHeaderBox) ModificationTime() time.Time {
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		secs := binary.BigEndian.Uint32(p[4:8])
		t := epochTimeMp4
		return t.Add(time.Duration(secs))
	case 1:
		secs := binary.BigEndian.Uint64(p[8:16])
		t := epochTimeMp4
		return t.Add(time.Duration(secs))
	}
	return time.Time{}
}

//Duration - Duration of the content
func (b *TrackHeaderBox) Duration() time.Duration {
	p := b.FullBox.getPayload()
	node, err := b.GetParentByName("moov")
	if err != nil {
		return 0 * time.Second
	}
	moofBox, err := node.GetCollectionBaseBox()
	if err != nil {
		return 0 * time.Second
	}
	node, err = moofBox.GetChildByName("mvhd")
	if err != nil {
		return 0 * time.Second
	}
	mvhdBox, err := node.GetMovieHeaderBox()
	if err != nil {
		return 0 * time.Second
	}
	scale := mvhdBox.Scale()
	if err != nil {
		return 0 * time.Second
	}
	switch b.FullBox.Version() {
	case 0:
		dur := binary.BigEndian.Uint32(p[16:20])
		if scale != 0 {
			return time.Duration(dur / scale)
		}
	case 1:
		dur := binary.BigEndian.Uint64(p[24:32])
		if scale != 0 {
			return time.Duration(dur / uint64(scale))
		}
	}
	return 0 * time.Second
}

//String - Display
func (b *TrackHeaderBox) String() string {
	var ret string
	ret += b.FullBox.String()
	ret += fmt.Sprintf("\n%v Creation:%v Modification:%v Duration:%v", b.leadString(), b.CreationTime(), b.ModificationTime(), b.Duration())
	return ret
}
