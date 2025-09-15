package main

import (
	task1 "dz1/internal/task_1"
	task2 "dz1/internal/task_2"
	task3 "dz1/internal/task_3"
	"fmt"
	"os"
	"path/filepath"
)

func example_1() {
	// Пример 1: нормальная работа
	result1, result2, err := task1.FilterCommonDigits(123, 456)
	if err != nil {
		println("Ошибка:", err.Error())
	} else {
		println("Результат:", result1, result2) // 123, 456
	}

	// Пример 2: общие цифры
	result1, result2, err = task1.FilterCommonDigits(123, 345)
	if err != nil {
		println("Ошибка:", err.Error())
	} else {
		println("Результат:", result1, result2) // 12, 45
	}

	// Пример 3: отрицательные числа
	_, _, err = task1.FilterCommonDigits(-123, 456)
	if err != nil {
		println("Ошибка (ожидаемая):", err.Error()) // negative numbers are not allowed
	}

	// Пример 4: пустой результат
	_, _, err = task1.FilterCommonDigits(111, 111)
	if err != nil {
		println("Ошибка (ожидаемая):", err.Error()) // resulting number is empty
	}
}

func createTestFiles() error {
	// Создаем директорию если её нет
	dir := "./internal/task_2/files/"
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	// Создаем тестовые файлы
	files := map[string]string{
		"file1.txt":  "hello world go programming test",
		"file2.txt":  "world test code go example",
		"file3.txt":  "apple banana cherry",
		"file4.txt":  "apple banana cherry",
		"file5.txt":  "python java c++",
		"file6.txt":  "javascript php ruby",
		"file7.txt":  "single file example words",
		"empty1.txt": "",
		"empty2.txt": "   \n\n\t  ",
	}

	for filename, content := range files {
		file, err := os.Create(filepath.Join(dir, filename))
		if err != nil {
			return err
		}
		_, err = file.WriteString(content)
		file.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func example_2() {
	// Задание 2
	err := createTestFiles()
	if err != nil {
		println("Ошибка создания файлов:", err.Error())
		return
	}

	// Пример 1: нормальная работа с общими словами
	err = task2.FindCommonWords("./internal/task_2/files/result1.txt",
		"./internal/task_2/files/file1.txt", "./internal/task_2/files/file2.txt")
	if err != nil {
		println("Ошибка:", err.Error())
	} else {
		println("Успешно записаны общие слова в result1.txt")
	}

	// Пример 2: все слова общие
	err = task2.FindCommonWords("./internal/task_2/files/result2.txt",
		"./internal/task_2/files/file3.txt", "./internal/task_2/files/file4.txt")
	if err != nil {
		println("Ошибка:", err.Error())
	} else {
		println("Успешно записаны все слова в result2.txt")
	}

	// Пример 3: нет общих слов
	err = task2.FindCommonWords("./internal/task_2/files/result3.txt",
		"./internal/task_2/files/file5.txt", "./internal/task_2/files/file6.txt")
	if err != nil {
		println("Ошибка:", err.Error())
	} else {
		println("Успешно создан пустой файл result3.txt")
	}

	// Пример 4: несуществующий файл
	err = task2.FindCommonWords("./internal/task_2/files/result4.txt",
		"./internal/task_2/files/nonexistent.txt", "./internal/task_2/files/file1.txt")
	if err != nil {
		println("Ошибка (ожидаемая):", err.Error()) // failed to open file
	}

	// Пример 5: только один файл
	err = task2.FindCommonWords("./internal/task_2/files/result5.txt",
		"./internal/task_2/files/file1.txt")
	if err != nil {
		println("Ошибка:", err.Error())
	} else {
		println("Успешно записаны слова из одного файла в result5.txt")
	}

	// Пример 6: без входных файлов
	err = task2.FindCommonWords("./internal/task_2/files/result6.txt")
	if err != nil {
		println("Ошибка:", err.Error())
	} else {
		println("Успешно создан пустой файл (нет входных файлов)")
	}
}

func example_3() {
	// Пример 1: Нормальное масштабирование
	slice1 := []int{1, 2, 3}
	err := task3.ScaleSlice(&slice1, 3)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("Результат: %v\n", slice1) // [1 2 3 1 2 3 1 2 3]
	}

	// Пример 2: Переполнение
	slice2 := make([]int, 1000000)        // 1,000,000 элементов
	err = task3.ScaleSlice(&slice2, 5000) // 1,000,000 * 5,000 = 5,000,000,000 > 4,294,967,295
	if err != nil {
		fmt.Printf("Ошибка (ожидаемая): %v\n", err) // ErrOverflow
	} else {
		fmt.Printf("Результат: %v\n", slice2)
	}

	// Пример 3: Коэффициент 0
	slice3 := []int{1, 2, 3}
	err = task3.ScaleSlice(&slice3, 0)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("Результат: %v\n", slice3) // []
	}

	// Пример 4: Пустой срез
	slice4 := []int{}
	err = task3.ScaleSlice(&slice4, 5)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("Результат: %v\n", slice4) // []
	}

	// Пример 5: Коэффициент 1
	slice5 := []int{1, 2, 3}
	err = task3.ScaleSlice(&slice5, 1)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("Результат: %v\n", slice5) // [1 2 3]
	}
}

// Пример использования
func main() {
	example_1()
	example_2()
	example_3()
}
