// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package schema

import (
	"github.com/marvin-hansen/typedb-client-go/data"
	"github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/marvin-hansen/typedb-client-go/test/client/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

const dbName = utils.DBName

func TestSchemaCreate(t *testing.T) {
	conf := v2.NewLocalConfig(dbName)
	client, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, client, utils.ClientError)

	testSchema := data.GetPhoneCallsSchema()
	err := client.CreateDatabaseSchema(dbName, testSchema)

	assert.NoError(t, err, "Should be no schema error")
}

func TestSchemaGet(t *testing.T) {
	conf := v2.NewLocalConfig(dbName)
	client, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, client, utils.ClientError)

	allEntries, err := client.GetDatabaseSchema(dbName)

	assert.NoError(t, err, "Should be no schema error")
	assert.NotNil(t, t, allEntries, "Should not be nil")

	if len(allEntries) > 0 {
		for _, item := range allEntries {
			println(item)
		}
		println()
	}
}
