package task3

// ScaleSlice увеличивает размер среза в scaleFactor раз, заполняя новые элементы копиями исходных данных
// Возвращает ErrOverflow если размер результирующего среза превышает максимальное значение uint32
func ScaleSlice(slice *[]int, scaleFactor uint32) error {
	if slice == nil {
		return nil
	}

	originalLen := len(*slice)

	// Проверка на переполнение
	if originalLen > 0 && scaleFactor > 0 {
		// Преобразуем в uint64 чтобы избежать переполнения при умножении
		resultSize := uint64(originalLen) * uint64(scaleFactor)
		if resultSize > 1<<32-1 { // 1<<32-1 = 4294967295 (max uint32)
			return ErrOverflow
		}
	}

	// Если scaleFactor = 0 или пустой срез - очищаем исходный срез
	if scaleFactor == 0 || originalLen == 0 {
		*slice = []int{}
		return nil
	}

	// Если scaleFactor = 1 - ничего не меняем
	if scaleFactor == 1 {
		return nil
	}

	// Создаем новый срез нужного размера
	newSize := originalLen * int(scaleFactor)
	newSlice := make([]int, newSize)

	// Заполняем новый срез копиями исходных данных
	for i := 0; i < int(scaleFactor); i++ {
		start := i * originalLen
		end := start + originalLen
		copy(newSlice[start:end], *slice)
	}

	// Заменяем исходный срез
	*slice = newSlice
	return nil
}
