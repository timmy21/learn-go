package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type SyntaxError struct {
	Num string
}

func (e *SyntaxError) Error() string {
	return "parsing " + strconv.Quote(e.Num) + ": invalid syntax"
}

func implements() {
	typeOfError := reflect.TypeOf((*error)(nil)).Elem()
	sErrPtr := reflect.TypeOf(&SyntaxError{Num: "a1"})
	sErr := reflect.TypeOf(SyntaxError{Num: "a1"})

	fmt.Println("*SyntaxError implements error:", sErrPtr.Implements(typeOfError))
	fmt.Println("SyntaxError implements error:", sErr.Implements(typeOfError))
}

func main() {
	implements()
}
