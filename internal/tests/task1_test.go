package tests

import (
	task1 "dz1/internal/task_1"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFilterCommonDigits(t *testing.T) {
	t.Parallel()

	t.Run("Обычный случай: нет общих цифр", func(t *testing.T) {
		t.Parallel()
		// Given
		a, b := 123, 456

		// When
		result1, result2, err := task1.FilterCommonDigits(a, b)

		// Then
		require.NoError(t, err)
		assert.Equal(t, 123, result1)
		assert.Equal(t, 456, result2)
	})

	t.Run("Есть общие цифры", func(t *testing.T) {
		t.Parallel()
		// Given
		a, b := 12345, 56789

		// When
		result1, result2, err := task1.FilterCommonDigits(a, b)

		// Then
		require.NoError(t, err)
		assert.Equal(t, 1234, result1)
		assert.Equal(t, 6789, result2)
	})

	t.Run("Все цифры общие - ошибка EmptyNum", func(t *testing.T) {
		t.Parallel()
		// Given
		a, b := 111, 111

		// When
		result1, result2, err := task1.FilterCommonDigits(a, b)

		// Then
		require.Error(t, err)
		assert.Equal(t, task1.ErrEmptyNum, err)
		assert.Equal(t, 0, result1)
		assert.Equal(t, 0, result2)
	})

	t.Run("Отрицательные числа - ошибка NegNums", func(t *testing.T) {
		t.Parallel()
		// Given
		a, b := -123, 456

		// When
		result1, result2, err := task1.FilterCommonDigits(a, b)

		// Then
		require.Error(t, err)
		assert.Equal(t, task1.ErrNegNums, err)
		assert.Equal(t, 0, result1)
		assert.Equal(t, 0, result2)
	})

	t.Run("Нулевые значения", func(t *testing.T) {
		t.Parallel()
		// Given
		a, b := 0, 123

		// When
		result1, result2, err := task1.FilterCommonDigits(a, b)

		// Then
		require.NoError(t, err)
		assert.Equal(t, 0, result1)
		assert.Equal(t, 123, result2)
	})

	t.Run("Большие числа с общими цифрами - ошибка EmptyNum", func(t *testing.T) {
		t.Parallel()
		// Given
		a, b := 987654321, 123456789

		// When
		result1, result2, err := task1.FilterCommonDigits(a, b)

		// Then
		require.Error(t, err)
		assert.Equal(t, task1.ErrEmptyNum, err)
		assert.Equal(t, 0, result1)
		assert.Equal(t, 0, result2)
	})

	t.Run("Оба числа отрицательные", func(t *testing.T) {
		t.Parallel()
		// Given
		a, b := -123, -456

		// When
		result1, result2, err := task1.FilterCommonDigits(a, b)

		// Then
		require.Error(t, err)
		assert.Equal(t, task1.ErrNegNums, err)
		assert.Equal(t, 0, result1)
		assert.Equal(t, 0, result2)
	})

	t.Run("Частично общие цифры", func(t *testing.T) {
		t.Parallel()
		// Given
		a, b := 123456, 456789

		// When
		result1, result2, err := task1.FilterCommonDigits(a, b)

		// Then
		require.NoError(t, err)
		assert.Equal(t, 123, result1)
		assert.Equal(t, 789, result2)
	})
}
