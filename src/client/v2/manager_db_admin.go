package v2

import (
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

	tx, newTxErr := NewTransaction(c.client, sessionID)
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
	req := requests.GetDefinedQueryReq(schema, requestId, &common.Options{}, metadata)

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

	closeErr := c.client.SessionManager.CloseSession(sessionID)
	if closeErr != nil {
		return fmt.Errorf("could not close session: %w", closeErr)
	}

	return nil
}

// GetDatabaseSchema returns the schema for the DB
func (c *DBManager) GetDatabaseSchema(dbName string) (allEntries []string, err error) {

	dbExistErr := c.client.dbCheck(dbName)
	if dbExistErr != nil {
		return nil, dbExistErr
	}

	sessionID, openErr := c.client.SessionManager.NewSession(dbName, common.Session_SCHEMA)
	if openErr != nil {
		return nil, fmt.Errorf("could not open schema session: %w", openErr)
	}

	tx, newTxErr := NewTransaction(c.client, sessionID)
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
	req := requests.GetMatchQueryReq(query, options, requestId, metadata)

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
			reqCont := requests.GetTransactionStreamReq(transactionId)
			// run query
			_, queryErr := c.client.runQuery(sessionID, reqCont, options)
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

	closeErr := c.client.SessionManager.CloseSession(sessionID)
	if closeErr != nil {
		return nil, fmt.Errorf("could not close session: %w", closeErr)
	}

	return allEntries, nil
}
