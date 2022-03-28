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

}

func TestSchemaGet(t *testing.T) {
	println("* Create Client")
	dbName := "TestDB"
	conf := v2.NewLocalConfig(dbName)
	client, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, client, ClientError)

}
