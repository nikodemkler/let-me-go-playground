package main

import (
	"fmt"
)

func isEvenEnded(number int) bool {
	strRepr := fmt.Sprintf("%d", number)

	lastChar := strRepr[len(strRepr)-1]
	firstChar := strRepr[0]

	isEvenEnded := false
	if lastChar == firstChar {
		isEvenEnded = true
	}

	return isEvenEnded
}

func main () {
	// e.g. for a & b -> 1001 x 1011 => 1012011 the result is even ended
	evenEndedCounter := 0
	for i := 1000; i <= 9999; i++ {
		for j := i; j <= 9999; j++ {
			multiplication := i * j
			if isEvenEnded(multiplication) {
				evenEndedCounter++
			}
		}
	}

	fmt.Println(evenEndedCounter)
}