# typedb-client-go

Go client for [TypeDB](https://vaticle.com/typedb)

Status: Early stage work in progress! 

Supported TypeDB versions:
* [2.6](https://github.com/vaticle/typedb/releases/tag/2.6.4)
* [2.7](https://github.com/vaticle/typedb/releases/tag/2.7.1)
* [2.8](https://github.com/vaticle/typedb/releases/tag/2.8.1)

Protocol: 
* [2.6.1](https://github.com/vaticle/typedb-protocol/releases/tag/2.6.1)

NOT supported:
* Cluster

Current status: Active Development

* DBManager fully implemented
* SessionManager fully implemented
* TransactionManager fully implemented
* QueryManager implemented

Implementation & testing details in [status document](Status.md)

## Usage

For more detailed usage, see:

* Example in [exa,ple folder](example)
* Tests in [test folder](test/client)

```Go
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
	gql, dataErr := data.GetPhoneCallsDataGql()
	checkError("could not load data", dataErr)

	println("Open session") // bundles transactions
	sessionID, sessionOpenErr := client.Session.NewSession(utils.DbName, common.Session_DATA)
	checkError("could not open session.", sessionOpenErr)

	println("Insert data into TypeDB")
	// for single insert, use Insert instead.
	insertError := client.Query.InsertBulk(sessionID, gql, typeDB.NewOptions())
	checkError("could not insert data into TypeDB.", insertError)

	println("Query TypeDB: Get all people")
	q1 := data.GetQueryAllPeople()
	q1Results, q1Err := client.Query.Match(sessionID, q1, typeDB.NewOptions())
	checkError("could not run query.", q1Err)
	printResult(q1Results, verbose)

	println("Query TypeDB: Get all phone numbers")
	q2 := data.GetQueryAllPhoneNumbers()
	q2Results, q2Err := client.Query.Match(sessionID, q2, typeDB.NewOptions())
	checkError("could not run query.", q2Err)
	printResult(q2Results, verbose)

	println("Query TypeDB: Get only people with a phone numbers")
	q3 := data.GetQueryPersonWithPhone()
	q3Results, q3Err := client.Query.Match(sessionID, q3, typeDB.NewOptions())
	checkError("could not run query.", q3Err)
	printResult(q3Results, verbose)

	println("Close session")
	closeSessionErr := client.Session.CloseSession(sessionID)
	checkError("could not close session.", closeSessionErr)

	// Delete DB, if exists. Uncomment to actually delete...
	// ok, err := client.DBManager.DeleteDatabase(dbName)

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

```

## Make reference

```bash 
Setup: 
    make check                  Checks whether all project requirements are present.
     
Dev: 
    make build                  Builds the code base incrementally (fast).
    make rebuild                Rebuilds dependencies, build files, & code base (slow). Use after go mod changes.
    make tests                  Runs all tests. See scripts/dev/test.sh for details.
    make stats                  Shows the latest project stats. 
```

## Author

* Marvin Hansen 
* GPG key ID: 210D39BC
* Github key ID: 369D5A0B210D39BC
* GPG Fingerprint: 4B18 F7B2 04B9 7A72 967E  663E 369D 5A0B 210D 39BC 
* Public key: [key](pubkey.txt)

## Licence 

* [MIT Licence](LICENSE)
* Software is "as is" without any warranty. 
