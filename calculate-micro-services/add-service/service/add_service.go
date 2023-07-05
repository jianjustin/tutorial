package service

import (
	"errors"
	"strings"
)

var ErrEmpty = errors.New("empty string")

type StringService interface {
	Uppercase(string) (string, error)
	Add(int) (int, error)
}

type BaseStringService struct{}

func (BaseStringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (BaseStringService) Add(a int) (int, error) {
	return a + 3, nil
}
