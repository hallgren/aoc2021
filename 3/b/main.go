package main

import (
	"fmt"

	"github.com/hallgren/aoc2021"
)

type List struct {
	One  []aoc2021.Decimal
	Zero []aoc2021.Decimal
}

func main() {
	var counter = make([]List, 12)
	valuesOxygen := aoc2021.Decimals("../input")
	valuesCO2 := aoc2021.Decimals("../input")
	for i := 0; i < len(counter); i++ {
		for _, value := range valuesOxygen {
			if value.Zero(i) {
				counter[i].Zero = append(counter[i].Zero, value)
			} else {
				counter[i].One = append(counter[i].One, value)
			}
		}
		if len(counter[i].One) >= len(counter[i].Zero) {
			valuesOxygen = counter[i].One
		} else {
			valuesOxygen = counter[i].Zero
		}
		if len(valuesOxygen) == 1 {
			break
		}
	}
	// CO2
	counter = make([]List, 12)
	for i := 0; i < len(counter); i++ {
		for _, value := range valuesCO2 {
			if value.Zero(i) {
				counter[i].Zero = append(counter[i].Zero, value)
			} else {
				counter[i].One = append(counter[i].One, value)
			}
		}
		if len(counter[i].One) >= len(counter[i].Zero) {
			valuesCO2 = counter[i].Zero
		} else {
			valuesCO2 = counter[i].One
		}
		if len(valuesCO2) == 1 {
			break
		}
	}
	fmt.Println(valuesOxygen[0].Dec() * valuesCO2[0].Dec())
}
