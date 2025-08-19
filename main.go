package main

import (
	"fmt"
)

const USD_RUB = 80.35
const USD_EUR = 0.85

func main() {
	amount, fromCurrency, toCurrency := getUserInput()
	result := calculate(amount, fromCurrency, toCurrency)
	fmt.Printf("Итог: %.2f %s\n", result, toCurrency)
}

func calculate(amount float64, from string, to string) float64 {
	// Конвертируем через USD как базовую валюту
	var amountInUSD float64
	
	// Сначала переводим в USD
	if from == "USD" {
		amountInUSD = amount
	} else if from == "RUB" {
		amountInUSD = amount / USD_RUB
	} else if from == "EUR" {
		amountInUSD = amount / USD_EUR
	} else {
		fmt.Println("Неизвестная исходная валюта")
		return 0
	}
	
	// Затем конвертируем из USD в целевую валюту
	if to == "USD" {
		return amountInUSD
	} else if to == "RUB" {
		return amountInUSD * USD_RUB
	} else if to == "EUR" {
		return amountInUSD * USD_EUR
	} else {
		fmt.Println("Неизвестная целевая валюта")
		return 0
	}
}

func getUserInput() (float64, string, string) {
	var amount float64
	var fromCurrency, toCurrency string
	
	fmt.Print("Введите сумму: ")
	fmt.Scan(&amount)

	fmt.Print("Выберите исходную валюту (USD, RUB, EUR): ")
	fmt.Scan(&fromCurrency)

	fmt.Print("Выберите целевую валюту (USD, RUB, EUR): ")
	fmt.Scan(&toCurrency)

	return amount, fromCurrency, toCurrency
}

