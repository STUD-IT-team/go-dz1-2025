package tests

import (
	task3 "dz1/internal/task_3"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScaleSlice(t *testing.T) {
	t.Parallel()

	t.Run("Нормальное масштабирование", func(t *testing.T) {
		t.Parallel()
		// Given
		slice := []int{1, 2, 3}

		// When
		err := task3.ScaleSlice(&slice, 3)

		// Then
		require.NoError(t, err)
		assert.Equal(t, []int{1, 2, 3, 1, 2, 3, 1, 2, 3}, slice)
	})

	t.Run("Переполнение - ошибка", func(t *testing.T) {
		t.Parallel()
		// Given
		slice := make([]int, 1000000) // 1,000,000 элементов

		// When
		err := task3.ScaleSlice(&slice, 5000) // 5,000,000,000 > 4,294,967,295

		// Then
		require.Error(t, err)
		assert.Equal(t, task3.ErrOverflow, err)
	})

	t.Run("Коэффициент 0 - очистка среза", func(t *testing.T) {
		t.Parallel()
		// Given
		slice := []int{1, 2, 3}

		// When
		err := task3.ScaleSlice(&slice, 0)

		// Then
		require.NoError(t, err)
		assert.Empty(t, slice)
	})

	t.Run("Пустой срез", func(t *testing.T) {
		t.Parallel()
		// Given
		slice := []int{}

		// When
		err := task3.ScaleSlice(&slice, 5)

		// Then
		require.NoError(t, err)
		assert.Empty(t, slice)
	})

	t.Run("Коэффициент 1 - без изменений", func(t *testing.T) {
		t.Parallel()
		// Given
		original := []int{1, 2, 3}
		slice := make([]int, len(original))
		copy(slice, original)

		// When
		err := task3.ScaleSlice(&slice, 1)

		// Then
		require.NoError(t, err)
		assert.Equal(t, original, slice)
	})

	t.Run("nil срез", func(t *testing.T) {
		t.Parallel()
		// When
		err := task3.ScaleSlice(nil, 5)

		// Then
		require.NoError(t, err)
	})
}
