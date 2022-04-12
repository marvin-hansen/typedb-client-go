// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package query

import (
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/marvin-hansen/typedb-client-go/test/client/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

const verbose = true

func TestMatchQuery(t *testing.T) {
	client, cancel := utils.GetClient()
	defer cancel()
	sessionID, sessionOpenErr := client.SessionManager.NewSession(utils.DbName, common.Session_DATA)
	assert.NoError(t, sessionOpenErr, "Should be no error")
	utils.TestPrint("* Create Session ")

	utils.TestPrint("* Get Test Query")
	query := utils.GetTestQuery()

	utils.TestPrint("* Query TypeDB")
	options := v2.CreateNewRequestOptions()
	queryResults, queryErr := client.RunMatchQuery(sessionID, query, options)
	if queryErr != nil {
		println(fmt.Errorf("could not create transaction: %w", queryErr))
	}
	assert.NoError(t, queryErr, "Should be no error")
	assert.NotNil(t, queryResults, "Query should return some results")

	printResult(queryResults, verbose)

	utils.TestPrint("* CloseSession Session ")
	closeSessionErr := client.SessionManager.CloseSession(sessionID)
	assert.NoError(t, closeSessionErr, "Should be no error")
	client.Close() // close client
}

func TestMatchQueryAllPhoneNumbers(t *testing.T) {
	client, cancel := utils.GetClient()
	defer cancel()
	sessionID, sessionOpenErr := client.SessionManager.NewSession(utils.DbName, common.Session_DATA)
	assert.NoError(t, sessionOpenErr, "Should be no error")
	utils.TestPrint("* Create Session ")

	utils.TestPrint("* Get Test Query")
	query := utils.GetTestQueryAllPhone()

	utils.TestPrint("* Query TypeDB")
	options := v2.CreateNewRequestOptions()
	queryResults, queryErr := client.RunMatchQuery(sessionID, query, options)
	if queryErr != nil {
		println(fmt.Errorf("could not create transaction: %w", queryErr))
	}
	assert.NoError(t, queryErr, "Should be no error")
	assert.NotNil(t, queryResults, "Query should return some results")

	printResult(queryResults, verbose)

	utils.TestPrint("* CloseSession Session ")
	closeSessionErr := client.SessionManager.CloseSession(sessionID)
	assert.NoError(t, closeSessionErr, "Should be no error")
	client.Close() // close client
}
