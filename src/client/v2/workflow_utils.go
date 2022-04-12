// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import (
	"encoding/hex"
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/common"
	"google.golang.org/grpc"
	"log"
)

func byteToString(sessionID []byte) string {
	return hex.EncodeToString(sessionID)
}

func checkPrintErr(err error, errorMsg string) {
	if err != nil {
		log.Println("Error:", err.Error())
		log.Println("Error Message: ", errorMsg)
	}
}

func checkConnection(conn *grpc.ClientConn, err error) {
	if err != nil {
		checkPrintErr(err, connErr)
	}
	if conn == nil {
		log.Fatal(connErr)
	}
}

// dbCheck checks if a DB with the name exists
func (c *Client) dbCheck(dbName string) (err error) {
	_, dbExistErr := c.DBManager.CheckDatabaseExists(dbName)
	if dbExistErr != nil {
		return fmt.Errorf("could not check if database exists. Ensure DB connection works. Error: %w", dbExistErr)
	}

	return nil
}

func CreateNewRequestOptions() *common.Options {
	return &common.Options{}
}

func getSchemaQuery() string {
	return `
	match
		$x sub thing;
	get
		$x;
	`
}
