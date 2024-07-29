package domain

import "errors"

var (
	ErrNotExist          = errors.New("row does not exist")
	ErrUpdateFailed      = errors.New("update failed")
	ErrDeleteFailed      = errors.New("delete failed")
	ErrInvalidId         = errors.New("invalid id")
	ErrTaskNotFound      = errors.New("task with such credentials not found")
	ErrTaskAlreadyExists = errors.New("task already exists")
)
