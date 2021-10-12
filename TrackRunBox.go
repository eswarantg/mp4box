package mp4box

import (
	"encoding/binary"
	"fmt"
)

//TrackRunBox -
/*
aligned(8) class TrackRunBox extends FullBox(‘trun’, 0, tr_flags) {
	unsigned int(32) sample_count;
	// the following are optional fields
	signed int(32) data_offset;
	unsigned int(32) first_sample_flags;
	// all fields in the following array are optional {
	unsigned int(32) sample_duration;
	unsigned int(32) sample_size;
	unsigned int(32) sample_flags
	unsigned int(32) sample_composition_time_offset;
   	}[ sample_count ]
}
*/
type TrackRunBox struct {
	FullBox
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *TrackRunBox) getLeafBox() Box {
	return b
}

//GetTrackRunBox - Implement Box method for this object
func (b *TrackRunBox) GetTrackRunBox() (*TrackRunBox, error) {
	return b, nil
}

//Interface methods Impl - End

//SampleCount - Returns SampleCount
func (b *TrackRunBox) SampleCount() uint32 {
	var ret uint32
	p := b.FullBox.getPayload()
	if len(p) >= 4 {
		return binary.BigEndian.Uint32(p[0:4])
	}
	//Improper box
	return ret
}

//DataOffset - Returns DataOffset
func (b *TrackRunBox) DataOffset() *uint32 {
	var ret *uint32
	flags := b.FullBox.Flags()
	if (flags & 0x000001) > 0 {
		p := b.FullBox.getPayload()
		if len(p) >= 8 {
			ret = new(uint32)
			*ret = binary.BigEndian.Uint32(p[4:8])
			return ret
		}
		//Improper box
	}
	return ret
}

//FirstSampleFlags - Returns FirstSampleFlags
func (b *TrackRunBox) FirstSampleFlags() *uint32 {
	var ret *uint32
	flags := b.FullBox.Flags()
	if (flags & 0x000004) > 0 {
		p := b.FullBox.getPayload()
		if len(p) >= 12 {
			ret = new(uint32)
			*ret = binary.BigEndian.Uint32(p[8:12])
			return ret
		}
		//Improper box
	}
	return ret
}

//SampleDuration - Get SampleDuration
func (b *TrackRunBox) SampleDuration() []int32 {
	var ret []int32
	flags := b.FullBox.Flags()
	if (flags & 0x000100) > 0 {
		count := b.SampleCount()
		p := b.FullBox.getPayload()
		offset := 12
		if count > 0 {
			if int64(len(p)) >= int64(offset)+(int64(count)*4) {
				ret = make([]int32, count)
				for i := 0; uint32(i) < count; i++ {
					ret[i] = int32(binary.BigEndian.Uint32(p[offset : offset+4]))
					offset += 4
				}
				return ret
			}
			//Improper Box
		}
	}
	return ret
}

//SampleSize - Get SampleSize
func (b *TrackRunBox) SampleSize() []int32 {
	var ret []int32
	flags := b.FullBox.Flags()
	if (flags & 0x000200) > 0 {
		count := b.SampleCount()
		p := b.FullBox.getPayload()
		//SampleDuration present
		offset := 12 + ((4 * count) * 1)
		if count > 0 {
			if int64(len(p)) >= int64(offset)+(int64(count)*4) {
				ret = make([]int32, count)
				for i := 0; uint32(i) < count; i++ {
					ret[i] = int32(binary.BigEndian.Uint32(p[offset : offset+4]))
					offset += 4
				}
				return ret
			}
			//Improper Box
		}
	}
	return ret
}

//SampleFlags - Get SampleFlags
func (b *TrackRunBox) SampleFlags() []int32 {
	var ret []int32
	flags := b.FullBox.Flags()
	if (flags & 0x000200) > 0 {
		count := b.SampleCount()
		p := b.FullBox.getPayload()
		//SampleDuration, SampleSize present
		offset := 12 + ((4 * count) * 2)
		if count > 0 {
			if int64(len(p)) >= int64(offset)+(int64(count)*4) {
				ret = make([]int32, count)
				for i := 0; uint32(i) < count; i++ {
					ret[i] = int32(binary.BigEndian.Uint32(p[offset : offset+4]))
					offset += 4
				}
				return ret
			}
			//Improper Box
		}
	}
	return ret
}

//SampleCompositionTimeOffset - Get SampleCompositionTimeOffset
func (b *TrackRunBox) SampleCompositionTimeOffset() []int32 {
	var ret []int32
	flags := b.FullBox.Flags()
	if (flags & 0x000200) > 0 {
		count := b.SampleCount()
		p := b.FullBox.getPayload()
		//SampleDuration, SampleSize, SampleFlags present
		offset := 12 + ((4 * count) * 3)
		if count > 0 {
			if int64(len(p)) >= int64(offset)+(int64(count)*4) {
				ret = make([]int32, count)
				for i := 0; uint32(i) < count; i++ {
					ret[i] = int32(binary.BigEndian.Uint32(p[offset : offset+4]))
					offset += 4
				}
				return ret
			}
			//Improper Box
		}
	}
	return ret
}

//TotalSampleDuration - return total of the sample durations
func (b *TrackRunBox) TotalSampleDuration() uint64 {
	var ret uint64
	flags := b.FullBox.Flags()
	if (flags & 0x000100) > 0 {
		count := b.SampleCount()
		p := b.FullBox.getPayload()
		offset := 12
		if count > 0 {
			if int64(len(p)) >= int64(offset)+(int64(count)*4) {
				for i := 0; uint32(i) < count; i++ {
					ret += uint64(binary.BigEndian.Uint32(p[offset : offset+4]))
					offset += 4
				}
				return ret
			}
			//Improper Box
		}
	}
	return ret
}

//String - Display
func (b *TrackRunBox) String() string {
	var ret string
	var v *uint32
	ret += b.FullBox.String()
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf(" TotalSampleDuration:%v ", b.TotalSampleDuration())
	ret += fmt.Sprintf(" SampleCount:%v ", b.SampleCount())
	v = b.DataOffset()
	if v != nil {
		ret += fmt.Sprintf(" ,SampleCount:%v ", *v)
	} else {
		ret += " SampleCount: <NIL> "
	}
	v = b.DataOffset()
	if v != nil {
		ret += fmt.Sprintf(" ,SampleCount:%v ", *v)
	} else {
		ret += " SampleCount: <NIL> "
	}
	v = b.FirstSampleFlags()
	if v != nil {
		ret += fmt.Sprintf(" ,FirstSampleFlags:%v ", *v)
	} else {
		ret += " FirstSampleFlags: <NIL> "
	}
	sampleDuration := b.SampleDuration()
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf("SD: %v", sampleDuration)
	sampleSize := b.SampleSize()
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf("SS: %v", sampleSize)
	sampleFlags := b.SampleFlags()
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf("SF: %v", sampleFlags)
	sampleCompositionTimeOffset := b.SampleCompositionTimeOffset()
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf("CO: %v", sampleCompositionTimeOffset)
	return ret
}
