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

//Interface methods Impl - End

//TotalDuration - Returns total duration of the content
func (b *MovieFragmentBox) TotalDuration() uint64 {
	var ret uint64
	var tbox Box
	var sampleDuration *uint32
	var err error

	//Get DefaultSampleDuration from tfhd box
	tbox, err = b.GetChildByName("tfhd")
	if err == nil {
		tfhdbox, err := tbox.GetTrackFragmentHeaderBox()
		if err == nil {
			dur := tfhdbox.DefaultSampleDuration()
			if dur > 0 {
				sampleDuration = new(uint32)
				*sampleDuration = dur
			}
		}
		//Get the trun box
		tbox, err = b.GetChildByName("trun")
		if err == nil {
			trunbox, err := tbox.GetTrackRunBox()
			if err == nil {
				//Get TotalSampleDuration = Sum of all SampleDuration
				dur := trunbox.TotalSampleDuration()
				if dur > 0 {
					return dur
				}
				//Get SampleCount
				count := trunbox.SampleCount()
				if count > 0 {
					if sampleDuration != nil {
						//Get DefaultSampleDuration
						return uint64(*sampleDuration) * uint64(count)
					}
				}
			}
		}
	}
	return ret
}

//String - Display
func (b *MovieFragmentBox) detailString() string {
	var ret string
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf(" TotalDuration:%v ", b.TotalDuration())
	return ret
}
