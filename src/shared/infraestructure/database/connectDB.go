package database

import (
	"database/sql"
	"os"
	"sync"

	"github.com/getsentry/sentry-go"
	_ "github.com/go-sql-driver/mysql"
)

type connectMySQL struct {
	conn *sql.DB
}

func (cm *connectMySQL) Conn() *sql.DB {
	return cm.conn
}

var lock = &sync.Mutex{}
var instance *connectMySQL

var user string = os.Getenv("DB_USER")
var pass string = os.Getenv("DB_PASS")
var host string = os.Getenv("DB_HOST")
var port string = os.Getenv("DB_PORT")
var database string = os.Getenv("DB_NAME")

func InstanceDB() *connectMySQL {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			strConnect := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + database
			conn, err := sql.Open("mysql", strConnect)
			if err != nil {
				sentry.CaptureException(err)
			}

			instance = &connectMySQL{
				conn: conn,
			}
		}
	}

	return instance
}
