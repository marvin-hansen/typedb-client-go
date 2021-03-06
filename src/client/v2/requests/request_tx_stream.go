package requests

import "github.com/marvin-hansen/typedb-client-go/common"

func GetTransactionStreamReq() (req *common.Transaction_Req) {
	reqID := createNewRequestID()
	r := &common.Transaction_Stream_Req{}
	return &common.Transaction_Req{Req: &common.Transaction_Req_StreamReq{StreamReq: r}, ReqId: reqID}
}
