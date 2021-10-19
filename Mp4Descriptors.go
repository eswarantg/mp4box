package mp4box

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"
)

/*
abstract aligned(8) expandable(228-1) class BaseDescriptor : bit(8) tag=0 { // empty. To be filled by classes extending this class.
}
*/

type BaseDescriptor struct {
	payload         *[]byte //Payload of the Descriptor
	payloadStOffset int
	size            int
	level           int
}

func (b *BaseDescriptor) leadString() string {
	var lead string
	for i := 0; i < b.level; i++ {
		lead += "\t"
	}
	return lead
}

func (b *BaseDescriptor) initData(payload *[]byte, level int) (int, error) {
	var err error
	if payload != nil {
		b.payload = payload
		b.level = level
		//consume everything
		b.size, b.payloadStOffset, err = b.descriptorSize()
		if b.size+b.payloadStOffset > len(*payload) {
			return len(*payload), fmt.Errorf("payload (%d) < Size (%d)", len(*payload), b.size+b.payloadStOffset)
		}
		x := (*b.payload)[0 : b.size+b.payloadStOffset]
		b.payload = &x
		return int(b.size + b.payloadStOffset), err
	}
	return 0, nil
}

//getPayload - Returns the payload excluding headers
func (b *BaseDescriptor) getPayload() []byte {
	var ret []byte
	if b.payload != nil {
		if len(*b.payload) > 2 {
			return (*b.payload)[b.payloadStOffset:]
		}
	}
	return ret
}

func (b *BaseDescriptor) DescriptorTag() uint8 {
	var ret uint8
	if len(*b.payload) > 1 {
		ret = (*b.payload)[0]
		return ret
	}
	//Improper box
	return ret
}

func (b *BaseDescriptor) DescriptorName() string {
	switch b.DescriptorTag() {
	case 0x03:
		return "ES_Descriptor"
	case 0x04:
		return "DecoderConfigDescrTag"
	case 0x05:
		return "DecSpecificInfoTag"
	case 0x06:
		return "SLConfigDescrTag"
	case 0x07:
		return "ContentIdentDescrTag"
	case 0x08:
		return "SupplContentIdentDescrTag"
	case 0x09:
		return "IPI_DescrPointerTag"
	case 0x0A:
		return "IPMP_DescrPointerTag"
	case 0x0B:
		return "IPMP_DescrTag"
	case 0x0C:
		return "QoS_DescrTag"
	case 0x0D:
		return "RegistrationDescrTag"
	case 0x0E:
		return "ES_ID_IncTag"
	case 0x0F:
		return "ES_ID_RefTag"
	case 0x10:
		return "MP4_IOD_Tag"
	case 0x11:
		return "MP4_OD_Tag"
	case 0x12:
		return "IPL_DescrPointerRefTag"
	case 0x13:
		return "ExtensionProfileLevelDescrTag"
	case 0x14:
		return "profileLevelIndicationIndexDescrTag"
	case 0x40:
		return "ContentClassificationDescrTag"
	case 0x41:
		return "KeyWordDescrTag"
	case 0x42:
		return "RatingDescrTag"
	case 0x43:
		return "LanguageDescrTag"
	case 0x44:
		return "ShortTextualDescrTag"
	case 0x45:
		return "ExpandedTextualDescrTag"
	case 0x46:
		return "ContentCreatorNameDescrTag"
	case 0x47:
		return "ContentCreationDateDescrTag"
	}
	return fmt.Sprintf("Unknown_%x", b.DescriptorTag())
}

func (b *BaseDescriptor) DescriptorSize() int {
	d, _, err := b.descriptorSize()
	if err != nil {
		return 0
	}
	return d
}

func (b *BaseDescriptor) descriptorSize() (int, int, error) {
	var ret int
	offset := 1
	for {
		if len(*b.payload) <= offset {
			err := fmt.Errorf("descriptor %s(%x) Size expandable not enough bytes %d %d",
				b.DescriptorName(), b.DescriptorTag(), offset, len(*b.payload))
			return ret, offset, err
		}
		val := (*b.payload)[offset]
		offset++
		ret = (ret << 7) | int(val&0x7F)
		if (val & 0x80) != 0x80 {
			break
		}
	}
	return ret, offset, nil
}

//String - Returns User Readable description of content
func (b *BaseDescriptor) String() string {
	var ret string
	hdump := hex.Dump(*b.payload)
	prefix := fmt.Sprintf("\n%d%v", b.level, b.leadString())
	ret += prefix + b.DescriptorName()
	ret += prefix + strings.ReplaceAll(hdump, "\n", prefix)
	ret += fmt.Sprintf("\n%d%sTag: %v(%v), Len: %v, P:%v, O:%v", b.level, b.leadString(), b.DescriptorName(), b.DescriptorTag(), b.DescriptorSize(), len(*b.payload), b.payloadStOffset)
	return ret
}

//DecoderSpecificInfo -
/*
abstract class DecoderSpecificInfo extends BaseDescriptor : bit(8) tag=DecSpecificInfoTag
{
   // empty. To be filled by classes extending this class.
}
*/
type DecoderSpecificInfo struct {
	BaseDescriptor
}

//String - Returns User Readable description of content
func (b *DecoderSpecificInfo) String() string {
	var ret string
	ret += b.BaseDescriptor.String()
	return ret
}

//DecoderConfigDescriptor -
/*
class DecoderConfigDescriptor extends BaseDescriptor : bit(8) tag=DecoderConfigDescrTag {
   bit(8) objectTypeIndication;
   bit(6) streamType;
   bit(1) upStream;
   const bit(1) reserved=1;
   bit(24) bufferSizeDB;
   bit(32) maxBitrate;
   bit(32) avgBitrate;
   DecoderSpecificInfo decSpecificInfo[0 .. 1];
   profileLevelIndicationIndexDescriptor profileLevelIndicationIndexDescr [0..255];
}
*/
type DecoderConfigDescriptor struct {
	BaseDescriptor
}

func (b *DecoderConfigDescriptor) DecoderSpecificInfos() ([]DecoderSpecificInfo, error) {
	var ret []DecoderSpecificInfo = make([]DecoderSpecificInfo, 0)
	p := b.BaseDescriptor.getPayload()
	if len(p) < 13 {
		return ret, nil
	}
	count := 0
	offset := 13
	for len(p) > offset && count < 2 {
		x := p[offset:]
		if x[0] != 0x05 {
			break
		}
		obj := DecoderSpecificInfo{}
		used, err := obj.initData(&x, b.level+1)
		if err != nil {
			return ret, err
		}
		fmt.Printf("\nDecoderSpecificInfo bytes left o:%v p:%v", offset+used, len(p))
		ret = append(ret, obj)
		offset += used
		count++
	}
	return ret, nil
}

func (b *DecoderConfigDescriptor) ProfileLevelIndicationIndexDescriptor() ([]BaseDescriptor, error) {
	var ret []BaseDescriptor = make([]BaseDescriptor, 0)
	p := b.BaseDescriptor.getPayload()
	if len(p) < 13 {
		return ret, nil
	}
	count := 0
	offset := 13
	for len(p) > offset && count < 2 {
		x := p[offset:]
		if x[0] != 0x05 {
			break
		}
		obj := DecoderSpecificInfo{}
		used, err := obj.initData(&x, b.level+1)
		fmt.Printf("\nDecoderSpecificInfo bytes left o:%v p:%v", offset+used, len(p))
		if err != nil {
			return ret, err
		}
		offset += used
		count++
	}
	for len(p) > offset {
		x := p[offset:]
		obj := BaseDescriptor{}
		used, err := obj.initData(&x, b.level+1)
		fmt.Printf("\nBaseDescriptor bytes left o:%v p:%v", offset+used, len(p))
		if err != nil {
			return ret, err
		}
		ret = append(ret, obj)
		offset += used
	}
	return ret, nil
}

//String - Returns User Readable description of content
func (b *DecoderConfigDescriptor) String() string {
	var ret string
	ret += b.BaseDescriptor.String()
	ret += fmt.Sprintf("\n%d%s", b.level, b.leadString())
	p := b.BaseDescriptor.getPayload()
	if len(p) >= 1 {
		ret += fmt.Sprintf(" objectTypeIndication:%v", uint8(p[0]))
		if len(p) >= 2 {
			ret += fmt.Sprintf(" ,streamType:%v, upStream:%v", uint8((p[1]&0xFC)>>2), uint8((p[1]&0x02)>>1))
			if len(p) >= 5 {
				x := append([]byte{0}, p[2:5]...)
				ret += fmt.Sprintf(" ,bufferSizeDB:%d", binary.BigEndian.Uint32(x))
				if len(p) >= 9 {
					ret += fmt.Sprintf(" ,maxBitrate:%d", binary.BigEndian.Uint32(p[5:9]))
					if len(p) >= 13 {
						ret += fmt.Sprintf(" ,avgBitrate:%d", binary.BigEndian.Uint32(p[9:13]))
						if len(p) > 13 {
							dsInfos, err := b.DecoderSpecificInfos()
							if err != nil {
								ret += fmt.Sprintf("\n%d%s", b.level, b.leadString())
								ret += fmt.Sprintf("\tDecoderSpecificInfos:%s", err.Error())
							}
							for _, dsInfo := range dsInfos {
								ret += dsInfo.String()
							}
							profLvlIndIdxDescs, err := b.ProfileLevelIndicationIndexDescriptor()
							if err != nil {
								ret += fmt.Sprintf("\n%d%s", b.level, b.leadString())
								ret += fmt.Sprintf("\tProfileLevelIndicationIndexDescriptor:%s", err.Error())
							}
							for _, profLvlIndIdxDesc := range profLvlIndIdxDescs {
								ret += profLvlIndIdxDesc.String()
							}
						}
					}
				}
			}
		}
	}
	return ret
}

/*
class ES_Descriptor extends BaseDescriptor : bit(8) tag=ES_DescrTag {
   bit(16) ES_ID;
   bit(1) streamDependenceFlag;
   bit(1) URL_Flag;
   bit(1) OCRstreamFlag;
   bit(5) streamPriority;
   if (streamDependenceFlag)
      bit(16) dependsOn_ES_ID;
   if (URL_Flag) {
      bit(8) URLlength;
      bit(8) URLstring[URLlength];
   }
   if (OCRstreamFlag)
      bit(16) OCR_ES_Id;
   DecoderConfigDescriptor decConfigDescr;
   if (ODProfileLevelIndication==0x01)
   {
      SLConfigDescriptor slConfigDescr;
   }
   else {
      SLConfigDescriptor slConfigDescr;
   }
   //no SL extension.
   // SL extension is possible.
   IPI_DescrPointer ipiPtr[0 .. 1];
   IP_IdentificationDataSet ipIDS[0 .. 255];
   IPMP_DescriptorPointer ipmpDescrPtr[0 .. 255];
   LanguageDescriptor langDescr[0 .. 255];
   QoS_Descriptor qosDescr[0 .. 1];
   RegistrationDescriptor regDescr[0 .. 1];
   ExtensionDescriptor extDescr[0 .. 255];
}
*/
type ES_Descriptor struct {
	BaseDescriptor
}

func (b *ES_Descriptor) ES_ID() uint16 {
	var ret uint16
	p := b.BaseDescriptor.getPayload()
	if len(p) >= 2 {
		ret = binary.BigEndian.Uint16(p[0:2])
		return ret
	}
	//Improper box
	return ret
}

func (b *ES_Descriptor) StreamDependenceFlag() bool {
	var ret bool
	p := b.BaseDescriptor.getPayload()
	if len(p) >= 2+1 {
		ret = (p[2] & 0x80) == 0x80
		return ret
	}
	//Improper box
	return ret
}

func (b *ES_Descriptor) URLFlag() bool {
	var ret bool
	p := b.BaseDescriptor.getPayload()
	if len(p) >= 2+1 {
		ret = (p[2] & 0x40) == 0x40
		return ret
	}
	//Improper box
	return ret
}

func (b *ES_Descriptor) OCRstreamFlag() bool {
	var ret bool
	p := b.BaseDescriptor.getPayload()
	if len(p) >= 2+1 {
		ret = (p[2] & 0x20) == 0x20
		return ret
	}
	//Improper box
	return ret
}

func (b *ES_Descriptor) StreamPriority() uint8 {
	var ret uint8
	p := b.BaseDescriptor.getPayload()
	if len(p) >= 2+1 {
		ret = (p[2] & 0x1F)
		return ret
	}
	//Improper box
	return ret
}

func (b *ES_Descriptor) DependsOnESID() *uint16 {
	var ret uint16
	if b.StreamDependenceFlag() {
		p := b.getPayload()
		if len(p) >= 2+1+2 {
			ret = binary.BigEndian.Uint16(p[3:5])
			return &ret
		}
	}
	//Improper box
	return nil
}

func (b *ES_Descriptor) URLstring() *string {
	var ret string
	if b.URLFlag() {
		p := b.BaseDescriptor.getPayload()
		offset := 2 + 1
		if b.StreamDependenceFlag() {
			offset += 2
		}
		if len(p) >= offset+1 {
			l := int(p[offset])
			offset++
			if l > 0 && len(p) >= offset+l {
				ret = string(p[offset : offset+l])
			}
			return &ret
		}
	}
	//Improper box
	return nil
}

func (b *ES_Descriptor) OCR_ES_Id() *uint16 {
	var ret uint16
	if b.OCRstreamFlag() {
		p := b.BaseDescriptor.getPayload()
		offset := 2 + 1
		if b.StreamDependenceFlag() {
			offset += 2
		}
		if b.StreamDependenceFlag() {
			l := int(p[offset])
			offset += 1 + l
		}
		if len(p) >= offset+2 {
			ret = binary.BigEndian.Uint16(p[offset : offset+2])
			return &ret
		}
	}
	//Improper box
	return nil
}

func (b *ES_Descriptor) SLConfigDescriptor() (*BaseDescriptor, error) {
	var err error
	var used int
	ret := BaseDescriptor{}
	p := b.BaseDescriptor.getPayload()
	offset := 2 + 1
	if b.StreamDependenceFlag() {
		offset += 2
	}
	if b.StreamDependenceFlag() {
		l := int(p[offset])
		offset += 1 + l
	}
	if b.OCRstreamFlag() {
		offset += 2
	}
	if len(p) > offset {
		x := p[offset:]
		decConfigDescr := DecoderConfigDescriptor{}
		used, err = decConfigDescr.initData(&x, b.level+1)
		fmt.Printf("\nDecoderConfigDescriptor bytes left o:%v p:%v", offset+used, len(p))
		if err != nil {
			return nil, err
		}
		offset += used
		if len(p) > offset {
			x := p[offset:]
			used, err = ret.initData(&x, b.level+1)
			fmt.Printf("\nSLConfigDescriptor bytes left o:%v p:%v", offset+used, len(p))
			if err == nil {
				return &ret, nil
			}
		}
	}
	return nil, err
}

func (b *ES_Descriptor) DecoderConfigDescriptor() (*DecoderConfigDescriptor, error) {
	var err error
	var used int
	ret := DecoderConfigDescriptor{}
	p := b.BaseDescriptor.getPayload()
	offset := 2 + 1
	if b.StreamDependenceFlag() {
		offset += 2
	}
	if b.StreamDependenceFlag() {
		l := int(p[offset])
		offset += 1 + l
	}
	if b.OCRstreamFlag() {
		offset += 2
	}
	if len(p) > offset {
		x := p[offset:]
		used, err = ret.initData(&x, b.level+1)
		fmt.Printf("\nDecoderConfigDescriptor bytes left o:%v p:%v", offset+used, len(p))
		if err == nil {
			return &ret, nil
		}
	}
	return nil, err
}

//String - Returns User Readable description of content
func (b *ES_Descriptor) String() string {
	var ret string
	ret += b.BaseDescriptor.String()
	ret += fmt.Sprintf("\n%d%s", b.level, b.leadString())
	ret += fmt.Sprintf(" ES_ID:%v ", b.ES_ID())
	if b.StreamDependenceFlag() {
		ret += fmt.Sprintf(" DependsOnESID:%v", b.DependsOnESID())
	}
	if b.URLFlag() {
		ret += fmt.Sprintf(" URL:%v", b.URLstring())
	}
	if b.OCRstreamFlag() {
		ret += fmt.Sprintf(" OCR_ES_Id:%v", b.URLstring())
	}
	dcDesc, err := b.DecoderConfigDescriptor()
	if err != nil {
		ret += fmt.Sprintf("\n%d%s", b.level, b.leadString())
		ret += fmt.Sprintf("DecoderConfigDescriptor unable to read : %s", err.Error())
	} else if dcDesc == nil {
		ret += fmt.Sprintf("\n%d%s", b.level, b.leadString())
		ret += "DecoderConfigDescriptor is nil"
	} else {
		ret += dcDesc.String()
	}
	slConfigDescr, err := b.SLConfigDescriptor()
	if err != nil {
		ret += fmt.Sprintf("\n%d%s", b.level, b.leadString())
		ret += fmt.Sprintf("SLConfigDescriptor unable to read : %s", err.Error())
	} else if slConfigDescr == nil {
		ret += fmt.Sprintf("\n%d%s", b.level, b.leadString())
		ret += "SLConfigDescriptor is nil"
	} else {
		ret += slConfigDescr.String()
	}
	return ret
}
