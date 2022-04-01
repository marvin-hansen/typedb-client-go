// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package db_admin

import (
	v2 "github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/marvin-hansen/typedb-client-go/test/client/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

const dbName = utils.DBName

func TestNewClient(t *testing.T) {
	// create new client with default localhost config
	conf := v2.NewLocalConfig(dbName)
	c, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, c, utils.ClientError)
}

func TestCreateDatabase(t *testing.T) {
	conf := v2.NewLocalConfig(dbName)
	c, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, c, utils.ClientError)

	ok, err := c.DBManager.CreateDatabase(dbName)
	assert.NoError(t, err, "Should be no error")
	assert.Equal(t, ok, true, "Should be true")
}

func TestExistsDatabase(t *testing.T) {
	conf := v2.NewLocalConfig(dbName)
	c, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, c, utils.ClientError)

	existsDatabase, err := c.DBManager.CheckDatabaseExists(dbName)
	assert.NoError(t, err, "Should be no error")
	expected := true
	actual := existsDatabase
	assert.Equal(t, expected, actual, "Should be true i.e. exists")
}

func TestDeleteDatabase(t *testing.T) {
	conf := v2.NewLocalConfig(dbName)
	c, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, c, utils.ClientError)

	// check if DB exists
	existsDatabase, err := c.DBManager.CheckDatabaseExists(dbName)
	assert.NoError(t, err, "Should be no error")
	expected := true
	actual := existsDatabase
	assert.Equal(t, expected, actual, "Should be true i.e. exists")

	// Delete DB if exists.
	// Notice, DeleteDatabase returns true if the DB doesn't exist without being deleted b/c it's already gone
	// AND returns true when the DB actually got deleted.
	ok, err := c.DBManager.DeleteDatabase(dbName)
	assert.NoError(t, err, "Should be no error")
	assert.Equal(t, ok, true, "Should be true i.e. exists")

	existsDatabase, err = c.DBManager.CheckDatabaseExists(dbName)
	assert.NoError(t, err, "Should be no error")
	expected = false
	actual = existsDatabase
	assert.Equal(t, expected, actual, "Should be true i.e. exists")
}
