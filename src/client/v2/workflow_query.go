// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

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

// Methods with singular return type i.e. one request -> one result

func (c *Client) RunDefinedQuery(requestId []byte, query string, metadata map[string]string, options *common.Options) (queryResponses *common.QueryManager_Define_Res, err error) {

	// Create a request and attach meta data & request ID
	r1 := getDefinedQueryReq(query, options, requestId, metadata)
	// Stuff req into slice/array
	var req []*common.Transaction_Req
	req[0] = r1

	// run query
	res, queryErr := c.runQuery(req)
	if queryErr != nil {
		return nil, queryErr
	} else {
		return res.GetDefineRes(), nil
	}
}

func (c *Client) RunUnDefinedQuery(requestId []byte, query string, metadata map[string]string, options *common.Options) (queryResponses *common.QueryManager_Undefine_Res, err error) {

	// Create a request and attach meta data & request ID
	r1 := getUndefinedQueryReq(query, options, requestId, metadata)
	// Stuff req into slice/array
	var req []*common.Transaction_Req
	req[0] = r1

	// run query
	res, queryErr := c.runQuery(req)
	if queryErr != nil {
		return nil, queryErr
	} else {
		return res.GetUndefineRes(), nil
	}
}

func (c *Client) RunDeleteQuery(requestId []byte, query string, metadata map[string]string, options *common.Options) (queryResponses *common.QueryManager_Delete_Res, err error) {

	// Create a request and attach meta data & request ID
	r1 := getDeleteQueryReq(query, options, requestId, metadata)
	// Stuff req into slice/array
	var req []*common.Transaction_Req
	req[0] = r1

	// run query
	res, queryErr := c.runQuery(req)
	if queryErr != nil {
		return nil, queryErr
	} else {
		return res.GetDeleteRes(), nil
	}
}

func (c *Client) RunMatchAggregateQuery(requestId []byte, query string, metadata map[string]string, options *common.Options) (queryResponses *common.QueryManager_MatchAggregate_Res, err error) {

	// Create a request and attach meta data & request ID
	r1 := getMatchAggregateQueryReq(query, options, requestId, metadata)
	// Stuff req into slice/array
	var req []*common.Transaction_Req
	req[0] = r1

	// run query
	res, queryErr := c.runQuery(req)

	if queryErr != nil {
		return nil, queryErr
	} else {
		return res.GetMatchAggregateRes(), nil
	}
}

//
// Methods with streaming results i.e one initial request -> multiple partial stream results.
//

func (c *Client) RunInsertQuery(requestId []byte, query string, metadata map[string]string, options *common.Options) (matchResponses []*common.QueryManager_Insert_ResPart, recErr error) {

	// Create a request and attach meta data & request ID
	r1 := getInsertQueryReq(query, options, requestId, metadata)

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

func (c *Client) RunUpdateQuery(requestId []byte, query string, metadata map[string]string, options *common.Options) (queryResponses *common.QueryManager_Res, err error) {

	// Create a request and attach meta data & request ID
	r1 := getUpdateQueryReq(query, options, requestId, metadata)
	// Stuff req into slice/array
	var req []*common.Transaction_Req
	req[0] = r1

	// run query
	queryResponses, queryErr := c.runQuery(req)
	if queryErr != nil {
		return nil, queryErr
	} else {
		return queryResponses, nil
	}
}

func (c *Client) RunExplainQuery(requestId []byte, explainableID int64, metadata map[string]string, options *common.Options) (queryResponses *common.QueryManager_Res, err error) {

	// Create a request and attach meta data & request ID
	r1 := getExplainQueryReq(explainableID, options, requestId, metadata)
	// Stuff req into slice/array
	var req []*common.Transaction_Req
	req[0] = r1

	// run query
	queryResponses, queryErr := c.runQuery(req)
	if queryErr != nil {
		return nil, queryErr
	} else {
		return queryResponses, nil
	}
}

// 	// FIXME: Match queries return type
//
// The old golang client implemented match queries using a for-loop over continue flag i.e.
// for {
//		transactionResponse, err := transactionClient.Recv()
//		if err != nil {
//			return nil, fmt.Errorf("could not receive query response: %w", err)
//		}
//
//		if transactionResponse.GetContinue() {
// ...
//
// code snippet above taken from:
// https://github.com/taliesins/typedb-client-go/blob/main/v2/client/client.go#L300
//
//
// The GetContinue() method isn't available in the 2.6.1 specification anymore.
//
// I suppose the for loop may have been replaced with the "stream" concept, but the spec's doesn't any indicator of that.
//
// The below implementation is just an initial proof of concept until things have been clarified.

func (c *Client) RunMatchQuery(requestId []byte, query string, metadata map[string]string, options *common.Options) (queryResponses *common.QueryManager_Res, err error) {

	// Create a request and attach meta data & request ID
	r1 := getMatchQueryReq(query, options, requestId, metadata)
	// Stuff req into slice/array
	var req []*common.Transaction_Req
	req[0] = r1

	// run query
	queryResponses, queryErr := c.runQuery(req)

	if queryErr != nil {
		return nil, queryErr
	} else {
		return queryResponses, nil
	}
}

func (c *Client) RunMatchGroupQuery(requestId []byte, query string, metadata map[string]string, options *common.Options) (queryResponses *common.QueryManager_Res, err error) {

	// Create a request and attach meta data & request ID
	r1 := getMatchGroupQueryReq(query, options, requestId, metadata)
	// Stuff req into slice/array
	var req []*common.Transaction_Req
	req[0] = r1

	// run query
	queryResponses, queryErr := c.runQuery(req)

	if queryErr != nil {
		return nil, queryErr
	} else {
		return queryResponses, nil
	}
}

func (c *Client) RunMatchGroupAggregateQuery(requestId []byte, query string, metadata map[string]string, options *common.Options) (queryResponses *common.QueryManager_Res, err error) {

	// Create a request and attach meta data & request ID
	r1 := getMatchGroupQueryAggregateQueryReq(query, options, requestId, metadata)
	// Stuff req into slice/array
	var req []*common.Transaction_Req
	req[0] = r1

	// run query
	queryResponses, queryErr := c.runQuery(req)

	if queryErr != nil {
		return nil, queryErr
	} else {
		return queryResponses, nil
	}
}
