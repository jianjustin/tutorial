package service

import (
	"errors"
)

var ErrEmpty = errors.New("empty string")

type MulService interface {
	Mul(int) (int, error)
}

type BaseMulService struct{}

func (BaseMulService) Mul(a int) (int, error) {
	return a * 3, nil
}
