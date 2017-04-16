//Package consterr is another way of implementing constant errors in Go
package consterr

//Error is another way to create errors
type Error string

//Error implements the error interface
func (e Error) Error() string {
	return string(e)
}
