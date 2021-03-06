package query_update

import (
	"github.com/marvin-hansen/typedb-client-go/common"
	v2 "github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/marvin-hansen/typedb-client-go/test/client/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateQuery(t *testing.T) {
	client, cancel := utils.GetClient()
	defer cancel()

	utils.TestPrint("* Create Session")
	sessionID, sessionOpenErr := client.Session.NewSession(utils.DbName, common.Session_DATA)
	assert.NoError(t, sessionOpenErr, "Should be no error")

	gql := utils.GetPersonTelUpdate()
	utils.TestPrint("* Get Test Update")
	println(gql)

	utils.TestPrint("* Update into TypeDB")
	options := v2.NewOptions()

	updateRes, updateErr := client.Query.Update(sessionID, gql, options)
	assert.NoError(t, updateErr, "Should be no error")
	assert.NotNil(t, updateRes, "Query should return some results")

	utils.TestPrint("* Close Session")
	closeSessionErr := client.Session.CloseSession(sessionID)
	assert.NoError(t, closeSessionErr, "Should be no error")

	utils.TestPrint("* Close Client")
	client.Close()
}
