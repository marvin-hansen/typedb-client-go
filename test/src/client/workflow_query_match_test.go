package client

import (
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/common"
	v2 "github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryMatch(t *testing.T) {

	// CREATE CLIENT
	dbName := "TestDB"
	conf := v2.NewLocalConfig(dbName)
	c, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, c, ClientError)

	// CREATE DATABASE
	ok, status, createErr := c.CreateDatabase(dbName)
	assert.NoError(t, createErr, "Should be no error")
	assert.Equal(t, ok, true, "Should be true")
	assert.Equal(t, int(status), v2.OK, "Should be OK == 0")

	// TEST MATCH QUERY

	// create requestID
	requestId := ksuid.New().Bytes()

	query := `match
				$x sub thing; 
			get 
				$x;`

	// create options
	options := &common.Options{
		InferOpt:              &common.Options_Infer{Infer: true},
		ExplainOpt:            &common.Options_Explain{Explain: true},
		TransactionTimeoutOpt: &common.Options_TransactionTimeoutMillis{TransactionTimeoutMillis: 500},
	}

	// create metadata
	metadata := map[string]string{}

	queryResults, queryErr := c.RunMatchQuery(requestId, query, metadata, options)
	if queryErr != nil {
		println(fmt.Errorf("could not create transaction: %w", queryErr))
	}

	assert.NoError(t, queryErr, "Should be no error")
	assert.NotNil(t, queryResults, "Query should return some results")

	println(queryResults)

	// DELETE DATABASE
	ok, status, queryErr = c.DeleteDatabase(dbName)
	assert.NoError(t, queryErr, "Should be no error")
	assert.Equal(t, ok, true, "Should be true i.e. exists")
	assert.Equal(t, int(status), v2.OK, "Should be OK == 0")

	// CHECK DB WAS DELETED
	existsDatabase, status, queryErr := c.CheckDatabaseExists(dbName)
	assert.NoError(t, queryErr, "Should be no error")
	assert.Equal(t, existsDatabase, false, "Should be true i.e. exists")
	assert.Equal(t, int(status), int(v2.DBNotExists), "Should be OK == 0")
}
