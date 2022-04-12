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
	conf := v2.NewLocalConfig(dbName)
	c, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, c, utils.ClientError)

	c.Close()
}

func TestCreateDatabase(t *testing.T) {
	conf := v2.NewLocalConfig(dbName)
	c, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, c, utils.ClientError)
	err := c.DBManager.CreateDatabase(dbName)
	assert.NoError(t, err, "Should be no error")
	c.Close()
}

func TestExistsDatabase(t *testing.T) {
	conf := v2.NewLocalConfig(dbName)
	c, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, c, utils.ClientError)
	_, err := c.DBManager.CheckDatabaseExists(dbName)
	assert.NoError(t, err, "Should be no error")
	c.Close()
}

func TestDeleteDatabase(t *testing.T) {
	conf := v2.NewLocalConfig(dbName)
	c, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, c, utils.ClientError)
	_, err := c.DBManager.CheckDatabaseExists(dbName)
	assert.NoError(t, err, "Should be no error")

	// Delete DB if exists.
	// Notice, DeleteDatabase returns true if the DB doesn't exist without being deleted b/c it's already gone
	// AND returns true when the DB actually got deleted.
	err = c.DBManager.DeleteDatabase(dbName)
	assert.NoError(t, err, "Should be no error")
	_, err = c.DBManager.CheckDatabaseExists(dbName)
	assert.NoError(t, err, "Should be no error")
	c.Close()
}
