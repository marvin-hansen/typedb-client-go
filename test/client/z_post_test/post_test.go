package z_post

import (
	v2 "github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/marvin-hansen/typedb-client-go/test/client/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

const dbName = utils.DBName

func TestDBDelete(t *testing.T) {
	conf := v2.NewLocalConfig(dbName)
	client, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, client, utils.ClientError)
	teardownErr := utils.DBTeardown(client, dbName)
	assert.NoError(t, teardownErr, "Should be no DB teardown error")
	client.Close()
}

func TestDoesNotExistsDatabase(t *testing.T) {
	conf := v2.NewLocalConfig(dbName)
	c, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, c, utils.ClientError)
	existsDatabase, err := c.DBManager.CheckDatabaseExists(dbName)
	assert.NoError(t, err, "Should be no error")
	expected := false
	actual := existsDatabase
	assert.Equal(t, expected, actual, "Should be true i.e. exists")
	c.Close()
}
