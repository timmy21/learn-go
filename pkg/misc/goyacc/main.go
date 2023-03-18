package main

import (
	"fmt"

	"github.com/timmy21/learn-go/pkg/misc/goyacc/sql"
)

func main() {
	stmt, err := sql.Parse("select a, b from log")
	fmt.Println(err)
	fmt.Println(stmt)
}
