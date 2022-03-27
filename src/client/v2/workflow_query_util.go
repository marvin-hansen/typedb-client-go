package v2

import "github.com/marvin-hansen/typedb-client-go/common"

//runQuery util used by all other specific query methods below
func (c *Client) runQuery(req []*common.Transaction_Req) (*common.QueryManager_Res, error) {

	// Create a Transaction
	tx, txErr := c.client.Transaction(c.ctx)
	if txErr != nil {
		return nil, txErr
	}

	// Send request through
	sendErr := tx.Send(getTransactionClient(req))
	if sendErr != nil {
		return nil, sendErr
	}

	// Close send -  I suppose?
	txCloseErr := tx.CloseSend()
	if txCloseErr != nil {
		return nil, txCloseErr
	}

	// get return value - I assume grpc blocks until the result arrives.
	recv, recErr := tx.Recv()
	if recErr != nil {
		return nil, recErr
	}

	// Extract query result
	res := recv.GetRes().GetQueryManagerRes()

	return res, nil
}
