package main

import (
	"fmt"
	"strconv"
	"strings"
	"math"
	"flag"
	"sort"
)

func greaterThan(a, b float64) bool {
	return a - b > 1e-9
}

func inputNumbers() []float64 {
	input := ""
	var mySlice []float64
	fmt.Println("Введите элементы построчно. Для завершения введите 'stop' :")
	for {
		fmt.Scan(&input)
		input = strings.TrimSpace(input)
		if input == "stop" {
			break
		}
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Ошибка преобразования в число:", err)
			continue
		}
		// округляем значение
		formattedNum := math.Round(num * 100) / 100
		if formattedNum > 100000 || formattedNum < -100000 {
			fmt.Println("Число выходит за границы от -100 000 до 100 000")
			continue
		}
		mySlice = append(mySlice, formattedNum)
	}
	return mySlice
}


func calculateMean(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0.0
	}
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	mean := sum / float64(len(numbers))

	return math.Round(mean * 100) / 100
}

func calculateMedian(numbers []float64) float64 {
	n := len(numbers)
    if n == 0 {
        return 0.0
    }
    sort.Slice(numbers, func(i, j int) bool {
        return numbers[i] < numbers[j]
    })
    if n % 2 == 1 {
        return math.Round(numbers[n/2] * 100) / 100
    } else { // Четное количество элементов
        mid1 := numbers[n/2-1]
        mid2 := numbers[n/2]
        average := (mid1 + mid2) / 2
        return math.Round(average * 100) / 100
    }
}


func calculateMode(numbers []float64) float64 {
	frequency := make(map[float64]float64)

	for _, num := range numbers {
		frequency[num]++
	}

	maxFrequency := 0.0
	mode := 0.0

	for num, freq := range frequency {
		if greaterThan(freq, maxFrequency) {
				maxFrequency = freq
				mode = num
			}
		}
	return mode
}

func regularStandardDeviation(numbers []float64) float64 {
	n := len(numbers)
	if n <= 1 {
		return 0.0
	}
  	mean := calculateMean(numbers)
	diffSquares := make([]float64, n)	

	for i, num := range numbers {
		diff := num - mean
		diffSquares[i] = diff * diff
	}

	variance := calculateMean(diffSquares)
	regularStandardDev := math.Sqrt(variance)
	return math.Round(regularStandardDev * 100) / 100
}

func main() {
	meanFlag := flag.Bool("mean", false, "Вывод среднего значения")
	medianFlag := flag.Bool("median", false, "Вывод медианы")
	modeFlag := flag.Bool("mode", false, "Вывод наиболее встречающего элемента")
	sdFlag := flag.Bool("sd", false, "Вывод стандратного отклоннения")

	flag.Parse()
	result := inputNumbers()

	if *meanFlag {
		mean := calculateMean(result)
		fmt.Println("Mean:", mean)
	}
	if *medianFlag {
		median := calculateMedian(result)
		fmt.Println("Median:", median)
	}
	if *modeFlag {
		mode := calculateMode(result)
		fmt.Println("Mode:", mode)
	}
	if *sdFlag {
		sd := regularStandardDeviation(result)
		fmt.Println("SD:", sd)
	}
}