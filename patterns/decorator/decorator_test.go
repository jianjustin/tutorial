package decorator_test

import (
	"log"
	"testing"

	"go.guide/patterns/decorator"
)

func TestLogDecorate(t *testing.T) {
	t.Parallel()

	decorator := decorator.LogDecorate(func(taskID int) int {
		log.Printf("Task with ID %v is running....", taskID)
		return 0
	})

	decorator(5)
}
