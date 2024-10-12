package ctype

import "encoding/json"

type ImageType int

const (
	Local ImageType = 1
	S3    ImageType = 2
)

func (s ImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s ImageType) String() string {
	var str string
	switch s {
	case Local:
		str = "local"
	case S3:
		str = "cloud"
	default:
		str = "other"
	}
	return str
}
