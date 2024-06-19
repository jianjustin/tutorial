package service

import (
	"errors"
)

var ErrEmpty = errors.New("empty string")

type AddService interface {
	Add(int) (int, error)
}

type BaseAddService struct{}

func (BaseAddService) Add(a int) (int, error) {
	return a + 3, nil
}
