package mp4box

import "fmt"

//DataEntryURLBox -
/*
aligned(8) class DataEntryUrlBox (bit(24) flags) extends FullBox(‘url ’, version = 0, flags) {
	 string location;
}
*/
type DataEntryURLBox struct {
	FullBox
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *DataEntryURLBox) getLeafBox() Box {
	return b
}

//GetDataEntryURLBox - Implement Box method for this object
func (b *DataEntryURLBox) GetDataEntryURLBox() (*DataEntryURLBox, error) {
	return b, nil
}

//Location - Returns location url
func (b *DataEntryURLBox) Location() string {
	var ret string
	p := b.FullBox.getPayload()
	if p != nil {
		if len(p) >= 2 {
			ret, _ = getNullTermString(p)
		}
	}
	return ret
}

//String - Display
func (b *DataEntryURLBox) String() string {
	var ret string
	ret += b.FullBox.String()
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf(" Location:%v ", b.Location())
	return ret
}

//Interface methods Impl - End

//DataEntryUrnBox -
/*
aligned(8) class DataEntryUrnBox (bit(24) flags) extends FullBox(‘urn ’, version = 0, flags) {
	string name;
	string location;
}
*/
type DataEntryUrnBox struct {
	FullBox
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *DataEntryUrnBox) getLeafBox() Box {
	return b
}

//GetDataEntryUrnBox - Implement Box method for this object
func (b *DataEntryUrnBox) GetDataEntryUrnBox() (*DataEntryUrnBox, error) {
	return b, nil
}

//Name - Returns Name url
func (b *DataEntryUrnBox) Name() string {
	var ret string
	p := b.FullBox.getPayload()
	if p != nil {
		if len(p) >= 2 {
			ret, _ = getNullTermString(p)
		}
	}
	return ret
}

//Location - Returns Location url
func (b *DataEntryUrnBox) Location() string {
	var ret string
	var i int
	p := b.FullBox.getPayload()
	if p != nil {
		if len(p) >= 2 {
			_, i = getNullTermString(p)
		}
		if i+2 < len(p) {
			ret, i = getNullTermString(p[i+1:])
		}
	}
	return ret
}

//Interface methods Impl - End

//String - Display
func (b *DataEntryUrnBox) String() string {
	var ret string
	ret += b.FullBox.String()
	ret += fmt.Sprintf("\n%d%v ", b.level, b.leadString())
	ret += fmt.Sprintf(" Name:%v Location:%v ", b.Name(), b.Location())
	return ret
}
