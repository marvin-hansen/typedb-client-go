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

func (c *Client) RunStreamTx(sessionID []byte, req *common.Transaction_Req, transactionType common.Transaction_Type, options *common.Options) (queryResults []*common.QueryManager_ResPart, err error) {
	mtd := "RunStreamTx"

	dbgPrint(mtd, " Create a Transaction")
	tx, newTxErr := c.TransactionManager.NewTransaction(sessionID, transactionType)
	if newTxErr != nil {
		return nil, fmt.Errorf("could not create a new transaction: %w", newTxErr)
	}

	dbgPrint(mtd, " Open new transaction ")
	latencyMillis := int32(100)
	openTxErr := tx.OpenTransaction(sessionID, options, latencyMillis)
	if openTxErr != nil {
		return nil, fmt.Errorf("could not open transaction: %w", openTxErr)
	}

	dbgPrint(mtd, " Execute Transaction ")
	readErr := tx.ExecuteTransaction(req)
	if readErr != nil {
		return nil, fmt.Errorf("could not exevute transaction: %w", readErr)
	}

	dbgPrint(mtd, " Enter loop ")
	for {
		dbgPrint(mtd, " Get return value ")
		recs, recErr := tx.tx.Recv()
		if recErr != nil {
			return nil, fmt.Errorf("could not receive query response: %w", recErr)
		}

		dbgPrint(mtd, " Determine if stream is done")
		done := c.isStreamDone(recs.GetResPart(), true)
		if done {
			dbgPrint(mtd, " Done. Break loop ")
			break // break loop when done
		} else {
			//
			dbgPrint(mtd, " Continue. Send continue request")
			reqCont := requests.GetTransactionStreamReq()
			sendErr := tx.tx.Send(requests.GetTransactionClient(reqCont))
			if sendErr != nil {
				return nil, fmt.Errorf("could not send query request iterator: %w", sendErr)
			}
		}

		dbgPrint(mtd, " Collect results ")
		part := recs.GetResPart().GetQueryManagerResPart()
		queryResults = append(queryResults, part)
	}

	// Only write transactions can be committed
	if transactionType == TX_WRITE {
		// Commit transaction
		dbgPrint(mtd, " Commit Transaction ")
		commitErr := tx.CommitTransaction()
		if commitErr != nil {
			// Rollback commit in case of error
			rollbackErr := tx.RollbackTransaction()
			if rollbackErr != nil {
				return nil, fmt.Errorf("could NOT roll back Commit from failed transaction: %w", commitErr)
			}

			return nil, fmt.Errorf("could commit into DB. Rolled back transaction: %w", commitErr)
		}
	}

	dbgPrint(mtd, " Close Transaction ")
	closeTxErr := tx.CloseTransaction()
	if closeTxErr != nil {
		return nil, fmt.Errorf("could not close transaction: %w", closeTxErr)
	}

	return queryResults, nil
}

func (c *Client) RunStreamQuery(sessionID []byte, req *common.Transaction_Req, transactionType common.Transaction_Type, options *common.Options) (queryResults *common.QueryManager_ResPart, err error) {
	mtd := "RunStreamQuery"

	dbgPrint(mtd, " SessionID: "+hex.EncodeToString(sessionID))
	dbgPrint(mtd, " Create a Transaction")
	tx, newTxErr := c.TransactionManager.NewTransaction(sessionID, transactionType)
	if newTxErr != nil {
		return nil, fmt.Errorf("could not create a new transaction: %w", newTxErr)
	}

	transactionId := tx.GetTransactionId()
	latencyMillis := int32(100)
	dbgPrint(mtd, " Transaction ID "+hex.EncodeToString(transactionId))
	dbgPrint(mtd, " Open new transaction ")
	openTxErr := tx.OpenTransaction(sessionID, options, latencyMillis)
	if openTxErr != nil {
		return nil, fmt.Errorf("could not open transaction: %w", openTxErr)
	}

	dbgPrint(mtd, " Execute Transaction ")
	sendErr := tx.ExecuteTransaction(req)
	if sendErr != nil {
		return nil, fmt.Errorf("could not send transaction to server: %w", sendErr)
	}

	dbgPrint(mtd, " Get return value ")
	// here things get stuck pretty hard!
	// for some reason, nothing returns during insert.
	recv, recErr := tx.tx.Recv()
	if recErr != nil {
		return nil, fmt.Errorf("could not receive query response: %w", recErr)
	}
	dbgPrint(mtd, recv.String())

	dbgPrint(mtd, " Extract query result ")
	queryResults = recv.GetResPart().GetQueryManagerResPart()

	// Only write transactions can be committed
	if transactionType == TX_WRITE {
		// Commit transaction
		commitErr := tx.CommitTransaction()
		if commitErr != nil {
			// Rollback commit in case of error
			rollbackErr := tx.RollbackTransaction()
			if rollbackErr != nil {
				return nil, fmt.Errorf("could NOT roll back Commit from failed transaction: %w", commitErr)
			}

			return nil, fmt.Errorf("could commit into DB. Rolled back transaction: %w", commitErr)
		}
	}

	dbgPrint(mtd, "Close transaction")
	closErr := tx.tx.CloseSend()
	if closErr != nil {
		return nil, fmt.Errorf("could not close query transaction: %w", closErr)
	}

	dbgPrint(mtd, "Return query result")
	return queryResults, nil
}
