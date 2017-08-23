package raindrops
import (
	"bytes"
	"fmt"
)
const testVersion = 3

func divisibleBy3(num int) bool {
	// A number is divisible by 3 if all its digits are divisible by 3
	sum := 0
	for num >= 10 {
		sum += num % 10
		num = num / 10
	}
	sum += num
	return sum%3 == 0
}

func divisibleBy5(num int) bool {
	// A number is divisible by 5 if it ends in 0 or 5
	remainder := num%10
	return (remainder == 0) || (remainder == 5)
}

func divisibleBy7(num int) bool {
	// A number is divisible by 7 if
	remainder := num%10
	num_prefix := num/10

	return (num_prefix-(remainder*2))%7 == 0
}
func Convert(num int) string {
	var buffer bytes.Buffer
	if divisibleBy3(num){
		buffer.WriteString("Pling")
	}
	if divisibleBy5(num){
		buffer.WriteString("Plang")
	}
	if divisibleBy7(num){
		buffer.WriteString("Plong")
	}
	if buffer.String() == ""{
		buffer.WriteString(fmt.Sprintf("%d", num))
	}
	return buffer.String()
}