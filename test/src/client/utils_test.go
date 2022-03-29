package client

import (
	"fmt"
	v2 "github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"log"
)

func dbSetup(client *v2.Client, dbName string) error {
	println("* Run DB setup")

	existsDatabase, dbExistErr := client.CheckDatabaseExists(dbName)
	if dbExistErr != nil {
		return fmt.Errorf("could not check if database exists. Ensure DB connection works. Error: %w", dbExistErr)
	}

	if !existsDatabase {
		_, err := client.CreateDatabase(dbName)
		if err != nil {
			log.Println(err.Error())
			return err
		} else {
			return nil
		}
	}

	return nil
}

func dbTeardown(client *v2.Client, dbName string) error {
	println("* Run DB teardown")

	existsDatabase, dbExistErr := client.CheckDatabaseExists(dbName)
	if dbExistErr != nil {
		return fmt.Errorf("could not check if database exists. Ensure DB connection works. Error: %w", dbExistErr)
	}

	if existsDatabase {
		_, err := client.DeleteDatabase(dbName)
		if err != nil {
			log.Println(err.Error())
			return err
		} else {
			return nil
		}
	}

	return nil
}
