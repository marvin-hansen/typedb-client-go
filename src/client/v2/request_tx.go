package v2

import (
	"github.com/marvin-hansen/go-typedb/common"
)

// Transaction
// https://github.com/vaticle/typedb-client-python/blob/master/typedb/common/rpc/request_builder.py

func getTransactionClient(reqs []*common.Transaction_Req) (req *common.Transaction_Client) {
	return &common.Transaction_Client{Reqs: reqs}
}

func getTransactionOpenReq(sessionID []byte, sessionType common.Transaction_Type, options *common.Options, netMillisecondLatency int32) (req *common.Transaction_Open_Req) {
	return &common.Transaction_Open_Req{
		SessionId:            sessionID,
		Type:                 sessionType,
		Options:              options,
		NetworkLatencyMillis: netMillisecondLatency,
	}
}

func getTransactionCommitReq() (req *common.Transaction_Commit_Req) {
	return &common.Transaction_Commit_Req{}
}

func getTransactionRollbackReq() (req *common.Transaction_Req_RollbackReq) {
	return &common.Transaction_Req_RollbackReq{}
}
