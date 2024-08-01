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
	dbHost   string
	dbHostOK bool
	dbPort   string
	dbPortOK bool
	dbUser   string
	dbUserOK bool
	dbPass   string
	dbPassOK bool
	dbName   string
	dbNameOK bool
	dbSSL    string
	dbSSLOK  bool
}

// InitDB initializes the environment variables for the postgres databases
//
// if the variables are not set, the application will panic
func (e *environmentPostgres) initDB(wg *sync.WaitGroup) {
	muDB.Lock()
	e.dbHost, e.dbHostOK = os.LookupEnv("DB_HOST")
	e.dbPort, e.dbPortOK = os.LookupEnv("DB_PORT")
	e.dbUser, e.dbUserOK = os.LookupEnv("DB_USER")
	e.dbPass, e.dbPassOK = os.LookupEnv("DB_PASSWORD")
	e.dbName, e.dbNameOK = os.LookupEnv("DB_NAME")
	e.dbSSL, e.dbSSLOK = os.LookupEnv("DB_SSL")
	muDB.Unlock()
	wg.Done()
}

// `DB_HOST`: Host of postgres DB.
//
// if Host is not set, the application will panic
func (e *environmentPostgres) getDBHost() (val string) {
	if e.dbHostOK {
		val = e.dbHost
	} else {
		log.Panic("PostgreSQL: DB_HOST not set")
	}
	return
}

// `DB_PORT`: Poert of postgres DB.
//
// if Port is not set, the application will panic
func (e *environmentPostgres) getDBPort() (val int) {
	if e.dbPortOK {
		value, err := strconv.Atoi(e.dbPort)

		if err != nil {
			log.Panicf("Env: Invalid BD_PORT %s, err: %s", e.dbPort, err.Error())
		}
		val = value
	} else {
		log.Panic("PostgreSQL: DB_PORT not set")
	}
	return
}

// `DB_USER`: Usuario of postgres DB.
//
// if User is not set, the application will panic
func (e *environmentPostgres) getDBUser() (val string) {
	if e.dbUserOK {
		val = e.dbUser
	} else {
		log.Panic("PostgreSQL: DB_USER not set")
	}
	return
}

// `DB_PASSWORD`: Contraseña of postgres DB.
//
// if Password is not set, the application will panic
func (e *environmentPostgres) getDBPass() (val string) {
	if e.dbPassOK {
		val = e.dbPass
	} else {
		log.Panic("PostgreSQL: DB_PASS not set")
	}
	return
}

// `DB_NAME`: Nombre of postgres DB.
//
// if Name is not set, the application will panic
func (e *environmentPostgres) getDBName() (val string) {
	if e.dbNameOK {
		val = e.dbName
	} else {
		log.Panic("PostgreSQL: DB_NAME not set")
	}
	return
}

// `DB_SSL`: SSL of postgres DB.
//
// By default is `disable`.
func (e *environmentPostgres) getDBSSL() (val string) {
	if e.dbSSLOK {
		val = e.dbSSL
	} else {
		val = "disable"
	}
	return
}
