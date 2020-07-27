package mp4box

import (
	"encoding/binary"
	"fmt"
)

//TrackFragmentHeaderBox -
/*
aligned(8) class TrackFragmentHeaderBox extends FullBox(‘tfhd’, 0, tf_flags){
	unsigned int(32) track_ID;
	// all the following are optional fields
	unsigned int(64) base_data_offset;
	unsigned int(32) sample_description_index;
	unsigned int(32) default_sample_duration;
	unsigned int(32) default_sample_size;
	unsigned int(32) default_sample_flags
}
*/
type TrackFragmentHeaderBox struct {
	FullBox
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *TrackFragmentHeaderBox) getLeafBox() Box {
	return b
}

//GetTrackFragmentHeaderBox - Implement Box method for this object
func (b *TrackFragmentHeaderBox) GetTrackFragmentHeaderBox() (*TrackFragmentHeaderBox, error) {
	return b, nil
}

//Interface methods Impl - End

//TrackID - Returns TrackID
func (b *TrackFragmentHeaderBox) TrackID() uint32 {
	var ret uint32
	p := b.FullBox.getPayload()
	if len(p) >= 4 {
		return binary.BigEndian.Uint32(p[0:4])
	}
	return ret
}

//BaseDataOffset - Return BaseDataOffset
func (b *TrackFragmentHeaderBox) BaseDataOffset() uint32 {
	var ret uint32
	p := b.FullBox.getPayload()
	if len(p) >= 8 {
		return binary.BigEndian.Uint32(p[4:8])
	}
	return ret
}

//SampleDescriptionIndex - Return SampleDescriptionIndex
func (b *TrackFragmentHeaderBox) SampleDescriptionIndex() uint32 {
	var ret uint32
	p := b.FullBox.getPayload()
	if len(p) >= 12 {
		return binary.BigEndian.Uint32(p[8:12])
	}
	return ret
}

//DefaultSampleDuration - Return DefaultSampleDuration
func (b *TrackFragmentHeaderBox) DefaultSampleDuration() uint32 {
	var ret uint32
	p := b.FullBox.getPayload()
	if len(p) >= 16 {
		return binary.BigEndian.Uint32(p[12:16])
	}
	return ret
}

//DefaultSampleSize - Return DefaultSampleSize
func (b *TrackFragmentHeaderBox) DefaultSampleSize() uint32 {
	var ret uint32
	p := b.FullBox.getPayload()
	if len(p) >= 20 {
		return binary.BigEndian.Uint32(p[16:20])
	}
	return ret
}

//DefaultSampleFlags - Return DefaultSampleFlags
func (b *TrackFragmentHeaderBox) DefaultSampleFlags() uint32 {
	var ret uint32
	p := b.FullBox.getPayload()
	if len(p) >= 24 {
		return binary.BigEndian.Uint32(p[20:24])
	}
	return ret
}

//String - Display
func (b *TrackFragmentHeaderBox) String() string {
	var ret string
	ret += b.FullBox.String()
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf(" TrackID:%v BaseDataOffset:%v SampleDescriptionIndex:%v DefaultSampleDuration:%v DefaultSampleSize:%v DefaultSampleFlags:%v",
		b.TrackID(), b.BaseDataOffset(), b.SampleDescriptionIndex(), b.DefaultSampleSize(), b.DefaultSampleSize(), b.DefaultSampleFlags(),
	)
	return ret
}
