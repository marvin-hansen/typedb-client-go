package v2

import (
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/segmentio/ksuid"
)

//runQuery util used by all other single return value query methods
func (c *Client) runQuery(sessionID []byte, req *common.Transaction_Req, options *common.Options) (*common.QueryManager_Res, error) {

	// Create a Transaction
	tx, txErr := c.client.Transaction(c.ctx)
	if txErr != nil {
		return nil, fmt.Errorf("could not create transaction: %w", txErr)
	}

	transactionId := ksuid.New().Bytes()

	// Create open transaction request
	transactionType := READ
	netMillisecondLatency := int32(150)
	openReq := getTransactionOpenReq(sessionID, transactionId, transactionType, options, netMillisecondLatency)
	// Stuff req into slice/array
	reqArray := []*common.Transaction_Req{openReq, req}

	// Send request through
	sendErr := tx.Send(getTransactionClient(reqArray))
	if sendErr != nil {
		return nil, fmt.Errorf("could not send transaction to server: %w", sendErr)
	}

	// Get return value
	recv, recErr := tx.Recv()
	if recErr != nil {
		return nil, fmt.Errorf("could not receive query response: %w", recErr)
	}

	// Extract query result
	res := recv.GetRes().GetQueryManagerRes()

	// Close transaction
	closErr := tx.CloseSend()
	if closErr != nil {
		return nil, fmt.Errorf("could not close query transaction: %w", closErr)
	}

	return res, nil
}
