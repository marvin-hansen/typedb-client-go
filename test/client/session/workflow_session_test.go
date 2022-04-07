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

func TestMultiSession(t *testing.T) {

	testPrint("Create client")
	client, cancel := getClient()
	defer cancel()
	assert.NotNil(t, client, utils.ClientError)

	testPrint("Create session 1")
	sessionOneID, err := client.SessionManager.NewSession(dbName, common.Session_DATA)
	assert.NoError(t, err, "Should be no session open error")
	assert.NotNil(t, sessionOneID, "Should not be nil")

	testPrint("SessionID " + byteToString(sessionOneID))

	testPrint("Waiting...")
	time.Sleep(time.Second * 1)

	testPrint("Create session 2")
	sessionTwoID, err := client.SessionManager.NewSession(dbName, common.Session_DATA)
	assert.NoError(t, err, "Should be no session open error")
	assert.NotNil(t, sessionTwoID, "Should not be nil")

	testPrint("SessionID " + byteToString(sessionTwoID))

	testPrint("Waiting...")
	time.Sleep(time.Second * 1)

	testPrint("Create session 3")
	sessionThreeID, err := client.SessionManager.NewSession(dbName, common.Session_DATA)
	assert.NoError(t, err, "Should be no session open error")
	assert.NotNil(t, sessionThreeID, "Should not be nil")

	testPrint("SessionID " + byteToString(sessionThreeID))

	testPrint("Waiting...")
	time.Sleep(time.Second * 1)

	testPrint("Close session 1")
	closeErr := client.SessionManager.CloseSession(sessionOneID)
	assert.NoError(t, closeErr, "Should be no session close error")

	testPrint("Close session 2")
	closeErr = client.SessionManager.CloseSession(sessionTwoID)
	assert.NoError(t, closeErr, "Should be no session close error")

	testPrint("Close session 3")
	closeErr = client.SessionManager.CloseSession(sessionThreeID)
	assert.NoError(t, closeErr, "Should be no session close error")

	testPrint("Close client")
	client.Close()

	testPrint("Success: Test passed! ")
}
