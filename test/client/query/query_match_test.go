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

func TestMatchQuery(t *testing.T) {
	client, cancel := getClient()
	defer cancel()

	testPrint("* Create Session")
	sessionID, sessionOpenErr := client.SessionManager.NewSession(dbName, common.Session_DATA)
	assert.NoError(t, sessionOpenErr, "Should be no error")

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

	testPrint("* CloseSession Session")
	closeSessionErr := client.SessionManager.CloseSession(sessionID)
	assert.NoError(t, closeSessionErr, "Should be no error")
}