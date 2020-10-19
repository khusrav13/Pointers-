package main

import "fmt"



func main()  {
	x1 := 2
	y1 := 5
	x2 := 6
	y2 := 8

	swap(&x1, &y1, &x2, &y2)

	fmt.Println(x1, y1, x2, y2)
}


///Another function for changing variables
func swap(x1, y1, x2, y2 *int)  {
	*x1, *y1 , *x2, *y2 = *x2, *y2, *x1, *y1
}