// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com
package v2

import (
	"github.com/marvin-hansen/typedb-client-go/common"
)

// Transaction
// https://github.com/vaticle/typedb-client-python/blob/master/typedb/common/rpc/request_builder.py

func getTransactionClient(reqs []*common.Transaction_Req) (req *common.Transaction_Client) {
	return &common.Transaction_Client{Reqs: reqs}
}

// getQueryTx coverts a QueryManager_Req into a Transaction_Req
func getTxReq(req *common.QueryManager_Req) *common.Transaction_Req {
	r := &common.Transaction_Req_QueryManagerReq{QueryManagerReq: req}
	return &common.Transaction_Req{Req: r}
}

func getTransactionOpenReq(sessionID []byte, sessionType common.Transaction_Type, options *common.Options, netMillisecondLatency int32) (req *common.Transaction_Req) {

	r := &common.Transaction_Open_Req{
		SessionId:            sessionID,
		Type:                 sessionType,
		Options:              options,
		NetworkLatencyMillis: netMillisecondLatency,
	}

	return &common.Transaction_Req{Req: &common.Transaction_Req_OpenReq{OpenReq: r}}
}

func getTransactionCommitReq() (req *common.Transaction_Req) {
	r := &common.Transaction_Commit_Req{}
	return &common.Transaction_Req{Req: &common.Transaction_Req_CommitReq{CommitReq: r}}
}

func getTransactionRollbackReq() (req *common.Transaction_Req) {
	r := &common.Transaction_Rollback_Req{}
	return &common.Transaction_Req{Req: &common.Transaction_Req_RollbackReq{RollbackReq: r}}

}
