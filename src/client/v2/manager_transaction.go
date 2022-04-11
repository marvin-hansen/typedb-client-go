// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import (
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/marvin-hansen/typedb-client-go/core"
	"github.com/marvin-hansen/typedb-client-go/src/client/v2/requests"
	"github.com/segmentio/ksuid"
)

func NewTransactionManager(client *Client) *TransactionManager {
	return &TransactionManager{
		client: client,
	}
}

type TransactionManager struct {
	client *Client
}

// NewTransaction creates one atomic transaction with all methods
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
func (t *TransactionManager) NewTransaction(sessionId []byte, transactionType common.Transaction_Type) (*Transaction, error) {

	transactionClient, err := newTx(t.client)
	if err != nil {
		return nil, err
	}

	transactionId := ksuid.New().Bytes()

	return &Transaction{
		client:          t.client,
		sessionId:       sessionId,
		transactionId:   transactionId,
		transactionType: transactionType,
		tx:              transactionClient,
		isOpen:          false,
	}, nil
}

// Transaction encapsulates one single transaction for a specific session.
// Note, a session may hold one or more transactions.
type Transaction struct {
	client          *Client
	tx              core.TypeDB_TransactionClient
	sessionId       []byte
	transactionId   []byte
	transactionType common.Transaction_Type
	isOpen          bool
}

func newTx(client *Client) (core.TypeDB_TransactionClient, error) {
	tx, txErr := client.client.Transaction(client.ctx)
	if txErr != nil {
		return nil, fmt.Errorf("could not create transaction: %w", txErr)
	} else {
		return tx, nil
	}
}

func (c Transaction) GetSessionId() []byte {
	return c.sessionId
}

// GetTransactionId returns the unique ID of the transaction
func (c Transaction) GetTransactionId() []byte {
	return c.transactionId
}

// OpenTransaction needs to be called first to initiate a transaction on the server.
func (c Transaction) OpenTransaction(sessionID []byte, options *common.Options, netMillisecondLatency int32) error {

	//mtd := "OpenTransaction"
	//dbgPrint(mtd, " Create a tx Open request for tx: "+byteToString(c.transactionId))
	req := requests.GetTransactionOpenReq(sessionID, c.transactionType, options, netMillisecondLatency)

	//dbgPrint(mtd, " Execute OPEN Request ")
	execErr := c.ExecuteTransaction(req)
	if execErr != nil {
		return fmt.Errorf("could not open transaction: %w", execErr)
	} else {

		// we have to pull the OpenTX Ack here as confirmation from the server.
		_, recErr := c.ReceiveResult()
		if recErr != nil {
			return fmt.Errorf("could not receive Tx Open response: %w", recErr)
		}
		//dbgPrint(mtd, " Receive open Ack for Tx: "+byteToString(recv.GetRes().ReqId))
		// dbgPrint(mtd, " Get open Ack for Tx: "+(recv.String()))

		c.isOpen = true
		return nil
	}
}

// ExecuteTransaction needs to be called each time to send a request to the server.
func (c *Transaction) ExecuteTransaction(req *common.Transaction_Req) error {
	// create new Tx client to wrap req once more...
	txCl := requests.GetTransactionClient(req)
	// Send wrapped request to server
	sendErr := c.tx.Send(txCl)
	if sendErr != nil {
		return fmt.Errorf("could not send transaction to server: %w", sendErr)
	}
	return nil
}

// ReceiveResult Returns the result of the TX. Call multiple times on case of a stream.
func (c *Transaction) ReceiveResult() (recv *common.Transaction_Server, err error) {
	recv, err = c.tx.Recv()
	if err != nil {
		return nil, err
	}
	return recv, nil
}

// CommitTransaction needs to be called only once to finalize all previous steps.
func (c Transaction) CommitTransaction() error {
	// Create a request with meta data & request ID
	req := requests.GetTransactionCommitReq()
	commitErr := c.ExecuteTransaction(req)
	if commitErr != nil {
		return fmt.Errorf("could not commit transaction: %w", commitErr)
	} else {
		return nil
	}
}

// RollbackTransaction needs to be called whenever a commit fails to restore the previous state.
func (c Transaction) RollbackTransaction() error {
	// Create a request with meta data & request ID
	req := requests.GetTransactionRollbackReq()
	rollbackErr := c.ExecuteTransaction(req)
	if rollbackErr != nil {
		return fmt.Errorf("could not rollback transaction: %w", rollbackErr)
	} else {
		return nil
	}
}

// CloseTransaction call only once at the end to close the transaction.
func (c Transaction) CloseTransaction() error {
	closeErr := c.tx.CloseSend()
	if closeErr != nil {
		return fmt.Errorf("could not close query transaction: %w", closeErr)
	}
	c.isOpen = false
	return nil
}
