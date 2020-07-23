package mp4box

import "io"

//AccessBoxType - Interface for accessing the specific Box Types
type AccessBoxType interface {
	//Internal: GetBox()
	getLevel() int
	//Internal: IsCollection()
	isCollection() bool
	//Returns BoxType of the Box
	Boxtype() string
	//Returns Size of the Box
	Size() int64
	//User Readable description of content
	String() string
	//Internal: Init the box with boxSize, boxType, payload
	initData(boxSize int64, boxType string, payload *[]byte, parent AccessBoxType) error
	//Internal: Write the box to a io stream
	writeBytes(io.Writer) error
	//GetParentByName() returns Box interface of Box by name in heirachy
	GetChildByName(boxType string) (AccessBoxType, error)
	//GetParentByName() returns Box interface of Box by name in heirachy
	GetParentByName(boxType string) (AccessBoxType, error)
	//Internal: getLeafBox() returns leaf object Box interface
	getLeafBox() AccessBoxType

	//External: Get specific Box
	GetCollectionBaseBox() (*CollectionBaseBox, error)
	GetCollectionFullBox() (*CollectionFullBox, error)
	GetCollectionFullBoxCounted() (*CollectionFullBoxCounted, error)
	GetFileBox() (*FileBox, error)
	GetMediaDataBox() (*MediaDataBox, error)
	GetMovieHeaderBox() (*MovieHeaderBox, error)
	GetTrackHeaderBox() (*TrackHeaderBox, error)
	GetMediaHeaderBox() (*MediaHeaderBox, error)
	GetHandlerBox() (*HandlerBox, error)
	GetVideoMediaHeaderBox() (*VideoMediaHeaderBox, error)
	GetSoundMediaHeaderBox() (*SoundMediaHeaderBox, error)
	GetHintMediaHeaderBox() (*HintMediaHeaderBox, error)
	GetNullMediaHeaderBox() (*NullMediaHeaderBox, error)
	GetDataEntryURLBox() (*DataEntryURLBox, error)
	GetDataEntryUrnBox() (*DataEntryUrnBox, error)
	GetTimeToSampleBox() (*TimeToSampleBox, error)
	GetCompositionOffsetBox() (*CompositionOffsetBox, error)
	GetSampleDescriptionBox() (*SampleDescriptionBox, error)
	GetSampleSizeBox() (*SampleSizeBox, error)
	GetCompactSampleSizeBox() (*CompactSampleSizeBox, error)
	GetSampleToChunkBox() (*SampleToChunkBox, error)
	GetChunkOffsetBox() (*ChunkOffsetBox, error)
	GetChunkLargeOffsetBox() (*ChunkLargeOffsetBox, error)
	GetSyncSampleBox() (*SyncSampleBox, error)
	GetShadowSyncSampleBox() (*ShadowSyncSampleBox, error)
	GetDegradationPriorityBox() (*DegradationPriorityBox, error)
	GetPaddingBitsBox() (*PaddingBitsBox, error)
	GetFreeSpaceBox() (*FreeSpaceBox, error)
	GetEditListBox() (*EditListBox, error)
	GetUserDataBox() (*UserDataBox, error)
	GetCopyrightBox() (*CopyrightBox, error)
	GetMovieExtendsHeaderBox() (*MovieExtendsHeaderBox, error)
	GetTrackExtendsBox() (*TrackExtendsBox, error)
	GetMovieFragmentHeaderBox() (*MovieFragmentHeaderBox, error)
	GetTrackFragmentHeaderBox() (*TrackFragmentHeaderBox, error)
	GetTrackRunBox() (*TrackRunBox, error)
	GetTrackFragmentRandomAccessBox() (*TrackFragmentRandomAccessBox, error)
	GetMovieFragmentRandomAccessOffsetBox() (*MovieFragmentRandomAccessOffsetBox, error)
	GetSampleDependencyTypeBox() (*SampleDependencyTypeBox, error)
	GetSampleToGroupBox() (*SampleToGroupBox, error)
	GetSampleGroupDescriptionBox() (*SampleGroupDescriptionBox, error)
	GetSampleScaleBox() (*SampleScaleBox, error)
	GetSubSampleInformationBox() (*SubSampleInformationBox, error)
	GetProgressiveDownloadInfoBox() (*ProgressiveDownloadInfoBox, error)
	GetBaseBox() (*BaseBox, error)
}
