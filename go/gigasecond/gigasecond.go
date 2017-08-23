// Package gigasecond provide functionality for gigasecond manipulations
package gigasecond

// import path for the time package from the standard library
import (
	"time"
)

const testVersion = 4
// GigaSecond duration.
const GigaSecond = time.Second * 1e9

// AddGigasecond add 10^9 seconds to a time.
func AddGigasecond(input time.Time) time.Time {
	return input.Add(GigaSecond)
}
