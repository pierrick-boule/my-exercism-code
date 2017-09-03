package triangle

import (
	"sort"
)

const testVersion = 3

func KindFromSides(a, b, c float64) Kind {
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}
	triangle := []float64{a, b, c}
	sort.Float64s(triangle)

	if !((triangle[0]+triangle[1])-triangle[2] >= 0) {
		return NaT
	}
	var count int
	if a == b {
		count++
	}
	if b == c {
		count++
	}
	if c == a {
		count++
	}
	if count == 1 {
		return Iso
	} else if count > 1 {
		return Equ
	}
	return Sca
}

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

// Pick values for the following identifiers used by the test program.
const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// Organize your code for readability.
