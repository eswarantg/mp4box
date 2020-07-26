package mp4box

//Source
//https://mpeg.chiariglione.org/standards/mpeg-4/iso-base-media-file-format/text-isoiec-14496-12-5th-edition

//MediaDataBox -
/*
aligned(8) class MediaDataBox extends Box(‘mdat’) { bit(8) data[];
}
*/
type MediaDataBox struct {
	BaseBox
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

//Source:
//https://www.etsi.org/deliver/etsi_ts/126200_126299/126244/10.02.00_60/ts_126244v100200p.pdf

//TrackFragmentMediaAdjustmentBox -
/*
aligned(8) class TrackFragmentMediaAdjustmentBox extends FullBox("tfma", version, 0) {
	unsigned int(32) entry_count;
	for (i=1; i <= entry_count; i++) {
		if (version==1) {
			unsigned int(64) segment_duration; int(64) media_time;
		} else { // version==0
			unsigned int(32) segment_duration; int(32) media_time;
		}
		int(16) media_rate_integer; int(16) media_rate_fraction = 0;
	}
}
*/
type TrackFragmentMediaAdjustmentBox struct {
	FullBox
}

//SegmentIndexBox -
/*
aligned(8) class SegmentIndexBox extends FullBox("sidx", version, 0) {
	unsigned int(32) reference_ID;
	unsigned int(32) timescale;
	if (version==0)
	{
		bit (1)				reference_type;
		unsigned int(31)	referenced_size;
		unsigned int(32)	subsegment_duration;
		bit(1)				starts_with_SAP;
		unsigned int(3)		SAP_type;
		unsigned int(28)	SAP_delta_time;
		unsigned int(32) earliest_presentation_time; unsigned int(32) first_offset;
	} else {
		unsigned int(64) earliest_presentation_time;
		unsigned int(64) first_offset; }
		unsigned int(16) reserved = 0; unsigned int(16) reference_count; for(i=1; i <= reference_count; i++) {
	}
}
*/
type SegmentIndexBox struct {
	FullBox
}
