package service

import (
	"errors"
)

var ErrEmpty = errors.New("empty string")

type StringService interface {
	Mul(int) (int, error)
}

type BaseStringService struct{}

func (BaseStringService) Mul(a int) (int, error) {
	return a * 3, nil
}
