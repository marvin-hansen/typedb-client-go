package client

import (
	v2 "github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryMatch(t *testing.T) {

	// CREATE CLIENT
	dbName := "TestDB"
	conf := v2.NewLocalConfig(dbName)
	c, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, c, ClientError)

	// CREATE DATABASE
	ok, status, err := c.CreateDatabase(dbName)
	assert.NoError(t, err, "Should be no error")
	assert.Equal(t, ok, true, "Should be true")
	assert.Equal(t, int(status), v2.OK, "Should be OK == 0")

	// TEST MATCH QUERY

	// DELETE DATABASE
	ok, status, err = c.DeleteDatabase(dbName)
	assert.NoError(t, err, "Should be no error")
	assert.Equal(t, ok, true, "Should be true i.e. exists")
	assert.Equal(t, int(status), v2.OK, "Should be OK == 0")

	// CHECK DB WAS DELETED
	existsDatabase, status, err := c.CheckDatabaseExists(dbName)
	assert.NoError(t, err, "Should be no error")
	assert.Equal(t, existsDatabase, false, "Should be true i.e. exists")
	assert.Equal(t, int(status), int(v2.DBNotExists), "Should be OK == 0")
}
