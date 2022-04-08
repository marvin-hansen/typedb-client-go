package query_explain

import (
	"github.com/marvin-hansen/typedb-client-go/common"
	v2 "github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExplainQuery(t *testing.T) {
	client, cancel := getClient()
	defer cancel()

	testPrint("* Create Session")
	sessionID, sessionOpenErr := client.SessionManager.NewSession(dbName, common.Session_DATA)
	assert.NoError(t, sessionOpenErr, "Should be no error")

	explainableID := int64(42)
	options := v2.CreateNewRequestOptions()

	explainResult, err := client.RunExplainQuery(sessionID, explainableID, options)
	assert.NoError(t, err, "Should be no error")
	assert.NotNil(t, explainResult, "Query should return some results")

	testPrint("* Close Session")
	closeSessionErr := client.SessionManager.CloseSession(sessionID)
	assert.NoError(t, closeSessionErr, "Should be no error")

	testPrint("* Close Client")
	client.Close()

}
