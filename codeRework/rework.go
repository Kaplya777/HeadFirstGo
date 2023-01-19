package main

import (
	"fmt"
)

func main() {
	// var pebbleweight float64 = 0.1
	// var rockweight float64 = 1.2
	// var boulderweight float64 = 502.4
	pebbleweight, rockweight, boulderweight := 0.1, 1.2, 502.4
	var totalWeight float64 = pebbleweight + rockweight + boulderweight
	fmt.Println(totalWeight)
}

//fas
