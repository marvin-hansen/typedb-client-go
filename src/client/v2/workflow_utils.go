// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import (
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/segmentio/ksuid"
)

// dbCheck checks if a DB with the name exists
func (c *Client) dbCheck(dbName string) (err error) {
	existsDatabase, dbExistErr := c.CheckDatabaseExists(dbName)
	if dbExistErr != nil {
		return fmt.Errorf("could not check if database exists. Ensure DB connection works. Error: %w", dbExistErr)
	}

	if !existsDatabase {
		return fmt.Errorf(" database does not exists: %w", dbExistErr)
	}

	return nil
}

func CreateNewRequestIDOptions() (requestId []byte, metadata map[string]string) {
	return CreateNewRequestID(), CreateNewRequestMetadata()
}

func CreateNewRequestIDMetaOptions() (requestId []byte, options *common.Options, metadata map[string]string) {
	return CreateNewRequestID(), CreateNewRequestOptions(), CreateNewRequestMetadata()
}

func CreateNewRequestOptions() *common.Options {
	return &common.Options{}
}

// CreateNewRequestMetadata returns a new map of empty metadata
func CreateNewRequestMetadata() map[string]string {
	return map[string]string{}
}

// CreateNewRequestID generates a new unique request ID to use in transactions
func CreateNewRequestID() []byte {
	return ksuid.New().Bytes()
}

// CreateNewStringRequestID generates a new unique request ID to use in transactions
func CreateNewStringRequestID() string {
	return ksuid.New().String()
}

func getSchemaQuery() string {
	return `
	match
		$x sub thing;
	get
		$x;
	`
}
