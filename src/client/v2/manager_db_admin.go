package v2

import (
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/src/client/v2/requests"
	"log"
)

func (c *DBManager) GetAllDatabases() (allDatabases []string, err error) {
	req := requests.GetAllDBReq()
	databaseAllResult, err := c.client.client.DatabasesAll(c.client.ctx, req)
	if err != nil {
		log.Println(err.Error())
		return allDatabases, fmt.Errorf("could not read all database. Error: %w", err)
	}
	return databaseAllResult.Names, nil
}

func (c *DBManager) CreateDatabase(dbName string) (ok bool, err error) {
	req := requests.GetCreateDBReq(dbName)
	databaseCreateRes, err := c.client.client.DatabasesCreate(c.client.ctx, req)
	if err != nil {
		log.Println(databaseCreateRes.String())
		log.Println(err.Error())
		return false, fmt.Errorf("could not create database. Error: %w", err)
	}
	return true, nil
}

func (c *DBManager) CheckDatabaseExists(dbName string) (exists bool, err error) {
	req := requests.GetContainsDBReq(dbName)
	databaseExistsRes, dbExistErr := c.client.client.DatabasesContains(c.client.ctx, req)
	if dbExistErr != nil {
		return false, fmt.Errorf("could not check if database exists. Ensure DB connection works. Error: %w", dbExistErr)
	}
	if databaseExistsRes.Contains {
		return true, nil
	} else {
		return false, nil
	}
}

func (c *DBManager) DeleteDatabase(dbName string) (ok bool, err error) {
	exists, err := c.CheckDatabaseExists(dbName)
	if err != nil {
		return false, err
	}
	if exists {
		req := requests.GetDeleteDBReq(dbName)
		databaseDeleteRes, dbDeleteErr := c.client.client.DatabaseDelete(c.client.ctx, req)
		if dbDeleteErr != nil {
			log.Println(databaseDeleteRes.String())
			log.Println(dbDeleteErr.Error())
			return false, fmt.Errorf("could not delete database. Error: %w", dbDeleteErr)
		}
		return true, nil
	} else {
		return true, nil
	}
}
