package mp4box

import (
	"bytes"
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
	SchemeIdUri      string
	SchemeIdVal      string
	MsgData          string
	TimeScale        []byte
	PresentTimeDelta []byte
	EvtDur           []byte
	IdVal            []byte
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *EmsgBox) getLeafBox() Box {
	return b
}

//GetEmsgBox - Implement Box method for this object
func (b *EmsgBox) GetEmsgBox() (*EmsgBox, error) {
	return b, nil
}

//Interface methods Impl - End

//GetPayload - Returns the payload excluding headers
func (b *EmsgBox) GetPayload() []byte {
	ret := b.FullBox.getPayload()
	return ret
}

//GetMsgData - Returns the message data from payload
func (b *EmsgBox) GetMsgData() string {
	return b.MsgData
}

//GetTimeScale - Returns the timescale data from payload
func (b *EmsgBox) GetTimeScale() []byte {
	return b.TimeScale
}

//GetSchemeInfo - Returns SchemeIdUri and value from payload
func (b *EmsgBox) GetSchemeInfo() (string, string) {
	return b.SchemeIdUri, b.SchemeIdVal
}

//GetEventDuration - Returns EventDuration bytes from payload
func (b *EmsgBox) GetEventDuration() []byte {
	return b.EvtDur
}

func (b *EmsgBox) ParseAllData() error {
	var err error
	payload := b.FullBox.getPayload()
	if b.Version() > 1 || b.Version() < 0 {
		err = fmt.Errorf("eMsg Box version invalid")
		return err
	}
	if b.Version() == 0 {
		if len(payload) > 16 {
			startPayload := payload
			schemeIdUriIdx := bytes.IndexByte(startPayload, 0)
			b.SchemeIdUri = string(startPayload[0:schemeIdUriIdx])

			schemeIdValPayload := startPayload[schemeIdUriIdx+1:]
			schemeIdValIdx := bytes.IndexByte(schemeIdValPayload, 0)
			b.SchemeIdVal = string(schemeIdValPayload[0:schemeIdValIdx])

			otherPayload := schemeIdValPayload[schemeIdValIdx+1:]
			b.TimeScale = otherPayload[0:4]
			b.PresentTimeDelta = otherPayload[4:8]
			b.EvtDur = otherPayload[8:12]
			b.IdVal = otherPayload[12:16]
			b.MsgData = string(otherPayload[16:])
		} else {
			err = fmt.Errorf("eMsg Box payload is invalid")
			return err
		}
	}
	if b.Version() == 1 {
		if len(payload) > 20 {
			otherPayload := payload
			b.TimeScale = otherPayload[0:4]
			b.PresentTimeDelta = otherPayload[4:12]
			b.EvtDur = otherPayload[12:16]
			b.IdVal = otherPayload[16:20]

			startPayload := payload[20:]
			schemeIdUriIdx := bytes.IndexByte(startPayload, 0)
			b.SchemeIdUri = string(startPayload[0:schemeIdUriIdx])

			schemeIdValPayload := startPayload[schemeIdUriIdx+1:]
			schemeIdValIdx := bytes.IndexByte(schemeIdValPayload, 0)
			b.SchemeIdVal = string(schemeIdValPayload[0:schemeIdValIdx])
		} else {
			err = fmt.Errorf("eMsg Box payload is invalid")
			return err
		}
	}
	return nil
}

//String - Returns User Readable description of content
func (b *EmsgBox) String() string {
	var ret string
	ret += b.FullBox.String()
	ret += fmt.Sprintf("\n%d%v ", b.FullBox.level, b.FullBox.leadString())
	ret += fmt.Sprintf(" Version: %v, Flags: %v", b.FullBox.Version(), b.FullBox.Flags())
	return ret
}
