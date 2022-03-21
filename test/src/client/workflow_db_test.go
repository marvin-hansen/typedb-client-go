package client

import (
	v2 "github.com/marvin-hansen/go-typedb/src/client/v2"
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

	ok, status, err := c.CreateDatabase(dbName)
	assert.NoError(t, err, "Should be no error")
	assert.Equal(t, ok, true, "Should be true")
	assert.Equal(t, int(status), v2.OK, "Should be OK == 0")
}

func TestExistsDatabase(t *testing.T) {
	dbName := "TestDB"
	conf := v2.NewLocalConfig(dbName)
	c, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, c, ClientError)

	existsDatabase, status, err := c.ExistsDatabase(dbName)
	assert.NoError(t, err, "Should be no error")
	assert.Equal(t, existsDatabase, true, "Should be true i.e. exists")
	assert.Equal(t, int(status), v2.OK, "Should be OK == 0")
}

func TestDeleteDatabase(t *testing.T) {
	dbName := "TestDB"
	conf := v2.NewLocalConfig(dbName)
	c, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, c, ClientError)

	ok, status, err := c.DeleteDatabase(dbName)
	assert.NoError(t, err, "Should be no error")
	assert.Equal(t, ok, true, "Should be true i.e. exists")
	assert.Equal(t, int(status), v2.OK, "Should be OK == 0")

	existsDatabase, status, err := c.ExistsDatabase(dbName)
	assert.NoError(t, err, "Should be no error")
	assert.Equal(t, existsDatabase, false, "Should be true i.e. exists")
	assert.Equal(t, int(status), int(v2.DBNotExists), "Should be OK == 0")
}
