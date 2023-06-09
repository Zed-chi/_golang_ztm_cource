package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	dice, sides := 2, 12
	rolls := 1
	for r := 1; r <= rolls; r++ {
		sum := 0
		for d := 1; d <= dice; d++ {
			rolled := roll(sides)
			sum += rolled
			fmt.Println(r, d, rolled)
		}
	}
}

func roll(sides int) int {
	return rand.Intn(sides) + 1
}
