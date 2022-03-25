// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import (
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
// err :=tx.CommitTransaction(...)
// if err != nil{
// 	tx.RollbackTransaction(...)
// }
// tx.CloseTransaction()
//
func NewTransaction(client *Client, sessionId []byte) (*TransactionClient, error) {

	t, err := newTx(client)
	if err != nil {
		return nil, err
	}

	return &TransactionClient{
		client:    client,
		sessionId: sessionId,
		tx:        t,
	}, nil
}

type TransactionClient struct {
	client    *Client
	tx        core.TypeDB_TransactionClient
	sessionId []byte
}

func newTx(client *Client) (core.TypeDB_TransactionClient, error) {
	tx, txErr := client.client.Transaction(client.ctx)
	if txErr != nil {
		return nil, txErr
	} else {
		return tx, nil
	}
}

func (c *TransactionClient) executeTX(req []*common.Transaction_Req) error {

	// Send request through
	sendErr := c.tx.Send(getTransactionClient(req))
	if sendErr != nil {
		return sendErr
	}

	// get return value - I assume grpc blocks until the result arrives.
	_, recErr := c.tx.Recv()
	if recErr != nil {
		return recErr
	}

	return nil
}

func (c TransactionClient) OpenTransaction(sessionID []byte, sessionType common.Transaction_Type, options *common.Options, netMillisecondLatency int32) error {

	// Create a request with options & request ID
	r1 := getTransactionOpenReq(sessionID, sessionType, options, netMillisecondLatency)

	// Stuff req into slice/array
	var req []*common.Transaction_Req
	req[0] = r1

	execErr := c.executeTX(req)
	if execErr != nil {
		return execErr
	} else {
		return nil // no execution error
	}

}

func (c TransactionClient) CommitTransaction(transactionId []byte, metadata map[string]string) error {
	// Create a request with meta data & request ID
	r1 := getTransactionCommitReq(transactionId, metadata)

	// Stuff req into slice/array
	var req []*common.Transaction_Req
	req[0] = r1

	execErr := c.executeTX(req)
	if execErr != nil {
		return execErr
	} else {
		return nil // no execution error
	}
}

func (c TransactionClient) RollbackTransaction(transactionId []byte, metadata map[string]string) error {
	// Create a request with meta data & request ID
	r1 := getTransactionRollbackReq(transactionId, metadata)

	// Stuff req into slice/array
	var req []*common.Transaction_Req
	req[0] = r1

	execErr := c.executeTX(req)
	if execErr != nil {
		return execErr
	} else {
		return nil // no execution error
	}
}

func (c TransactionClient) CloseTransaction() error {

	// Close transaction
	txCloseErr := c.tx.CloseSend()
	if txCloseErr != nil {
		return txCloseErr
	}
	return nil
}
