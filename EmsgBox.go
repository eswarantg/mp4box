package mp4box

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

//EmsgBox -
/*
aligned(8) class DASHEventMessageBox extends FullBox('emsg', version, flags = 0) {
    if (version==0) {
        string              scheme_id_uri;
        string              value;
        unsigned int(32)    timescale;
        unsigned int(32)    presentation_time_delta;
        unsigned int(32)    event_duration;
        unsigned int(32)    id;
    } else if (version==1) {
        unsigned int(32)    timescale;
        unsigned int(64)    presentation_time;
        unsigned int(32)    event_duration;
        unsigned int(32)    id;
        string              scheme_id_uri;
        string              value;
    }
    unsigned int(8) message_data[];
}
*/

type EmsgBox struct {
	FullBox
	parsed           bool
	SchemeIdUri      string
	Value            string
	TimeScale        uint32
	PresentTimeDelta uint32
	PresentTime      uint64
	EvtDur           uint32
	IdVal            uint32
	MsgPayload       []byte
}

// Interface methods Impl - Begin
// getLeafBox() returns leaf object Box interface
func (b *EmsgBox) getLeafBox() Box {
	return b
}

// GetEmsgBox - Implement Box method for this object
func (b *EmsgBox) GetEmsgBox() (*EmsgBox, error) {
	if b.parsed {
		err := b.parseAllData()
		if err != nil {
			return nil, err
		}
		b.parsed = true
	}
	return b, nil
}

//Interface methods Impl - End

// getPayload - Returns the payload excluding headers
func (b *EmsgBox) GetPayload() []byte {
	return b.MsgPayload
}

// GetTimeScale - Returns the timescale data from payload
func (b *EmsgBox) GetTimeScale() uint32 {
	return b.TimeScale
}

// GetSchemeInfo - Returns SchemeIdUri and value from payload
func (b *EmsgBox) GetSchemeInfo() (string, string) {
	return b.SchemeIdUri, b.Value
}

// GetEventDuration - Returns EventDuration bytes from payload
func (b *EmsgBox) GetEventDuration() uint32 {
	return b.EvtDur
}

func (b *EmsgBox) parseAllData() error {
	var err error
	payload := b.FullBox.getPayload()
	switch b.Version() {
	case 0:
		if len(payload) > 16 {
			startPayload := payload
			schemeIdUriIdx := bytes.IndexByte(startPayload, 0)
			b.SchemeIdUri = string(startPayload[0:schemeIdUriIdx])

			schemeIdValPayload := startPayload[schemeIdUriIdx+1:]
			schemeIdValIdx := bytes.IndexByte(schemeIdValPayload, 0)
			b.Value = string(schemeIdValPayload[0:schemeIdValIdx])

			otherPayload := schemeIdValPayload[schemeIdValIdx+1:]
			b.TimeScale = binary.BigEndian.Uint32(otherPayload[0:4])
			b.PresentTimeDelta = binary.BigEndian.Uint32(otherPayload[4:8])
			b.EvtDur = binary.BigEndian.Uint32(otherPayload[8:12])
			b.IdVal = binary.BigEndian.Uint32(otherPayload[12:16])
			b.MsgPayload = otherPayload[16:]
		} else {
			err = fmt.Errorf("eMsg Box payload is invalid")
			return err
		}
	case 1:
		if len(payload) > 20 {
			otherPayload := payload
			b.TimeScale = binary.BigEndian.Uint32(otherPayload[0:4])
			b.PresentTime = binary.BigEndian.Uint64(otherPayload[4:12])
			b.EvtDur = binary.BigEndian.Uint32(otherPayload[12:16])
			b.IdVal = binary.BigEndian.Uint32(otherPayload[16:20])

			startPayload := payload[20:]
			schemeIdUriIdx := bytes.IndexByte(startPayload, 0)
			b.SchemeIdUri = string(startPayload[0:schemeIdUriIdx])

			schemeIdValPayload := startPayload[schemeIdUriIdx+1:]
			schemeIdValIdx := bytes.IndexByte(schemeIdValPayload, 0)
			b.Value = string(schemeIdValPayload[0:schemeIdValIdx])
			b.MsgPayload = schemeIdValPayload[schemeIdValIdx:]
		} else {
			err = fmt.Errorf("eMsg Box payload is invalid")
			return err
		}
	default:
		err = fmt.Errorf("eMsg Box version invalid")
		return err
	}
	return nil
}

// String - Returns User Readable description of content
func (b *EmsgBox) String() string {
	var ret string
	ret += b.FullBox.String()
	ret += fmt.Sprintf("\n%d%v ", b.FullBox.level, b.FullBox.leadString())
	ret += fmt.Sprintf(" SchemeIdUri: %v, Value: %v, TimeScale: %v, PresentationTimeDelta:%v, PresentationTime:%v, EventDur:%v, IdVal:%v",
		b.SchemeIdUri, b.Value, b.TimeScale,
		b.PresentTimeDelta, b.PresentTime, b.EvtDur, b.IdVal)
	ret += fmt.Sprintf("\n%d%v ", b.FullBox.level, b.FullBox.leadString())
	hdump := hex.Dump(b.MsgPayload)
	ret += fmt.Sprintf("%v", hdump)
	return ret
}
