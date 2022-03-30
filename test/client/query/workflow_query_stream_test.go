// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package query

import (
	"context"
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/data"
	"github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/marvin-hansen/typedb-client-go/test/client/utils"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	dbName  = utils.DBName
	verbose = true
)

func testPrint(msg string) {
	if verbose {
		println(msg)
	}
}

func getClient() (*v2.Client, context.CancelFunc) {
	conf := v2.NewLocalConfig(dbName)
	client, cancel := v2.NewClient(conf)
	return client, cancel
}

func TestInsertQuery(t *testing.T) {
	client, cancel := getClient()
	defer cancel()

	testPrint("* Get Test Data")
	gql, dataErr := data.GetPhoneCallsDataGql()
	assert.NoError(t, dataErr, "could not get phone calls data gql: %w", dataErr)
	assert.NotNil(t, gql, "Data should not be nil")

	testPrint("* Create Session")
	session, sessionOpenErr := client.OpenNewDataSession(dbName)
	assert.NoError(t, sessionOpenErr, "Should be no error")

	//testPrint("* Create session & request ID")
	//sessionID := session.GetSessionId()
	//requestId := ksuid.New().Bytes()

	testPrint("* Close Session")
	closeSessionErr := client.CloseSession(session.SessionId)
	if closeSessionErr != nil {
		return
	}

}

func TestMatchQuery(t *testing.T) {
	client, cancel := getClient()
	defer cancel()

	testPrint("* Create Session")
	session, sessionOpenErr := client.OpenNewDataSession(dbName)
	assert.NoError(t, sessionOpenErr, "Should be no error")

	testPrint("* Create session & request ID")
	sessionID := session.GetSessionId()
	requestId := ksuid.New().Bytes()

	// TEST MATCH QUERY
	query := utils.GetTestQuery()

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
