package client

import (
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryMatch(t *testing.T) {

	println("* Create Client")
	dbName := "TestDB"
	conf := v2.NewLocalConfig(dbName)
	client, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, client, ClientError)

	println("* Create Session")
	session, sessionOpenErr := client.OpenNewDataSession(dbName)
	if sessionOpenErr != nil {
		return
	}

	// TEST MATCH QUERY
	query := `
match 
	$x sub thing; 
get 
	$x;
`

	println("* Create Options")
	options := &common.Options{
		InferOpt:              &common.Options_Infer{Infer: true},
		ExplainOpt:            &common.Options_Explain{Explain: true},
		TransactionTimeoutOpt: &common.Options_TransactionTimeoutMillis{TransactionTimeoutMillis: 500},
	}

	println("* Create Metadata")
	metadata := map[string]string{}

	println("* Create session & request ID")
	sessionID := session.GetSessionId()
	requestId := ksuid.New().Bytes()

	println("* Query TypeDB")
	queryResults, queryErr := client.RunMatchQuery(sessionID, requestId, query, metadata, options)
	if queryErr != nil {
		println(fmt.Errorf("could not create transaction: %w", queryErr))
	}
	assert.NoError(t, queryErr, "Should be no error")
	assert.NotNil(t, queryResults, "Query should return some results")
	println(queryResults)

	println("* Close Session")
	closeSessionErr := client.CloseSession(session.SessionId)
	if closeSessionErr != nil {
		return
	}
}
