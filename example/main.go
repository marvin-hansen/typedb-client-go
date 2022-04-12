package main

import (
	"github.com/marvin-hansen/typedb-client-go/data"
	typeDB "github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/marvin-hansen/typedb-client-go/test/client/utils"
	"log"
)

const dbName = utils.DBName

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

	// Delete DB if exists. Uncomment to actually delete...
	// ok, err := client.DBManager.DeleteDatabase(dbName)
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
