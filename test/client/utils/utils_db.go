// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package utils

import (
	"fmt"
	v2 "github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"log"
)

func DBSetup(client *v2.Client, dbName string) error {
	println("* Run DB setup")

	_, dbExistErr := client.DBManager.CheckDatabaseExists(dbName)
	if dbExistErr != nil {
		return fmt.Errorf("could not check if database exists. Ensure DB connection works. Error: %w", dbExistErr)
	}

	println("Create Database: " + dbName)
	err := client.DBManager.CreateDatabase(dbName)
	if err != nil {
		log.Println(err.Error())
		return err
	} else {
		return nil
	}
}

func DBTeardown(client *v2.Client, dbName string) error {
	println("* Run DB teardown")
	_, dbExistErr := client.DBManager.CheckDatabaseExists(dbName)
	if dbExistErr != nil {
		log.Println(dbExistErr.Error())
		return dbExistErr
	} else {
		err := client.DBManager.DeleteDatabase(dbName)
		if err != nil {
			log.Println(err.Error())
			return err
		} else {
			return nil
		}
	}
}
