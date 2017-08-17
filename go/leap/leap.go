// Package leap contain function utility for leap year.
package leap

const testVersion = 3

// IsLeapYear is a function that take a year as argument and return true if it's a leap year.
func IsLeapYear(year int) bool {
	return year%400 == 0 || (year%4 == 0 && !(year%100 == 0))
}
