// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import (
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/marvin-hansen/typedb-client-go/src/client/v2/requests"
)

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

func (c *Client) RunDefineQuery(sessionID []byte, query string, options *common.Options) (queryResponses *common.QueryManager_Define_Res, err error) {
	// Create a request and attach meta data & request ID
	req := requests.GetDefinedQueryReq(query, options)
	// run query
	res, queryErr := c.runQuery(sessionID, req, options)
	if queryErr != nil {
		return nil, queryErr
	} else {
		return res.GetDefineRes(), nil
	}
}

func (c *Client) RunUndefineQuery(sessionID []byte, query string, options *common.Options) (queryResponses *common.QueryManager_Undefine_Res, err error) {
	// Create a request and attach meta data & request ID
	req := requests.GetUndefinedQueryReq(query, options)
	// run query
	res, queryErr := c.runQuery(sessionID, req, options)
	if queryErr != nil {
		return nil, queryErr
	} else {
		return res.GetUndefineRes(), nil
	}
}

func (c *Client) RunMatchAggregateQuery(sessionID []byte, query string, options *common.Options) (queryResponses *common.QueryManager_MatchAggregate_Res, err error) {
	// Create a request and attach meta data & request ID
	req := requests.GetMatchAggregateQueryReq(query, options)
	// run query
	res, queryErr := c.runQuery(sessionID, req, options)

	if queryErr != nil {
		return nil, queryErr
	} else {
		return res.GetMatchAggregateRes(), nil
	}
}

func (c *Client) RunDeleteQuery(sessionID []byte, query string, options *common.Options) (queryResponses *common.QueryManager_Delete_Res, err error) {
	// Create a request and attach meta data & request ID
	req := requests.GetDeleteQueryReq(query, options)
	// run query
	res, queryErr := c.runQuery(sessionID, req, options)
	if queryErr != nil {
		return nil, queryErr
	} else {
		return res.GetDeleteRes(), nil
	}
}
