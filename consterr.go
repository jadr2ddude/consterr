//Package consterr is another way of implementing
package consterr

//Error is another way to create errors
type Error string

func (e Error) Error() string {
	return string(e)
}
