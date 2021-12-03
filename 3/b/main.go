package main

import (
	"fmt"

	"github.com/hallgren/aoc2021"
)

type List struct {
	One  []string
	Zero []string
}

func main() {
	var counter = make([]List, 12)
	valuesOxygen := aoc2021.Lines("../input")
	valuesCO2 := aoc2021.Lines("../input")
	for i := 0; i < len(counter); i++ {
		for _, value := range valuesOxygen {
			if value[i] == 48 {
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
			if value[i] == 48 {
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
	fmt.Println(aoc2021.StringBinToDec(valuesOxygen[0]) * aoc2021.StringBinToDec(valuesCO2[0]))
}
