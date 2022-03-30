// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package client

import (
	v2 "github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewClient(t *testing.T) {
	conf := v2.NewLocalConfig("TestDB")
	c, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, c, ClientError)
}

func TestCreateDatabase(t *testing.T) {

	dbName := "TestDB"
	conf := v2.NewLocalConfig(dbName)
	c, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, c, ClientError)

	ok, err := c.CreateDatabase(dbName)
	assert.NoError(t, err, "Should be no error")
	assert.Equal(t, ok, true, "Should be true")
}

func TestExistsDatabase(t *testing.T) {
	dbName := "TestDB"
	conf := v2.NewLocalConfig(dbName)
	c, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, c, ClientError)

	existsDatabase, err := c.CheckDatabaseExists(dbName)
	assert.NoError(t, err, "Should be no error")
	assert.Equal(t, existsDatabase, true, "Should be true i.e. exists")
}

func TestDeleteDatabase(t *testing.T) {
	dbName := "TestDB"
	conf := v2.NewLocalConfig(dbName)
	c, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, c, ClientError)

	ok, err := c.DeleteDatabase(dbName)
	assert.NoError(t, err, "Should be no error")
	assert.Equal(t, ok, true, "Should be true i.e. exists")

	existsDatabase, err := c.CheckDatabaseExists(dbName)
	assert.NoError(t, err, "Should be no error")
	assert.Equal(t, existsDatabase, false, "Should be true i.e. exists")
}
