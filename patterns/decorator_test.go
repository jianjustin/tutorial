package patterns_test

import (
	"go.guide/patterns"
	"log"
	"testing"
)

func TestLogDecorate(t *testing.T) {
	t.Parallel()

	decorator := patterns.LogDecorate(func(taskID int) int {
		log.Printf("Task with ID %v is running....", taskID)
		return 0
	})

	decorator(5)
}
