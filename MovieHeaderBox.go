package mp4box

import (
	"encoding/binary"
	"fmt"
	"time"
)

//MovieHeaderBox -
/*
aligned(8)
class MovieHeaderBox extends FullBox(‘mvhd’, version, 0)
{
	if (version==1) {
		unsigned int(64)  creation_time; //(in seconds since midnight, Jan. 1, 1904, in UTC time)
		unsigned int(64)  modification_time; //(in seconds since midnight, Jan. 1, 1904, in UTC time)
		unsigned int(32)  timescale;
		unsigned int(64)  duration;
 	} else { // version==0
		unsigned int(32)  creation_time; //(in seconds since midnight, Jan. 1, 1904, in UTC time)
		unsigned int(32)  modification_time; //(in seconds since midnight, Jan. 1, 1904, in UTC time)
		unsigned int(32)  timescale;
		unsigned int(32)  duration;
	}
	template int(32) rate = 0x00010000; // typically 1.0
	template int(16) volume = 0x0100; // typically, full volume const bit(16) reserved = 0;
	const unsigned int(32)[2] reserved = 0;
	template int(32)[9] matrix =
	{ 0x00010000,0,0,0,0x00010000,0,0,0,0x40000000 };
		// Unity matrix
	bit(32)[6]  pre_defined = 0;
	unsigned int(32)  next_track_ID;
}*/
type MovieHeaderBox struct {
	FullBox
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *MovieHeaderBox) getLeafBox() AccessBoxType {
	return b
}

//GetMovieHeaderBox - Implement AccessBoxType method for this object
func (b *MovieHeaderBox) GetMovieHeaderBox() (*MovieHeaderBox, error) {
	return b, nil
}

//Interface methods Impl - End

//CreationTime - CreationTime of the content
func (b *MovieHeaderBox) CreationTime() time.Time {
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		secs := binary.BigEndian.Uint32(p[0:4])
		t := epochTimeMp4
		return t.Add(time.Duration(secs) * time.Second)
	case 1:
		secs := binary.BigEndian.Uint64(p[0:8])
		t := epochTimeMp4
		return t.Add(time.Duration(secs) * time.Second)
	}
	return time.Time{}
}

//ModificationTime - ModificationTime of the content
func (b *MovieHeaderBox) ModificationTime() time.Time {
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		secs := binary.BigEndian.Uint32(p[4:8])
		t := epochTimeMp4
		return t.Add(time.Duration(secs) * time.Second)
	case 1:
		secs := binary.BigEndian.Uint64(p[8:16])
		t := epochTimeMp4
		return t.Add(time.Duration(secs) * time.Second)
	}
	return time.Time{}
}

//Scale - Ticks per second for all Timing info
func (b *MovieHeaderBox) Scale() uint32 {
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		scale := binary.BigEndian.Uint32(p[8:12])
		return scale
	case 1:
		scale := binary.BigEndian.Uint32(p[16:20])
		return scale
	}
	return 0
}

//Duration - Duration of the content
func (b *MovieHeaderBox) Duration() time.Duration {
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		scale := binary.BigEndian.Uint32(p[8:12])
		dur := binary.BigEndian.Uint32(p[12:16])
		if scale != 0 {
			secs := float64(dur) / float64(scale)
			return time.Duration(secs*1000000) * time.Microsecond
		}
	case 1:
		scale := binary.BigEndian.Uint32(p[16:20])
		dur := binary.BigEndian.Uint64(p[20:28])
		if scale != 0 {
			secs := float64(dur) / float64(scale)
			return time.Duration(secs*1000000) * time.Microsecond
		}
	}
	return 0 * time.Second
}

//String - Display
func (b *MovieHeaderBox) String() string {
	var ret string
	ret += b.FullBox.String()
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf(" Creation:%v Modification:%v Duration:%v", b.CreationTime(), b.ModificationTime(), b.Duration())
	return ret
}
