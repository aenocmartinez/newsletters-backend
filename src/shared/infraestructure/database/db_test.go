package database

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnDb(t *testing.T) {

	db := InstanceDB()
	conn := db.Conn()
	err := conn.Ping()
	if err == nil {
		t.Fatal(err)
	}
}
