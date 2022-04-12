package query_delete

import (
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/marvin-hansen/typedb-client-go/data"
	v2 "github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/marvin-hansen/typedb-client-go/test/client/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteQuery(t *testing.T) {
	client, cancel := utils.GetClient()
	defer cancel()

	utils.TestPrint("* Create Session")
	sessionID, sessionOpenErr := client.Session.NewSession(utils.DbName, common.Session_DATA)
	assert.NoError(t, sessionOpenErr, "Should be no error")

	gql := data.GetQueryDeletePerson()
	utils.TestPrint("* Get Test Delete statement")
	println(gql)

	utils.TestPrint("* Delete entry from TypeDB")
	options := v2.NewOptions()

	_, deleteErr := client.Query.Delete(sessionID, gql, options)
	if deleteErr != nil {
		println(deleteErr.Error())
	}
	assert.NoError(t, deleteErr, "Should be no error")

	utils.TestPrint("* Close Session")
	closeSessionErr := client.Session.CloseSession(sessionID)
	assert.NoError(t, closeSessionErr, "Should be no error")

	utils.TestPrint("* Close Client")
	client.Close()
}
