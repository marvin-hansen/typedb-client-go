package query

import (
	"github.com/marvin-hansen/typedb-client-go/common"
	v2 "github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/marvin-hansen/typedb-client-go/test/client/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TODO / FIXME: Timeout after 300 sec
// Error: could not receive query response: rpc error: code =
// Internal desc = [SRV28] Invalid Server Operation: Transaction exceeded maximum configured duration of '300' seconds.
func TestInsertQuery(t *testing.T) {
	client, cancel := utils.GetClient()
	defer cancel()

	utils.TestPrint("* Create Session")
	sessionID, sessionOpenErr := client.SessionManager.NewSession(utils.DbName, common.Session_DATA)
	assert.NoError(t, sessionOpenErr, "Should be no error")

	gql := utils.GetCompanyInsert()
	utils.TestPrint("* Get Test Insert")
	println(gql)

	utils.TestPrint("* Insert into TypeDB")
	options := v2.CreateNewRequestOptions()

	// this one fails
	insertResults, insertError := client.RunInsertQuery(sessionID, gql, options)

	assert.NoError(t, insertError, "Should be no error")
	assert.NotNil(t, insertResults, "Query should return some results")

	utils.TestPrint("* Close Session")
	closeSessionErr := client.SessionManager.CloseSession(sessionID)
	assert.NoError(t, closeSessionErr, "Should be no error")

	utils.TestPrint("* Close Client")
	client.Close()
}