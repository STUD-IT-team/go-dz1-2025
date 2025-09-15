package task1

import (
	"strconv"
)

// Определяем ошибки

func FilterCommonDigits(a, b int) (int, int, error) {
	// Проверка на отрицательные числа
	if a < 0 || b < 0 {
		return 0, 0, ErrNegNums
	}

	// Преобразуем числа в строки для работы с цифрами
	strA := strconv.Itoa(a)
	strB := strconv.Itoa(b)

	// Создаем множества цифр для каждого числа
	digitsA := make(map[rune]bool)
	digitsB := make(map[rune]bool)

	// Заполняем множества
	for _, digit := range strA {
		digitsA[digit] = true
	}
	for _, digit := range strB {
		digitsB[digit] = true
	}

	// Находим общие цифры
	commonDigits := make(map[rune]bool)
	for digit := range digitsA {
		if digitsB[digit] {
			commonDigits[digit] = true
		}
	}

	// Фильтруем цифры из первого числа
	filteredA := ""
	for _, digit := range strA {
		if !commonDigits[digit] {
			filteredA += string(digit)
		}
	}

	// Фильтруем цифры из второго числа
	filteredB := ""
	for _, digit := range strB {
		if !commonDigits[digit] {
			filteredB += string(digit)
		}
	}

	// Проверка на пустые результаты
	if filteredA == "" || filteredB == "" {
		return 0, 0, ErrEmptyNum
	}

	// Преобразуем обратно в числа
	resultA, _ := strconv.Atoi(filteredA)
	resultB, _ := strconv.Atoi(filteredB)

	return resultA, resultB, nil
}
