package utils

import (
	"github.com/marvin-hansen/typedb-client-go/data"
	v2 "github.com/marvin-hansen/typedb-client-go/src/client/v2"
)

func SchemaSetup(client *v2.Client, dbName string) (err error) {

	ok, dbExistErr := client.DBManager.CheckDatabaseExists(dbName)
	if dbExistErr != nil {
		return dbExistErr
	}

	if !ok {
		err = DBSetup(client, dbName)
		if err != nil {
			return err
		}
	}

	_, schemaErr := client.DBManager.GetDatabaseSchema(dbName)
	if schemaErr == nil {
		return schemaErr
	}

	testSchema := data.GetPhoneCallsSchema()
	createSchemaErr := client.DBManager.CreateDatabaseSchema(dbName, testSchema)
	if createSchemaErr != nil {
		return createSchemaErr
	}

	return nil
}
