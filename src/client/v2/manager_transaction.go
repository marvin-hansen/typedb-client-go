// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import (
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/marvin-hansen/typedb-client-go/core"
	"github.com/marvin-hansen/typedb-client-go/src/client/v2/requests"
	"github.com/segmentio/ksuid"
)

// NewTransaction creates one atomic transaction options with all methods
// operating on that transaction. Each transaction must be opened & closed.
// Notice, one session may run multiple transactions.
//
// The recommended transaction workflow:
//
// tx := NewTransaction(client, sessionID)
// tx.OpenTransaction(...)
// tx.ExecuteTransaction(...)
// err := tx.CommitTransaction(...)
// if err != nil{
// 	tx.RollbackTransaction(...)
// }
// tx.CloseTransaction()
//
func NewTransaction(client *Client, sessionId []byte) (*TransactionManager, error) {

	transactionClient, err := newTx(client)
	if err != nil {
		return nil, err
	}

	transactionId := ksuid.New().Bytes()

	return &TransactionManager{
		client:        client,
		sessionId:     sessionId,
		transactionId: transactionId,
		tx:            transactionClient,
	}, nil
}

// TransactionManager encapsulates one single transaction for a specific session.
// Note, a session may hold one or more transactions.
type TransactionManager struct {
	client        *Client
	tx            core.TypeDB_TransactionClient
	sessionId     []byte
	transactionId []byte
}

func newTx(client *Client) (core.TypeDB_TransactionClient, error) {
	tx, txErr := client.client.Transaction(client.ctx)
	if txErr != nil {
		return nil, fmt.Errorf("could not create transaction: %w", txErr)
	} else {
		return tx, nil
	}
}

func (c TransactionManager) GetSessionId() []byte {
	return c.sessionId
}

// GetTransactionId returns the unique ID of the transaction
func (c TransactionManager) GetTransactionId() []byte {
	return c.transactionId
}

// OpenTransaction needs to be called first to initiate a transaction on the server.
func (c TransactionManager) OpenTransaction(sessionID, transactionId []byte, sessionType common.Transaction_Type, options *common.Options, netMillisecondLatency int32) error {
	// Create a request with options & request ID
	req := requests.GetTransactionOpenReq(sessionID, transactionId, sessionType, options, netMillisecondLatency)
	// Stuff req into slice/array

	execErr := c.ExecuteTransaction(req)
	if execErr != nil {
		return fmt.Errorf("could not open transaction: %w", execErr)
	} else {
		return nil
	}
}

// ExecuteTransaction needs to be called each time to send a request to the server.
func (c *TransactionManager) ExecuteTransaction(req ...*common.Transaction_Req) error {
	// Send request through
	sendErr := c.tx.Send(requests.GetTransactionClient(req))
	if sendErr != nil {
		return fmt.Errorf("could not send transaction to server: %w", sendErr)
	}

	// check for return value error
	_, recErr := c.tx.Recv()
	if recErr != nil {
		return fmt.Errorf("could not receive transaction response: %w", recErr)
	}
	return nil
}

// CommitTransaction needs to be called only once to finalize all previous steps.
func (c TransactionManager) CommitTransaction(transactionId []byte) error {
	// Create a request with meta data & request ID
	metadata := map[string]string{}
	req := requests.GetTransactionCommitReq(transactionId, metadata)

	commitErr := c.ExecuteTransaction(req)
	if commitErr != nil {
		return fmt.Errorf("could not commit transaction: %w", commitErr)
	} else {
		return nil
	}
}

// RollbackTransaction needs to be called whenever a commit fails to restore the previous state.
func (c TransactionManager) RollbackTransaction(transactionId []byte, metadata map[string]string) error {
	// Create a request with meta data & request ID
	req := requests.GetTransactionRollbackReq(transactionId, metadata)
	rollbackErr := c.ExecuteTransaction(req)
	if rollbackErr != nil {
		return fmt.Errorf("could not rollback transaction: %w", rollbackErr)
	} else {
		return nil
	}
}

// CloseTransaction call only once at the end to close the transaction.
func (c TransactionManager) CloseTransaction() error {
	closeErr := c.tx.CloseSend()
	if closeErr != nil {
		return fmt.Errorf("could not close query transaction: %w", closeErr)
	}
	return nil
}
