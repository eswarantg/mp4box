package mp4box

import (
	"bytes"
	"fmt"
)

/*
aligned(8) class ESDBox
 extends FullBox(‘esds’, version = 0, 0) {
 ES_Descriptor ES;
}
*/
type ESDBox struct {
	FullBox
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *ESDBox) getLeafBox() Box {
	return b
}

//GetSampleDescriptionBox - Implement Box method for this object
func (b *ESDBox) GetESDBox() (*ESDBox, error) {
	return b, nil
}

//Interface methods Impl - End
func (b *ESDBox) ES_Descriptor() (ES_Descriptor, error) {
	p := b.FullBox.getPayload()
	ret := ES_Descriptor{}
	used, err := ret.initData(&p, b.level+1)
	fmt.Printf("\nES_Descriptor bytes left o:%v p:%v", used, len(p))
	return ret, err
}

//String - Display
func (b *ESDBox) String() string {
	var ret string
	ret += b.FullBox.String()
	desc, err := b.ES_Descriptor()
	if err == nil {
		ret += desc.String()
	}
	return ret
}

/*
 // Visual Streams
class MP4VisualSampleEntry() extends VisualSampleEntry ('mp4v'){
 ESDBox ES;
}
*/
type MP4VisualSampleEntry struct {
	VisualSampleEntry
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *MP4VisualSampleEntry) getLeafBox() Box {
	return b
}

//GetSampleDescriptionBox - Implement Box method for this object
func (b *MP4VisualSampleEntry) GetMP4VisualSampleEntry() (*MP4VisualSampleEntry, error) {
	return b, nil
}

//Interface methods Impl - End
func (b *MP4VisualSampleEntry) ESDBox() *ESDBox {
	p := b.VisualSampleEntry.getPayload()
	if p != nil {
		r := bytes.NewReader(p)
		reader := newBoxReader(r, b)
		box, err := reader.NextBox()
		if err == nil {
			esdsbox, err := box.GetESDBox()
			if err == nil {
				return esdsbox
			}
		}
	}
	return nil
}

//String - Display
func (b *MP4VisualSampleEntry) String() string {
	var ret string
	ret += b.VisualSampleEntry.String()
	esdsbox := b.ESDBox()
	if esdsbox != nil {
		ret += esdsbox.String()
	}
	return ret
}

/*
 // Audio Streams
class MP4AudioSampleEntry() extends AudioSampleEntry ('mp4a'){
 ESDBox ES;
}
*/
type MP4AudioSampleEntry struct {
	AudioSampleEntry
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *MP4AudioSampleEntry) getLeafBox() Box {
	return b
}

//GetSampleDescriptionBox - Implement Box method for this object
func (b *MP4AudioSampleEntry) GetMP4AudioSampleEntry() (*MP4AudioSampleEntry, error) {
	return b, nil
}

//Interface methods Impl - End
func (b *MP4AudioSampleEntry) ESDBox() (*ESDBox, error) {
	var err error
	p := b.AudioSampleEntry.getPayload()
	if p != nil {
		r := bytes.NewReader(p)
		reader := newBoxReader(r, b)
		box, err := reader.NextBox()
		if err == nil {
			esdsbox, err := box.GetESDBox()
			if err == nil {
				return esdsbox, nil
			}
		}
	}
	return nil, err
}

//String - Display
func (b *MP4AudioSampleEntry) String() string {
	var ret string
	ret += b.AudioSampleEntry.String()
	esdsbox, err := b.ESDBox()
	if err != nil {
		ret += fmt.Sprintf("\n%d%v  %v", b.level, b.leadString(), err.Error())
	}
	if esdsbox != nil {
		ret += esdsbox.String()
	}
	return ret
}

/*
 // all other Mpeg stream types
class MpegSampleEntry() extends SampleEntry ('mp4s'){
 ESDBox ES;
}
*/
type MpegSampleEntry struct {
	SampleEntry
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *MpegSampleEntry) getLeafBox() Box {
	return b
}

//GetSampleDescriptionBox - Implement Box method for this object
func (b *MpegSampleEntry) GetMpegSampleEntry() (*MpegSampleEntry, error) {
	return b, nil
}

//Interface methods Impl - End

//String - Display
func (b *MpegSampleEntry) String() string {
	var ret string
	ret += b.SampleEntry.String()
	return ret
}
