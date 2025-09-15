package tests

import (
	task2 "dz1/internal/task_2"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindCommonWords(t *testing.T) {
	t.Parallel()

	t.Run("Обычный случай с существующими файлами", func(t *testing.T) {
		t.Parallel()
		// Given
		tempDir := t.TempDir()

		file1 := filepath.Join(tempDir, "file1.txt")
		file2 := filepath.Join(tempDir, "file2.txt")
		file3 := filepath.Join(tempDir, "file3.txt")
		outputFile := filepath.Join(tempDir, "result.txt")

		require.NoError(t, os.WriteFile(file1, []byte("hello world go programming"), 0644))
		require.NoError(t, os.WriteFile(file2, []byte("world test code go"), 0644))
		require.NoError(t, os.WriteFile(file3, []byte("go language world example"), 0644))

		// When
		err := task2.FindCommonWords(outputFile, file1, file2, file3)

		// Then
		require.NoError(t, err)

		content, err := os.ReadFile(outputFile)
		require.NoError(t, err)

		result := string(content)
		assert.Contains(t, result, "world")
		assert.Contains(t, result, "go")
		assert.NotContains(t, result, "hello")
		assert.NotContains(t, result, "test")
	})

	t.Run("Несуществующий файл", func(t *testing.T) {
		t.Parallel()
		// Given
		tempDir := t.TempDir()

		file1 := filepath.Join(tempDir, "file1.txt")
		nonexistent := filepath.Join(tempDir, "nonexistent.txt")
		outputFile := filepath.Join(tempDir, "result.txt")

		require.NoError(t, os.WriteFile(file1, []byte("hello world"), 0644))

		// When
		err := task2.FindCommonWords(outputFile, file1, nonexistent)

		// Then
		require.Error(t, err)
		assert.Equal(t, task2.ErrOpenFile, err)
	})

	t.Run("Пустой список файлов", func(t *testing.T) {
		t.Parallel()
		// Given
		tempDir := t.TempDir()
		outputFile := filepath.Join(tempDir, "result.txt")

		// When
		err := task2.FindCommonWords(outputFile)

		// Then
		require.NoError(t, err)
	})

	t.Run("Файлы без общих слов", func(t *testing.T) {
		t.Parallel()
		// Given
		tempDir := t.TempDir()

		file1 := filepath.Join(tempDir, "file1.txt")
		file2 := filepath.Join(tempDir, "file2.txt")
		outputFile := filepath.Join(tempDir, "result.txt")

		require.NoError(t, os.WriteFile(file1, []byte("python java c++"), 0644))
		require.NoError(t, os.WriteFile(file2, []byte("javascript php ruby"), 0644))

		// When
		err := task2.FindCommonWords(outputFile, file1, file2)

		// Then
		require.NoError(t, err)

		content, err := os.ReadFile(outputFile)
		require.NoError(t, err)
		assert.Empty(t, string(content))
	})

	t.Run("Один файл", func(t *testing.T) {
		t.Parallel()
		// Given
		tempDir := t.TempDir()

		file1 := filepath.Join(tempDir, "single.txt")
		outputFile := filepath.Join(tempDir, "result.txt")

		require.NoError(t, os.WriteFile(file1, []byte("single file example words"), 0644))

		// When
		err := task2.FindCommonWords(outputFile, file1)

		// Then
		require.NoError(t, err)

		content, err := os.ReadFile(outputFile)
		require.NoError(t, err)

		result := string(content)
		assert.Contains(t, result, "single")
		assert.Contains(t, result, "file")
		assert.Contains(t, result, "example")
		assert.Contains(t, result, "words")
	})

	t.Run("Пустые файлы", func(t *testing.T) {
		t.Parallel()
		// Given
		tempDir := t.TempDir()

		file1 := filepath.Join(tempDir, "empty1.txt")
		file2 := filepath.Join(tempDir, "empty2.txt")
		outputFile := filepath.Join(tempDir, "result.txt")

		require.NoError(t, os.WriteFile(file1, []byte(""), 0644))
		require.NoError(t, os.WriteFile(file2, []byte("   \n\n\t  "), 0644))

		// When
		err := task2.FindCommonWords(outputFile, file1, file2)

		// Then
		require.NoError(t, err)

		content, err := os.ReadFile(outputFile)
		require.NoError(t, err)
		assert.Empty(t, string(content))
	})

	t.Run("Файлы с одинаковым содержимым", func(t *testing.T) {
		t.Parallel()
		// Given
		tempDir := t.TempDir()

		file1 := filepath.Join(tempDir, "file1.txt")
		file2 := filepath.Join(tempDir, "file2.txt")
		outputFile := filepath.Join(tempDir, "result.txt")

		require.NoError(t, os.WriteFile(file1, []byte("apple banana cherry"), 0644))
		require.NoError(t, os.WriteFile(file2, []byte("apple banana cherry"), 0644))

		// When
		err := task2.FindCommonWords(outputFile, file1, file2)

		// Then
		require.NoError(t, err)

		content, err := os.ReadFile(outputFile)
		require.NoError(t, err)

		result := string(content)
		assert.Contains(t, result, "apple")
		assert.Contains(t, result, "banana")
		assert.Contains(t, result, "cherry")
	})

	t.Run("Учет регистра", func(t *testing.T) {
		t.Parallel()
		// Given
		tempDir := t.TempDir()

		file1 := filepath.Join(tempDir, "file1.txt")
		file2 := filepath.Join(tempDir, "file2.txt")
		outputFile := filepath.Join(tempDir, "result.txt")

		require.NoError(t, os.WriteFile(file1, []byte("Hello World"), 0644))
		require.NoError(t, os.WriteFile(file2, []byte("hello world"), 0644))

		// When
		err := task2.FindCommonWords(outputFile, file1, file2)

		// Then
		require.NoError(t, err)

		content, err := os.ReadFile(outputFile)
		require.NoError(t, err)
		assert.Empty(t, string(content))
	})

	t.Run("Ошибка записи в выходной файл", func(t *testing.T) {
		t.Parallel()
		// Given
		tempDir := t.TempDir()

		file1 := filepath.Join(tempDir, "file1.txt")
		invalidOutput := filepath.Join(tempDir, "invalid_dir", "result.txt")

		require.NoError(t, os.WriteFile(file1, []byte("test"), 0644))

		// When
		err := task2.FindCommonWords(invalidOutput, file1)

		// Then
		require.Error(t, err)
		assert.Equal(t, task2.ErrOpenFile, err)
	})
}
