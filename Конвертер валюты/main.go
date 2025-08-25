package main

import (
	"fmt"
)

// Определяем тип для валюты
type Currency string

// Создаем map для хранения курсов валют относительно USD
var rates = map[Currency]float64{
	"USD": 1.0,   // 1 USD = 1 USD
	"RUB": 80.35, // 1 USD = 80.35 RUB
	"EUR": 0.85,  // 1 USD = 0.85 EUR
}

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

func calculate(amount float64, from Currency, to Currency) float64 {
	// Конвертируем через USD как базовую валюту
	var amountInUSD float64

	// Сначала переводим в USD, используя map
	amountInUSD = amount / rates[from]

	// Затем конвертируем из USD в целевую валюту, используя map
	return amountInUSD * rates[to]
}

// Функция для ввода и проверки исходной валюты
func getFromCurrency() Currency {
	var input string

	for {
		fmt.Println("Доступные валюты: USD, RUB, EUR")
		fmt.Print("Введите исходную валюту: ")
		fmt.Scan(&input)

		// Преобразуем ввод в тип Currency
		currency := Currency(input)

		// Проверяем корректность ввода через map
		if _, exists := rates[currency]; exists {
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
func getToCurrency(fromCurrency Currency) Currency {
	var input string

	for {
		// Подсказываем доступные варианты в зависимости от исходной валюты
		switch fromCurrency {
		case "RUB":
			fmt.Println("Доступные валюты: USD, EUR")
		case "EUR":
			fmt.Println("Доступные валюты: USD, RUB")
		case "USD":
			fmt.Println("Доступные валюты: RUB, EUR")
		}
		
		fmt.Print("Введите целевую валюту: ")
		fmt.Scan(&input)

		// Преобразуем ввод в тип Currency
		currency := Currency(input)

		// Проверяем, что целевая валюта корректна и отличается от исходной
		if _, exists := rates[currency]; exists && currency != fromCurrency {
			return currency
		}
		fmt.Println("Ошибка: неправильно введена целевая валюта. Попробуйте снова.")
		fmt.Println()
	}
}

// Основная функция для получения всех входных данных
func getUserInput() (float64, Currency, Currency) {
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