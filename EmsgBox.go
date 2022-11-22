package mp4box

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//FullBox - Base box holding the bytes
/*
aligned(8) class FullBox(unsigned int(32) boxtype, unsigned int(8) v, bit(24) f) extends Box(boxtype) {
	unsigned int(8) version = v;
	bit(24) flags = f;
}
*/
type EmsgBox struct {
	BaseBox
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
	b.SchemeIdUri, b.SchemeIdVal, b.MsgData, b.TimeScale, b.PresentTimeDelta, b.EvtDur, b.IdVal = b.GetAllData()
	return b, nil
}

//Interface methods Impl - End

//GetPayload - Returns the payload excluding headers
func (b *EmsgBox) GetPayload() []byte {
	ret := b.BaseBox.getPayload()
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

func (b *EmsgBox) GetAllData() (schemeIdUri string, schemeIdVal string, msgData string,
	timeScale []byte, presentTimeDelta []byte, evtDur []byte, idVal []byte) {
	payload := b.BaseBox.getPayload()
	if b.Version() == 0 && len(payload) > 16 {
		startPayload := payload[4:]
		schemeIdUriIdx := bytes.IndexByte(startPayload, 0)
		schemeIdUri = string(startPayload[0:schemeIdUriIdx])

		schemeIdValPayload := startPayload[schemeIdUriIdx+1:]
		schemeIdValIdx := bytes.IndexByte(schemeIdValPayload, 0)
		schemeIdVal = string(schemeIdValPayload[0:schemeIdValIdx])

		otherPayload := schemeIdValPayload[schemeIdValIdx+1:]
		timeScale = otherPayload[0:4]
		presentTimeDelta = otherPayload[4:8]
		evtDur = otherPayload[8:12]
		idVal = otherPayload[12:16]
		msgData = string(otherPayload[16:])
	}
	if b.Version() == 1 && len(payload) > 20 {
		otherPayload := payload[4:]
		timeScale = otherPayload[0:4]
		presentTimeDelta = otherPayload[4:12]
		evtDur = otherPayload[12:16]
		idVal = otherPayload[16:20]

		startPayload := payload[20:]
		schemeIdUriIdx := bytes.IndexByte(startPayload, 0)
		schemeIdUri = string(startPayload[0:schemeIdUriIdx])

		schemeIdValPayload := startPayload[schemeIdUriIdx+1:]
		schemeIdValIdx := bytes.IndexByte(schemeIdValPayload, 0)
		schemeIdVal = string(schemeIdValPayload[0:schemeIdValIdx])

	}
	return
}

//Version - returns the version of the box
func (b *EmsgBox) Version() int8 {
	var ret int8
	p := b.BaseBox.getPayload()
	if len(p) >= 1 {
		return int8(p[0])
	}
	//Improper Box
	return ret
}

//Flags - Returns flags
// bit(24) - fit to lower 24 bits of uint32
func (b *EmsgBox) Flags() uint32 {
	var ret uint32
	p := b.BaseBox.getPayload()
	if len(p) >= 4 {
		buf := []byte{p[1], p[2], p[3], 0}
		ret = binary.BigEndian.Uint32(buf)
	}
	//Improper Box
	return ret
}

//String - Returns User Readable description of content
func (b *EmsgBox) String() string {
	var ret string
	ret += b.BaseBox.String()
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf(" Version: %v, Flags: %v", b.Version(), b.Flags())
	return ret
}
