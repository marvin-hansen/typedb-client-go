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
//
//     MatchGroupAggregate.ResPart match_group_aggregate_res_part = 102;
//      Insert.ResPart insert_res_part = 103;
//      Update.ResPart update_res_part = 104;
//      Explain.ResPart explain_res_part = 105;
//    }
//  }

func (c *Client) RunMatchQuery(sessionID, requestId []byte, query string, metadata map[string]string, options *common.Options) (queryResults []*common.QueryManager_Match_ResPart, err error) {

	// Create query request
	req := getMatchQueryReq(query, options, requestId, metadata)

	// run query
	streamQuery, queryErr := c.runStreamQuery(sessionID, READ, req, options)
	if queryErr != nil {
		return nil, queryErr
	}

	// extract match results and stuff into queryResults collection
	for _, item := range streamQuery {
		queryResults = append(queryResults, item.GetMatchResPart())
	}

	return queryResults, nil
}

func (c *Client) RunMatchGroupQuery(sessionID, requestId []byte, query string, metadata map[string]string, options *common.Options) (queryResults []*common.QueryManager_MatchGroup_ResPart, err error) {

	// Create query request
	req := getMatchGroupQueryReq(query, options, requestId, metadata)

	// run query
	streamQuery, queryErr := c.runStreamQuery(sessionID, READ, req, options)
	if queryErr != nil {
		return nil, queryErr
	}

	// extract match results and stuff into queryResults collection
	for _, item := range streamQuery {
		queryResults = append(queryResults, item.GetMatchGroupResPart())
	}

	return queryResults, nil
}

func (c *Client) RunMatchGroupAggregateQuery(sessionID, requestId []byte, query string, metadata map[string]string, options *common.Options) (queryResults []*common.QueryManager_MatchGroupAggregate_ResPart, err error) {

	// Create query request
	req := getMatchGroupQueryReq(query, options, requestId, metadata)

	// run query
	streamQuery, queryErr := c.runStreamQuery(sessionID, READ, req, options)
	if queryErr != nil {
		return nil, queryErr
	}

	// extract match results and stuff into queryResults collection
	for _, item := range streamQuery {
		queryResults = append(queryResults, item.GetMatchGroupAggregateResPart())
	}

	return queryResults, nil
}
