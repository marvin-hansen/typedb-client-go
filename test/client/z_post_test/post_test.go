package a_pre_test

import (
	v2 "github.com/marvin-hansen/typedb-client-go/src/client/v2"
	utils "github.com/marvin-hansen/typedb-client-go/test/client/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

const dbName = utils.DBName

func TestDBDelete(t *testing.T) {
	conf := v2.NewLocalConfig(dbName)
	client, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, client, utils.ClientError)

	teardownErr := DBTeardown(client, dbName)
	assert.NoError(t, teardownErr, "Should be no DB teardown error")

}
