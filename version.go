package covargo

import (
	"log"
	"time"
)

const (
	internal_BUILD_TIMESTAMP = 1528420599
	internal_BUILD_NUMBER    = 12
	internal_VERSION_STRING  = "0.1.3"
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

func About() {
	log.Printf("covargo (v%v, build %v, build date:%v)", Version(), BuildNumber(), BuildDate())
}
