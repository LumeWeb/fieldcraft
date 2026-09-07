package fieldcraft

import "errors"

// ErrUnresolvedField is returned (wrapped by FieldError) when a required field
// has no resolved value and the run cannot prompt for one.
var ErrUnresolvedField = errors.New("fieldcraft: required field is unresolved in non-interactive mode")
