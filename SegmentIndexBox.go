package mp4box

import (
	"encoding/binary"
	"fmt"
)

//SegmentIndexBox -
/*
aligned(8) class SegmentIndexBox extends FullBox("sidx", version, 0) {
	unsigned int(32) reference_ID;
	unsigned int(32) timescale;
	if (version==0)
	{
		bit (1)				reference_type;
		unsigned int(31)	referenced_size;
		unsigned int(32)	subsegment_duration;
		bit(1)				starts_with_SAP;
		unsigned int(3)		SAP_type;
		unsigned int(28)	SAP_delta_time;
		unsigned int(32) earliest_presentation_time;
		unsigned int(32) first_offset;
	} else {
		unsigned int(64) earliest_presentation_time;
		unsigned int(64) first_offset;
	}
	unsigned int(16) reserved = 0;
	unsigned int(16) reference_count;
	for(i=1; i <= reference_count; i++) {
	}
}
*/
type SegmentIndexBox struct {
	FullBox
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *SegmentIndexBox) getLeafBox() Box {
	return b
}

//GetSegmentIndexBox - Implement Box method for this object
func (b *SegmentIndexBox) GetSegmentIndexBox() (*SegmentIndexBox, error) {
	return b, nil
}

//Interface methods Impl - End

//ReferenceID - Returns ReferenceID
func (b *SegmentIndexBox) ReferenceID() uint32 {
	var ret uint32
	p := b.FullBox.getPayload()
	if len(p) >= 4 {
		return binary.BigEndian.Uint32(p[0:4])
	}
	//Improper box
	return ret
}

//TimeScale - Return TimeScale
func (b *SegmentIndexBox) TimeScale() uint32 {
	var ret uint32
	p := b.FullBox.getPayload()
	if len(p) >= 8 {
		return binary.BigEndian.Uint32(p[4:8])
	}
	//Improper box
	return ret
}

//EarliestPresentationTime - Return EarliestPresentationTime
func (b *SegmentIndexBox) EarliestPresentationTime() uint64 {
	var ret uint64
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		//3 * 4 = 12 skip
		if len(p) >= 24 {
			return uint64(binary.BigEndian.Uint32(p[20:24]))
		}
	default:
		if len(p) >= 16 {
			return binary.BigEndian.Uint64(p[8:16])
		}
	}
	//Improper box
	return ret
}

//FirstOffset - Return FirstOffset
func (b *SegmentIndexBox) FirstOffset() uint64 {
	var ret uint64
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		//3 * 4 = 12 skip + 4
		if len(p) >= 28 {
			return uint64(binary.BigEndian.Uint32(p[24:28]))
		}
	default:
		if len(p) >= 24 {
			return binary.BigEndian.Uint64(p[16:24])
		}
	}
	//Improper box
	return ret
}

//String - Display
func (b *SegmentIndexBox) String() string {
	var ret string
	ret += b.FullBox.String()
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf(" ReferenceID:%v TimeScale:%v EarliestPresentationTime:%v FirstOffset:%v",
		b.ReferenceID(), b.TimeScale(), b.EarliestPresentationTime(), b.FirstOffset(),
	)
	return ret
}
