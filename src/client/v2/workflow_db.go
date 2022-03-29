// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import (
	"fmt"
	"log"
)

func (c *Client) GetAllDatabases() (allDatabases []string, err error) {
	req := getAllDBReq()
	databaseAllResult, err := c.client.DatabasesAll(c.ctx, req)
	if err != nil {
		log.Println(err.Error())
		return allDatabases, fmt.Errorf("could not read all database. Error: %w", err)
	}
	return databaseAllResult.Names, nil
}

func (c *Client) CreateDatabase(dbName string) (ok bool, err error) {
	req := getCreateDBReq(dbName)
	databaseCreateRes, err := c.client.DatabasesCreate(c.ctx, req)
	if err != nil {
		log.Println(databaseCreateRes.String())
		log.Println(err.Error())
		return false, fmt.Errorf("could not create database. Error: %w", err)
	}
	return true, nil
}

func (c *Client) CheckDatabaseExists(dbName string) (exists bool, err error) {
	req := getContainsDBReq(dbName)
	databaseExistsRes, dbExistErr := c.client.DatabasesContains(c.ctx, req)
	if dbExistErr != nil {
		return false, fmt.Errorf("could not check if database exists. Ensure DB connection works. Error: %w", dbExistErr)
	}
	if databaseExistsRes.Contains {
		return true, nil
	} else {
		return false, fmt.Errorf(" database does not exists: %w", dbExistErr)
	}
}

func (c *Client) DeleteDatabase(dbName string) (ok bool, err error) {
	exists, err := c.CheckDatabaseExists(dbName)
	if err != nil {
		return false, err
	}
	if exists {
		req := getDeleteDBReq(dbName)
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
