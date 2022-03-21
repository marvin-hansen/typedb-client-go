package v2

import (
	pb "github.com/marvin-hansen/go-typedb/proto/core"
	"log"
)

func (c *Client) GetAllDatabases() (allDatabases []string, err error) {
	databaseAllResult, err := c.client.DatabasesAll(c.ctx, &pb.CoreDatabaseManager_All_Req{})
	if err != nil {
		log.Println(err.Error())
		return allDatabases, err
	}

	return databaseAllResult.Names, nil
}

func (c *Client) CreateDatabase(dbName string) (res string, err error) {
	databaseCreateRes, err := c.client.DatabasesCreate(c.ctx, &pb.CoreDatabaseManager_Create_Req{Name: dbName})
	if err != nil {
		log.Println(err.Error())
		return res, err
	}

	return databaseCreateRes.String(), nil
}

func (c *Client) ExistsDatabase(dbName string) (exists bool, err error) {

	databaseExistsRes, err := c.client.DatabasesContains(c.ctx, &pb.CoreDatabaseManager_Contains_Req{Name: dbName})
	if err != nil {
		log.Println("could not get database: %w", err)
		return false, err
	}
	if databaseExistsRes.Contains {
		return true, nil
	} else {
		return false, nil
	}
}

func (c *Client) DeleteDatabase(dbName string) (res string, err error) {

	exists, err := c.ExistsDatabase(dbName)
	if err != nil {
		return "", err
	}
	if exists {
		databaseDeleteRes, err := c.client.DatabaseDelete(c.ctx, &pb.CoreDatabase_Delete_Req{Name: dbName})
		if err != nil {
			log.Println(err.Error())
			return res, err
		}

		return databaseDeleteRes.String(), nil
	} else {
		return "", nil
	}
}
