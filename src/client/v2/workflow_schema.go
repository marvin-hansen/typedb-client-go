// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import (
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/common"
)

func (c *Client) CreateDatabaseSchema(dbName, schema string) (allEntries []string, status DBStatusType, err error) {

	session, openErr := NewSession(c, dbName, common.Session_SCHEMA)
	if openErr != nil {
		return allEntries, SessionOpenError, fmt.Errorf("could not open schema session: %w", openErr)
	}

	sessionID := session.GetSessionId()

	tx, newTxErr := NewTransaction(c, sessionID)
	if newTxErr != nil {
		return allEntries, ErrorCreateTransaction, fmt.Errorf("could not create a new transaction: %w", newTxErr)
	}

	transactionId := tx.GetTransactionId()
	openTxErr := tx.OpenTransaction(sessionID, transactionId, WRITE, nil, 250)
	if openTxErr != nil {
		return nil, ErrorOpenTransaction, fmt.Errorf("could not open transaction: %w", openTxErr)
	}

	requestId := session.GetNewUID()
	metadata := map[string]string{}
	req := getDefinedQueryReq(schema, requestId, &common.Options{}, metadata)

	writeErr := tx.ExecuteTransaction(req)
	if writeErr != nil {
		return nil, ErrorWriteSchema, fmt.Errorf("could not write schema into DB: %w", writeErr)
	}

	commitErr := tx.CommitTransaction(transactionId)
	if commitErr != nil {
		rollbackErr := tx.RollbackTransaction(transactionId, metadata)
		if rollbackErr != nil {
			return nil, ErrorRollbackSchemaTransaction, fmt.Errorf("could commit schema into DB AND could NOT Rolled back transaction: %w", commitErr)
		}

		return nil, ErrorCommitSchemaTransaction, fmt.Errorf("could commit schema into DB. Rolled back transaction: %w", commitErr)
	}

	closeTxErr := tx.CloseTransaction()
	if closeTxErr != nil {
		return nil, ErrorCloseSchemaTransaction, fmt.Errorf("could not close transaction: %w", closeTxErr)
	}

	closeErr := session.Close()
	if closeErr != nil {
		return allEntries, SchemaReadError, fmt.Errorf("could not close session: %w", closeErr)
	}
	return allEntries, OK, nil
}

func (c *Client) GetDatabaseSchema(dbName string, sessionID []byte) (allEntries []string, status DBStatusType, err error) {

	openReq := getSessionOpenReq(dbName, common.Session_SCHEMA, &common.Options{})
	schemaSession, openErr := c.client.SessionOpen(c.ctx, openReq)
	if openErr != nil {
		return allEntries, SessionOpenError, openErr
	}

	//query := getSchemaQuery()
	//infer := c.config.Infer
	//explain := c.config.Explain
	//batchSize := int32(0)
	//
	//schemaTransaction, openTXErr := c.client.Transaction(c.ctx)
	//if openErr != nil {
	//	return allEntries, ErrorCreateTransaction, openTXErr
	//}
	//
	//transactionId := ksuid.New().String()
	//

	closeReq := getSessionCloseReq(schemaSession.SessionId)
	_, closeErr := c.client.SessionClose(c.ctx, closeReq)
	if closeErr != nil {
		return allEntries, SchemaReadError, closeErr
	}

	return allEntries, SessionCloseError, nil
}
