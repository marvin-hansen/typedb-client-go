// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package query

import (
	"context"
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/marvin-hansen/typedb-client-go/data"
	"github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/marvin-hansen/typedb-client-go/test/client/utils"
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

// TODO / FIXME: Timeout after 300 sec
func TestInsertBulkQuery(t *testing.T) {
	client, cancel := getClient()
	defer cancel()

	testPrint("* Get Test Data")
	gql, dataErr := data.GetPhoneCallsDataGql()
	assert.NoError(t, dataErr, "could not get phone calls data gql: %w", dataErr)
	assert.NotNil(t, gql, "Data should not be nil")

	testPrint("* Create Session")
	session, sessionOpenErr := v2.NewSession(client, dbName, common.Session_DATA) //client.OpenNewDataSession(dbName)
	assert.NoError(t, sessionOpenErr, "Should be no error")

	testPrint("* Insert into TypeDB")
	sessionID := session.GetSessionId()
	options := v2.CreateNewRequestOptions()
	insertResults, insertError := client.RunInsertBulkQuery(sessionID, gql, options)
	assert.NoError(t, insertError, "Should be no error")
	assert.NotNil(t, insertResults, "Query should return some results")

	if verbose {
		println("Print results")
		for _, item := range insertResults {
			println(item.String())
		}
	}

	testPrint("* Close Session")
	closeSessionErr := session.Close() //client.CloseSession(session.GetSessionId())
	assert.NoError(t, closeSessionErr, "Should be no error")
}

// TODO / FIXME: Timeout after 300 sec
func TestInsertQuery(t *testing.T) {
	client, cancel := getClient()
	defer cancel()

	testPrint("* Create Session")
	session, sessionOpenErr := v2.NewSession(client, dbName, common.Session_DATA)
	assert.NoError(t, sessionOpenErr, "Should be no error")

	testPrint("* Get Test Insert")
	gql := utils.GetCompanyInsert()

	testPrint("* Insert into TypeDB")
	sessionID := session.GetSessionId()
	options := v2.CreateNewRequestOptions()

	insertResults, insertError := client.RunInsertQuery(sessionID, gql, options)
	assert.NoError(t, insertError, "Should be no error")
	assert.NotNil(t, insertResults, "Query should return some results")

	testPrint("* Close Session")
	closeSessionErr := session.Close()
	assert.NoError(t, closeSessionErr, "Should be no error")
}

func TestMatchQuery(t *testing.T) {
	client, cancel := getClient()
	defer cancel()

	testPrint("* Create Session")
	session, sessionOpenErr := v2.NewSession(client, dbName, common.Session_DATA) //client.OpenNewDataSession(dbName)
	assert.NoError(t, sessionOpenErr, "Should be no error")

	testPrint("* Create session & request ID")
	sessionID := session.GetSessionId()

	// TEST MATCH QUERY
	query := utils.GetTestQuery()

	testPrint("* Query TypeDB")
	options := v2.CreateNewRequestOptions()
	queryResults, queryErr := client.RunMatchQuery(sessionID, query, options)
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
	closeSessionErr := session.Close() //client.CloseSession(session.SessionId)
	assert.NoError(t, closeSessionErr, "Should be no error")
}
