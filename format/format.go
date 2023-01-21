package main

import "fmt"

func paintNedeed(width float64, heigth float64) float64 { //объявление функции
	area := width * heigth
    return area / 10.0
}

func main() {
    var amount, total float64
    amount = paintNedeed(4.2, 3.0)
    fmt.Printf("%0.2f Litres nedeed\n", amount)
    total += amount
    amount = paintNedeed(5.2, 3.5)
    fmt.Printf("%0.2f Litres nedeed\n", amount)
    total += amount
    fmt.Printf("Total: %0.2f liters\n", total)
}