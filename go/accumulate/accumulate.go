package accumulate

const testVersion = 1

func Accumulate(array []string, operation func(string) string) []string {
	for i, e := range array {
		array[i] = operation(e)
	}
	return array
}
