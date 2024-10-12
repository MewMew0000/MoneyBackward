package res

type ErrorCode int

const (
	SettingsError ErrorCode = 1001 // system error
	ArgumentError ErrorCode = 1002 // parameter error
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "system error",
		ArgumentError: "pparameter error",
	}
)
