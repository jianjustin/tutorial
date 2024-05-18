package patterns

import (
	"testing"
)

func TestHandle_SecondHanlderHandlesRequest(t *testing.T) {
	t.Parallel()
	barry := NewHandler("Barry", nil, 2)
	paul := NewHandler("Paul", barry, 1)
	paul.Handle()

}
