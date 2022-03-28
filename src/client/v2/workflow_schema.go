// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import (
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/common"
)

// dbCheck checks if a DB with the name exists
func (c *Client) dbCheck(dbName string) (status DBStatusType, err error) {
	existsDatabase, status, dbExistErr := c.CheckDatabaseExists(dbName)
	if dbExistErr != nil {
		return CheckExistsError, fmt.Errorf("could not check if database exists. Ensure DB connection works. Error: %w", dbExistErr)
	}
	if status != OK {
		return CheckExistsError, fmt.Errorf("error checking if database exists: %w", dbExistErr)
	}

	if !existsDatabase {
		return DBNotExists, fmt.Errorf(" database does not exists: %w", dbExistErr)
	}

	return OK, nil
}

// CreateDatabaseSchema creates schema fo the given DB
func (c *Client) CreateDatabaseSchema(dbName, schema string) (status DBStatusType, err error) {

	status, dbExistErr := c.dbCheck(dbName)
	if dbExistErr != nil {
		return status, dbExistErr
	}

	session, openErr := NewSession(c, dbName, common.Session_SCHEMA)
	if openErr != nil {
		return SessionOpenError, fmt.Errorf("could not open schema session: %w", openErr)
	}

	sessionID := session.GetSessionId()

	tx, newTxErr := NewTransaction(c, sessionID)
	if newTxErr != nil {
		return ErrorCreateTransaction, fmt.Errorf("could not create a new transaction: %w", newTxErr)
	}

	transactionId := tx.GetTransactionId()
	latencyMillis := int32(100)
	options := &common.Options{InferOpt: &common.Options_Infer{Infer: true}, ExplainOpt: &common.Options_Explain{Explain: true}}
	openTxErr := tx.OpenTransaction(sessionID, transactionId, TX_WRITE, options, latencyMillis)
	if openTxErr != nil {
		return ErrorOpenTransaction, fmt.Errorf("could not open transaction: %w", openTxErr)
	}

	requestId := tx.CreateNewRequestID()
	metadata := tx.CreateNewRequestMetadata()
	req := getDefinedQueryReq(schema, requestId, &common.Options{}, metadata)

	writeErr := tx.ExecuteTransaction(req)
	if writeErr != nil {
		return ErrorWriteSchema, fmt.Errorf("could not write schema into DB: %w", writeErr)
	}

	commitErr := tx.CommitTransaction(transactionId)
	if commitErr != nil {
		rollbackErr := tx.RollbackTransaction(transactionId, metadata)
		if rollbackErr != nil {
			return ErrorRollbackSchemaTransaction, fmt.Errorf("could commit schema into DB AND could NOT Rolled back transaction: %w", commitErr)
		}

		return ErrorCommitSchemaTransaction, fmt.Errorf("could commit schema into DB. Rolled back transaction: %w", commitErr)
	}

	closeTxErr := tx.CloseTransaction()
	if closeTxErr != nil {
		return ErrorCloseSchemaTransaction, fmt.Errorf("could not close transaction: %w", closeTxErr)
	}

	closeErr := session.Close()
	if closeErr != nil {
		return SchemaReadError, fmt.Errorf("could not close session: %w", closeErr)
	}

	return OK, nil
}

// GetDatabaseSchema returns the schema for the DB
func (c *Client) GetDatabaseSchema(dbName string) (allEntries []string, status DBStatusType, err error) {

	status, dbExistErr := c.dbCheck(dbName)
	if dbExistErr != nil {
		return nil, status, dbExistErr
	}

	session, openErr := NewSession(c, dbName, common.Session_SCHEMA)
	if openErr != nil {
		return nil, SessionOpenError, fmt.Errorf("could not open schema session: %w", openErr)
	}

	sessionID := session.GetSessionId()

	tx, newTxErr := NewTransaction(c, sessionID)
	if newTxErr != nil {
		return nil, ErrorCreateTransaction, fmt.Errorf("could not create a new transaction: %w", newTxErr)
	}

	transactionId := tx.GetTransactionId()
	latencyMillis := int32(100)
	options := &common.Options{InferOpt: &common.Options_Infer{Infer: true}, ExplainOpt: &common.Options_Explain{Explain: true}}
	openTxErr := tx.OpenTransaction(sessionID, transactionId, TX_READ, options, latencyMillis)
	if openTxErr != nil {
		return nil, ErrorOpenTransaction, fmt.Errorf("could not open transaction: %w", openTxErr)
	}

	query := getSchemaQuery()
	requestId := tx.CreateNewRequestID()
	metadata := tx.CreateNewRequestMetadata()
	req := getMatchQueryReq(query, options, requestId, metadata)

	readErr := tx.ExecuteTransaction(req)
	if readErr != nil {
		return nil, ErrorReadSchema, fmt.Errorf("could not read schema: %w", openTxErr)
	}

	queryResults, queryErr := c.runSchemaQuery(tx.tx, sessionID)
	if queryErr != nil {
		return nil, ErrorQuerySchema, fmt.Errorf("could not query schema: %w", openTxErr)

	}

	closeTxErr := tx.CloseTransaction()
	if closeTxErr != nil {
		return nil, ErrorCloseSchemaTransaction, fmt.Errorf("could not close transaction: %w", closeTxErr)
	}

	closeErr := session.Close()
	if closeErr != nil {
		return nil, SessionCloseError, fmt.Errorf("could not close session: %w", closeErr)
	}

	// converting each query result item into its string representation
	for _, item := range queryResults {
		allEntries = append(allEntries, item.String())
	}

	return allEntries, OK, nil
}
