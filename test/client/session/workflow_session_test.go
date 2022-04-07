package session

import (
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/marvin-hansen/typedb-client-go/test/client/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSingleSession(t *testing.T) {

	testPrint("Create client")
	client, cancel := getClient()
	defer cancel()
	assert.NotNil(t, client, utils.ClientError)

	testPrint("Create session")
	sessionID, err := client.SessionManager.NewSession(dbName, common.Session_DATA)
	assert.NoError(t, err, "Should be no session open error")
	assert.NotNil(t, sessionID, "Should not be nil")

	testPrint("SessionID " + byteToString(sessionID))

	testPrint("Waiting...")
	time.Sleep(time.Second * 3)

	testPrint("Close session")
	closeErr := client.SessionManager.CloseSession(sessionID)
	assert.NoError(t, closeErr, "Should be no session close error")

	testPrint("Close client")
	client.Close()

	testPrint("Success: Test passed! ")
}
