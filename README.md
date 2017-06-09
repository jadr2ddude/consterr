# consterr [![GoDoc](https://godoc.org/github.com/panux/consterr?status.svg)](https://godoc.org/github.com/panux/consterr)[![Build Status](https://travis-ci.org/panux/consterr.svg?branch=master)](https://travis-ci.org/panux/consterr)[![Coverage Status](https://coveralls.io/repos/github/panux/consterr/badge.svg?branch=master)](https://coveralls.io/github/panux/consterr?branch=master)
A better way to implement constant errors in Go

## Usage
Simply import the package and use consterr.Error as the type for your constants, assigning a string value. For example:
```
import "github.com/panux/consterr"
const (
	example consterr.Error = "Example error"
)
```
