package task2

import (
	"bufio"
	"os"
	"strings"
)

func FindCommonWords(outputFilename string, inputFilenames ...string) error {
	if len(inputFilenames) == 0 {
		return nil
	}

	// Создаем слайс для хранения множеств слов из каждого файла
	wordSets := make([]map[string]bool, len(inputFilenames))

	// Читаем слова из каждого файла
	for i, filename := range inputFilenames {
		file, err := os.Open(filename)
		if err != nil {
			return ErrOpenFile
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		wordSet := make(map[string]bool)
		for scanner.Scan() {
			word := strings.TrimSpace(scanner.Text())
			if word != "" {
				wordSet[word] = true
			}
		}

		if err := scanner.Err(); err != nil {
			return ErrOpenFile
		}

		wordSets[i] = wordSet
	}

	// Находим пересечение всех множеств
	commonWords := findIntersection(wordSets)

	// Записываем результат в выходной файл
	outputFile, err := os.Create(outputFilename)
	if err != nil {
		return ErrOpenFile
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for word := range commonWords {
		_, err := writer.WriteString(word + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

// Вспомогательная функция для нахождения пересечения множеств
func findIntersection(wordSets []map[string]bool) map[string]bool {
	if len(wordSets) == 0 {
		return make(map[string]bool)
	}

	// Начинаем с первого множества
	result := make(map[string]bool)
	for word := range wordSets[0] {
		result[word] = true
	}

	// Постепенно находим пересечение с остальными множествами
	for i := 1; i < len(wordSets); i++ {
		temp := make(map[string]bool)
		for word := range result {
			if wordSets[i][word] {
				temp[word] = true
			}
		}
		result = temp
	}

	return result
}
