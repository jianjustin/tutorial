package err

import "errors"

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("empty string")
