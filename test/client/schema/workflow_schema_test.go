// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package schema

import (
	"github.com/marvin-hansen/typedb-client-go/src/client/v2"
	utils2 "github.com/marvin-hansen/typedb-client-go/test/client/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

const dbName = "TestDB"

func TestSchemaCreate(t *testing.T) {
	conf := v2.NewLocalConfig(dbName)
	client, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, client, utils2.ClientError)

	dbErr := DBSetup(client, dbName)
	assert.NoError(t, dbErr, "Should be no DB setup error")

	testSchema := utils2.GetPhoneCallsSchema()
	err := client.CreateDatabaseSchema(dbName, testSchema)

	assert.NoError(t, err, "Should be no schema error")
}

func TestSchemaGet(t *testing.T) {
	conf := v2.NewLocalConfig(dbName)
	client, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, client, utils2.ClientError)

	dbErr := DBSetup(client, dbName)
	assert.NoError(t, dbErr, "Should be no DB error")

	allEntries, err := client.GetDatabaseSchema(dbName)

	assert.NoError(t, err, "Should be no schema error")
	assert.NotNil(t, t, allEntries, "Should not be nil")

	if len(allEntries) > 0 {
		for _, item := range allEntries {
			println(item)
		}
		println()
	}

	teardownErr := DBTeardown(client, dbName)
	assert.NoError(t, teardownErr, "Should be no DB teardown error")
}
