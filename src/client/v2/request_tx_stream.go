package v2

import "github.com/marvin-hansen/typedb-client-go/common"

func getTransactionStreamReq() (req *common.Transaction_Req) {
	r := &common.Transaction_Stream_Req{}
	return &common.Transaction_Req{Req: &common.Transaction_Req_StreamReq{StreamReq: r}}
}
