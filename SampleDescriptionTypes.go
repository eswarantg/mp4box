package mp4box

import (
	"encoding/binary"
	"fmt"
)

//SampleEntry -
/*
aligned(8) abstract class SampleEntry (unsigned int(32) format) extends Box(format){
	const unsigned int(8)[6] reserved = 0;
	unsigned int(16) data_reference_index;
}
*/
type SampleEntry struct {
	BaseBox
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *SampleEntry) getLeafBox() Box {
	return b
}

//GetSampleDescriptionBox - Implement Box method for this object
func (b *SampleEntry) GetSampleEntryBox() (*SampleEntry, error) {
	return b, nil
}

//getPayload - Returns the payload excluding headers
func (b *SampleEntry) getPayload() []byte {
	var ret []byte
	p := b.BaseBox.getPayload()
	if len(p) > 6+2 {
		return p[8:]
	}
	return ret
}

//Interface methods Impl - End

func (b *SampleEntry) DataReferenceIndex() uint16 {
	var ret uint16
	p := b.BaseBox.getPayload()
	if len(p) >= 6+2 {
		ret = binary.BigEndian.Uint16(p[6:8])
		return ret
	}
	//Improper box
	return ret
}

//String - Display
func (b *SampleEntry) String() string {
	var ret string
	ret += b.BaseBox.String()
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf(" data_reference_index: %v", b.DataReferenceIndex())
	return ret
}

//HintSampleEntry -
/*
class HintSampleEntry() extends SampleEntry (protocol) { unsigned int(8) data [];
}
*/
type HintSampleEntry struct {
	SampleEntry
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *HintSampleEntry) getLeafBox() Box {
	return b
}

//GetSampleDescriptionBox - Implement Box method for this object
func (b *HintSampleEntry) GetHintSampleEntryBox() (*HintSampleEntry, error) {
	return b, nil
}

//Interface methods Impl - End

//VisualSampleEntry -
/*
class VisualSampleEntry(codingname) extends SampleEntry (codingname){
	unsigned int(16) pre_defined = 0;
	const unsigned int(16) reserved = 0;
	unsigned int(32)[3] pre_defined = 0;
	unsigned int(16) width;
	unsigned int(16) height;
	template unsigned int(32) horizresolution = 0x00480000; // 72 dpi template
	unsigned int(32) vertresolution = 0x00480000; // 72 dpi
	const unsigned int(32) reserved = 0;
	template unsigned int(16) frame_count = 1;
	string[32] compressorname;
	template unsigned int(16) depth = 0x0018;
	int(16) pre_defined = -1;
}
*/
type VisualSampleEntry struct {
	SampleEntry
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *VisualSampleEntry) getLeafBox() Box {
	return b
}

//getPayload - Returns the payload excluding headers
func (b *VisualSampleEntry) getPayload() []byte {
	var ret []byte
	p := b.SampleEntry.getPayload()
	if len(p) >= 2+2+(4*3)+2+2+4+4+4+2+4+4 {
		return p[42:]
	}
	return ret
}

//GetSampleDescriptionBox - Implement Box method for this object
func (b *VisualSampleEntry) GetVisualSampleEntryBox() (*VisualSampleEntry, error) {
	return b, nil
}

//Interface methods Impl - End
func (b *VisualSampleEntry) Width() uint16 {
	var ret uint16
	p := b.SampleEntry.getPayload()
	if len(p) >= 2+2+(4*3)+2 {
		ret = binary.BigEndian.Uint16(p[16:18])
		return ret
	}
	//Improper box
	return ret
}
func (b *VisualSampleEntry) Height() uint16 {
	var ret uint16
	p := b.SampleEntry.getPayload()
	if len(p) >= 2+2+(4*3)+2+2 {
		ret = binary.BigEndian.Uint16(p[18:20])
		return ret
	}
	//Improper box
	return ret
}
func (b *VisualSampleEntry) HorizResolution() uint32 {
	var ret uint32
	p := b.SampleEntry.getPayload()
	if len(p) >= 2+2+(4*3)+2+2+4 {
		ret = binary.BigEndian.Uint32(p[20:24])
		return ret
	}
	//Improper box
	return ret
}
func (b *VisualSampleEntry) VertResolution() uint32 {
	var ret uint32
	p := b.SampleEntry.getPayload()
	if len(p) >= 2+2+(4*3)+2+2+4+4 {
		ret = binary.BigEndian.Uint32(p[24:28])
		return ret
	}
	//Improper box
	return ret
}
func (b *VisualSampleEntry) FrameCount() uint16 {
	var ret uint16
	p := b.SampleEntry.getPayload()
	if len(p) >= 2+2+(4*3)+2+2+4+4+4+2 {
		ret = binary.BigEndian.Uint16(p[32:34])
		return ret
	}
	//Improper box
	return ret
}

func (b *VisualSampleEntry) CompressorName() string {
	var ret string
	p := b.SampleEntry.getPayload()
	if len(p) >= 2+2+(4*3)+2+2+4+4+4+2+4 {
		ret = string(p[34:38])
	}
	//Improper box
	return ret
}

//String - Display
func (b *VisualSampleEntry) String() string {
	var ret string
	ret += b.SampleEntry.String()
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf(" Width: %v, Height: %v, HorizResolution:%v, VertResolution:%v, FrameCount:%v, CompressorName:%v ", b.Width(), b.Height(), b.HorizResolution(), b.VertResolution(), b.FrameCount(), b.CompressorName())
	return ret
}

//AudioSampleEntry -
/*
class AudioSampleEntry(codingname) extends SampleEntry (codingname){
	const unsigned int(32)[2] reserved = 0;
	template unsigned int(16) channelcount = 2;
	template unsigned int(16) samplesize = 16;
	unsigned int(16) pre_defined = 0;
	const unsigned int(16) reserved = 0 ;
	template unsigned int(32) samplerate = {timescale of media}<<16;
}
*/
type AudioSampleEntry struct {
	SampleEntry
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *AudioSampleEntry) getLeafBox() Box {
	return b
}

//getPayload - Returns the payload excluding headers
func (b *AudioSampleEntry) getPayload() []byte {
	var ret []byte
	p := b.SampleEntry.getPayload()
	if len(p) > (4*2)+2+2+2+2+4 {
		return p[20:]
	}
	return ret
}

//GetSampleDescriptionBox - Implement Box method for this object
func (b *AudioSampleEntry) GetAudioSampleEntryBox() (*AudioSampleEntry, error) {
	return b, nil
}

//Interface methods Impl - End

func (b *AudioSampleEntry) ChannelCount() uint16 {
	var ret uint16
	p := b.SampleEntry.getPayload()
	if len(p) >= (4*2)+2 {
		ret = binary.BigEndian.Uint16(p[8:10])
		return ret
	}
	//Improper box
	return ret
}

func (b *AudioSampleEntry) SampleSize() uint16 {
	var ret uint16
	p := b.SampleEntry.getPayload()
	if len(p) >= (4*2)+2+2 {
		ret = binary.BigEndian.Uint16(p[10:12])
		return ret
	}
	//Improper box
	return ret
}

func (b *AudioSampleEntry) SampleRate() uint32 {
	var ret uint32
	p := b.SampleEntry.getPayload()
	if len(p) >= (4*2)+2+2+2+2+4 {
		ret = binary.BigEndian.Uint32(p[16:20])
		return (ret >> 16)
	}
	//Improper box
	return ret
}

//String - Display
func (b *AudioSampleEntry) String() string {
	var ret string
	ret += b.SampleEntry.String()
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf(" ChannelCount: %v, SampleSize: %v, SampleRate:%v ", b.ChannelCount(), b.SampleSize(), b.SampleRate())
	return ret
}
