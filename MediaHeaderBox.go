package mp4box

import (
	"encoding/binary"
	"fmt"
	"time"
)

//MediaHeaderBox -
/*
aligned(8) class MediaHeaderBox extends FullBox(‘mdhd’, version, 0) { if (version==1) {
      unsigned int(64)  creation_time;
      unsigned int(64)  modification_time;
      unsigned int(32)  timescale;
      unsigned int(64)  duration;
   } else { // version==0
      unsigned int(32)  creation_time;
      unsigned int(32)  modification_time;
      unsigned int(32)  timescale;
      unsigned int(32)  duration;
}
bit(1) pad=0;
unsigned int(5)[3] language; // ISO-639-2/T language code unsigned int(16) pre_defined = 0;
}
*/
type MediaHeaderBox struct {
	FullBox
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *MediaHeaderBox) getLeafBox() Box {
	return b
}

//GetMediaHeaderBox - Implement Box method for this object
func (b *MediaHeaderBox) GetMediaHeaderBox() (*MediaHeaderBox, error) {
	return b, nil
}

//Interface methods Impl - End

//CreationTime - CreationTime of the content
func (b *MediaHeaderBox) CreationTime() time.Time {
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
func (b *MediaHeaderBox) ModificationTime() time.Time {
	var ret time.Time = epochTimeMp4
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		if len(p) >= 8 {
			secs := binary.BigEndian.Uint32(p[4:8])
			return ret.Add(time.Duration(secs) * time.Second)
		}
	case 1:
		if len(p) >= 16 {
			secs := binary.BigEndian.Uint64(p[8:16])
			return ret.Add(time.Duration(secs) * time.Second)
		}
	}
	//Improper box
	return ret
}

//Scale - Ticks per second for all Timing info
func (b *MediaHeaderBox) Scale() uint32 {
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
func (b *MediaHeaderBox) Duration() time.Duration {
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

//Language - Language of the content
func (b *MediaHeaderBox) Language() string {
	var ret string
	p := b.FullBox.getPayload()
	var bytePos int
	switch b.FullBox.Version() {
	case 0:
		bytePos = 16
	case 1:
		bytePos = 28
	}
	if len(p) >= bytePos+2 {
		//int(5)[3] = 15 bits = 2 bytes
		//p[bytePos : bytePos+2]
		b1 := (p[bytePos] & 0x7C) >> 2                                  //0111 1100
		b2 := ((p[bytePos] & 0x03) << 3) | ((p[bytePos+1] & 0xE0) >> 5) //0000 0011, 1110 0000
		b3 := ((p[bytePos+1] & 0x1F) >> 0)                              //0001 1111
		return string([]byte{b1 + 0x60, b2 + 0x60, b3 + 0x60})
	}
	//Improper box
	return ret
}

//String - Display
func (b *MediaHeaderBox) String() string {
	var ret string
	ret += b.FullBox.String()
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf(" Creation:%v Modification:%v Duration:%v Language:%v",
		b.CreationTime(), b.ModificationTime(), b.Duration(),
		b.Language(),
	)
	return ret
}
