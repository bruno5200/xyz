package util

import "fmt"

// Annotate: adds context to an error, if it is not nil.
//
// Example usage:
//	err := someFunction(input)
//	if err != nil {
//		return util.Annotate(err, "someFunction failed with input %d", input)
//	}
func Annotate(err error, format string, args ...any) (result error) {
	if err != nil {
		result = fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), err)
	}
	return
}
