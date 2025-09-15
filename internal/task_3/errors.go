package task3

import "errors"

var (
	ErrOverflow = errors.New("resulting slice size exceeds uint32 maximum")
)
