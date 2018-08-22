/**
 *  author: lim
 *  data  : 18-8-22 下午9:27
 */

package router

import "fmt"

// ParserError: To be deprecated.
// TODO(sougou): deprecate.
type ParserError struct {
	Message string
}

func NewParserError(format string, args ...interface{}) ParserError {
	return ParserError{fmt.Sprintf(format, args...)}
}

func (err ParserError) Error() string {
	return err.Message
}

func handleError(err *error) {
	if x := recover(); x != nil {
		*err = x.(error)
	}
}
