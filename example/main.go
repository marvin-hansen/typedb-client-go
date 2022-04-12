package main

import (
	"github.com/marvin-hansen/typedb-client-go/data"
	typeDB "github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/marvin-hansen/typedb-client-go/test/client/utils"
	"log"
)

const dbName = utils.DBName

func main() {
	// Create new client with default localhost config
	conf := typeDB.NewLocalConfig(dbName)
	client, cancel := typeDB.NewClient(conf)
	defer cancel()

	// Create a new DB
	_, err := client.DBManager.CreateDatabase(dbName)

	// Check if DB has been created
	_, err = client.DBManager.CheckDatabaseExists(dbName)
	if err != nil {
		log.Fatal("DB Doesn't exists", err)
	}

	// Delete DB if exists. Uncomment to actually delete...
	// ok, err := client.DBManager.DeleteDatabase(dbName)

	// Load a TypeDB schema. See data folder
	testSchema := data.GetPhoneCallsSchema()

	// Write schema into TypeDB
	err = client.DBManager.CreateDatabaseSchema(dbName, testSchema)
	if err != nil {
		log.Fatal("could not create DB schema", err)
	}

	// Load the Schema from the DB
	allEntries, err := client.DBManager.GetDatabaseSchema(dbName)
	if err != nil {
		log.Fatal("could not load schema from DB", err)
	}

	// Print schema to console
	if len(allEntries) > 0 {
		for _, item := range allEntries {
			println(item)
		}
		println()
	}

}
