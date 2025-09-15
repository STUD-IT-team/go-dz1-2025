package task1

import "errors"

var (
	ErrNegNums  = errors.New("negative numbers are not allowed")
	ErrEmptyNum = errors.New("resulting number is empty")
)
