package driver

import (
	"database/sql"
	"os"
	"time"
)

//DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}

var dbconn = &DB{}

const naxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute

func ConnectSQL() (*DB, error) {
	connectionString := "host=ms-db port=5432 dbname=" + os.Getenv("POSTGRES_DB") + " user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD")
	db, err := newDatabase(connectionString)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(naxOpenDbConn)
	db.SetMaxIdleConns(maxIdleDbConn)
	db.SetConnMaxLifetime(maxDbLifetime)
	dbconn.SQL = db
	err = testDb(db)
	if err != nil {
		return nil, err
	}
	return dbconn, nil
}

//try to ping database
func testDb(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}
	return nil
}

//Create new database
func newDatabase(dns string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dns)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
