// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import "github.com/marvin-hansen/typedb-client-go/common"

//
// Methods with singular return type i.e. one request -> one result
//
// From the proto spec in common/query.proto @line 47
//
//   message Res {
//    oneof res {
//      Define.Res define_res = 100;
//      Undefine.Res undefine_res = 101;
//      MatchAggregate.Res match_aggregate_res = 102;
//      Delete.Res delete_res = 104;
//    }
//  }
//

func (c *Client) RunDefineQuery(requestId []byte, query string, metadata map[string]string, options *common.Options) (queryResponses *common.QueryManager_Define_Res, err error) {
	// Create a request and attach meta data & request ID
	r1 := getDefinedQueryReq(query, options, requestId, metadata)
	// Stuff req into slice/array
	req := []*common.Transaction_Req{r1}
	// run query
	res, queryErr := c.runQuery(req)
	if queryErr != nil {
		return nil, queryErr
	} else {
		return res.GetDefineRes(), nil
	}
}

func (c *Client) RunUndefineQuery(requestId []byte, query string, metadata map[string]string, options *common.Options) (queryResponses *common.QueryManager_Undefine_Res, err error) {
	// Create a request and attach meta data & request ID
	r1 := getUndefinedQueryReq(query, options, requestId, metadata)
	// Stuff req into slice/array
	req := []*common.Transaction_Req{r1}
	// run query
	res, queryErr := c.runQuery(req)
	if queryErr != nil {
		return nil, queryErr
	} else {
		return res.GetUndefineRes(), nil
	}
}

func (c *Client) RunMatchAggregateQuery(requestId []byte, query string, metadata map[string]string, options *common.Options) (queryResponses *common.QueryManager_MatchAggregate_Res, err error) {
	// Create a request and attach meta data & request ID
	r1 := getMatchAggregateQueryReq(query, options, requestId, metadata)
	// Stuff req into slice/array
	req := []*common.Transaction_Req{r1}
	// run query
	res, queryErr := c.runQuery(req)

	if queryErr != nil {
		return nil, queryErr
	} else {
		return res.GetMatchAggregateRes(), nil
	}
}

func (c *Client) RunDeleteQuery(requestId []byte, query string, metadata map[string]string, options *common.Options) (queryResponses *common.QueryManager_Delete_Res, err error) {
	// Create a request and attach meta data & request ID
	r1 := getDeleteQueryReq(query, options, requestId, metadata)
	// Stuff req into slice/array
	req := []*common.Transaction_Req{r1}
	// run query
	res, queryErr := c.runQuery(req)
	if queryErr != nil {
		return nil, queryErr
	} else {
		return res.GetDeleteRes(), nil
	}
}
