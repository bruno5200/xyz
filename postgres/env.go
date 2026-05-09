package postgres

import (
	"log"
	"os"
	"strconv"
	"sync"
)

var (
	e    *environmentPostgres = &environmentPostgres{}
	muDB                      = &sync.Mutex{}
)

// environmentPostgres contains the environment variables for the postgres databases
type environmentPostgres struct {
	dbURL   string
	dbURLOK bool
	dbHost  string
	dbPort  string
	dbUser  string
	dbPass  string
	dbName  string
	dbSSL   string
}

// InitDB initializes the environment variables for the postgres databases
//
// if the variables are not set, the application will panic
func (e *environmentPostgres) initDB() {
	muDB.Lock()
	defer muDB.Unlock()
	e.dbURL, e.dbURLOK = os.LookupEnv("DATABASE_URL")
	e.dbHost = e.mustLookup("DB_HOST")
	e.dbPort = e.mustLookup("DB_PORT")
	e.dbUser = e.mustLookup("DB_USER")
	e.dbPass = e.mustLookup("DB_PASSWORD")
	e.dbName = e.mustLookup("DB_NAME")
	e.dbSSL = os.Getenv("DB_SSL")
}

// `DATABASE_URL`: 
func (e *environmentPostgres) getDBURL() (val string) {
	return e.dbURL
}

// mustLookup fetches an environment variable or panics if it is missing.
func (e *environmentPostgres) mustLookup(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Panicf("PostgreSQL: %s not set", key)
	}
	return val
}

// `DB_HOST`: Host of postgres DB.
func (e *environmentPostgres) getDBHost() (val string) {
	return e.dbHost
}

// `DB_PORT`: Poert of postgres DB.
func (e *environmentPostgres) getDBPort() (val int) {
	val, err := strconv.Atoi(e.dbPort)
	if err != nil {
		log.Panicf("Env: Invalid DB_PORT %s, err: %s", e.dbPort, err)
	}
	return
}

// `DB_USER`: Usuario of postgres DB.
func (e *environmentPostgres) getDBUser() (val string) {
	return e.dbUser
}

// `DB_PASSWORD`: Contraseña of postgres DB.
func (e *environmentPostgres) getDBPass() (val string) {
	return e.dbPass
}

// `DB_NAME`: Nombre of postgres DB.
func (e *environmentPostgres) getDBName() (val string) {
	return e.dbName
}

// `DB_SSL`: SSL of postgres DB.
//
// By default is `disable`.
func (e *environmentPostgres) getDBSSL() (val string) {
	if e.dbSSL == "" {
		return "disable"
	}
	return e.dbSSL
}
