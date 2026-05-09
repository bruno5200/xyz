package postgres

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvironment_InitDB(t *testing.T) {
	t.Run("PanicOnMissingRequiredVars", func(t *testing.T) {
		os.Clearenv()
		assert.Panics(t, func() {
			e.initDB()
		})
	})

	t.Run("SuccessfulInitialization", func(t *testing.T) {
		t.Setenv("DB_HOST", "localhost")
		t.Setenv("DB_PORT", "5432")
		t.Setenv("DB_USER", "testuser")
		t.Setenv("DB_PASSWORD", "testpass")
		t.Setenv("DB_NAME", "testdb")
		t.Setenv("DB_SSL", "require")

		assert.NotPanics(t, func() {
			e.initDB()
		})

		assert.Equal(t, "localhost", e.getDBHost())
		assert.Equal(t, 5432, e.getDBPort())
		assert.Equal(t, "testuser", e.getDBUser())
		assert.Equal(t, "testdb", e.getDBName())
		assert.Equal(t, "require", e.getDBSSL())
	})
}

func TestGetDBPort_Validation(t *testing.T) {
	env := &environmentPostgres{}

	t.Run("ValidPort", func(t *testing.T) {
		env.dbPort = "5432"
		assert.Equal(t, 5432, env.getDBPort())
	})

	t.Run("InvalidPortFormat", func(t *testing.T) {
		env.dbPort = "not-a-number"
		assert.Panics(t, func() {
			env.getDBPort()
		})
	})
}

func TestGetDBSSL_Defaults(t *testing.T) {
	env := &environmentPostgres{}
	t.Run("DefaultToDisable", func(t *testing.T) {
		env.dbSSL = ""
		assert.Equal(t, "disable", env.getDBSSL())
	})

	t.Run("CustomValue", func(t *testing.T) {
		env.dbSSL = "verify-full"
		assert.Equal(t, "verify-full", env.getDBSSL())
	})
}

func TestConnection_DSNConstruction(t *testing.T) {
	t.Setenv("DB_HOST", "localhost")
	t.Setenv("DB_PORT", "5432")
	t.Setenv("DB_USER", "user")
	t.Setenv("DB_PASSWORD", "pass")
	t.Setenv("DB_NAME", "db")

	t.Run("ConstructStandardDSN", func(t *testing.T) {
		os.Unsetenv("DATABASE_URL")
		t.Setenv("DB_SSL", "disable")

		dsn := connection()
		expected := "user=user password=pass host=localhost dbname=db port=5432 sslmode=disable"
		assert.Equal(t, expected, dsn)
	})

	t.Run("PrecedenceOfDatabaseURL", func(t *testing.T) {
		url := "postgres://admin:secret@remote:5433/proddb?sslmode=verify-ca"
		t.Setenv("DATABASE_URL", url)

		dsn := connection()
		assert.Equal(t, url, dsn)
	})
}

func TestPublicFunctions_InitialState(t *testing.T) {
	t.Run("PostgresDBReturnsNilBeforeConnect", func(t *testing.T) {
		assert.Nil(t, PostgresDB())
	})

	t.Run("NewConnectionPanicsOnInvalidConfig", func(t *testing.T) {
		t.Setenv("DB_PORT", "0")
		assert.Panics(t, func() {
			NewConnection()
		})
	})
}
