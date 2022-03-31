// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package requests

import (
	"github.com/marvin-hansen/typedb-client-go/common"
)

// QueryManager
// https://github.com/vaticle/typedb-client-python/blob/master/typedb/common/rpc/request_builder.py

// getQueryTx coverts a QueryManager_Req into a Transaction_Req
func getQueryTx(req *common.QueryManager_Req, requestId []byte, metadata map[string]string) *common.Transaction_Req {
	r := &common.Transaction_Req_QueryManagerReq{QueryManagerReq: req}
	return &common.Transaction_Req{Req: r, ReqId: requestId, Metadata: metadata}
}

func GetInsertQueryReq(query string, options *common.Options, requestId []byte, metadata map[string]string) *common.Transaction_Req {
	r := &common.QueryManager_Req_InsertReq{InsertReq: &common.QueryManager_Insert_Req{Query: query}}
	req := &common.QueryManager_Req{Options: options, Req: r}
	return getQueryTx(req, requestId, metadata)
}

func GetDefinedQueryReq(query string, requestId []byte, options *common.Options, metadata map[string]string) *common.Transaction_Req {
	r := &common.QueryManager_Req_DefineReq{DefineReq: &common.QueryManager_Define_Req{Query: query}}
	req := &common.QueryManager_Req{Options: options, Req: r}
	return getQueryTx(req, requestId, metadata)
}

func GetUndefinedQueryReq(query string, options *common.Options, requestId []byte, metadata map[string]string) *common.Transaction_Req {
	r := &common.QueryManager_Req_UndefineReq{UndefineReq: &common.QueryManager_Undefine_Req{Query: query}}
	req := &common.QueryManager_Req{Options: options, Req: r}
	return getQueryTx(req, requestId, metadata)
}

func GetUpdateQueryReq(query string, options *common.Options, requestId []byte, metadata map[string]string) *common.Transaction_Req {
	r := &common.QueryManager_Req_UpdateReq{UpdateReq: &common.QueryManager_Update_Req{Query: query}}
	req := &common.QueryManager_Req{Options: options, Req: r}
	return getQueryTx(req, requestId, metadata)
}

func GetExplainQueryReq(explainableID int64, options *common.Options, requestId []byte, metadata map[string]string) *common.Transaction_Req {
	r := &common.QueryManager_Req_ExplainReq{ExplainReq: &common.QueryManager_Explain_Req{ExplainableId: explainableID}}
	req := &common.QueryManager_Req{Options: options, Req: r}
	return getQueryTx(req, requestId, metadata)
}

func GetDeleteQueryReq(query string, options *common.Options, requestId []byte, metadata map[string]string) *common.Transaction_Req {
	r := &common.QueryManager_Req_DeleteReq{DeleteReq: &common.QueryManager_Delete_Req{Query: query}}
	req := &common.QueryManager_Req{Options: options, Req: r}
	return getQueryTx(req, requestId, metadata)
}

func GetMatchQueryReq(query string, options *common.Options, requestId []byte, metadata map[string]string) *common.Transaction_Req {
	r := &common.QueryManager_Req_MatchReq{MatchReq: &common.QueryManager_Match_Req{Query: query}}
	req := &common.QueryManager_Req{Options: options, Req: r}
	return getQueryTx(req, requestId, metadata)
}

func GetMatchAggregateQueryReq(query string, options *common.Options, requestId []byte, metadata map[string]string) *common.Transaction_Req {
	r := &common.QueryManager_Req_MatchAggregateReq{MatchAggregateReq: &common.QueryManager_MatchAggregate_Req{Query: query}}
	req := &common.QueryManager_Req{Options: options, Req: r}
	return getQueryTx(req, requestId, metadata)
}

func GetMatchGroupQueryReq(query string, options *common.Options, requestId []byte, metadata map[string]string) *common.Transaction_Req {
	r := &common.QueryManager_Req_MatchGroupReq{MatchGroupReq: &common.QueryManager_MatchGroup_Req{Query: query}}
	req := &common.QueryManager_Req{Options: options, Req: r}
	return getQueryTx(req, requestId, metadata)
}

func GetMatchGroupQueryAggregateQueryReq(query string, options *common.Options, requestId []byte, metadata map[string]string) *common.Transaction_Req {
	r := &common.QueryManager_Req_MatchGroupAggregateReq{MatchGroupAggregateReq: &common.QueryManager_MatchGroupAggregate_Req{Query: query}}
	req := &common.QueryManager_Req{Options: options, Req: r}
	return getQueryTx(req, requestId, metadata)
}
