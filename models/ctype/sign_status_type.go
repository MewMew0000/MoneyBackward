package ctype

import "encoding/json"

type SignStatus int

const (
	SignLine   SignStatus = 1
	SignWechat SignStatus = 2
	SignEmail  SignStatus = 3
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	var str string
	switch s {
	case SignLine:
		str = "Line"
	case SignWechat:
		str = "Wechat"
	case SignEmail:
		str = "Email"
	default:
		str = "other"
	}
	return str
}
