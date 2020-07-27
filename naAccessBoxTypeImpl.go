package mp4box

//naAccessBoxTypeImpl - Returns nil for all types of Box
type naAccessBoxTypeImpl struct {
}

//CollectionBaseBox -
func (*naAccessBoxTypeImpl) GetCollectionBaseBox() (*CollectionBaseBox, error) {
	return nil, ErrBoxNotFound
}

//CollectionFullBox -
func (*naAccessBoxTypeImpl) GetCollectionFullBox() (*CollectionFullBox, error) {
	return nil, ErrBoxNotFound
}

//CollectionFullBoxCounted -
func (*naAccessBoxTypeImpl) GetCollectionFullBoxCounted() (*CollectionFullBoxCounted, error) {
	return nil, ErrBoxNotFound
}

//FileBox -
func (*naAccessBoxTypeImpl) GetFileBox() (*FileBox, error) { return nil, ErrBoxNotFound }

//MediaDataBox -
func (*naAccessBoxTypeImpl) GetMediaDataBox() (*MediaDataBox, error) { return nil, ErrBoxNotFound }

//MovieHeaderBox -
func (*naAccessBoxTypeImpl) GetMovieHeaderBox() (*MovieHeaderBox, error) { return nil, ErrBoxNotFound }

//TrackHeaderBox -
func (*naAccessBoxTypeImpl) GetTrackHeaderBox() (*TrackHeaderBox, error) { return nil, ErrBoxNotFound }

//MediaHeaderBox -
func (*naAccessBoxTypeImpl) GetMediaHeaderBox() (*MediaHeaderBox, error) { return nil, ErrBoxNotFound }

//HandlerBox -
func (*naAccessBoxTypeImpl) GetHandlerBox() (*HandlerBox, error) { return nil, ErrBoxNotFound }

//VideoMediaHeaderBox -
func (*naAccessBoxTypeImpl) GetVideoMediaHeaderBox() (*VideoMediaHeaderBox, error) {
	return nil, ErrBoxNotFound
}

//SoundMediaHeaderBox -
func (*naAccessBoxTypeImpl) GetSoundMediaHeaderBox() (*SoundMediaHeaderBox, error) {
	return nil, ErrBoxNotFound
}

//HintMediaHeaderBox -
func (*naAccessBoxTypeImpl) GetHintMediaHeaderBox() (*HintMediaHeaderBox, error) {
	return nil, ErrBoxNotFound
}

//NullMediaHeaderBox -
func (*naAccessBoxTypeImpl) GetNullMediaHeaderBox() (*NullMediaHeaderBox, error) {
	return nil, ErrBoxNotFound
}

//DataEntryURLBox -
func (*naAccessBoxTypeImpl) GetDataEntryURLBox() (*DataEntryURLBox, error) {
	return nil, ErrBoxNotFound
}

//DataEntryUrnBox -
func (*naAccessBoxTypeImpl) GetDataEntryUrnBox() (*DataEntryUrnBox, error) {
	return nil, ErrBoxNotFound
}

//TimeToSampleBox -
func (*naAccessBoxTypeImpl) GetTimeToSampleBox() (*TimeToSampleBox, error) {
	return nil, ErrBoxNotFound
}

//CompositionOffsetBox -
func (*naAccessBoxTypeImpl) GetCompositionOffsetBox() (*CompositionOffsetBox, error) {
	return nil, ErrBoxNotFound
}

//SampleDescriptionBox -
func (*naAccessBoxTypeImpl) GetSampleDescriptionBox() (*SampleDescriptionBox, error) {
	return nil, ErrBoxNotFound
}

//SampleSizeBox -
func (*naAccessBoxTypeImpl) GetSampleSizeBox() (*SampleSizeBox, error) { return nil, ErrBoxNotFound }

//CompactSampleSizeBox -
func (*naAccessBoxTypeImpl) GetCompactSampleSizeBox() (*CompactSampleSizeBox, error) {
	return nil, ErrBoxNotFound
}

//SampleToChunkBox -
func (*naAccessBoxTypeImpl) GetSampleToChunkBox() (*SampleToChunkBox, error) {
	return nil, ErrBoxNotFound
}

//ChunkOffsetBox -
func (*naAccessBoxTypeImpl) GetChunkOffsetBox() (*ChunkOffsetBox, error) { return nil, ErrBoxNotFound }

//ChunkLargeOffsetBox -
func (*naAccessBoxTypeImpl) GetChunkLargeOffsetBox() (*ChunkLargeOffsetBox, error) {
	return nil, ErrBoxNotFound
}

//SyncSampleBox -
func (*naAccessBoxTypeImpl) GetSyncSampleBox() (*SyncSampleBox, error) { return nil, ErrBoxNotFound }

//ShadowSyncSampleBox -
func (*naAccessBoxTypeImpl) GetShadowSyncSampleBox() (*ShadowSyncSampleBox, error) {
	return nil, ErrBoxNotFound
}

//DegradationPriorityBox -
func (*naAccessBoxTypeImpl) GetDegradationPriorityBox() (*DegradationPriorityBox, error) {
	return nil, ErrBoxNotFound
}

//PaddingBitsBox -
func (*naAccessBoxTypeImpl) GetPaddingBitsBox() (*PaddingBitsBox, error) { return nil, ErrBoxNotFound }

//FreeSpaceBox -
func (*naAccessBoxTypeImpl) GetFreeSpaceBox() (*FreeSpaceBox, error) { return nil, ErrBoxNotFound }

//EditListBox -
func (*naAccessBoxTypeImpl) GetEditListBox() (*EditListBox, error) { return nil, ErrBoxNotFound }

//UserDataBox -
func (*naAccessBoxTypeImpl) GetUserDataBox() (*UserDataBox, error) { return nil, ErrBoxNotFound }

//CopyrightBox -
func (*naAccessBoxTypeImpl) GetCopyrightBox() (*CopyrightBox, error) { return nil, ErrBoxNotFound }

//MovieExtendsHeaderBox -
func (*naAccessBoxTypeImpl) GetMovieExtendsHeaderBox() (*MovieExtendsHeaderBox, error) {
	return nil, ErrBoxNotFound
}

//TrackExtendsBox -
func (*naAccessBoxTypeImpl) GetTrackExtendsBox() (*TrackExtendsBox, error) {
	return nil, ErrBoxNotFound
}

//MovieFragmentHeaderBox -
func (*naAccessBoxTypeImpl) GetMovieFragmentHeaderBox() (*MovieFragmentHeaderBox, error) {
	return nil, ErrBoxNotFound
}

//TrackFragmentHeaderBox -
func (*naAccessBoxTypeImpl) GetTrackFragmentHeaderBox() (*TrackFragmentHeaderBox, error) {
	return nil, ErrBoxNotFound
}

//TrackRunBox -
func (*naAccessBoxTypeImpl) GetTrackRunBox() (*TrackRunBox, error) { return nil, ErrBoxNotFound }

//TrackFragmentRandomAccessBox -
func (*naAccessBoxTypeImpl) GetTrackFragmentRandomAccessBox() (*TrackFragmentRandomAccessBox, error) {
	return nil, ErrBoxNotFound
}

//MovieFragmentRandomAccessOffsetBox -
func (*naAccessBoxTypeImpl) GetMovieFragmentRandomAccessOffsetBox() (*MovieFragmentRandomAccessOffsetBox, error) {
	return nil, ErrBoxNotFound
}

//SampleDependencyTypeBox -
func (*naAccessBoxTypeImpl) GetSampleDependencyTypeBox() (*SampleDependencyTypeBox, error) {
	return nil, ErrBoxNotFound
}

//SampleToGroupBox -
func (*naAccessBoxTypeImpl) GetSampleToGroupBox() (*SampleToGroupBox, error) {
	return nil, ErrBoxNotFound
}

//SampleGroupDescriptionBox -
func (*naAccessBoxTypeImpl) GetSampleGroupDescriptionBox() (*SampleGroupDescriptionBox, error) {
	return nil, ErrBoxNotFound
}

//SampleScaleBox -
func (*naAccessBoxTypeImpl) GetSampleScaleBox() (*SampleScaleBox, error) { return nil, ErrBoxNotFound }

//SubSampleInformationBox -
func (*naAccessBoxTypeImpl) GetSubSampleInformationBox() (*SubSampleInformationBox, error) {
	return nil, ErrBoxNotFound
}

//ProgressiveDownloadInfoBox -
func (*naAccessBoxTypeImpl) GetProgressiveDownloadInfoBox() (*ProgressiveDownloadInfoBox, error) {
	return nil, ErrBoxNotFound
}

func (*naAccessBoxTypeImpl) GetSegmentBox() (*SegmentBox, error) {
	return nil, ErrBoxNotFound
}

func (*naAccessBoxTypeImpl) GetTrackFragmentMediaAdjustmentBox() (*TrackFragmentMediaAdjustmentBox, error) {
	return nil, ErrBoxNotFound
}
func (*naAccessBoxTypeImpl) GetSegmentIndexBox() (*SegmentIndexBox, error) {
	return nil, ErrBoxNotFound
}
func (*naAccessBoxTypeImpl) GetTrackFragmentBaseMediaDecodeTimeBox() (*TrackFragmentBaseMediaDecodeTimeBox, error) {
	return nil, ErrBoxNotFound
}

func (*naAccessBoxTypeImpl) GetMovieFragmentBox() (*MovieFragmentBox, error) {
	return nil, ErrBoxNotFound
}
func (*naAccessBoxTypeImpl) GetMovieBox() (*MovieBox, error) {
	return nil, ErrBoxNotFound
}

//BaseBox -
func (*naAccessBoxTypeImpl) GetBaseBox() (*BaseBox, error) { return nil, ErrBoxNotFound }
