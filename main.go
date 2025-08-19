

package main

import (
	"fmt"
)

func main()  {


	var USD_EUR float64
    var USD_RUB float64
	EUR_RUB := USD_RUB / USD_EUR	
	fmt.Print(EUR_RUB)
}

func getUserInput() (float64, float64) {
	var userUSD float64
	var userKg float64
	fmt.Print("Введите количество USD: ")
	fmt.Scan(&userUSD)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(&userKg)
	return userUSD, userKg
}