package v2

import "github.com/marvin-hansen/go-typedb/proto/common"

// QueryManager
// https://github.com/vaticle/typedb-client-python/blob/master/typedb/common/rpc/request_builder.py

func getDefinedQueryReq(query string, options *common.Options) *common.QueryManager_Req {
	req := &common.QueryManager_Req_DefineReq{DefineReq: &common.QueryManager_Define_Req{Query: query}}
	return &common.QueryManager_Req{
		Options: options,
		Req:     req,
	}
}

func getUndefinedQueryReq(query string, options *common.Options) *common.QueryManager_Req {
	req := &common.QueryManager_Req_UndefineReq{}
	req.UndefineReq.Query = query
	return &common.QueryManager_Req{
		Options: options,
		Req:     req,
	}
}

func getMatchQueryReq(query string, options *common.Options) *common.QueryManager_Req {
	req := &common.QueryManager_Req_MatchReq{}
	req.MatchReq.Query = query
	return &common.QueryManager_Req{
		Options: options,
		Req:     req,
	}
}

func getMatchAggregateQueryReq(query string, options *common.Options) *common.QueryManager_Req {
	req := &common.QueryManager_Req_MatchAggregateReq{}
	req.MatchAggregateReq.Query = query
	return &common.QueryManager_Req{
		Options: options,
		Req:     req,
	}
}

func getMatchGroupQueryReq(query string, options *common.Options) *common.QueryManager_Req {
	req := &common.QueryManager_Req_MatchGroupReq{}
	req.MatchGroupReq.Query = query
	return &common.QueryManager_Req{
		Options: options,
		Req:     req,
	}
}

func getMatchGroupQueryAggregateReq(query string, options *common.Options) *common.QueryManager_Req {
	req := &common.QueryManager_Req_MatchGroupAggregateReq{}
	req.MatchGroupAggregateReq.Query = query
	return &common.QueryManager_Req{
		Options: options,
		Req:     req,
	}
}

func getInsertQueryReq(query string, options *common.Options) *common.QueryManager_Req {
	req := &common.QueryManager_Req_InsertReq{}
	req.InsertReq.Query = query
	return &common.QueryManager_Req{
		Options: options,
		Req:     req,
	}
}

func getUpdateQueryReq(query string, options *common.Options) *common.QueryManager_Req {
	req := &common.QueryManager_Req_UpdateReq{}
	req.UpdateReq.Query = query
	return &common.QueryManager_Req{
		Options: options,
		Req:     req,
	}
}

func getDeleteReq(query string, options *common.Options) *common.QueryManager_Req {
	req := &common.QueryManager_Req_DeleteReq{}
	req.DeleteReq.Query = query
	return &common.QueryManager_Req{
		Options: options,
		Req:     req,
	}
}

func getExplainReq(explainableID int64, options *common.Options) *common.QueryManager_Req {
	req := &common.QueryManager_Req_ExplainReq{}
	req.ExplainReq.ExplainableId = explainableID
	return &common.QueryManager_Req{
		Options: options,
		Req:     req,
	}
}
