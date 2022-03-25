// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import (
	"github.com/marvin-hansen/typedb-client-go/common"
)

func (c *Client) runQuery(req []*common.Transaction_Req) (*common.QueryManager_Res, error) {

	// Create a Transaction
	tx, txErr := c.client.Transaction(c.ctx)
	if txErr != nil {
		return nil, txErr
	}

	// Send request through
	sendErr := tx.Send(&common.Transaction_Client{Reqs: req})
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

func (c *Client) RunInsertQuery(requestId string,
	query string,
	metadata map[string]string,
	explain bool,
	batchSize int32,
	latencyMillis int32) (matchResponses []*common.QueryManager_Insert_ResPart, recErr error) {

	if batchSize == 0 {
		batchSize = 2147483647
	}

	// @TODO: Construct req options. Skipped for now.
	var options *common.Options = nil

	// Create a Request i.e. insert request. cast it into an slice/array
	var req []*common.Transaction_Req
	req[0] = getInsertQueryReq(query, options)

	// run query
	res, queryErr := c.runQuery(req)
	if queryErr != nil {
		return nil, queryErr
	}

	// FIXME:
	// res is of type QueryManager_Res but for whatever reason I can only
	// extract the types listed below..

	res.GetDefineRes()
	res.GetDeleteRes()
	res.GetMatchAggregateRes()
	res.GetUndefineRes()

	// however there is no GetInsertRes() or similar to

	return matchResponses, recErr

}
