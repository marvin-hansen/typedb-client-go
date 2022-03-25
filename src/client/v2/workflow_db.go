// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import (
	"log"
)

func (c *Client) GetAllDatabases() (allDatabases []string, status DBStatusType, err error) {
	req := getAllDBReq()
	databaseAllResult, err := c.client.DatabasesAll(c.ctx, req)
	if err != nil {
		log.Println(err.Error())
		return allDatabases, ReadAllDBError, err
	}
	return databaseAllResult.Names, OK, nil
}

func (c *Client) CreateDatabase(dbName string) (ok bool, status DBStatusType, err error) {
	req := getCreateDBReq(dbName)
	databaseCreateRes, err := c.client.DatabasesCreate(c.ctx, req)
	if err != nil {
		log.Println(databaseCreateRes.String())
		log.Println(err.Error())
		return false, CreateError, err
	}
	return true, OK, nil
}

func (c *Client) CheckDatabaseExists(dbName string) (exists bool, status DBStatusType, err error) {
	req := getContainsDBReq(dbName)
	databaseExistsRes, err := c.client.DatabasesContains(c.ctx, req)
	if err != nil {
		log.Println("could not get database: %w", err)
		return false, CheckExistsError, err
	}
	if databaseExistsRes.Contains {
		return true, OK, nil
	} else {
		return false, DBNotExists, nil
	}
}

func (c *Client) DeleteDatabase(dbName string) (ok bool, status DBStatusType, err error) {
	exists, status, err := c.CheckDatabaseExists(dbName)
	if err != nil {
		return false, status, err
	}
	if exists {
		req := getDeleteDBReq(dbName)
		databaseDeleteRes, err := c.client.DatabaseDelete(c.ctx, req)
		if err != nil {
			log.Println(databaseDeleteRes.String())
			log.Println(err.Error())
			return false, DeleteError, err
		}
		return true, OK, nil
	} else {
		return true, DBNotExists, nil
	}
}
