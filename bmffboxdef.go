package mp4box

import (
	"encoding/binary"
	"fmt"
	"log"
	"time"
)

var epochTimeMp4 time.Time

func init() {
	var err error
	epochTimeMp4, err = time.Parse(time.RFC3339, "1904-01-01T00:00:00Z")
	if err != nil {
		panic("ERROR in time format")
	}
}

//FileBox - Base box holding the bytes
/*
aligned(8) class FileTypeBox
extends Box(‘ftyp’) {
unsigned int(32) major_brand; unsigned int(32) minor_version; unsigned int(32) compatible_brands[];
}
*/
type FileBox struct {
	BaseBox
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *FileBox) getLeafBox() AccessBoxType {
	return b
}

//GetFileBox - Implement AccessBoxType method for this object
func (b *FileBox) GetFileBox() (*FileBox, error) {
	return b, nil
}

//Interface methods Impl - End

//MajorBrand - returns the major brand of the file
func (b *FileBox) MajorBrand() string {
	p := b.BaseBox.getPayload()
	if p != nil && len(p) >= 4 {
		ret := string(p[0:4])
		return ret
	}
	return ""
}

//MinorVersion - returns the minor version of the file
func (b *FileBox) MinorVersion() uint32 {
	p := b.BaseBox.getPayload()
	if p != nil && len(p) >= 8 {
		ret := binary.BigEndian.Uint32(p[4:8])
		return ret
	}
	return 0
}

//CompatibleBrands - returns the Compatible brands
func (b *FileBox) CompatibleBrands() []string {
	var ret []string
	p := b.BaseBox.getPayload()
	if p != nil && len(p) >= 12 {
		nEntries := (len(p) - 12) / 4
		ret = make([]string, nEntries)
		bytesRead := 8
		for i := 0; i < nEntries; i++ {
			ret[i] = string(p[bytesRead : bytesRead+4])
			bytesRead += 4
		}
	}
	return ret
}

func (b *FileBox) String() string {
	var ret string
	ret += b.BaseBox.String()
	ret += fmt.Sprintf("\n%v MajorBrand: %v, MinorVersion: %v, CompatibleBrands:%v ", b.leadString(), b.MajorBrand(), b.MinorVersion(), b.CompatibleBrands())
	return ret
}

//MediaDataBox -
/*
aligned(8) class MediaDataBox extends Box(‘mdat’) { bit(8) data[];
}
*/
type MediaDataBox struct {
	BaseBox
}

//MovieHeaderBox -
/*
aligned(8)
class MovieHeaderBox extends FullBox(‘mvhd’, version, 0)
{
	if (version==1) {
		unsigned int(64)  creation_time; //(in seconds since midnight, Jan. 1, 1904, in UTC time)
		unsigned int(64)  modification_time; //(in seconds since midnight, Jan. 1, 1904, in UTC time)
		unsigned int(32)  timescale;
		unsigned int(64)  duration;
 	} else { // version==0
		unsigned int(32)  creation_time; //(in seconds since midnight, Jan. 1, 1904, in UTC time)
		unsigned int(32)  modification_time; //(in seconds since midnight, Jan. 1, 1904, in UTC time)
		unsigned int(32)  timescale;
		unsigned int(32)  duration;
	}
	template int(32) rate = 0x00010000; // typically 1.0
	template int(16) volume = 0x0100; // typically, full volume const bit(16) reserved = 0;
	const unsigned int(32)[2] reserved = 0;
	template int(32)[9] matrix =
	{ 0x00010000,0,0,0,0x00010000,0,0,0,0x40000000 };
		// Unity matrix
	bit(32)[6]  pre_defined = 0;
	unsigned int(32)  next_track_ID;
}*/
type MovieHeaderBox struct {
	FullBox
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *MovieHeaderBox) getLeafBox() AccessBoxType {
	return b
}

//GetMovieHeaderBox - Implement AccessBoxType method for this object
func (b *MovieHeaderBox) GetMovieHeaderBox() (*MovieHeaderBox, error) {
	return b, nil
}

//Interface methods Impl - End

//CreationTime - CreationTime of the content
func (b *MovieHeaderBox) CreationTime() time.Time {
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		secs := binary.BigEndian.Uint32(p[0:4])
		t := epochTimeMp4
		return t.Add(time.Duration(secs))
	case 1:
		secs := binary.BigEndian.Uint64(p[0:8])
		t := epochTimeMp4
		return t.Add(time.Duration(secs))
	}
	return time.Time{}
}

//ModificationTime - ModificationTime of the content
func (b *MovieHeaderBox) ModificationTime() time.Time {
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		secs := binary.BigEndian.Uint32(p[4:8])
		t := epochTimeMp4
		log.Printf("Secs %v ", secs)
		return t.Add(time.Duration(secs))
	case 1:
		secs := binary.BigEndian.Uint64(p[8:16])
		t := epochTimeMp4
		log.Printf("Secs %v ", secs)
		return t.Add(time.Duration(secs))
	}
	return time.Time{}
}

//Scale - Ticks per second for all Timing info
func (b *MovieHeaderBox) Scale() uint32 {
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		scale := binary.BigEndian.Uint32(p[8:12])
		return scale
	case 1:
		scale := binary.BigEndian.Uint32(p[16:20])
		return scale
	}
	return 0
}

//Duration - Duration of the content
func (b *MovieHeaderBox) Duration() time.Duration {
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		scale := binary.BigEndian.Uint32(p[8:12])
		dur := binary.BigEndian.Uint32(p[12:16])
		log.Printf("Scale %v Dur %v", scale, dur)
		if scale != 0 {
			return time.Duration(dur / scale)
		}
	case 1:
		scale := binary.BigEndian.Uint32(p[16:20])
		dur := binary.BigEndian.Uint64(p[20:28])
		log.Printf("Scale %v Dur %v", scale, dur)
		if scale != 0 {
			return time.Duration(dur / uint64(scale))
		}
	}
	return 0 * time.Second
}

//String - Display
func (b *MovieHeaderBox) String() string {
	var ret string
	log.Printf("MovieHeaderBox String Check1")
	ret += b.FullBox.String()
	log.Printf("MovieHeaderBox String Check2")
	ret += fmt.Sprintf("\n%v Creation:%v Modification:%v Duration:%v", b.leadString(), b.CreationTime(), b.ModificationTime(), b.Duration())
	log.Printf("MovieHeaderBox String Check3")
	log.Print(ret)
	return ret
}

//TrackHeaderBox -
/*
aligned(8)
class TrackHeaderBox extends FullBox(‘tkhd’, version, flags)
{ if (version==1) {
      unsigned int(64)  creation_time;
      unsigned int(64)  modification_time;
      unsigned int(32)  track_ID;
      const unsigned int(32)  reserved = 0;
      unsigned int(64)  duration;
   } else { // version==0
      unsigned int(32)  creation_time;
      unsigned int(32)  modification_time;
      unsigned int(32)  track_ID;
      const unsigned int(32)  reserved = 0;
      unsigned int(32)  duration;
}
const unsigned int(32)[2] reserved = 0;
template int(16) layer = 0;
template int(16) alternate_group = 0;
template int(16) volume = {if track_is_audio 0x0100 else 0}; const unsigned int(16) reserved = 0;
template int(32)[9] matrix=
{ 0x00010000,0,0,0,0x00010000,0,0,0,0x40000000 };
      // unity matrix
   unsigned int(32) width;
   unsigned int(32) height;
}
*/
type TrackHeaderBox struct {
	FullBox
}

//Interface methods Impl - Begin
//getLeafBox() returns leaf object Box interface
func (b *TrackHeaderBox) getLeafBox() AccessBoxType {
	return b
}

//GetMTrackHeaderBox - Implement AccessBoxType method for this object
func (b *TrackHeaderBox) GetMTrackHeaderBox() (*TrackHeaderBox, error) {
	return b, nil
}

//Interface methods Impl - End

//CreationTime - CreationTime of the content
func (b *TrackHeaderBox) CreationTime() time.Time {
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		secs := binary.BigEndian.Uint32(p[0:4])
		t := epochTimeMp4
		return t.Add(time.Duration(secs))
	case 1:
		secs := binary.BigEndian.Uint64(p[0:8])
		t := epochTimeMp4
		return t.Add(time.Duration(secs))
	}
	return time.Time{}
}

//ModificationTime - ModificationTime of the content
func (b *TrackHeaderBox) ModificationTime() time.Time {
	p := b.FullBox.getPayload()
	switch b.FullBox.Version() {
	case 0:
		secs := binary.BigEndian.Uint32(p[4:8])
		t := epochTimeMp4
		log.Printf("Secs %v ", secs)
		return t.Add(time.Duration(secs))
	case 1:
		secs := binary.BigEndian.Uint64(p[8:16])
		t := epochTimeMp4
		log.Printf("Secs %v ", secs)
		return t.Add(time.Duration(secs))
	}
	return time.Time{}
}

//Duration - Duration of the content
func (b *TrackHeaderBox) Duration() time.Duration {
	p := b.FullBox.getPayload()
	node, err := b.GetParentByName("moov")
	if err != nil {
		return 0 * time.Second
	}
	moofBox, err := node.GetCollectionBaseBox()
	if err != nil {
		return 0 * time.Second
	}
	node, err = moofBox.GetChildByName("mvhd")
	if err != nil {
		return 0 * time.Second
	}
	mvhdBox, err := node.GetMovieHeaderBox()
	if err != nil {
		return 0 * time.Second
	}
	scale := mvhdBox.Scale()
	if err != nil {
		return 0 * time.Second
	}
	switch b.FullBox.Version() {
	case 0:
		dur := binary.BigEndian.Uint32(p[16:20])
		log.Printf("Scale %v Dur %v", scale, dur)
		if scale != 0 {
			return time.Duration(dur / scale)
		}
	case 1:
		dur := binary.BigEndian.Uint64(p[24:32])
		log.Printf("Scale %v Dur %v", scale, dur)
		if scale != 0 {
			return time.Duration(dur / uint64(scale))
		}
	}
	return 0 * time.Second
}

//String - Display
func (b *TrackHeaderBox) String() string {
	var ret string
	ret += b.FullBox.String()
	ret += fmt.Sprintf("\n%v Creation:%v Modification:%v Duration:%v", b.leadString(), b.CreationTime(), b.ModificationTime(), b.Duration())
	return ret
}

//MediaHeaderBox -
/*
aligned(8) class MediaHeaderBox extends FullBox(‘mdhd’, version, 0) { if (version==1) {
      unsigned int(64)  creation_time;
      unsigned int(64)  modification_time;
      unsigned int(32)  timescale;
      unsigned int(64)  duration;
   } else { // version==0
      unsigned int(32)  creation_time;
      unsigned int(32)  modification_time;
      unsigned int(32)  timescale;
      unsigned int(32)  duration;
}
bit(1) pad=0;
unsigned int(5)[3] language; // ISO-639-2/T language code unsigned int(16) pre_defined = 0;
}
*/
type MediaHeaderBox struct {
	FullBox
}

//HandlerBox -
/*
aligned(8) class HandlerBox extends FullBox(‘hdlr’, version = 0, 0) { unsigned int(32) pre_defined = 0;
unsigned int(32) handler_type;
const unsigned int(32)[3] reserved = 0;
   string   name;
}
*/
type HandlerBox struct {
	FullBox
}

//VideoMediaHeaderBox -
/*
aligned(8) class VideoMediaHeaderBox
extends FullBox(‘vmhd’, version = 0, 1) {
template unsigned int(16) graphicsmode = 0; // copy, see below template unsigned int(16)[3] opcolor = {0, 0, 0};
}
*/
type VideoMediaHeaderBox struct {
	FullBox
}

//SoundMediaHeaderBox -
/*
aligned(8) class SoundMediaHeaderBox
   extends FullBox(‘smhd’, version = 0, 0) {
   template int(16) balance = 0;
   const unsigned int(16)  reserved = 0;
}
*/
type SoundMediaHeaderBox struct {
	FullBox
}

//HintMediaHeaderBox -
/*
aligned(8) class HintMediaHeaderBox
   extends FullBox(‘hmhd’, version = 0, 0) {
   unsigned int(16)  maxPDUsize;
   unsigned int(16)  avgPDUsize;
   unsigned int(32)  maxbitrate;
   unsigned int(32)  avgbitrate;
   unsigned int(32)  reserved = 0;
}
*/
type HintMediaHeaderBox struct {
	FullBox
}

//NullMediaHeaderBox -
/*
aligned(8) class NullMediaHeaderBox
extends FullBox(’nmhd’, version = 0, flags) {
}
*/
type NullMediaHeaderBox struct {
	FullBox
}

//DataEntryURLBox -
/*
aligned(8) class DataEntryUrlBox (bit(24) flags) extends FullBox(‘url ’, version = 0, flags) { string location;
}
*/
type DataEntryURLBox struct {
	FullBox
}

//DataEntryUrnBox -
/*
aligned(8) class DataEntryUrnBox (bit(24) flags) extends FullBox(‘urn ’, version = 0, flags) { string name;
string location;
}
*/
type DataEntryUrnBox struct {
	FullBox
}

//TimeToSampleBox -
/*
aligned(8) class TimeToSampleBox
   extends FullBox(’stts’, version = 0, 0) {
   unsigned int(32)  entry_count;
      int i;
   for (i=0; i < entry_count; i++) {
      unsigned int(32)  sample_count;
      unsigned int(32)  sample_delta;
   }
}
*/
type TimeToSampleBox struct {
	FullBox
}

//CompositionOffsetBox -
/*
aligned(8) class CompositionOffsetBox extends FullBox(‘ctts’, version = 0, 0) { unsigned int(32) entry_count;
      int i;
   for (i=0; i < entry_count; i++) {
      unsigned int(32)  sample_count;
      unsigned int(32)  sample_offset;
   }
}
*/
type CompositionOffsetBox struct {
	FullBox
}

//SampleDescriptionBox -
/*
aligned(8) class SampleDescriptionBox (unsigned int(32) handler_type) extends FullBox('stsd', 0, 0){
	int i ;
	unsigned int(32) entry_count;
	for (i = 1 ; i u entry_count ; i++){
		switch (handler_type){
			case ‘soun’: // for audio tracks
			AudioSampleEntry();
			break;
		case ‘vide’: // for video tracks
			VisualSampleEntry();
			break;
		case ‘hint’: // Hint track
			HintSampleEntry();
			break;
		}
	}
}
*/
type SampleDescriptionBox struct {
	FullBox
}

//SampleSizeBox -
/*
aligned(8) class SampleSizeBox extends FullBox(‘stsz’, version = 0, 0) { unsigned int(32) sample_size;
	unsigned int(32) sample_count;
	if (sample_size==0) {
	for (i=1; i u sample_count; i++) {
		  unsigned int(32)  entry_size;
	} }
	}
*/
type SampleSizeBox struct {
	FullBox
}

//CompactSampleSizeBox -
/*
aligned(8) class CompactSampleSizeBox extends FullBox(‘stz2’, version = 0, 0) { unsigned int(24) reserved = 0;
unisgned int(8) field_size;
unsigned int(32) sample_count;
for (i=1; i u sample_count; i++) { unsigned int(field_size) entry_size;
} }
*/
type CompactSampleSizeBox struct {
	FullBox
}

//SampleToChunkBox -
/*
aligned(8) class SampleToChunkBox
extends FullBox(‘stsc’, version = 0, 0) { unsigned int(32) entry_count;
for (i=1; i u entry_count; i++) {
unsigned int(32) first_chunk;
unsigned int(32) samples_per_chunk; unsigned int(32) sample_description_index;
} }
*/
type SampleToChunkBox struct {
	FullBox
}

//ChunkOffsetBox -
/*
aligned(8) class ChunkOffsetBox
extends FullBox(‘stco’, version = 0, 0) { unsigned int(32) entry_count;
for (i=1; i u entry_count; i++) {
      unsigned int(32)  chunk_offset;
   }
}
*/
type ChunkOffsetBox struct {
	FullBox
}

//ChunkLargeOffsetBox -
/*
aligned(8) class ChunkLargeOffsetBox
extends FullBox(‘co64’, version = 0, 0) { unsigned int(32) entry_count;
for (i=1; i u entry_count; i++) {
      unsigned int(64)  chunk_offset;
   }
}
*/
type ChunkLargeOffsetBox struct {
	FullBox
}

//SyncSampleBox -
/*
aligned(8) class SyncSampleBox
   extends FullBox(‘stss’, version = 0, 0) {
   unsigned int(32)  entry_count;
   int i;
   for (i=0; i < entry_count; i++) {
      unsigned int(32)  sample_number;
   }
}
*/
type SyncSampleBox struct {
	FullBox
}

//ShadowSyncSampleBox -
/*
aligned(8) class ShadowSyncSampleBox
   extends FullBox(‘stsh’, version = 0, 0) {
   unsigned int(32)  entry_count;
   int i;
   for (i=0; i < entry_count; i++) {
unsigned int(32) shadowed_sample_number;
      unsigned int(32)  sync_sample_number;
   }
}
*/
type ShadowSyncSampleBox struct {
	FullBox
}

//DegradationPriorityBox -
/*
aligned(8) class DegradationPriorityBox extends FullBox(‘stdp’, version = 0, 0) { int i;
	for (i=0; i < sample_count; i++) {
		  unsigned int(16)  priority;
	   }
	}
*/
type DegradationPriorityBox struct {
	FullBox
}

//PaddingBitsBox -
/*
aligned(8) class PaddingBitsBox extends FullBox(‘padb’, version = 0, 0) { unsigned int(32) sample_count;
	int i;
	for (i=0; i < ((sample_count + 1)/2); i++) {
		  bit(1)   reserved = 0;
		  bit(3)   pad1;
		  bit(1)   reserved = 0;
		  bit(3)   pad2;
	} }
*/
type PaddingBitsBox struct {
	FullBox
}

//FreeSpaceBox -
/*
free_type may be ‘free’ or ‘skip’.
aligned(8) class FreeSpaceBox extends Box(free_type) { unsigned int(8) data[];
}
*/
type FreeSpaceBox struct {
	BaseBox
}

//EditListBox -
/*
aligned(8) class EditListBox extends FullBox(‘elst’, version, 0) { unsigned int(32) entry_count;
for (i=1; i <= entry_count; i++) {
} }
if (version==1) {
   unsigned int(64) segment_duration;
   int(64) media_time;
} else { // version==0
   unsigned int(32) segment_duration;
   int(32)  media_time;
}
int(16) media_rate_integer;
int(16) media_rate_fraction = 0;
*/
type EditListBox struct {
	FullBox
}

//UserDataBox -
/*
aligned(8) class UserDataBox extends Box(‘udta’) { }
*/
type UserDataBox struct {
	BaseBox
}

//CopyrightBox -
/*
aligned(8) class CopyrightBox
extends FullBox(‘cprt’, version = 0, 0) {
const bit(1) pad = 0;
unsigned int(5)[3] language; // ISO-639-2/T language code string notice;
}
*/
type CopyrightBox struct {
	FullBox
}

//MovieExtendsHeaderBox -
/*
aligned(8) class MovieExtendsHeaderBox extends FullBox(‘mehd’, version, 0) { if (version==1) {
	unsigned int(64)  fragment_duration;
 } else { // version==0
	unsigned int(32)  fragment_duration;
 }
}
*/
type MovieExtendsHeaderBox struct {
	FullBox
}

//TrackExtendsBox -
/*
aligned(8) class TrackExtendsBox extends FullBox(‘trex’, 0, 0){
	unsigned int(32) track_ID;
	unsigned int(32) default_sample_description_index;
	unsigned int(32) default_sample_duration;
	unsigned int(32) default_sample_size;
	unsigned int(32) default_sample_flags
}
*/
type TrackExtendsBox struct {
	FullBox
}

//MovieFragmentHeaderBox -
/*
aligned(8) class MovieFragmentHeaderBox extends FullBox(‘mfhd’, 0, 0){
	unsigned int(32)  sequence_number;
 }
*/
type MovieFragmentHeaderBox struct {
	FullBox
}

//TrackFragmentHeaderBox -
/*
aligned(8) class TrackFragmentHeaderBox extends FullBox(‘tfhd’, 0, tf_flags){
unsigned int(32) track_ID;
// all the following are optional fields unsigned int(64) base_data_offset; unsigned int(32) sample_description_index; unsigned int(32) default_sample_duration; unsigned int(32) default_sample_size; unsigned int(32) default_sample_flags
}
*/
type TrackFragmentHeaderBox struct {
	FullBox
}

//TrackRunBox -
/*
aligned(8) class TrackRunBox
         extends FullBox(‘trun’, 0, tr_flags) {
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

//TrackFragmentRandomAccessBox -
/*
aligned(8)
class TrackFragmentRandomAccessBox extends FullBox(‘tfra’, version, 0) {
	unsigned int(32)  track_ID;
	const unsigned int(26)  reserved = 0;
	unsigned int(2) length_size_of_traf_num;
	unsigned int(2) length_size_of_trun_num;
	unsigned int(2) length_size_of_sample_num;
	unsigned int(32)  number_of_entry;
	for(i=1; i • number_of_entry; i++){
		if(version==1){
			unsigned int(64)  time;
			unsigned int(64)  moof_offset;
		}else{
			unsigned int(32)  time;
			unsigned int(32)  moof_offset;
		}
	}
	unsignedint((length_size_of_traf_num+1)*8) traf_number;
	unsignedint((length_size_of_trun_num+1)*8) trun_number;
	unsigned int((length_size_of_sample_num+1) * 8)sample_number;
}
*/
type TrackFragmentRandomAccessBox struct {
	FullBox
}

//MovieFragmentRandomAccessOffsetBox -
/*
aligned(8) class MovieFragmentRandomAccessOffsetBox extends FullBox(‘mfro’, version, 0) {
   unsigned int(32)  size;
}
*/
type MovieFragmentRandomAccessOffsetBox struct {
	FullBox
}

//SampleDependencyTypeBox -
/*
aligned(8) class SampleDependencyTypeBox extends FullBox(‘sdtp’, version = 0, 0) { for (i=0; i < sample_count; i++){
unsigned int(2) reserved = 0;
unsigned int(2) sample_depends_on; unsigned int(2) sample_is_depended_on; unsigned int(2) sample_has_redundancy;
} }
*/
type SampleDependencyTypeBox struct {
	FullBox
}

//SampleToGroupBox -
/*
aligned(8) class SampleToGroupBox
   extends FullBox(‘sbgp’, version = 0, 0)
{
   unsigned int(32)  grouping_type;
   unsigned int(32)  entry_count;
   for (i=1; i <= entry_count; i++)
   {
      unsigned int(32)  sample_count;
unsigned int(32) group_description_index; }
}
*/
type SampleToGroupBox struct {
	FullBox
}

//SampleGroupDescriptionBox -
/*
aligned(8) class SampleGroupDescriptionBox (unsigned int(32) handler_type) extends FullBox('sgpd', 0, 0){
unsigned int(32) grouping_type;
unsigned int(32) entry_count;
   int i;
   for (i = 1 ; i <= entry_count ; i++){
      switch (handler_type){
         case ‘vide’: // for video tracks
} }
}
   VisualSampleGroupEntry ();
   break;
case ‘soun’: // for audio tracks
   AudioSampleGroupEntry();
   break;
case ‘hint’: // for hint tracks
   HintSampleGroupEntry();
   break;
*/
type SampleGroupDescriptionBox struct {
	FullBox
}

//SampleScaleBox -
/*
aligned(8) class SampleScaleBox extends FullBox(‘stsl’, version = 0, 0) { bit(7) reserved = 0;
bit(1) constraint_flag;
unsigned int(8) scale_method;
   int(16) display_center_x;
   int(16) display_center_y;
}
*/
type SampleScaleBox struct {
	FullBox
}

//SubSampleInformationBox -
/*
aligned(8) class SubSampleInformationBox extends FullBox(‘subs’, version, 0) { unsigned int(32) entry_count;
int i,j;
   for (i=0; i < entry_count; i++) {
      unsigned int(32) sample_delta;
      unsigned int(16) subsample_count;
      if (subsample_count > 0) {
         for (j=0; j < subsample_count; j++) {
            if(version == 1)
            {
               unsigned int(32) subsample_size;
            }
else {
               unsigned int(16) subsample_size;
            }
            unsigned int(8) subsample_priority;
            unsigned int(8) discardable;
            unsigned int(32) reserved = 0;
} }
*/
type SubSampleInformationBox struct {
	FullBox
}

//ProgressiveDownloadInfoBox -
/*
aligned(8) class ProgressiveDownloadInfoBox extends FullBox(‘pdin’, version = 0, 0) {
	for(i=0;;i++){ //toendofbox unsigned int(32) rate;
	unsigned int(32) initial_delay;
	} }
*/
type ProgressiveDownloadInfoBox struct {
	FullBox
}
