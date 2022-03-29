package client

import (
	"github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSchemaCreate(t *testing.T) {
	println("* Create Client")
	dbName := "TestDB"
	conf := v2.NewLocalConfig(dbName)
	client, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, client, ClientError)

	dbErr := dbSetup(client, dbName)
	assert.NoError(t, dbErr, "Should be no DB setup error")

	println("* Write Schema")
	testSchema := getPhoneCallsSchema()
	status, err := client.CreateDatabaseSchema(dbName, testSchema)

	assert.NoError(t, err, "Should be no schema error")
	assert.Equal(t, int(v2.OK), int(status), "Should be OK == 0")
}

func TestSchemaGet(t *testing.T) {
	println("* Create Client")
	dbName := "TestDB"
	conf := v2.NewLocalConfig(dbName)
	client, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, client, ClientError)

	dbErr := dbSetup(client, dbName)
	assert.NoError(t, dbErr, "Should be no DB error")

	println("* Get Schema")
	allEntries, status, err := client.GetDatabaseSchema(dbName)

	assert.NoError(t, err, "Should be no schema error")
	assert.Equal(t, int(v2.OK), int(status), "Should be OK == 0")
	assert.NotNil(t, t, allEntries, "Should not be nil")

	if len(allEntries) > 0 {
		for _, item := range allEntries {
			println(item)
		}
		println()
	}

	teardownErr := dbTeardown(client, dbName)
	assert.NoError(t, teardownErr, "Should be no DB teardown error")
}
