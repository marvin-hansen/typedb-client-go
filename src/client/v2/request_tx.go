package v2

import (
	common2 "github.com/marvin-hansen/go-typedb/common"
)

// Transaction
// https://github.com/vaticle/typedb-client-python/blob/master/typedb/common/rpc/request_builder.py

func getTransactionClient(reqs []*common2.Transaction_Req) (req *common2.Transaction_Client) {
	return &common2.Transaction_Client{Reqs: reqs}
}

func getTransactionOpenReq(sessionID []byte, sessionType common2.Transaction_Type, options *common2.Options, netMillisecondLatency int32) (req *common2.Transaction_Open_Req) {
	return &common2.Transaction_Open_Req{
		SessionId:            sessionID,
		Type:                 sessionType,
		Options:              options,
		NetworkLatencyMillis: netMillisecondLatency,
	}
}

func getTransactionCommitReq() (req *common2.Transaction_Commit_Req) {
	return &common2.Transaction_Commit_Req{}
}

func getTransactionRollbackReq() (req *common2.Transaction_Req_RollbackReq) {
	return &common2.Transaction_Req_RollbackReq{}
}
