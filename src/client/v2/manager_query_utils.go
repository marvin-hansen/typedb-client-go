package v2

import (
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

func (c *Client) runStreamTx(sessionID []byte, req *common.Transaction_Req, transactionType common.Transaction_Type, options *common.Options, singleResPart bool) (queryResults []*common.QueryManager_ResPart, err error) {
	mtd := "runStreamTx"

	tx, newTxErr := c.Transaction.NewTransaction(sessionID, transactionType)
	if newTxErr != nil {
		return nil, fmt.Errorf("could not create a new transaction: %w", newTxErr)
	}
	dbgPrint(mtd, " Create a new Transaction: "+byteToString(tx.GetTransactionId()))

	dbgPrint(mtd, " Open new transaction ")
	latencyMillis := int32(1000)
	openTxErr := tx.OpenTransaction(sessionID, options, latencyMillis)
	if openTxErr != nil {
		return nil, fmt.Errorf("could not open transaction: %w", openTxErr)
	}

	dbgPrint(mtd, " Execute Transaction ")
	readErr := tx.ExecuteTransaction(req)
	if readErr != nil {
		return nil, fmt.Errorf("could not exevute transaction: %w", readErr)
	}

	if singleResPart {
		recv, recErr := tx.ReceiveResult()
		if recErr != nil {
			return nil, fmt.Errorf("could not receive query response: %w", recErr)
		}
		dbgPrint(mtd, " Get return value: "+recv.String())

		dbgPrint(mtd, " Extract query result ")
		dbgPrint(mtd, " Collect results ")
		part := recv.GetResPart().GetQueryManagerResPart()
		queryResults = append(queryResults, part)

	} else {

		dbgPrint(mtd, " Enter loop ")
		for {
			dbgPrint(mtd, " Get return value ")
			recs, recErr := tx.ReceiveResult() //
			if recErr != nil {
				return nil, fmt.Errorf("could not receive query response: %w", recErr)
			}

			dbgPrint(mtd, " Collect results ")
			part := recs.GetResPart().GetQueryManagerResPart()
			queryResults = append(queryResults, part)

			dbgPrint(mtd, " Determine if stream is done")
			done := c.isStreamDone(recs.GetResPart(), true)
			if done {
				dbgPrint(mtd, " Done. Break loop ")
				break // break loop when done
			} else {
				//
				dbgPrint(mtd, " Continue loop. Send continue request")
				reqCont := requests.GetTransactionStreamReq()
				sendErr := tx.tx.Send(requests.GetTransactionClient(reqCont))
				if sendErr != nil {
					return nil, fmt.Errorf("could not send query request iterator: %w", sendErr)
				}
			}
		}
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

func (c *Client) isStreamDone(reqResult *common.Transaction_ResPart, stream bool) (loopBreak bool) {
	if stream {
		state := reqResult.GetStreamResPart().GetState()
		if state == DONE {
			return true
		}
		if state == CONTINUE {
			return false
		}
	}
	return false // no stream
}

//RunQuery util used by all other single return value query methods
func (c *Client) runQuery(sessionID []byte, req *common.Transaction_Req, transactionType common.Transaction_Type, options *common.Options) (*common.QueryManager_Res, error) {

	// Create a Transaction
	tx, newTxErr := c.Transaction.NewTransaction(sessionID, transactionType)
	if newTxErr != nil {
		return nil, fmt.Errorf("could not create a new transaction: %w", newTxErr)
	}

	// open new transaction
	latencyMillis := int32(1000)
	openTxErr := tx.OpenTransaction(sessionID, options, latencyMillis)
	if openTxErr != nil {
		return nil, fmt.Errorf("could not open transaction: %w", openTxErr)
	}

	//  Execute transaction
	sendErr := tx.ExecuteTransaction(req)
	if sendErr != nil {
		return nil, fmt.Errorf("could not send transaction to server: %w", sendErr)
	}

	// Only write transactions can be committed
	if transactionType == TX_WRITE {
		// Create request meta data

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

	// Close transaction
	closeTxErr := tx.CloseTransaction()
	if closeTxErr != nil {
		return nil, fmt.Errorf("could not close transaction: %w", closeTxErr)
	}

	// Get return value
	recv, recErr := tx.ReceiveResult()
	if recErr != nil {
		return nil, fmt.Errorf("could not receive query response: %w", recErr)
	}

	// Extract query result
	res := recv.GetRes().GetQueryManagerRes()

	return res, nil
}
