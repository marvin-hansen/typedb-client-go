package v2

import (
	"encoding/hex"
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/marvin-hansen/typedb-client-go/src/client/v2/requests"
)

const (
	CONTINUE = common.Transaction_Stream_CONTINUE
	DONE     = common.Transaction_Stream_DONE
	TX_READ  = common.Transaction_READ
	TX_WRITE = common.Transaction_WRITE
)

func (c *Client) runStreamTx(sessionID []byte, transactionType common.Transaction_Type, req *common.Transaction_Req, options *common.Options) (queryResults []*common.QueryManager_ResPart, err error) {

	// Create a Transaction
	tx, newTxErr := c.TransactionManager.NewTransaction(sessionID)
	if newTxErr != nil {
		return nil, fmt.Errorf("could not create a new transaction: %w", newTxErr)
	}

	// open new transaction
	transactionId := tx.GetTransactionId()
	latencyMillis := int32(100)
	openTxErr := tx.OpenTransaction(sessionID, transactionId, transactionType, options, latencyMillis)
	if openTxErr != nil {
		return nil, fmt.Errorf("could not open transaction: %w", openTxErr)
	}

	readErr := tx.ExecuteTransaction(req)
	if readErr != nil {
		return nil, fmt.Errorf("could not exevute transaction: %w", readErr)
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
			reqCont := requests.GetTransactionStreamReq(transactionId)
			reqArray := []*common.Transaction_Req{reqCont}
			sendErr := tx.tx.Send(requests.GetTransactionClient(reqArray))
			if sendErr != nil {
				return nil, fmt.Errorf("could not send query request iterator: %w", sendErr)
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

	// Only write transactions can be committed
	if transactionType == TX_WRITE {
		// Create request meta data
		metadata := CreateNewRequestMetadata()

		// Commit transaction
		commitErr := tx.CommitTransaction(transactionId)
		if commitErr != nil {
			// Rollback commit in case of error
			rollbackErr := tx.RollbackTransaction(transactionId, metadata)
			if rollbackErr != nil {
				return nil, fmt.Errorf("could NOT roll back Commit from failed transaction: %w", commitErr)
			}

			return nil, fmt.Errorf("could commit into DB. Rolled back transaction: %w", commitErr)
		}
	}

	// CloseSession transaction
	closeTxErr := tx.CloseTransaction()
	if closeTxErr != nil {
		return nil, fmt.Errorf("could not close transaction: %w", closeTxErr)
	}

	return queryResults, nil
}

func (c *Client) runStreamQuery(sessionID []byte, transactionType common.Transaction_Type, req *common.Transaction_Req, options *common.Options) (queryResults *common.QueryManager_ResPart, err error) {
	mtd := "runStreamQuery"

	dbgPrint(mtd, " sessionID: "+hex.EncodeToString(sessionID))
	dbgPrint(mtd, " Create a Transaction")
	tx, newTxErr := c.TransactionManager.NewTransaction(sessionID)
	if newTxErr != nil {
		return nil, fmt.Errorf("could not create a new transaction: %w", newTxErr)
	}

	transactionId := tx.GetTransactionId()
	latencyMillis := int32(100)
	dbgPrint(mtd, " new transaction "+hex.EncodeToString(transactionId))
	dbgPrint(mtd, " open new transaction ")
	openTxErr := tx.OpenTransaction(sessionID, transactionId, transactionType, options, latencyMillis)
	if openTxErr != nil {
		return nil, fmt.Errorf("could not open transaction: %w", openTxErr)
	}

	dbgPrint(mtd, " Send request through ")
	sendErr := tx.ExecuteTransaction(req)
	if sendErr != nil {
		return nil, fmt.Errorf("could not send transaction to server: %w", sendErr)
	}

	dbgPrint(mtd, " Get return value ")
	// here things get stuck pretty hard!
	// for some reason, nothing returns
	recv, recErr := tx.tx.Recv()
	if recErr != nil {
		return nil, fmt.Errorf("could not receive query response: %w", recErr)
	}

	dbgPrint(mtd, recv.String())

	dbgPrint(mtd, " Extract query result ")
	queryResults = recv.GetResPart().GetQueryManagerResPart()

	dbgPrint(mtd, "Close transaction")
	closErr := tx.tx.CloseSend()
	if closErr != nil {
		return nil, fmt.Errorf("could not close query transaction: %w", closErr)
	}

	dbgPrint(mtd, "Return query result")
	return queryResults, nil
}
