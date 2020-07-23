package mp4box

import (
	"errors"
)

//ErrBoxNotFound - box searched is not found
var ErrBoxNotFound error = errors.New("BoxNotFound")

//Box - Interface for ISO BMFF Box
type Box interface {
	//Returns BoxType of the Box
	Boxtype() string
	//Returns Size of the Box
	Size() int64
	//User Readable description of content
	String() string
	//GetParentByName() returns Box interface of Box by name in heirachy
	GetChildByName(boxType string) (AccessBoxType, error)
	//GetParentByName() returns Box interface of Box by name in heirachy
	GetParentByName(boxType string) (AccessBoxType, error)
}
