package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test for the Image function
func TestImage(t *testing.T) {
	// Working image
	result1, err := Image("postgres")
	assert.NoError(t, err)
	assert.Equal(
		t,
		result1,
		ImageData{URL: "https://hub.docker.com/_/postgres", Variables: []{{Name: "POSTGRES_INITDB_ARGS", Description: "Used to send arguments to `postgres initdb`. The value is a space separated string of arguments as `postgres initdb` would expect them. This is useful for adding functionality like data page checksums: -e `POSTGRES_INITDB_ARGS=\"--data-checksums\"`", Default: "", Required: false, Uncommon: true}, {Name: "POSTGRES_INITDB_WALDIR", Description: "Used to define another location for the Postgres transaction log.", Default: "", Required: false, Uncommon: true}, {Name: "POSTGRES_HOST_AUTH_METHOD", Description: "Used to control the `auth-method` for `host` connections for `all` databases, `all` users, and `all` addresses.", Default: "", Required: false, Uncommon: true}, {Name: "POSTGRES_PASSWORD", Description: "Sets the password for the default user", Default: "", Required: true, Uncommon: false}, {Name: "POSTGRES_USER", Description: "The username for the default user", Default: "postgres", Required: false, Uncommon: false}, {Name: "POSTGRES_DB", Description: "The default database name. Defaults to the value of `POSTGRES_USER`", Default: "", Required: false, Uncommon: false}}},
	)

	// Testing failing case
	_, err = Image("fakeimage")
	assert.Error(t, err)
}
