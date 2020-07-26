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
func (b *TrackHeaderBox) getLeafBox() Box {
	return b
}

//GetTrackHeaderBox - Implement Box method for this object
func (b *TrackHeaderBox) GetTrackHeaderBox() (*TrackHeaderBox, error) {
	return b, nil
}

//Interface methods Impl - End

//CreationTime - CreationTime of the content
func (b *TrackHeaderBox) CreationTime() time.Time {
	var ret time.Time = epochTimeMp4
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		if len(p) >= 5 {
			secs := binary.BigEndian.Uint32(p[0:4])
			return ret.Add(time.Duration(secs) * time.Second)
		}
	case 1:
		if len(p) >= 8 {
			secs := binary.BigEndian.Uint64(p[0:8])
			return ret.Add(time.Duration(secs) * time.Second)
		}
	}
	//Improper box
	return ret
}

//ModificationTime - ModificationTime of the content
func (b *TrackHeaderBox) ModificationTime() time.Time {
	var ret time.Time = epochTimeMp4
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		secs := binary.BigEndian.Uint32(p[4:8])
		return ret.Add(time.Duration(secs) * time.Second)
	case 1:
		secs := binary.BigEndian.Uint64(p[8:16])
		return ret.Add(time.Duration(secs) * time.Second)
	}
	//Improper box
	return ret
}

//Duration - Duration of the content
func (b *TrackHeaderBox) Duration(scale uint32) time.Duration {
	var ret time.Duration
	p := b.FullBox.getPayload()
	if scale != 0 {
		switch b.FullBox.Version() {
		case 0:
			if len(p) >= 20 {
				dur := binary.BigEndian.Uint32(p[16:20])
				secs := float64(dur) / float64(scale)
				return time.Duration(secs*1000000) * time.Microsecond
			}
		case 1:
			if len(p) >= 32 {
				dur := binary.BigEndian.Uint64(p[24:32])
				secs := float64(dur) / float64(scale)
				return time.Duration(secs*1000000) * time.Microsecond
			}
		}
	}
	//Improper box
	return ret
}

//Layer - Layer of the content
func (b *TrackHeaderBox) Layer() uint16 {
	var ret uint16
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		//20 + (2*4) = 28
		if len(p) >= 30 {
			return binary.BigEndian.Uint16(p[28:30])
		}
	case 1:
		//32 + (2*4) = 40
		if len(p) >= 42 {
			return binary.BigEndian.Uint16(p[40:42])
		}
	}
	//Improper box
	return ret
}

//AlernateGroup - Alternate Group
func (b *TrackHeaderBox) AlernateGroup() uint16 {
	var ret uint16
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		//20 + (2*4) + 2 = 30
		if len(p) >= 32 {
			return binary.BigEndian.Uint16(p[30:32])
		}
	case 1:
		//32 + (2*4) + 2 = 42
		if len(p) >= 44 {
			return binary.BigEndian.Uint16(p[42:44])
		}
	}
	//Improper box
	return ret
}

//Volume - Volume of the content
func (b *TrackHeaderBox) Volume() uint16 {
	var ret uint16
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		//20 + (2*4) + 2 + 2 = 32
		if len(p) >= 34 {
			return binary.BigEndian.Uint16(p[32:34])
		}
	case 1:
		//32 + (2*4) + 2 + 2 = 44
		if len(p) >= 46 {
			return binary.BigEndian.Uint16(p[44:46])
		}
	}
	//Improper box
	return ret
}

//UnityMatrix - matrix
func (b *TrackHeaderBox) UnityMatrix() []uint32 {
	var ret []uint32
	p := b.FullBox.getPayload()
	var bytePos int
	switch b.FullBox.Version() {
	case 0:
		//20 + (2*4) + 2 + 2 + 2 = 34
		bytePos = 34
	case 1:
		//32 + (2*4) + 2 + 2 + 2 = 46
		bytePos = 46
	}
	if len(p) >= bytePos+(9*4) {
		ret = []uint32{0, 0, 0, 0, 0, 0, 0, 0, 0}
		for i := 0; i < 9; i++ {
			ret[i] = binary.BigEndian.Uint32(p[bytePos : bytePos+4])
			bytePos += 4
		}
		return ret
	}
	//Improper box
	return ret
}

//Width - Width
func (b *TrackHeaderBox) Width() uint32 {
	var ret uint32
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		//20 + (2*4) + 2 + 2 + 2 + (9*4)= 70
		if len(p) >= 74 {
			return binary.BigEndian.Uint32(p[70:74])
		}
	case 1:
		//32 + (2*4) + 2 + 2 + 2 + (9*4)= 82
		if len(p) >= 86 {
			return binary.BigEndian.Uint32(p[82:86])
		}
	}
	//Improper box
	return ret
}

//Height - Height
func (b *TrackHeaderBox) Height() uint32 {
	var ret uint32
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		//20 + (2*4) + 2 + 2 + 2 + (9*4) + 4 = 74
		if len(p) >= 78 {
			return binary.BigEndian.Uint32(p[74:78])
		}

	case 1:
		//32 + (2*4) + 2 + 2 + 2 + (9*4) + 4 = 86
		if len(p) >= 90 {
			return binary.BigEndian.Uint32(p[86:90])
		}
	}
	//Improper box
	return ret
}

//String - Display
func (b *TrackHeaderBox) String() string {
	var ret string
	ret += b.FullBox.String()
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf(" Creation:%v Modification:%v Duration:%v Layer:%v, AlternateGroup: %v, Volume:%v, UnityMatrix:%v, Width:%v, Height:%v",
		b.CreationTime(), b.ModificationTime(), b.Duration(90000),
		b.Layer(), b.AlernateGroup(), b.Volume(), b.UnityMatrix(), b.Width(), b.Height(),
	)
	return ret
}
