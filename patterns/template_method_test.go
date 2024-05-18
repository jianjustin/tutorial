package patterns_test

import (
	templatemethod "go.guide/patterns"
	"testing"
)

func TestTemplateMethod(t *testing.T) {
	worker := templatemethod.NewWorker(&templatemethod.PostMan{})
	worker.DailyRoutine()
}
