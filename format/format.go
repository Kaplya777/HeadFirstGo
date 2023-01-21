package main

import "fmt"

func paintNedeed(width float64, heigth float64) { //объявление функции
	area := width * heigth
	fmt.Printf("%.2f liters needed\n", area/10.0)

}

func main() {
	paintNedeed(4.2, 3.0)
	paintNedeed(5.2, 3.5)
	paintNedeed(5.0, 3.3)
}
