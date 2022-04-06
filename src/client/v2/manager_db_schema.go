package v2

import (
	"encoding/hex"
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/marvin-hansen/typedb-client-go/src/client/v2/requests"
)

func NewDBManager(client *Client) *DBManager {
	return &DBManager{client: client}
}

type DBManager struct {
	client *Client
}

// CreateDatabaseSchema creates schema fo the given DB
func (c *DBManager) CreateDatabaseSchema(dbName, schema string) (err error) {

	dbExistErr := c.client.dbCheck(dbName)
	if dbExistErr != nil {
		return dbExistErr
	}

	sessionID, openErr := c.client.SessionManager.NewSession(dbName, common.Session_SCHEMA)
	if openErr != nil {
		return fmt.Errorf("could not open schema session: %w", openErr)
	}

	tx, newTxErr := c.client.TransactionManager.NewTransaction(sessionID, TX_WRITE)
	if newTxErr != nil {
		return fmt.Errorf("could not create a new transaction: %w", newTxErr)
	}

	latencyMillis := int32(100)
	options := &common.Options{}

	openTxErr := tx.OpenTransaction(sessionID, options, latencyMillis)
	if openTxErr != nil {
		return fmt.Errorf("could not open transaction: %w", openTxErr)
	}

	req := requests.GetDefinedQueryReq(schema, options)

	writeErr := tx.ExecuteTransaction(req)
	if writeErr != nil {
		return fmt.Errorf("could not write schema into DB: %w", writeErr)
	}

	commitErr := tx.CommitTransaction()
	if commitErr != nil {
		rollbackErr := tx.RollbackTransaction()
		if rollbackErr != nil {
			return fmt.Errorf("could commit schema into DB AND could NOT Rolled back transaction: %w", commitErr)
		}

		return fmt.Errorf("could commit schema into DB. Rolled back transaction: %w", commitErr)
	}

	closeTxErr := tx.CloseTransaction()
	if closeTxErr != nil {
		return fmt.Errorf("could not close transaction: %w", closeTxErr)
	}

	closeErr := c.client.SessionManager.CloseSession(sessionID)
	if closeErr != nil {
		return fmt.Errorf("could not close session: %w", closeErr)
	}

	return nil
}

func (c *DBManager) GetDatabaseSchema(dbName string) (allEntries []string, err error) {
	mtd := "GetDatabaseSchema"

	dbgPrint(mtd, " Check if DB exists")
	dbExistErr := c.client.dbCheck(dbName)
	if dbExistErr != nil {
		return nil, dbExistErr
	}

	sessionID, openErr := c.client.SessionManager.NewSession(dbName, common.Session_SCHEMA)
	if openErr != nil {
		return nil, fmt.Errorf("could not open schema session: %w", openErr)
	}
	dbgPrint(mtd, " Create new session: "+hex.EncodeToString(sessionID))

	dbgPrint(mtd, " Create new GetSchemaQuery")
	query := getSchemaQuery()
	options := &common.Options{}
	req := requests.GetMatchQueryReq(query, options)

	dbgPrint(mtd, " RUN GetSchemaQuery")
	stream, err := c.client.RunStreamTx(sessionID, req, TX_READ, options)
	if err != nil {
		return nil, err
	}

	closeErr := c.client.SessionManager.CloseSession(sessionID)
	if closeErr != nil {
		return nil, fmt.Errorf("could not close session: %w", closeErr)
	}
	dbgPrint(mtd, " Close session: "+hex.EncodeToString(sessionID))

	for _, part := range stream {
		allEntries = append(allEntries, part.GetMatchResPart().String())
	}

	return allEntries, nil
}
