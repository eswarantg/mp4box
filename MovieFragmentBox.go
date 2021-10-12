package mp4box

import (
	"fmt"
)

//MovieFragmentBox -
/*
aligned(8) class MovieFragmentBox extends Box(‘moof’){
}
*/
type MovieFragmentBox struct {
	CollectionBaseBox
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *MovieFragmentBox) getLeafBox() Box {
	return b
}

//GetMovieFragmentBox - Implement Box method for this object
func (b *MovieFragmentBox) GetMovieFragmentBox() (*MovieFragmentBox, error) {
	return b, nil
}

func (b *MovieFragmentBox) initData(boxSize int64, boxType string, payload *[]byte, parent Box) error {
	return b.CollectionBaseBox.initData(boxSize, boxType, payload, parent)
}

//GetChildrenByName - Get child by name
func (b *MovieFragmentBox) GetChildrenByName(boxType string) ([]Box, error) {
	return b.CollectionBaseBox.GetChildrenByName(boxType)
}

//Interface methods Impl - End

//Summary - Returns summary of parameters
func (b *MovieFragmentBox) Summary() (sequenceNumber *uint32, baseMediaDecodeTime *uint64, trackID *uint32, timescale *uint32) {
	var tboxes []Box
	var err error

	//Get Sequence No from tfhd box
	tboxes, err = b.GetChildrenByName("mfhd")
	if tboxes != nil && err == nil {
		var mfhdbox *MovieFragmentHeaderBox
		for _, tbox := range tboxes {
			mfhdbox, err = tbox.GetMovieFragmentHeaderBox()
			if mfhdbox != nil && err == nil {
				sequenceNumber = new(uint32)
				*sequenceNumber = mfhdbox.SequenceNumber()
			}
		}
	}
	//Get the tfdt box
	tboxes, err = b.GetChildrenByName("tfdt")
	if tboxes != nil && err == nil {
		var tfdtbox *TrackFragmentBaseMediaDecodeTimeBox
		for _, tbox := range tboxes {
			tfdtbox, err = tbox.GetTrackFragmentBaseMediaDecodeTimeBox()
			if tfdtbox != nil && err == nil {
				baseMediaDecodeTime = new(uint64)
				*baseMediaDecodeTime = tfdtbox.BaseMediaDecodeTime()
			}
		}
	}

	//Get DefaultSampleDuration from tfhd box
	tboxes, err = b.GetChildrenByName("tfhd")
	if tboxes != nil && err == nil {
		var tfhdbox *TrackFragmentHeaderBox
		for _, tbox := range tboxes {
			tfhdbox, err = tbox.GetTrackFragmentHeaderBox()
			if tfhdbox != nil && err == nil {
				trackID = new(uint32)
				*trackID = tfhdbox.TrackID()
			}
		}
	}
	return
}

//String - Display
func (b *MovieFragmentBox) detailString() string {
	var ret string

	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())

	sequenceNumber, baseMediaDecodeTime, trackID, timescale := b.Summary()
	ret += " Summary:"
	if sequenceNumber != nil {
		ret += fmt.Sprintf(" SequenceNumber:%v", *sequenceNumber)
	}
	if baseMediaDecodeTime != nil {
		ret += fmt.Sprintf(" BaseMediaDecodeTime:%v", *baseMediaDecodeTime)
	}
	if trackID != nil {
		ret += fmt.Sprintf(" TrackID:%v", *trackID)
	}
	if timescale != nil {
		ret += fmt.Sprintf(" Timescale:%v", *timescale)
	}
	return ret
}
