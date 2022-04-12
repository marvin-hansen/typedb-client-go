package main

import (
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/marvin-hansen/typedb-client-go/data"
	typeDB "github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/marvin-hansen/typedb-client-go/test/client/utils"
	"log"
)

const dbName = utils.DBName
const verbose = true // prints out all query results. False disables printout

func main() {
	println("Create new client with default localhost config")
	conf := typeDB.NewLocalConfig(dbName)
	client, cancel := typeDB.NewClient(conf)
	defer cancel()

	println("Create a new DB")
	createDBErr := client.DBManager.CreateDatabase(dbName)
	checkError("Could not create DB: "+dbName, createDBErr)

	println("Check if DB has been created")
	_, checkDBErr := client.DBManager.CheckDatabaseExists(dbName)
	checkError("DB Doesn't exists: "+dbName, checkDBErr)

	// See data folder for schema & data definitions
	println("Load a TypeDB schema.")
	testSchema := data.GetPhoneCallsSchema()

	println(" Write schema into TypeDB")
	createSchemaErr := client.DBManager.CreateDatabaseSchema(dbName, testSchema)
	checkError("could not create DB schema", createSchemaErr)

	println("Load Schema from the DB & print to console")
	schema, getSchemaErr := client.DBManager.GetDatabaseSchema(dbName)
	checkError("could not load schema from DB", getSchemaErr)
	printSchema(schema)

	println("Load data")
	gql, dataErr := data.GetPhoneExampleData()
	checkError("could not load data", dataErr)

	println("Open session") // bundles 1 ...n transactions in a session
	sessionID, sessionOpenErr := client.Session.NewSession(utils.DbName, common.Session_DATA)
	checkError("could not open session.", sessionOpenErr)

	println("Insert data into TypeDB")
	// for single insert, use Insert instead.
	insertError := client.Query.InsertBulk(sessionID, gql, newOptions())
	checkError("could not insert data into TypeDB.", insertError)

	println("Query: Get all people")
	q1 := data.GetQueryAllPeople()
	q1Results, q1Err := client.Query.Match(sessionID, q1, newOptions())
	checkError("could not run query.", q1Err)
	printResult(q1Results, verbose)

	println("Query: Get all phone numbers")
	q2 := data.GetQueryAllPhoneNumbers()
	q2Results, q2Err := client.Query.Match(sessionID, q2, newOptions())
	checkError("could not run query.", q2Err)
	printResult(q2Results, verbose)

	println("Query: Get only people with a phone numbers")
	q3 := data.GetQueryPersonWithPhone()
	q3Results, q3Err := client.Query.Match(sessionID, q3, newOptions())
	checkError("could not run query.", q3Err)
	printResult(q3Results, verbose)

	println("Query: Since September 14th, which customers called person X?")
	q4 := data.GetQuerySept14()
	q4Results, q4Err := client.Query.Match(sessionID, q4, newOptions())
	checkError("could not run query.", q4Err)
	printResult(q4Results, verbose)

	println("Query: Who are the people who have received a call from a London customer aged over 50 who has previously called someone aged under 20??")
	q5 := data.GetQueryLondon()
	q5Results, q5Err := client.Query.Match(sessionID, q5, newOptions())
	checkError("could not run query.", q5Err)
	printResult(q5Results, verbose)

	println("Query: Who are the common contacts of customers X and Y?")
	q6 := data.GetQueryCommonContacts()
	q6Results, q6Err := client.Query.Match(sessionID, q6, newOptions())
	checkError("could not run query.", q6Err)
	printResult(q6Results, verbose)

	println("Query: Who are the customers who 1) have all called each other and 2) have all called person X at least once?\n")
	q7 := data.GetQueryLead()
	q7Results, q7Err := client.Query.Match(sessionID, q7, newOptions())
	checkError("could not run query.", q7Err)
	printResult(q7Results, verbose)

	println("Query: How does the average call duration among customers aged under 20 compare with those aged over 40?")
	println("Query A: Under 20")
	q8 := data.GetQueryCallDurationUnder20()
	q8Results, q8Err := client.Query.Match(sessionID, q8, newOptions())
	checkError("could not run query.", q8Err)
	printResult(q8Results, verbose)

	println("Query B: Over 40")
	q9 := data.GetQueryCallDurationOver40()
	q9Results, q9Err := client.Query.Match(sessionID, q9, newOptions())
	checkError("could not run query.", q9Err)
	printResult(q9Results, verbose)

	println("Close session")
	closeSessionErr := client.Session.CloseSession(sessionID)
	checkError("could not close session.", closeSessionErr)

	println("Delete DB")
	delErr := client.DBManager.DeleteDatabase(dbName)
	checkError("could not delete DB.", delErr)

	client.Close()
}

func checkError(errMsg string, err error) {
	if err != nil {
		log.Fatal(errMsg, err)
	}
}

func printSchema(schema []string) {
	// Print schema to console
	if len(schema) > 0 {
		for _, item := range schema {
			println(item)
		}
		println()
	}
}

func printResult(queryResults []*common.QueryManager_Match_ResPart, verbose bool) {
	if verbose {
		utils.TestPrint("* Print Query Results: ")
		println("* Print results: ")
		for _, item := range queryResults {
			println(item.String())
		}
	}
}

func newOptions() *common.Options {
	return typeDB.NewOptions()
}
