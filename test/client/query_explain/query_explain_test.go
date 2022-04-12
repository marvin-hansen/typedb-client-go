package query_explain

import (
	"github.com/marvin-hansen/typedb-client-go/common"
	v2 "github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/marvin-hansen/typedb-client-go/test/client/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExplainQuery(t *testing.T) {
	client, cancel := utils.GetClient()
	defer cancel()

	utils.TestPrint("* Create Session")
	sessionID, sessionOpenErr := client.Session.NewSession(utils.DbName, common.Session_DATA)
	assert.NoError(t, sessionOpenErr, "Should be no error")

	explainableID := int64(42)
	options := v2.NewOptions()

	explainResult, err := client.RunExplainQuery(sessionID, explainableID, options)
	assert.NoError(t, err, "Should be no error")
	assert.NotNil(t, explainResult, "Query should return some results")

	utils.TestPrint("* Close Session")
	closeSessionErr := client.Session.CloseSession(sessionID)
	assert.NoError(t, closeSessionErr, "Should be no error")

	utils.TestPrint("* Close Client")
	client.Close()

}
