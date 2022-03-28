// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import (
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/marvin-hansen/typedb-client-go/core"
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

	t, err := newTx(client)
	if err != nil {
		return nil, err
	}

	return &TransactionManager{
		client:    client,
		sessionId: sessionId,
		tx:        t,
	}, nil
}

type TransactionManager struct {
	client    *Client
	tx        core.TypeDB_TransactionClient
	sessionId []byte
}

func newTx(client *Client) (core.TypeDB_TransactionClient, error) {
	tx, txErr := client.client.Transaction(client.ctx)
	if txErr != nil {
		return nil, fmt.Errorf("could not create transaction: %w", txErr)
	} else {
		return tx, nil
	}
}

func (c TransactionManager) OpenTransaction(sessionID []byte, sessionType common.Transaction_Type, options *common.Options, netMillisecondLatency int32) error {
	// Create a request with options & request ID
	r1 := getTransactionOpenReq(sessionID, sessionType, options, netMillisecondLatency)
	// Stuff req into slice/array
	req := []*common.Transaction_Req{r1}

	execErr := c.ExecuteTransaction(req)
	if execErr != nil {
		return fmt.Errorf("could not open transaction: %w", execErr)
	} else {
		return nil
	}
}

func (c *TransactionManager) ExecuteTransaction(req []*common.Transaction_Req) error {
	// Send request through
	sendErr := c.tx.Send(getTransactionClient(req))
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

func (c TransactionManager) CommitTransaction(transactionId []byte, metadata map[string]string) error {
	// Create a request with meta data & request ID
	r1 := getTransactionCommitReq(transactionId, metadata)
	// Stuff req into slice/array
	req := []*common.Transaction_Req{r1}

	commitErr := c.ExecuteTransaction(req)
	if commitErr != nil {
		return fmt.Errorf("could not commit transaction: %w", commitErr)
	} else {
		return nil
	}
}

func (c TransactionManager) RollbackTransaction(transactionId []byte, metadata map[string]string) error {
	// Create a request with meta data & request ID
	r1 := getTransactionRollbackReq(transactionId, metadata)
	// Stuff req into slice/array
	req := []*common.Transaction_Req{r1}
	rollbackErr := c.ExecuteTransaction(req)
	if rollbackErr != nil {
		return fmt.Errorf("could not rollback transaction: %w", rollbackErr)
	} else {
		return nil
	}
}

func (c TransactionManager) CloseTransaction() error {
	closeErr := c.tx.CloseSend()
	if closeErr != nil {
		return fmt.Errorf("could not close query transaction: %w", closeErr)
	}
	return nil
}
