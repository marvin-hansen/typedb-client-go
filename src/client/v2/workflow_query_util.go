package v2

import (
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/common"
)

//runQuery util used by all other single return value query methods
func (c *Client) runQuery(req []*common.Transaction_Req) (*common.QueryManager_Res, error) {

	// Create a Transaction
	tx, txErr := c.client.Transaction(c.ctx)
	if txErr != nil {
		return nil, fmt.Errorf("could not create transaction: %w", txErr)
	}

	// Send request through
	sendErr := tx.Send(getTransactionClient(req))
	if sendErr != nil {
		return nil, fmt.Errorf("could not send transaction to server: %w", sendErr)
	}

	// get return value
	recv, recErr := tx.Recv()
	if recErr != nil {
		return nil, fmt.Errorf("could not receive query response: %w", recErr)
	}

	// Extract query result
	res := recv.GetRes().GetQueryManagerRes()

	// close transaction
	closErr := tx.CloseSend()
	if closErr != nil {
		return nil, fmt.Errorf("could not close query transaction: %w", closErr)
	}

	return res, nil
}
