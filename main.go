package main

import (
	"fmt"
)

const USD_RUB = 80.35
const USD_EUR = 0.85

func main() {
	for {
		fmt.Println("=== Конвертер валют ===")
		amount, fromCurrency, toCurrency := getUserInput()
		result := calculate(amount, fromCurrency, toCurrency)
		fmt.Printf("Итог: %.2f %s\n", result, toCurrency)

		if !checkRepeat() {
			break
		}
	}
}

func calculate(amount float64, from string, to string) float64 {
	// Конвертируем через USD как базовую валюту
	var amountInUSD float64

	// Сначала переводим в USD
	switch from {
	case "USD", "usd":
		amountInUSD = amount
	case "RUB", "rub":
		amountInUSD = amount / USD_RUB
	case "EUR", "eur":
		amountInUSD = amount / USD_EUR
	}

	// Затем конвертируем из USD в целевую валюту
	switch to {
	case "USD", "usd":
		return amountInUSD
	case "RUB", "rub":
		return amountInUSD * USD_RUB
	case "EUR", "eur":
		return amountInUSD * USD_EUR
	}

	return 0
}

// Функция для ввода и проверки исходной валюты
func getFromCurrency() string {
	var currency string

	for {
		fmt.Println("Доступные валюты: USD, RUB, EUR")
		fmt.Print("Введите исходную валюту: ")
		fmt.Scan(&currency)

		// Проверяем корректность ввода
		if currency == "USD" || currency == "usd" ||
			currency == "RUB" || currency == "rub" ||
			currency == "EUR" || currency == "eur" {
			return currency
		}
		fmt.Println("Ошибка: неправильно введена исходная валюта. Попробуйте снова.")
		fmt.Println()
	}
}

// Функция для ввода и проверки суммы
func getAmount() float64 {
	var amount float64

	for {
		fmt.Print("Введите сумму: ")
		_, err := fmt.Scan(&amount)
		if err == nil && amount > 0 {
			return amount
		}
		fmt.Println("Ошибка: некорректно введено количество валюты. Попробуйте снова.")
		fmt.Println()
	}
}

// Функция для ввода и проверки целевой валюты
func getToCurrency(fromCurrency string) string {
	var currency string

	for {
		// Подсказываем доступные варианты в зависимости от исходной валюты
		switch fromCurrency {
		case "RUB", "rub":
			fmt.Println("Доступные валюты: USD, EUR")
			fmt.Print("Введите целевую валюту: ")
		case "EUR", "eur":
			fmt.Println("Доступные валюты: USD, RUB")
			fmt.Print("Введите целевую валюту: ")
		case "USD", "usd":
			fmt.Println("Доступные валюты: RUB, EUR")
			fmt.Print("Введите целевую валюту: ")
		}

		fmt.Scan(&currency)

		// Проверяем, что целевая валюта корректна и отличается от исходной
		validCurrency := false
		if currency == "USD" || currency == "usd" ||
			currency == "RUB" || currency == "rub" ||
			currency == "EUR" || currency == "eur" {
			validCurrency = true
		}

		// Проверяем, что не выбрана та же самая валюта
		sameCurrency := false
		if (fromCurrency == "USD" || fromCurrency == "usd") && (currency == "USD" || currency == "usd") {
			sameCurrency = true
		} else if (fromCurrency == "RUB" || fromCurrency == "rub") && (currency == "RUB" || currency == "rub") {
			sameCurrency = true
		} else if (fromCurrency == "EUR" || fromCurrency == "eur") && (currency == "EUR" || currency == "eur") {
			sameCurrency = true
		}

		if validCurrency && !sameCurrency {
			return currency
		}
		fmt.Println("Ошибка: неправильно введена целевая валюта. Попробуйте снова.")
		fmt.Println()
	}
}

// Основная функция для получения всех входных данных
func getUserInput() (float64, string, string) {
	fmt.Println("=== Шаг 1: Выбор исходной валюты ===")
	fromCurrency := getFromCurrency()

	fmt.Println("\n=== Шаг 2: Ввод суммы ===")
	amount := getAmount()

	fmt.Println("\n=== Шаг 3: Выбор целевой валюты ===")
	toCurrency := getToCurrency(fromCurrency)

	return amount, fromCurrency, toCurrency
}

func checkRepeat() bool {
	var answer string
	fmt.Print("\nХотите выполнить еще одну конвертацию? (yes/no): ")
	fmt.Scan(&answer)
	return answer == "yes" || answer == "Yes" || answer == "YES"
}
