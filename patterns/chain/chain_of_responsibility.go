package chain

import (
	"fmt"
	"strconv"
)

type Handler interface {
	Handle() error
}

type handler struct {
	name     string
	next     Handler
	handleID int
}

func NewHandler(name string, next Handler, handleID int) Handler {
	return &handler{name, next, handleID}
}

func (h *handler) Handle() error {
	fmt.Println(h.name + " passed " + strconv.Itoa(h.handleID))
	if h.next != nil {
		return h.next.Handle()
	}
	return nil
}
