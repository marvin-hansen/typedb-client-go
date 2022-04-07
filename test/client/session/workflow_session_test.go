package session

import (
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/marvin-hansen/typedb-client-go/test/client/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// Important, to make the testing time reasonable & visible,
//* Switch the heartBeatInterval to 1 (File: src/client/v2/manager_session_monitor)
//* Uncomment the debug lines in the sendPulseRequest method to see a console log

func TestSingleSession(t *testing.T) {
	// testing a single session. Good citizen closes session & client

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
	// testing multiple session. Good citizen closes all sessions & client

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

func TestSessionNotClosed(t *testing.T) {
	// testing multiple session. Sloppy citizen forgets to close a session.
	// Good client will close it anyway:-0

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

	testPrint("Close session 1")
	closeErr := client.SessionManager.CloseSession(sessionOneID)
	assert.NoError(t, closeErr, "Should be no session close error")

	testPrint("NOT closing session 2. Sits idle for now but still sends heartbeats")

	time.Sleep(time.Second * 1)

	testPrint("Close client")
	testPrint("Client closes all idling sessions.")

	client.Close()

	// Lazy citizen is on the beach already. Thanks client ;-0
	testPrint("Success: Test passed! ")
}
