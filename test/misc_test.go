package test

import (
	"fmt"
	"test_custom_db_pool/db"
	"testing"
)

func TestR(t *testing.T) {
	db.InitConnPool()
	dbConn, dbConnErr := db.AcquireFromConnPool()
	if dbConnErr != nil {
		fmt.Println(fmt.Sprintf("Failed to Acquire Connection: %s", dbConnErr))
		return
	}
	var sum int
	resErr := dbConn.QueryRow("select 1 + 1 as sum").Scan(&sum)
	if resErr != nil {
		fmt.Println(fmt.Sprintf("Error Occurred During Query: %s", resErr))
		return
	}
	fmt.Println(fmt.Sprintf("Sum Successfully Fetched: %d", sum))
}
