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
* SessionManager implemented 
* TransactionManager implemented 
* QueryManager under development

Implementation & testing details in [status document](Status.md) 

## Usage 

For more detailed usage, see test files in [test folder](test/client)

```Go
package main

import (
	"github.com/marvin-hansen/typedb-client-go/data"
	v2 "github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/marvin-hansen/typedb-client-go/test/client/utils"
	"log"
)

const dbName = utils.DBName

    func main(){
		// Create new client with default localhost config
		conf := v2.NewLocalConfig(dbName)
		client, cancel := v2.NewClient(conf)
		defer cancel()
		
		// Create a new DB 
		_, err := client.DBManager.CreateDatabase(dbName)

		// Check if DB has been created 
		_, err = client.DBManager.CheckDatabaseExists(dbName)
		if err != nil{
			log.Fatal("DB Doesn't exists", err)
		}

		// Delete DB if exists. Uncomment to actually delete...
		// ok, err := client.DBManager.DeleteDatabase(dbName)
		
		// Load a TypeDB schema. See data folder  
		testSchema := data.GetPhoneCallsSchema()
		
		// Write schema into TypeDB 
		err = client.DBManager.CreateDatabaseSchema(dbName, testSchema)
		if err != nil{
			log.Fatal("could not create DB schema", err)
		}
		
		// Load the Schema from the DB 
		allEntries, err := client.DBManager.GetDatabaseSchema(dbName)
		if err != nil{
			log.Fatal("could not load schema from DB", err)
		}

		// Print schema to console 
		if len(allEntries) > 0 {
			for _, item := range allEntries {
				println(item)
			}
			println() 
		}
		
		// TODO: insert & query data
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
