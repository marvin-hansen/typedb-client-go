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

func getTransactionOpenReq(sessionID, transactionId []byte, sessionType common.Transaction_Type, options *common.Options, netMillisecondLatency int32) (req *common.Transaction_Req) {

	if options == nil {
		options = &common.Options{}
	}

	r := &common.Transaction_Open_Req{
		SessionId:            sessionID,
		Type:                 sessionType,
		Options:              options,
		NetworkLatencyMillis: netMillisecondLatency,
	}

	req = &common.Transaction_Req{
		ReqId: transactionId,
		Req:   &common.Transaction_Req_OpenReq{OpenReq: r},
	}
	return req
}

func getTransactionCommitReq(transactionId []byte, metadata map[string]string) (req *common.Transaction_Req) {
	req = &common.Transaction_Req{
		ReqId:    transactionId,
		Metadata: metadata,
		Req:      &common.Transaction_Req_CommitReq{CommitReq: &common.Transaction_Commit_Req{}},
	}
	return req
}

func getTransactionRollbackReq(transactionId []byte, metadata map[string]string) (req *common.Transaction_Req) {
	req = &common.Transaction_Req{
		ReqId:    transactionId,
		Metadata: metadata,
		Req:      &common.Transaction_Req_RollbackReq{RollbackReq: &common.Transaction_Rollback_Req{}},
	}
	return req
}
