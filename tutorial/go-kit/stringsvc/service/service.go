package service

import (
	"go.guide/tutorial/go-kit/stringsvc/err"
	"strings"
)

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type StringServiceImpl struct{}

func (StringServiceImpl) Uppercase(s string) (string, error) {
	if s == "" {
		return "", err.ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (StringServiceImpl) Count(s string) int {
	return len(s)
}
