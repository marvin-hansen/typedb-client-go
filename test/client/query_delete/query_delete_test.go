package query_delete

import (
	"github.com/marvin-hansen/typedb-client-go/common"
	v2 "github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/marvin-hansen/typedb-client-go/test/client/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteQuery(t *testing.T) {
	client, cancel := utils.GetClient()
	defer cancel()

	utils.TestPrint("* Create Session")
	sessionID, sessionOpenErr := client.SessionManager.NewSession(utils.DbName, common.Session_DATA)
	assert.NoError(t, sessionOpenErr, "Should be no error")

	gql := utils.GetPersonDelete()
	utils.TestPrint("* Get Test Delete statement")
	println(gql)

	utils.TestPrint("* Delete entry from TypeDB")
	options := v2.CreateNewRequestOptions()

	deleteResult, deleteErr := client.RunDeleteQuery(sessionID, gql, options)
	assert.NoError(t, deleteErr, "Should be no error")
	assert.NotNil(t, deleteResult, "Query should return some results")

	utils.TestPrint("* Close Session")
	closeSessionErr := client.SessionManager.CloseSession(sessionID)
	assert.NoError(t, closeSessionErr, "Should be no error")

	utils.TestPrint("* Close Client")
	client.Close()
}
