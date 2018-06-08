package covargo

import "time"

const (
	internal_BUILD_TIMESTAMP = 1528420598
	internal_BUILD_NUMBER    = 10
	internal_VERSION_STRING  = "0.1.1"
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
