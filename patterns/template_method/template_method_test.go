package templatemethod_test

import (
	"testing"

	templatemethod "go.guide/patterns/template_method"
)

func TestTemplateMethod(t *testing.T) {
	worker := templatemethod.NewWorker(&templatemethod.PostMan{})
	worker.DailyRoutine()
}
