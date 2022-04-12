package query

import (
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/marvin-hansen/typedb-client-go/data"
	v2 "github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/marvin-hansen/typedb-client-go/test/client/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetup(t *testing.T) {

	// we have to destroy the entire DB, build a new one with schema before bulk inserts
	// to prevent double entries left-over from previous tests.

	client, cancel := utils.GetClient()
	defer cancel()

	DBTeardownErr := utils.DBTeardown(client, utils.DbName)
	assert.NoError(t, DBTeardownErr, "Should be no error")

	dbSetupErr := utils.DBSetup(client, utils.DbName)
	assert.NoError(t, dbSetupErr, "Should be no error")

	schemaErr := utils.SchemaSetup(client, utils.DbName)
	assert.NoError(t, schemaErr, "Should be no error")
}

func TestInsertBulkQuery(t *testing.T) {
	client, cancel := utils.GetClient()
	defer cancel()

	utils.TestPrint("* Create Session")
	sessionID, sessionOpenErr := client.Session.NewSession(utils.DbName, common.Session_DATA)
	assert.NoError(t, sessionOpenErr, "Should be no error")

	utils.TestPrint("* Get data to Insert")
	gql, dataErr := data.GetPhoneCallsDataGql()
	assert.NoError(t, dataErr, "Should be no data error")

	utils.TestPrint("* Insert into TypeDB")
	options := v2.NewOptions()
	insertError := client.Query.InsertBulk(sessionID, gql, options)
	assert.NoError(t, insertError, "Should be no error")

	utils.TestPrint("* Close Session")
	closeSessionErr := client.Session.CloseSession(sessionID)
	assert.NoError(t, closeSessionErr, "Should be no error")

	utils.TestPrint("* Close Client")
	client.Close()
}
