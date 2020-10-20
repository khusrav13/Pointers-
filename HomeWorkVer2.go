package main

import (
	"fmt"
)

func main()  {
	////I use uint (without negative numbers)
	newArray([]uint{122,65,23,80,16,70,100,500,1235,70})
}

func newArray(arrays  []uint) {


	min := arrays[0]

	max := arrays[0]


	///The smallest number
	for _, newValue := range arrays {
		if newValue < min {
			min = newValue
		}
	}

	///The biggest number
	for _, newValue := range arrays {
		if newValue > max {
			max = newValue
		}
	}



	fmt.Println("Our smallest number is : ", min, "and the biggest is:", max)

}