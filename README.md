# typedb-client-go

Go client for [TypeDB](https://vaticle.com/typedb). TypeDB is an innovative database using a static type system over a
graph DB. Life science research organization use TypeDB for drug discovery, genomic and protein data. In absence of a Go
client, I decided to write one for evaluation purpose. I've documented my experience with
TypeDB [here](considerations.md).

Supported TypeDB versions:
* [2.6](https://github.com/vaticle/typedb/releases/tag/2.6.4)
* [2.7](https://github.com/vaticle/typedb/releases/tag/2.7.1)
* [2.8](https://github.com/vaticle/typedb/releases/tag/2.8.1)

Protocol: 
* [2.6.1](https://github.com/vaticle/typedb-protocol/releases/tag/2.6.1)

NOT supported:
* Cluster

Status: Not maintained!

* DBManager fully implemented
* SessionManager fully implemented
* TransactionManager fully implemented
* QueryManager implemented, but not fully tested. 

Implementation & testing details in [status document](Status.md)

## Setup

Ensure a TypeDB server is running on your local machine. Consult
the [TypeDB documentation](https://docs.vaticle.com/docs/general/introduction) for details.

For simple testing, run TypeDB locally without an external volume:

```Bash
docker run --name typedb -d -p 1729:1729 vaticle/typedb:latest
```

Notice, this instance does not store data on disk. All data will be lost when the instance ends.

For development, run TypeDB with an external volume that stores data on the mounded folder on disk.

```Bash
docker run --name typedb -d -v /path/to/volume:/typedb-all-linux/server/data/ -p 1729:1729 vaticle/typedb:latest
```

## Usage

For more detailed usage, see:

* Tests in [test folder](test/client)
* Example code in [example folder](example)
* Example [schema](data/schema.go)
* Example [data](data/phone_example_data.go)

```Go
package main

import (
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/marvin-hansen/typedb-client-go/data"
	typeDB "github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/marvin-hansen/typedb-client-go/test/client/utils"
	"log"
)

const dbName = "TestDB" 
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

## Acknowledgment

I express my sincere gratefulness to entire [Humans of Julia Team](https://github.com/Humans-of-Julia) that developed
the Julia client for TypeDB, which I used extensively as an inspiration for this Go client. I also owe much of my
gratitude to [Frank Urbach](https://github.com/FrankUrbach) for his in-depth technical support. Without his deep
expertise of theTypeDB protocol, this Golang client form would not be possible.

## Author

* Marvin Hansen
* GPG key ID: 210D39BC
* Github key ID: 369D5A0B210D39BC
* GPG Fingerprint: 4B18 F7B2 04B9 7A72 967E 663E 369D 5A0B 210D 39BC
* Public key: [key](pubkey.txt)

## Licence 

* [MIT Licence](LICENSE)
* Software is "as is" without any warranty. 
