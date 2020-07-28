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
func (b *MovieBox) Summary() (duration *uint64, timescale *uint32, trackID *uint32) {
	var err error
	var tboxes []Box

	tboxes, err = b.GetChildrenByName("mvhd")
	if tboxes != nil && err == nil {
		var mvhdbox *MovieHeaderBox
		for _, tbox := range tboxes {
			mvhdbox, err = tbox.GetMovieHeaderBox()
			if mvhdbox != nil && err == nil {
				d := mvhdbox.Duration()
				if d > 0 && d != 0xFFFFFFFFFFFFFFFF {
					duration = new(uint64)
					*duration = d
				}
				ts := mvhdbox.TimeScale()
				if ts > 0 {
					timescale = new(uint32)
					*timescale = ts
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
				if d > 0 && d != 0xFFFFFFFFFFFFFFFF {
					if duration != nil {
						duration = new(uint64)
					}
					*duration = d
				}
				trackID = new(uint32)
				*trackID = tkhdbox.TrackID()
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
	ret += fmt.Sprintf(" Summary:")
	if duration != nil {
		ret += fmt.Sprintf(" Duration:%v", *duration)
	}
	if timescale != nil {
		ret += fmt.Sprintf(" Timescale:%v", *timescale)
	}
	if trackid != nil {
		ret += fmt.Sprintf(" TrackID:%v", *trackid)
	}
	return ret
}
