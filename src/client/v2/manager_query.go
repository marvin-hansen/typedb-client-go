// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import (
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/marvin-hansen/typedb-client-go/src/client/v2/requests"
)

func newQueryManager(client *Client) *QueryManager {
	return &QueryManager{client: client}
}

type QueryManager struct {
	client *Client
}

func (q *QueryManager) Insert(sessionID []byte, query string, options *common.Options) (err error) {
	c := q.client
	req := requests.GetInsertQueryReq(query, options)
	_, queryErr := c.runStreamTx(sessionID, req, TX_WRITE, options, true)
	if queryErr != nil {
		return queryErr
	}
	return nil
}

func (q *QueryManager) InsertBulk(sessionID []byte, queries []string, options *common.Options) (err error) {
	c := q.client
	for _, query := range queries {
		req := requests.GetInsertQueryReq(query, options)
		_, queryErr := c.runStreamTx(sessionID, req, TX_WRITE, options, true)
		if queryErr != nil {
			return queryErr
		}
	}
	return nil
}

func (q *QueryManager) Update(sessionID []byte, query string, options *common.Options) (queryResults []*common.QueryManager_Update_ResPart, err error) {
	c := q.client
	// Update is a sequence of match, delete & insert => TX_WRITE
	req := requests.GetUpdateQueryReq(query, options)
	streamQuery, queryErr := c.runStreamTx(sessionID, req, TX_WRITE, options, true)
	if queryErr != nil {
		return nil, queryErr
	}

	// extract match results and stuff into queryResults collection
	for _, item := range streamQuery {
		queryResults = append(queryResults, item.GetUpdateResPart())
	}

	return queryResults, nil
}

func (q *QueryManager) Explain(sessionID []byte, explainableID int64, options *common.Options) (queryResults []*common.QueryManager_Explain_ResPart, err error) {
	c := q.client
	req := requests.GetExplainQueryReq(explainableID, options)
	streamQuery, queryErr := c.runStreamTx(sessionID, req, TX_READ, options, true)
	if queryErr != nil {
		return nil, queryErr
	}

	// extract match results and stuff into queryResults collection
	for _, item := range streamQuery {
		queryResults = append(queryResults, item.GetExplainResPart())
	}

	return queryResults, nil
}

func (q *QueryManager) Match(sessionID []byte, query string, options *common.Options) (queryResults []*common.QueryManager_Match_ResPart, err error) {
	c := q.client
	req := requests.GetMatchQueryReq(query, options)
	streamQuery, queryErr := c.runStreamTx(sessionID, req, TX_READ, options, false)
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

func (q *QueryManager) MatchGroup(sessionID []byte, query string, options *common.Options) (queryResults []*common.QueryManager_MatchGroup_ResPart, err error) {
	c := q.client
	req := requests.GetMatchGroupQueryReq(query, options)
	streamQuery, queryErr := c.runStreamTx(sessionID, req, TX_READ, options, false)
	if queryErr != nil {
		return nil, queryErr
	}

	// extract match results and stuff into queryResults collection
	for _, item := range streamQuery {
		queryResults = append(queryResults, item.GetMatchGroupResPart())
	}
	return queryResults, nil
}

func (q *QueryManager) MatchGroupAggregate(sessionID []byte, query string, options *common.Options) (queryResults []*common.QueryManager_MatchGroupAggregate_ResPart, err error) {
	c := q.client
	req := requests.GetMatchGroupQueryAggregateQueryReq(query, options)
	streamQuery, queryErr := c.runStreamTx(sessionID, req, TX_READ, options, false)
	if queryErr != nil {
		return nil, queryErr
	}

	// extract match results and stuff into queryResults collection
	for _, item := range streamQuery {
		queryResults = append(queryResults, item.GetMatchGroupAggregateResPart())
	}

	return queryResults, nil
}

func (q *QueryManager) Define(sessionID []byte, query string, options *common.Options) (queryResponses *common.QueryManager_Define_Res, err error) {
	c := q.client
	req := requests.GetDefinedQueryReq(query, options)
	res, queryErr := c.runQuery(sessionID, req, TX_READ, options)
	if queryErr != nil {
		return nil, queryErr
	} else {
		return res.GetDefineRes(), nil
	}
}

func (q *QueryManager) Undefine(sessionID []byte, query string, options *common.Options) (queryResponses *common.QueryManager_Undefine_Res, err error) {
	c := q.client
	req := requests.GetUndefinedQueryReq(query, options)
	res, queryErr := c.runQuery(sessionID, req, TX_READ, options)
	if queryErr != nil {
		return nil, queryErr
	} else {
		return res.GetUndefineRes(), nil
	}
}

func (q *QueryManager) MatchAggregate(sessionID []byte, query string, options *common.Options) (queryResponses *common.QueryManager_MatchAggregate_Res, err error) {
	c := q.client
	req := requests.GetMatchAggregateQueryReq(query, options)
	res, queryErr := c.runQuery(sessionID, req, TX_READ, options)
	if queryErr != nil {
		return nil, queryErr
	} else {
		return res.GetMatchAggregateRes(), nil
	}
}

func (q *QueryManager) Delete(sessionID []byte, query string, options *common.Options) (queryResponses *common.QueryManager_Delete_Res, err error) {
	c := q.client
	req := requests.GetDeleteQueryReq(query, options)
	res, queryErr := c.runQuery(sessionID, req, TX_WRITE, options)
	if queryErr != nil {
		return nil, queryErr
	} else {
		return res.GetDeleteRes(), nil
	}
}
