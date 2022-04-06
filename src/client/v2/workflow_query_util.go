package v2

import (
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/common"
)

//RunQuery util used by all other single return value query methods
func (c *Client) RunQuery(sessionID []byte, req *common.Transaction_Req, transactionType common.Transaction_Type, options *common.Options) (*common.QueryManager_Res, error) {

	// Create a Transaction
	tx, newTxErr := c.TransactionManager.NewTransaction(sessionID, TX_READ)
	if newTxErr != nil {
		return nil, fmt.Errorf("could not create a new transaction: %w", newTxErr)
	}

	// open new transaction
	latencyMillis := int32(100)
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
	recv, recErr := tx.tx.Recv()
	if recErr != nil {
		return nil, fmt.Errorf("could not receive query response: %w", recErr)
	}

	// Extract query result
	res := recv.GetRes().GetQueryManagerRes()

	return res, nil
}
