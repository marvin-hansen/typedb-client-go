package v2

import (
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/common"
)

const (
	CONTINUE = common.Transaction_Stream_CONTINUE
	DONE     = common.Transaction_Stream_DONE
	TX_READ  = common.Transaction_READ
	TX_WRITE = common.Transaction_WRITE
)

func (c *Client) runStreamTx(sessionID []byte, transactionType common.Transaction_Type, req *common.Transaction_Req, options *common.Options) (queryResults []*common.QueryManager_ResPart, err error) {

	// Create a Transaction
	tx, newTxErr := NewTransaction(c, sessionID)
	if newTxErr != nil {
		return nil, fmt.Errorf("could not create a new transaction: %w", newTxErr)
	}

	// Create request meta data
	metadata := CreateNewRequestMetadata()

	// run request in query
	streamQuery, queryErr := c.runStreamQuery(tx, transactionType, req, options)
	if queryErr != nil {
		return nil, queryErr
	}

	// commit transaction
	transactionId := tx.GetSessionId()
	commitErr := tx.CommitTransaction(transactionId)
	if commitErr != nil {
		rollbackErr := tx.RollbackTransaction(transactionId, metadata)
		if rollbackErr != nil {
			return nil, fmt.Errorf("could commit insert into DB AND could NOT Rolled back transaction: %w", commitErr)
		}

		return nil, fmt.Errorf("could commit insert into DB. Rolled back transaction: %w", commitErr)
	}

	// close transaction
	closeTxErr := tx.CloseTransaction()
	if closeTxErr != nil {
		return nil, fmt.Errorf("could not close transaction: %w", closeTxErr)
	}

	return streamQuery, nil
}

//runStreamQuery util used by all other streaming return value query methods
func (c *Client) runStreamQuery(tx *TransactionManager, transactionType common.Transaction_Type, req *common.Transaction_Req, options *common.Options) (queryResults []*common.QueryManager_ResPart, err error) {

	sessionID := tx.GetSessionId()
	transactionId := tx.GetTransactionId()

	// Create open transaction request
	netMillisecondLatency := int32(150)
	openReq := getTransactionOpenReq(sessionID, transactionId, transactionType, options, netMillisecondLatency)

	// Send request through
	sendErr := tx.ExecuteTransaction(openReq, req)
	if sendErr != nil {
		return nil, fmt.Errorf("could not send transaction to server: %w", sendErr)
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
			queryResults = append(queryResults, part)
		}
	}

	// Close transaction after the last part returned
	closErr := tx.CloseTransaction()
	if closErr != nil {
		return nil, fmt.Errorf("could not close query transaction: %w", closErr)
	}

	return queryResults, nil
}
