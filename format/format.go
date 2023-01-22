package main

import "fmt"

func paintNedeed(width float64, heigth float64) (float64, error) {
    if width < 0 {
        return 0, fmt.Errorf("a width of %0.2f is invalid", width)
    }
    if heigth < 0 {
        return 0, fmt.Errorf("a height of %0.2f is invalid", heigth)
    }
	area := width * heigth
    return area / 10.0, nil
}

func main() {
    amount, err := paintNedeed(4.2, -3.0)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("%0.2f Litres nedeed\n", amount)
    }
    //fmt.Println(err)
    //fmt.Printf("%0.2f Litres nedeed\n", amount)
    /*var amount, total float64
    amount = paintNedeed(4.2, 3.0)
    fmt.Printf("%0.2f Litres nedeed\n", amount)
    total += amount
    amount = paintNedeed(5.2, 3.5)
    fmt.Printf("%0.2f Litres nedeed\n", amount)
    total += amount */
}