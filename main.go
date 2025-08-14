// Курс актуален на 14.08.2025

package main

import (
	"fmt"
)

func main()  {
	const USD_EUR = 0.859
    const USD_RUB = 79.77
	EUR_RUB := USD_RUB / USD_EUR	
	fmt.Print(EUR_RUB)
}