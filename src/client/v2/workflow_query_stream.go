package v2

import (
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/marvin-hansen/typedb-client-go/src/client/v2/requests"
)

func (c *Client) RunInsertQuery(sessionID []byte, query string, options *common.Options) (err error) {
	req := requests.GetInsertQueryReq(query, options)
	_, queryErr := c.RunStreamTx(sessionID, req, TX_WRITE, options, true)
	if queryErr != nil {
		return queryErr
	}
	return nil
}

func (c *Client) RunInsertBulkQuery(sessionID []byte, queries []string, options *common.Options) (err error) {
	for _, query := range queries {
		req := requests.GetInsertQueryReq(query, options)
		_, queryErr := c.RunStreamTx(sessionID, req, TX_WRITE, options, true)
		if queryErr != nil {
			return queryErr
		}
	}

	return nil
}

func (c *Client) RunUpdateQuery(sessionID []byte, query string, options *common.Options) (queryResults []*common.QueryManager_Update_ResPart, err error) {

	// Update is a sequence of match, delete & insert => TX_WRITE
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

	req := requests.GetMatchQueryReq(query, options)
	streamQuery, queryErr := c.RunStreamTx(sessionID, req, TX_READ, options, false)
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
