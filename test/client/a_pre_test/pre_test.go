package a_pre_test

import (
	v2 "github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/marvin-hansen/typedb-client-go/test/client/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

const dbName = utils.DBName

func TestDBCreate(t *testing.T) {
	conf := v2.NewLocalConfig(dbName)
	client, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, client, utils.ClientError)

	dbErr := DBSetup(client, dbName)
	assert.NoError(t, dbErr, "Should be no DB setup error")

	client.Close()
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

	c.Close()
}
