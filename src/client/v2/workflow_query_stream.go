package v2

import (
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/marvin-hansen/typedb-client-go/src/client/v2/requests"
)

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
//
//      Insert.ResPart insert_res_part = 103;
//      Update.ResPart update_res_part = 104;
//      Explain.ResPart explain_res_part = 105;
//    }
//  }

func (c *Client) RunInsertQuery(sessionID []byte, query string, options *common.Options) (queryResults []*common.ConceptMap, err error) {

	// run request
	req := requests.GetInsertQueryReq(query, options)
	streamQuery, queryErr := c.RunStreamTx(sessionID, req, TX_WRITE, options, true)
	if queryErr != nil {
		return nil, queryErr
	}

	for _, item := range streamQuery {
		queryResults = item.GetInsertResPart().GetAnswers()

		// queryResults = append(queryResults, item.String())

	}

	return queryResults, nil
}

func (c *Client) RunUpdateQuery(sessionID []byte, query string, options *common.Options) (queryResults []*common.QueryManager_Update_ResPart, err error) {

	// Update is a sequence of match, delte & insert. Thus TX_WRITE
	req := requests.GetUpdateQueryReq(query, options)
	streamQuery, queryErr := c.RunStreamTx(sessionID, req, TX_WRITE, options, true)
	if queryErr != nil {
		return nil, queryErr
	}

	// extract match results and stuff into queryResults collection
	for _, item := range streamQuery {
		queryResults = append(queryResults, item.GetUpdateResPart())
	}

	return queryResults, nil
}

func (c *Client) RunExplainQuery(sessionID []byte, explainableID int64, options *common.Options) (queryResults []*common.QueryManager_Explain_ResPart, err error) {

	// run query
	req := requests.GetExplainQueryReq(explainableID, options)
	streamQuery, queryErr := c.RunStreamTx(sessionID, req, TX_READ, options, true)
	if queryErr != nil {
		return nil, queryErr
	}

	// extract match results and stuff into queryResults collection
	for _, item := range streamQuery {
		queryResults = append(queryResults, item.GetExplainResPart())
	}

	return queryResults, nil
}

func (c *Client) RunMatchQuery(sessionID []byte, query string, options *common.Options) (queryResults []*common.QueryManager_Match_ResPart, err error) {

	// Create query request
	req := requests.GetMatchQueryReq(query, options)

	// run query
	streamQuery, queryErr := c.RunStreamTx(sessionID, req, TX_READ, options, true)
	if queryErr != nil {
		return nil, queryErr
	}

	// extract match results and stuff into queryResults collection
	for _, item := range streamQuery {
		queryResults = append(queryResults, item.GetMatchResPart())
		item.GetMatchResPart().GetAnswers()
	}

	return queryResults, nil
}

func (c *Client) RunMatchGroupQuery(sessionID []byte, query string, options *common.Options) (queryResults []*common.QueryManager_MatchGroup_ResPart, err error) {

	// Create query request
	req := requests.GetMatchGroupQueryReq(query, options)

	// run query
	streamQuery, queryErr := c.RunStreamTx(sessionID, req, TX_READ, options, false)
	if queryErr != nil {
		return nil, queryErr
	}

	// extract match results and stuff into queryResults collection
	for _, item := range streamQuery {
		queryResults = append(queryResults, item.GetMatchGroupResPart())
	}

	return queryResults, nil
}

func (c *Client) RunMatchGroupAggregateQuery(sessionID []byte, query string, options *common.Options) (queryResults []*common.QueryManager_MatchGroupAggregate_ResPart, err error) {

	// Create query request
	req := requests.GetMatchGroupQueryAggregateQueryReq(query, options)

	// run query
	streamQuery, queryErr := c.RunStreamTx(sessionID, req, TX_READ, options, false)
	if queryErr != nil {
		return nil, queryErr
	}

	// extract match results and stuff into queryResults collection
	for _, item := range streamQuery {
		queryResults = append(queryResults, item.GetMatchGroupAggregateResPart())
	}

	return queryResults, nil
}
