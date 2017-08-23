package acronym

import "strings"

const testVersion = 3
const space = " "
func Abbreviate(s string) string {
	var abr string
	s = strings.Replace(s,"-"," ", -1)
	for _,e := range strings.Split(s, space) {
		abr += string(e[0])
	}
	return strings.ToUpper(abr)
}
