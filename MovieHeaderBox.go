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
func (b *MovieHeaderBox) getLeafBox() Box {
	return b
}

//GetMovieHeaderBox - Implement Box method for this object
func (b *MovieHeaderBox) GetMovieHeaderBox() (*MovieHeaderBox, error) {
	return b, nil
}

//Interface methods Impl - End

//CreationTime - CreationTime of the content
func (b *MovieHeaderBox) CreationTime() time.Time {
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
func (b *MovieHeaderBox) ModificationTime() time.Time {
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

//Scale - Ticks per second for all Timing info
func (b *MovieHeaderBox) Scale() uint32 {
	var ret uint32
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		if len(p) >= 12 {
			return binary.BigEndian.Uint32(p[8:12])
		}
	case 1:
		if len(p) >= 20 {
			return binary.BigEndian.Uint32(p[16:20])
		}
	}
	//Improper box
	return ret
}

//Duration - Duration of the content
func (b *MovieHeaderBox) Duration() time.Duration {
	var ret time.Duration
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		if len(p) >= 16 {
			scale := binary.BigEndian.Uint32(p[8:12])
			dur := binary.BigEndian.Uint32(p[12:16])
			if scale != 0 {
				secs := float64(dur) / float64(scale)
				return time.Duration(secs*1000000) * time.Microsecond
			}
		}
	case 1:
		if len(p) >= 28 {
			scale := binary.BigEndian.Uint32(p[16:20])
			dur := binary.BigEndian.Uint64(p[20:28])
			if scale != 0 {
				secs := float64(dur) / float64(scale)
				return time.Duration(secs*1000000) * time.Microsecond
			}
		}
	}
	//Improper box
	return ret
}

//Rate - Rate of the content
func (b *MovieHeaderBox) Rate() uint32 {
	var ret uint32
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		//16
		if len(p) >= 20 {
			return binary.BigEndian.Uint32(p[16:20])
		}
	case 1:
		//28
		if len(p) >= 32 {
			return binary.BigEndian.Uint32(p[28:32])
		}
	}
	//Improper box
	return ret
}

//Volume - Volume of the content
func (b *MovieHeaderBox) Volume() uint16 {
	var ret uint16
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		//16 + 4 = 20
		if len(p) >= 22 {
			return binary.BigEndian.Uint16(p[20:22])
		}
	case 1:
		//28 + 4 = 32
		if len(p) >= 34 {
			return binary.BigEndian.Uint16(p[32:34])
		}
	}
	//Improper box
	return ret
}

//UnityMatrix - matrix
func (b *MovieHeaderBox) UnityMatrix() []uint32 {
	var ret []uint32
	p := b.FullBox.getPayload()
	var bytePos int
	switch b.FullBox.Version() {
	case 0:
		//16 + 4 + 2 + 2 + (2*4) = 32
		bytePos = 32

	case 1:
		//28 + 4 + 2 + 2 + (2*4) = 44
		bytePos = 44
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

//NextTrackID - matrix
func (b *MovieHeaderBox) NextTrackID() uint32 {
	var ret uint32
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		//16 + 4 + 2 + 2 + (2*4) + (4*9) + (4*6) = 92
		if len(p) >= 96 {
			return binary.BigEndian.Uint32(p[92:96])
		}
	case 1:
		//28 + 4 + 2 + 2 + (2*4) + (4*9) + (4*6) = 104
		if len(p) >= 108 {
			return binary.BigEndian.Uint32(p[104:108])
		}
	}
	//Improper box
	return ret
}

//String - Display
func (b *MovieHeaderBox) String() string {
	var ret string
	ret += b.FullBox.String()
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf(" Creation:%v Modification:%v Duration:%v Rate:%v Volume:%v UnityMatrix:%v NextTrackID:%v",
		b.CreationTime(), b.ModificationTime(), b.Duration(), b.Rate(),
		b.Volume(), b.UnityMatrix(), b.NextTrackID())
	return ret
}
