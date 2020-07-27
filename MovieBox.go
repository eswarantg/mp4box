package mp4box

import (
	"fmt"
	"time"
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

//Interface methods Impl - End

//TotalDuration - Returns total duration of the content
func (b *MovieBox) TotalDuration() time.Duration {
	var ret time.Duration
	return ret
}

//String - Display
func (b *MovieBox) detailString() string {
	var ret string
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf(" TotalDuration:%v ", b.TotalDuration())
	return ret
}
