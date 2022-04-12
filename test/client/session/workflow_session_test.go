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

	utils.TestPrint("Create client")
	client, cancel := utils.GetClient()
	defer cancel()
	assert.NotNil(t, client, utils.ClientError)

	utils.TestPrint("Create session")
	sessionID, err := client.Session.NewSession(utils.DbName, common.Session_DATA)
	assert.NoError(t, err, "Should be no session open error")
	assert.NotNil(t, sessionID, "Should not be nil")

	utils.TestPrint("SessionID " + utils.ByteToString(sessionID))

	utils.TestPrint("Waiting...")
	time.Sleep(time.Second * 3)

	utils.TestPrint("Close session")
	closeErr := client.Session.CloseSession(sessionID)
	assert.NoError(t, closeErr, "Should be no session close error")

	utils.TestPrint("Close client")
	client.Close()

	utils.TestPrint("Success: Test passed! ")
}

func TestMultiSession(t *testing.T) {
	// testing multiple session. Good citizen closes all sessions & client

	utils.TestPrint("Create client")
	client, cancel := utils.GetClient()
	defer cancel()
	assert.NotNil(t, client, utils.ClientError)

	utils.TestPrint("Create session 1")
	sessionOneID, err := client.Session.NewSession(utils.DbName, common.Session_DATA)
	assert.NoError(t, err, "Should be no session open error")
	assert.NotNil(t, sessionOneID, "Should not be nil")

	utils.TestPrint("SessionID " + utils.ByteToString(sessionOneID))

	utils.TestPrint("Waiting...")
	time.Sleep(time.Second * 1)

	utils.TestPrint("Create session 2")
	sessionTwoID, err := client.Session.NewSession(utils.DbName, common.Session_DATA)
	assert.NoError(t, err, "Should be no session open error")
	assert.NotNil(t, sessionTwoID, "Should not be nil")

	utils.TestPrint("SessionID " + utils.ByteToString(sessionTwoID))

	utils.TestPrint("Waiting...")
	time.Sleep(time.Second * 1)

	utils.TestPrint("Create session 3")
	sessionThreeID, err := client.Session.NewSession(utils.DbName, common.Session_DATA)
	assert.NoError(t, err, "Should be no session open error")
	assert.NotNil(t, sessionThreeID, "Should not be nil")

	utils.TestPrint("SessionID " + utils.ByteToString(sessionThreeID))

	utils.TestPrint("Waiting...")
	time.Sleep(time.Second * 1)

	utils.TestPrint("Close session 1")
	closeErr := client.Session.CloseSession(sessionOneID)
	assert.NoError(t, closeErr, "Should be no session close error")

	utils.TestPrint("Close session 2")
	closeErr = client.Session.CloseSession(sessionTwoID)
	assert.NoError(t, closeErr, "Should be no session close error")

	utils.TestPrint("Close session 3")
	closeErr = client.Session.CloseSession(sessionThreeID)
	assert.NoError(t, closeErr, "Should be no session close error")

	utils.TestPrint("Close client")
	client.Close()

	utils.TestPrint("Success: Test passed! ")
}

func TestSessionNotClosed(t *testing.T) {
	// testing multiple session. Sloppy citizen forgets to close a session.
	// Good client will close it anyway:-0

	utils.TestPrint("Create client")
	client, cancel := utils.GetClient()
	defer cancel()
	assert.NotNil(t, client, utils.ClientError)

	utils.TestPrint("Create session 1")
	sessionOneID, err := client.Session.NewSession(utils.DbName, common.Session_DATA)
	assert.NoError(t, err, "Should be no session open error")
	assert.NotNil(t, sessionOneID, "Should not be nil")

	utils.TestPrint("SessionID " + utils.ByteToString(sessionOneID))

	utils.TestPrint("Waiting...")
	time.Sleep(time.Second * 1)

	utils.TestPrint("Create session 2")
	sessionTwoID, err := client.Session.NewSession(utils.DbName, common.Session_DATA)
	assert.NoError(t, err, "Should be no session open error")
	assert.NotNil(t, sessionTwoID, "Should not be nil")

	utils.TestPrint("SessionID " + utils.ByteToString(sessionTwoID))

	utils.TestPrint("Waiting...")
	time.Sleep(time.Second * 1)

	utils.TestPrint("Close session 1")
	closeErr := client.Session.CloseSession(sessionOneID)
	assert.NoError(t, closeErr, "Should be no session close error")

	utils.TestPrint("NOT closing session 2. Sits idle for now but still sends heartbeats")

	time.Sleep(time.Second * 1)

	utils.TestPrint("Close client")
	utils.TestPrint("Client closes all idling sessions.")

	client.Close()

	// Lazy citizen is on the beach already. Thanks client ;-0
	utils.TestPrint("Success: Test passed! ")
}
