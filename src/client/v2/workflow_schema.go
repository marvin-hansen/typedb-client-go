// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import (
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/common"
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

// CreateDatabaseSchema creates schema fo the given DB
func (c *Client) CreateDatabaseSchema(dbName, schema string) (err error) {

	dbExistErr := c.dbCheck(dbName)
	if dbExistErr != nil {
		return dbExistErr
	}

	session, openErr := NewSession(c, dbName, common.Session_SCHEMA)
	if openErr != nil {
		return fmt.Errorf("could not open schema session: %w", openErr)
	}

	sessionID := session.GetSessionId()

	tx, newTxErr := NewTransaction(c, sessionID)
	if newTxErr != nil {
		return fmt.Errorf("could not create a new transaction: %w", newTxErr)
	}

	transactionId := tx.GetTransactionId()
	latencyMillis := int32(100)
	options := &common.Options{}
	openTxErr := tx.OpenTransaction(sessionID, transactionId, TX_WRITE, options, latencyMillis)
	if openTxErr != nil {
		return fmt.Errorf("could not open transaction: %w", openTxErr)
	}

	requestId := CreateNewRequestID()
	metadata := CreateNewRequestMetadata()
	req := getDefinedQueryReq(schema, requestId, &common.Options{}, metadata)

	writeErr := tx.ExecuteTransaction(req)
	if writeErr != nil {
		return fmt.Errorf("could not write schema into DB: %w", writeErr)
	}

	commitErr := tx.CommitTransaction(transactionId)
	if commitErr != nil {
		rollbackErr := tx.RollbackTransaction(transactionId, metadata)
		if rollbackErr != nil {
			return fmt.Errorf("could commit schema into DB AND could NOT Rolled back transaction: %w", commitErr)
		}

		return fmt.Errorf("could commit schema into DB. Rolled back transaction: %w", commitErr)
	}

	closeTxErr := tx.CloseTransaction()
	if closeTxErr != nil {
		return fmt.Errorf("could not close transaction: %w", closeTxErr)
	}

	closeErr := session.Close()
	if closeErr != nil {
		return fmt.Errorf("could not close session: %w", closeErr)
	}

	return nil
}

// GetDatabaseSchema returns the schema for the DB
func (c *Client) GetDatabaseSchema(dbName string) (allEntries []string, err error) {

	dbExistErr := c.dbCheck(dbName)
	if dbExistErr != nil {
		return nil, dbExistErr
	}

	session, openErr := NewSession(c, dbName, common.Session_SCHEMA)
	if openErr != nil {
		return nil, fmt.Errorf("could not open schema session: %w", openErr)
	}

	sessionID := session.GetSessionId()

	tx, newTxErr := NewTransaction(c, sessionID)
	if newTxErr != nil {
		return nil, fmt.Errorf("could not create a new transaction: %w", newTxErr)
	}

	transactionId := tx.GetTransactionId()
	latencyMillis := int32(100)
	options := &common.Options{}
	openTxErr := tx.OpenTransaction(sessionID, transactionId, TX_READ, options, latencyMillis)
	if openTxErr != nil {
		return nil, fmt.Errorf("could not open transaction: %w", openTxErr)
	}

	query := getSchemaQuery()
	requestId := CreateNewRequestID()
	metadata := CreateNewRequestMetadata()
	req := getMatchQueryReq(query, options, requestId, metadata)

	readErr := tx.ExecuteTransaction(req)
	if readErr != nil {
		return nil, fmt.Errorf("could not read schema: %w", readErr)
	}

	for {
		// Get return value
		recs, recErr := tx.tx.Recv()
		if recErr != nil {
			return nil, fmt.Errorf("could not receive query response: %w", recErr)
		}

		// Extract state of current partial result
		state := recs.GetResPart().GetStreamResPart().GetState()

		// When the server sends a Stream.ResPart with state = CONTINUE
		// it indicates that there are more answers to fetch,
		// so the client should respond with Stream.Req
		if state == CONTINUE {
			// Create a request and attach meta data & request ID
			reqCont := getTransactionStreamReq()
			// run query
			_, queryErr := c.runQuery(sessionID, reqCont, options)
			if queryErr != nil {
				return nil, fmt.Errorf("could not send query request iterator: %w", queryErr)
			}
		}

		// If the Stream.ResPart has state = DONE,
		// it indicates that there are no more answers to fetch,
		// so the client doesn't need to respond.
		if state == DONE {
			break

		} else {
			part := recs.GetResPart().GetQueryManagerResPart()
			allEntries = append(allEntries, part.String())
		}
	}

	closeTxErr := tx.CloseTransaction()
	if closeTxErr != nil {
		return nil, fmt.Errorf("could not close transaction: %w", closeTxErr)
	}

	closeErr := session.Close()
	if closeErr != nil {
		return nil, fmt.Errorf("could not close session: %w", closeErr)
	}

	return allEntries, nil
}
