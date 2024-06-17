package service

import (
	"go.guide/tutorial/go-kit/stringsvc/err"
	"strings"
)

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

// ServiceMiddleware is a chainable behavior modifier for StringService.
type ServiceMiddleware func(StringService) StringService

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
