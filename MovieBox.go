package mp4box

import (
	"fmt"
)

//MovieBox -
/*
aligned(8) class MovieBox extends Box(‘moov’){
}
*/
type MovieBox struct {
	CollectionBaseBox
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *MovieBox) getLeafBox() Box {
	return b
}

//GetMovieBox - Implement Box method for this object
func (b *MovieBox) GetMovieBox() (*MovieBox, error) {
	return b, nil
}

func (b *MovieBox) initData(boxSize int64, boxType string, payload *[]byte, parent Box) error {
	return b.CollectionBaseBox.initData(boxSize, boxType, payload, parent)
}

//GetChildrenByName - Get child by name
func (b *MovieBox) GetChildrenByName(boxType string) ([]Box, error) {
	return b.CollectionBaseBox.GetChildrenByName(boxType)
}

//Interface methods Impl - End

//TrackID - Returns TrackID
func (b *MovieBox) TrackID() uint32 {
	var ret uint32
	return ret
}

//TimeScale - Returns TimeScale
func (b *MovieBox) TimeScale() uint32 {
	var ret uint32
	return ret
}

//Summary - Returns summary of parameters
func (b *MovieBox) Summary() (duration uint64, timescale uint32, trackID uint32) {
	var err error
	var tboxes []Box

	duration = 0
	timescale = 0
	trackID = 0

	tboxes, err = b.GetChildrenByName("mvhd")
	if tboxes != nil && err == nil {
		var mvhdbox *MovieHeaderBox
		for _, tbox := range tboxes {
			mvhdbox, err = tbox.GetMovieHeaderBox()
			if mvhdbox != nil && err == nil {
				d := mvhdbox.Duration()
				if d > 0 {
					duration = d
				}
				ts := mvhdbox.TimeScale()
				if ts > 0 {
					timescale = ts
				}
			}
		}
	}
	tboxes, err = b.GetChildrenByName("tkhd")
	if tboxes != nil && err == nil {
		var tkhdbox *TrackHeaderBox
		for _, tbox := range tboxes {
			tkhdbox, err = tbox.GetTrackHeaderBox()
			if tkhdbox != nil && err == nil {
				d := tkhdbox.Duration()
				if d > 0 {
					duration = d
				}
				trackID = tkhdbox.TrackID()
			}
		}
	}
	tboxes, err = b.GetChildrenByName("mdhd")
	if tboxes != nil && err == nil {
		var mdhdbox *MediaHeaderBox
		for _, tbox := range tboxes {
			mdhdbox, err = tbox.GetMediaHeaderBox()
			if mdhdbox != nil && err == nil {
				d := mdhdbox.Duration()
				if d > 0 {
					duration = d
				}
				ts := mdhdbox.TimeScale()
				if ts > 0 {
					timescale = ts
				}
			}
		}
	}
	return
}

//String - Display
func (b *MovieBox) detailString() string {
	var ret string
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	duration, timescale, trackid := b.Summary()
	ret += fmt.Sprintf(" Summary:%v %v %v", duration, timescale, trackid)
	return ret
}
