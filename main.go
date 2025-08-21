package main

import (

	"fmt"
)

const USD_RUB = 80.35
const USD_EUR = 0.85

func main() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("Ошибка:", r)
		}
	}()
	
	for {
		fmt.Println("=== Конвертер валют ===")
		amount, fromCurrency, toCurrency := getUserInput()
		result := calculate(amount, fromCurrency, toCurrency)
		fmt.Printf("Итог: %.2f %s\n", result, toCurrency)
		
		isRepeat := checkRepeat()
		if !isRepeat {
			break
		}
	}
}

func calculate(amount float64, from string, to string) float64 {
	// Конвертируем через USD как базовую валюту
	var amountInUSD float64
	
	// Сначала переводим в USD
	if from == "USD" || from == "usd" {
		amountInUSD = amount
	} else if from == "RUB" || from == "rub" {
		amountInUSD = amount / USD_RUB
	} else if from == "EUR" || from == "eur" {
		amountInUSD = amount / USD_EUR
	} 
	
	// Затем конвертируем из USD в целевую валюту
	if to == "USD" || to == "usd" {
		return amountInUSD
	} else if to == "RUB" || to == "rub" {
		return amountInUSD * USD_RUB
	} else if to == "EUR" || to == "eur" {
		return amountInUSD * USD_EUR
	}
	
	return 0
}

func getUserInput() (float64, string, string) {
	var fromCurrency string
	var amount float64
	var toCurrency string

	// Шаг 1: Ввод исходной валюты
	fmt.Println("=== Шаг 1: Выбор исходной валюты ===")
	for {
		fmt.Println("Доступные валюты: USD, RUB, EUR")
		fmt.Print("Введите исходную валюту: ")
		fmt.Scan(&fromCurrency)
		
		// Проверяем корректность ввода прямо здесь
		if fromCurrency == "USD" || fromCurrency == "usd" || 
		   fromCurrency == "RUB" || fromCurrency == "rub" || 
		   fromCurrency == "EUR" || fromCurrency == "eur" {
			break
		}
		// Если валюта некорректна, вызываем panic
		panic("Неправильно введена исходная валюта")
	}

	// Шаг 2: Ввод суммы
	fmt.Println("\n=== Шаг 2: Ввод суммы ===")
	for {
		fmt.Print("Введите сумму: ")
		_, err := fmt.Scan(&amount)
		if err == nil && amount > 0 {
			break
		}
		// Если сумма некорректна, вызываем panic
		panic("Некорректно введено количество валюты")
	}

	// Шаг 3: Ввод целевой валюты
	fmt.Println("\n=== Шаг 3: Выбор целевой валюты ===")
	for {
		// Подсказываем доступные варианты в зависимости от исходной валюты
		if fromCurrency == "RUB" || fromCurrency == "rub" {
			fmt.Println("Доступные валюты: USD, EUR")
			fmt.Print("Введите целевую валюту: ")
		} else if fromCurrency == "EUR" || fromCurrency == "eur" {
			fmt.Println("Доступные валюты: USD, RUB")
			fmt.Print("Введите целевую валюту: ")
		} else if fromCurrency == "USD" || fromCurrency == "usd" {
			fmt.Println("Доступные валюты: RUB, EUR")
			fmt.Print("Введите целевую валюту: ")
		}
		
		fmt.Scan(&toCurrency)
		
		// Проверяем, что целевая валюта корректна и отличается от исходной
		validCurrency := false
		if toCurrency == "USD" || toCurrency == "usd" || 
		   toCurrency == "RUB" || toCurrency == "rub" || 
		   toCurrency == "EUR" || toCurrency == "eur" {
			validCurrency = true
		}
		
		// Проверяем, что не выбрана та же самая валюта
		sameCurrency := false
		if (fromCurrency == "USD" || fromCurrency == "usd") && (toCurrency == "USD" || toCurrency == "usd") {
			sameCurrency = true
		} else if (fromCurrency == "RUB" || fromCurrency == "rub") && (toCurrency == "RUB" || toCurrency == "rub") {
			sameCurrency = true
		} else if (fromCurrency == "EUR" || fromCurrency == "eur") && (toCurrency == "EUR" || toCurrency == "eur") {
			sameCurrency = true
		}
		
		if validCurrency && !sameCurrency {
			break
		}
		// Если целевая валюта некорректна, вызываем panic
		panic("Неправильно введена целевая валюта")
	}

	return amount, fromCurrency, toCurrency
}

func checkRepeat() bool {
	var answer string
	fmt.Print("\nХотите выполнить еще одну конвертацию? (yes/no): ")
	fmt.Scan(&answer)
	return answer == "yes" || answer == "Yes" || answer == "YES"
}