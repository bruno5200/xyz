package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	// is a pointer to the database connection
	db   *sql.DB
	once sync.Once

	user     string
	password string
	host     string
	name     string
	port     int
)

// Connect creates a connection to database
//
// if the connection fails, the application will panic
//
// # Connect should be called only once
//
// # Connect is thread safe
//
// Connect uses the following environment variables:
//
// - `DB_HOST`: Host of postgres DB.
//
// - `DB_PORT`: Poert of postgres DB.
//
// - `DB_USER`: User of postgres DB.
//
// - `DB_PASSWORD`: Password of postgres DB.
//
// - `DB_NAME`: Name of postgres DB.
//
// - `DB_SSL`: SSL mode of postgres DB. (optional, default: disable)
func NewConnection() {

	once.Do(func() {

		wg := sync.WaitGroup{}
		wg.Add(1)
		go e.initDB(&wg)
		wg.Wait()

		connStr := connection()

		var err error

		if db, err = sql.Open("postgres", connStr); err != nil {
			log.Printf("Error connecting to database %s", err)
		}

		if err := db.Ping(); err != nil {
			log.Panicf("Error pinging database %s", err)
		} else {
			fmt.Println("Connected to postgres!")
		}
	})
}

func connection() string {

	muDB.Lock()
	host = e.getDBHost()
	port = e.getDBPort()
	name = e.getDBName()
	user = e.getDBUser()
	password = e.getDBPass()
	muDB.Unlock()

	return fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=disable", user, password, host, name, port)
}

func PostgresDB() *sql.DB {
	return db
}

func Close() {
	db.Close()
}
