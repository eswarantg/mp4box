package mp4box

// BoxFactory - Create BoxObject
type BoxFactory struct {
}

// MakeEmptyBoxObject - make object by boxType
func (BoxFactory) MakeEmptyBoxObject(boxType string) Box {
	//collectionBaseBoxList
	switch boxType {
	//CollectionBaseBox - Begin
	case "moov":
		return &MovieBox{}
	case "moof":
		return &MovieFragmentBox{}
	case "trak", "mdia", "minf", "dinf",
		"stbl", "mvex", "traf":
		return &CollectionBaseBox{}
	//CollectionBaseBox - End
	//CollectionFullBox - Begin
	case "XXXX":
		return &CollectionFullBox{}
	//CollectionFullBox - End
	//CollectionFullBoxCounted - Begin
	case "stsd", "dref":
		return &CollectionFullBoxCounted{}
	//CollectionFullBoxCounted - End
	//Standalone boxes - Begin
	case "ftyp":
		return new(FileBox)
	case "mdat":
		return new(MediaDataBox)
	case "mvhd":
		return new(MovieHeaderBox)
	case "tkhd":
		return new(TrackHeaderBox)
	case "mdhd":
		return new(MediaHeaderBox)
	case "hdlr":
		return new(HandlerBox)
	case "vmhd":
		return new(VideoMediaHeaderBox)
	case "smhd":
		return new(SoundMediaHeaderBox)
	case "hmhd":
		return new(HintMediaHeaderBox)
	case "nmhd":
		return new(NullMediaHeaderBox)
	case "url ":
		return new(DataEntryURLBox)
	case "urn ":
		return new(DataEntryUrnBox)
	case "stts":
		return new(TimeToSampleBox)
	case "ctts":
		return new(CompositionOffsetBox)
	case "stsz":
		return new(SampleSizeBox)
	case "stz2":
		return new(CompactSampleSizeBox)
	case "stsc":
		return new(SampleToChunkBox)
	case "stco":
		return new(ChunkOffsetBox)
	case "c064":
		return new(ChunkLargeOffsetBox)
	case "stss":
		return new(SyncSampleBox)
	case "stsh":
		return new(ShadowSyncSampleBox)
	case "stdp":
		return new(DegradationPriorityBox)
	case "padb":
		return new(PaddingBitsBox)
	case "free":
	case "skip":
		return new(FreeSpaceBox)
	case "elst":
		return new(EditListBox)
	case "udta":
		return new(UserDataBox)
	case "cprt":
		return new(CopyrightBox)
	case "mehd":
		return new(MovieExtendsHeaderBox)
	case "trex":
		return new(TrackExtendsBox)
	case "mfhd":
		return new(MovieFragmentHeaderBox)
	case "tfhd":
		return new(TrackFragmentHeaderBox)
	case "trun":
		return new(TrackRunBox)
	case "tfra":
		return new(TrackFragmentRandomAccessBox)
	case "mfro":
		return new(MovieFragmentRandomAccessOffsetBox)
	case "sdtp":
		return new(SampleDependencyTypeBox)
	case "sbgp":
		return new(SampleToGroupBox)
	case "sgpd":
		return new(SampleGroupDescriptionBox)
	case "stsl":
		return new(SampleScaleBox)
	case "subs":
		return new(SubSampleInformationBox)
	case "pdin":
		return new(ProgressiveDownloadInfoBox)
		//Standalone boxes - End
	case "styp":
		return new(SegmentBox)
	case "tfma":
		return new(TrackFragmentMediaAdjustmentBox)
	case "sidx":
		return new(SegmentIndexBox)
	case "tfdt":
		return new(TrackFragmentBaseMediaDecodeTimeBox)
		//mp4boxes
	case "mp4v":
		return new(MP4VisualSampleEntry)
	case "mp4a":
		return new(MP4AudioSampleEntry)
	case "mp4s":
		return new(MpegSampleEntry)
	case "esds":
		return new(ESDBox)
	case "emsg":
		return new(EmsgBox)
	}
	return new(BaseBox)
}
