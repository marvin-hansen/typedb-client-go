package v2

import (
	pb "github.com/marvin-hansen/go-typedb/proto/core"
	"log"
)

func (c *Client) GetAllDatabases() (allDatabases []string, status DBStatusType, err error) {
	databaseAllResult, err := c.client.DatabasesAll(c.ctx, &pb.CoreDatabaseManager_All_Req{})
	if err != nil {
		log.Println(err.Error())
		return allDatabases, ReadAllDBError, err
	}
	return databaseAllResult.Names, OK, nil
}

func (c *Client) CreateDatabase(dbName string) (res string, status DBStatusType, err error) {
	databaseCreateRes, err := c.client.DatabasesCreate(c.ctx, &pb.CoreDatabaseManager_Create_Req{Name: dbName})
	if err != nil {
		log.Println(err.Error())
		return res, CreateError, err
	}
	return databaseCreateRes.String(), OK, nil
}

func (c *Client) ExistsDatabase(dbName string) (exists bool, status DBStatusType, err error) {

	databaseExistsRes, err := c.client.DatabasesContains(c.ctx, &pb.CoreDatabaseManager_Contains_Req{Name: dbName})
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

	exists, status, err := c.ExistsDatabase(dbName)
	if err != nil {
		return false, status, err
	}
	if exists {
		databaseDeleteRes, err := c.client.DatabaseDelete(c.ctx, &pb.CoreDatabase_Delete_Req{Name: dbName})
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
