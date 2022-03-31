// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import (
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/src/client/v2/requests"
	"log"
)

func (c *Client) GetAllDatabases() (allDatabases []string, err error) {
	req := requests.GetAllDBReq()
	databaseAllResult, err := c.client.DatabasesAll(c.ctx, req)
	if err != nil {
		log.Println(err.Error())
		return allDatabases, fmt.Errorf("could not read all database. Error: %w", err)
	}
	return databaseAllResult.Names, nil
}

func (c *Client) CreateDatabase(dbName string) (ok bool, err error) {
	req := requests.GetCreateDBReq(dbName)
	databaseCreateRes, err := c.client.DatabasesCreate(c.ctx, req)
	if err != nil {
		log.Println(databaseCreateRes.String())
		log.Println(err.Error())
		return false, fmt.Errorf("could not create database. Error: %w", err)
	}
	return true, nil
}

func (c *Client) CheckDatabaseExists(dbName string) (exists bool, err error) {
	req := requests.GetContainsDBReq(dbName)
	databaseExistsRes, dbExistErr := c.client.DatabasesContains(c.ctx, req)
	if dbExistErr != nil {
		return false, fmt.Errorf("could not check if database exists. Ensure DB connection works. Error: %w", dbExistErr)
	}
	if databaseExistsRes.Contains {
		return true, nil
	} else {
		return false, nil
	}
}

func (c *Client) DeleteDatabase(dbName string) (ok bool, err error) {
	exists, err := c.CheckDatabaseExists(dbName)
	if err != nil {
		return false, err
	}
	if exists {
		req := requests.GetDeleteDBReq(dbName)
		databaseDeleteRes, dbDeleteErr := c.client.DatabaseDelete(c.ctx, req)
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
