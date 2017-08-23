package hamming

import (
	"errors"
)

const testVersion = 6

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0 ,errors.New("length mismatch")
	}
	var distance int
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}