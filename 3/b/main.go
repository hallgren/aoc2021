package main

import (
	"fmt"

	"github.com/hallgren/aoc2021"
)

type List struct {
	One  Bucket
	Zero Bucket
}

type Bucket struct {
	Counter int
	List    []string
}

func main() {
	var counter = make([]List, 12)
	valuesOxygen := aoc2021.Lines("input.txt")
	valuesCO2 := aoc2021.Lines("input.txt")
	for i := 0; i < len(counter); i++ {
		for _, value := range valuesOxygen {
			if value[i] == 48 {
				counter[i].Zero.Counter++
				counter[i].Zero.List = append(counter[i].Zero.List, value)
			} else {
				counter[i].One.Counter++
				counter[i].One.List = append(counter[i].One.List, value)
			}
		}
		if counter[i].One.Counter >= counter[i].Zero.Counter {
			valuesOxygen = counter[i].One.List
		} else {
			valuesOxygen = counter[i].Zero.List
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
				counter[i].Zero.Counter++
				counter[i].Zero.List = append(counter[i].Zero.List, value)
			} else {
				counter[i].One.Counter++
				counter[i].One.List = append(counter[i].One.List, value)
			}
		}
		if counter[i].One.Counter >= counter[i].Zero.Counter {
			valuesCO2 = counter[i].Zero.List
		} else {
			valuesCO2 = counter[i].One.List
		}
		if len(valuesCO2) == 1 {
			break
		}
	}
	fmt.Println(aoc2021.StringToDec(valuesOxygen[0]) * aoc2021.StringToDec(valuesCO2[0]))
}
