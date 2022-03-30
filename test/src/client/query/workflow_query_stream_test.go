// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package query

import (
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/marvin-hansen/typedb-client-go/test/src/client/utils"
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
	assert.NotNil(t, client, utils.ClientError)

	println("* Create Session")
	session, sessionOpenErr := client.OpenNewDataSession(dbName)
	if sessionOpenErr != nil {
		return
	}

	// TEST MATCH QUERY
	query := utils.GetTestQuery()

	println("* Create session & request ID")
	sessionID := session.GetSessionId()
	requestId := ksuid.New().Bytes()

	println("* Query TypeDB")
	queryResults, queryErr := client.RunMatchQuery(sessionID, requestId, query)
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
