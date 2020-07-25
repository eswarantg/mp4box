package mp4box

import (
	"errors"
	"io"
)

//ErrBoxNotFound - box searched is not found
var ErrBoxNotFound error = errors.New("BoxNotFound")

//Box - Interface for ISO BMFF Box
type Box interface {
	//Internal: GetBox()
	getLevel() int
	//Internal: IsCollection()
	isCollection() bool
	//Internal: Init the box with boxSize, boxType, payload
	initData(boxSize int64, boxType string, payload *[]byte, parent Box) error
	//Internal: getLeafBox() returns leaf object Box interface
	getLeafBox() Box

	//External Interface - Outside Package

	//Write - Writes the box bytes to io.Writer
	Write(io.Writer) error
	//Returns BoxType of the Box
	Boxtype() string
	//User Readable description of content
	String() string
	//GetParentByName() returns Box interface of Box by name in heirachy
	GetChildByName(boxType string) (Box, error)
	//GetParentByName() returns Box interface of Box by name in heirachy
	GetParentByName(boxType string) (Box, error)

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
	GetSegmentBox() (*SegmentBox, error)
	GetTrackFragmentMediaAdjustmentBox() (*TrackFragmentMediaAdjustmentBox, error)
	GetSegmentIndexBox() (*SegmentIndexBox, error)
	GetTrackFragmentBaseMediaDecodeTimeBox() (*TrackFragmentBaseMediaDecodeTimeBox, error)
	//Default
	GetBaseBox() (*BaseBox, error)
}
