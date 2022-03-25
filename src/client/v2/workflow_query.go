// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import "github.com/marvin-hansen/typedb-client-go/common"

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

func (c *Client) RunInsertQuery(requestId []byte, query string, metadata map[string]string, explain, infer, parallel bool,
) (matchResponses []*common.QueryManager_Insert_ResPart, recErr error) {

	// construct req options
	var options = &common.Options{
		InferOpt:    &common.Options_Infer{Infer: infer},
		ExplainOpt:  &common.Options_Explain{Explain: explain},
		ParallelOpt: &common.Options_Parallel{Parallel: parallel},
	}

	// Create a request and attach meta data & request ID
	r1 := getInsertQueryReq(query, options)
	r1.Metadata = metadata
	r1.ReqId = requestId

	// Stuff req into slice/array
	var req []*common.Transaction_Req
	req[0] = r1

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

	// however, there is no GetInsertRes() or similar
	// should insert query just return QueryManager_Res ?

	return matchResponses, recErr
}

func (c *Client) RunDefineQuery(requestId []byte, query string, metadata map[string]string, explain, infer, parallel bool) (queryResponses *common.QueryManager_Res, err error) {

	// construct req options
	var options = &common.Options{
		InferOpt:    &common.Options_Infer{Infer: infer},
		ExplainOpt:  &common.Options_Explain{Explain: explain},
		ParallelOpt: &common.Options_Parallel{Parallel: parallel},
	}

	// Create a request and attach meta data & request ID
	r1 := getDefinedQueryReq(query, options)
	r1.Metadata = metadata
	r1.ReqId = requestId

	// Stuff req into slice/array
	var req []*common.Transaction_Req
	req[0] = r1

	// run query
	queryResponses, queryErr := c.runQuery(req)
	if queryErr != nil {
		return nil, queryErr
	}

	return queryResponses, nil
}

func (c *Client) RunDeleteQuery(requestId []byte, query string, metadata map[string]string, explain, infer, parallel bool) (queryResponses *common.QueryManager_Res, err error) {

	// construct req options
	var options = &common.Options{
		InferOpt:    &common.Options_Infer{Infer: infer},
		ExplainOpt:  &common.Options_Explain{Explain: explain},
		ParallelOpt: &common.Options_Parallel{Parallel: parallel},
	}

	// Create a request and attach meta data & request ID
	r1 := getDeleteQueryReq(query, options)
	r1.Metadata = metadata
	r1.ReqId = requestId

	// Stuff req into slice/array
	var req []*common.Transaction_Req
	req[0] = r1

	// run query
	queryResponses, queryErr := c.runQuery(req)
	if queryErr != nil {
		return nil, queryErr
	}

	return queryResponses, nil
}
