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
func (b *MovieFragmentBox) Summary() (duration uint64, timescale uint32, trackID uint32, baseMediaDecodeTime uint64) {
	var tboxes []Box
	var sampleDuration *uint32
	var err error

	duration = 0
	timescale = 0
	trackID = 0
	baseMediaDecodeTime = 0

	//Get DefaultSampleDuration from tfhd box
	tboxes, err = b.GetChildrenByName("tfhd")
	if tboxes != nil && err == nil {
		var tfhdbox *TrackFragmentHeaderBox
		for _, tbox := range tboxes {
			tfhdbox, err = tbox.GetTrackFragmentHeaderBox()
			if tfhdbox != nil && err == nil {
				dur := tfhdbox.DefaultSampleDuration()
				if dur > 0 {
					sampleDuration = new(uint32)
					*sampleDuration = dur
				}
				trackID = tfhdbox.TrackID()
			}
		}
	}
	//Get the tfdt box
	tboxes, err = b.GetChildrenByName("tdft")
	if tboxes != nil && err == nil {
		var tfdtbox *TrackFragmentBaseMediaDecodeTimeBox
		for _, tbox := range tboxes {
			tfdtbox, err = tbox.GetTrackFragmentBaseMediaDecodeTimeBox()
			if tfdtbox != nil && err == nil {
				baseMediaDecodeTime = tfdtbox.BaseMediaDecodeTime()
			}
		}
	}
	//Get the trun box
	tboxes, err = b.GetChildrenByName("trun")
	if tboxes != nil && err == nil {
		var trunbox *TrackRunBox
		for _, tbox := range tboxes {
			trunbox, err = tbox.GetTrackRunBox()
			if err == nil {
				//Get TotalSampleDuration = Sum of all SampleDuration
				dur := trunbox.TotalSampleDuration()
				if dur > 0 {
					duration = dur
				}
				//Get SampleCount
				count := trunbox.SampleCount()
				if count > 0 {
					if sampleDuration != nil {
						//Get DefaultSampleDuration
						duration = uint64(*sampleDuration) * uint64(count)
					}
				}
			}
		}
	}
	return
}

//String - Display
func (b *MovieFragmentBox) detailString() string {
	var ret string
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	duration, timescale, trackID, baseMediaDecodeTime := b.Summary()
	ret += fmt.Sprintf(" Duration:%v TimeScale:%v TrackID:%v BaseMediaDecodeTime:%v",
		duration, timescale, trackID, baseMediaDecodeTime,
	)
	return ret
}
