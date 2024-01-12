package main

import (
	"database/sql"
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	_, _ = sql.Open("postgres", "")
}
