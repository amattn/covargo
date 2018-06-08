package covargo

import "time"

const (
	internal_BUILD_TIMESTAMP = 1528417661
	internal_BUILD_NUMBER    = 5
	internal_VERSION_STRING  = "0.1.0"
)

func BuildDate() time.Time {
	return time.Unix(internal_BUILD_TIMESTAMP, 0)
}
func BuildNumber() int64 {
	return internal_BUILD_NUMBER
}
func Version() string {
	return internal_VERSION_STRING
}
