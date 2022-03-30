// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package query

import (
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/marvin-hansen/typedb-client-go/test/client/utils"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	dbName  = utils.DBName
	verbose = false
)

func testPrint(msg string) {
	if verbose {
		println(msg)
	}
}

func TestQueryMatch(t *testing.T) {

	conf := v2.NewLocalConfig(dbName)
	client, cancel := v2.NewClient(conf)
	defer cancel()
	assert.NotNil(t, client, utils.ClientError)

	testPrint("* Create Session")
	session, sessionOpenErr := client.OpenNewDataSession(dbName)
	assert.NoError(t, sessionOpenErr, "Should be no error")

	// TEST MATCH QUERY
	query := utils.GetTestQuery()

	testPrint("* Create session & request ID")
	sessionID := session.GetSessionId()
	requestId := ksuid.New().Bytes()

	testPrint("* Query TypeDB")
	queryResults, queryErr := client.RunMatchQuery(sessionID, requestId, query)
	if queryErr != nil {
		println(fmt.Errorf("could not create transaction: %w", queryErr))
	}
	assert.NoError(t, queryErr, "Should be no error")
	assert.NotNil(t, queryResults, "Query should return some results")

	if verbose {
		println("Print results")
		for _, item := range queryResults {
			println(item.String())
		}
	}

	testPrint("* Close Session")
	closeSessionErr := client.CloseSession(session.SessionId)
	if closeSessionErr != nil {
		return
	}
}
