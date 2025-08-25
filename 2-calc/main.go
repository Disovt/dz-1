package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    
    fmt.Print("Введите операцию (SUM, AVG, MED): ")
    operation, _ := reader.ReadString('\n')
    operation = strings.TrimSpace(operation)
    
    fmt.Print("Введите числа через запятую: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    numbers, err := parseNumbers(input)
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }
    
    switch strings.ToUpper(operation) {
    case "SUM":
        fmt.Println("Сумма:", sum(numbers))
    case "AVG":
        fmt.Printf("Среднее: %.2f\n", avg(numbers))
    case "MED":
        fmt.Printf("Медиана: %.2f\n", med(numbers))
    default:
        fmt.Println("Неизвестная операция:", operation)
    }
}

func parseNumbers(input string) ([]int, error) {
    // Удаляем все пробелы из строки
    input = strings.ReplaceAll(input, " ", "")
    // Разбиваем по запятым
    parts := strings.Split(input, ",")
    var numbers []int
    
    for _, part := range parts {
        // Пропускаем пустые части
        if part == "" {
            continue
        }
        
        num, err := strconv.Atoi(part)
        if err != nil {
            return nil, fmt.Errorf("не удалось преобразовать '%s' в число", part)
        }
        numbers = append(numbers, num)
    }
    
    return numbers, nil
}

func sum(numbers []int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

func avg(numbers []int) float64 {
    if len(numbers) == 0 {
        return 0
    }
    return float64(sum(numbers)) / float64(len(numbers))
}

func med(numbers []int) float64 {
    if len(numbers) == 0 {
        return 0
    }
    
    // Создаем копию массива для сортировки
    sorted := make([]int, len(numbers))
    copy(sorted, numbers)
    sort.Ints(sorted)
    
    n := len(sorted)
    if n%2 == 1 {
        return float64(sorted[n/2])
    }
    return float64(sorted[n/2-1]+sorted[n/2]) / 2.0
}