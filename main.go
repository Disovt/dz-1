package main

import (
	"errors"
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

func getUserInputSourceCurrency(string, string) (float64, error) {
	
	var fromCurrency, toCurrency string
	
	fmt.Println("Выберите исходную валюту (USD, RUB, EUR): ")
	fmt.Scan(&fromCurrency)
	if fromCurrency != "USD" || fromCurrency != "usd" || fromCurrency != "EUR" || fromCurrency != "eur" || fromCurrency != "RUB" || fromCurrency != "rub"  {
		return 0, errors.New("Неправильно введена исходная валюта")
	}

	switch {
	case fromCurrency == "RUB" || fromCurrency == "rub":
		fmt.Println("Выберите целевую валюту (USD, EUR): ")
	case fromCurrency == "EUR" || fromCurrency == "eur":
		fmt.Println("Выберите целевую валюту (USD, RUB): ")
	case fromCurrency == "USD" || fromCurrency == "usd":
		fmt.Println("Выберите целевую валюту (RUB, EUR): ")
	}
	
	fmt.Scan(&toCurrency)

	return fromCurrency, toCurrency
}
	func getUserInputAmount(float64) (float64, error) {
	var amount float64
		fmt.Println("Введите сумму: ")
	fmt.Scan(&amount)

	if amount > 0 {
		return amount
		
	} else { 
		return 1, errors.New("Некорректно введено количество валюты") 
	}
	}

	

	

