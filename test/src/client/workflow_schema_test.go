package client

import (
	"github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSchemaCreate(t *testing.T) {

	println("* Create Client")
	dbName := "TestDB"
	conf := v2.NewLocalConfig(dbName)
	client, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, client, ClientError)

	println("* Create Session")
	session, sessionOpenErr := client.OpenNewSchemaSession(dbName)
	if sessionOpenErr != nil {
		return
	}

	println("* Close Session")
	closeSessionErr := client.CloseSession(session.SessionId)
	if closeSessionErr != nil {
		return
	}

}
