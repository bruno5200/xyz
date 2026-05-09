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
// - `DATABASE_URL`: A full database connection URL (optional, takes precedence over individual variables).
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
		connStr := connection()
		var err error

		if db, err = sql.Open("postgres", connStr); err != nil {
			log.Printf("Error connecting to database %s", err)
		}

		if err := db.Ping(); err != nil {
			log.Panicf("Error pinging database %s", err)
		}

		fmt.Println("Connected to postgres!")
	})
}

func connection() string {
	e.initDB()
	muDB.Lock()
	defer muDB.Unlock()

	if e.getDBURL() != "" {
		return e.getDBURL()
	}

	return fmt.Sprintf("user=%s password=%s host=%s dbname=%s port=%d sslmode=%s", e.getDBUser(), e.getDBPass(), e.getDBHost(), e.getDBName(), e.getDBPort(), e.getDBSSL())
}

func PostgresDB() *sql.DB {
	return db
}

func Close() {
	db.Close()
}
