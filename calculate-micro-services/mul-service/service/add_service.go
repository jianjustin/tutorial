package service

type AddServiceMiddleware func(AddService) AddService

type AddService interface {
	Add(int) (int, error)
}
