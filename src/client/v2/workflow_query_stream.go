package v2

import "github.com/marvin-hansen/typedb-client-go/common"

//
// Methods with streaming results i.e one initial request -> multiple partial stream results.
//
// From the proto spec in common/query.proto @line 56
//
//   message ResPart {
//    oneof res {
//      Match.ResPart match_res_part = 100;
//      MatchGroup.ResPart match_group_res_part = 101;
//      MatchGroupAggregate.ResPart match_group_aggregate_res_part = 102;
//      Insert.ResPart insert_res_part = 103;
//      Update.ResPart update_res_part = 104;
//      Explain.ResPart explain_res_part = 105;
//    }
//  }

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
