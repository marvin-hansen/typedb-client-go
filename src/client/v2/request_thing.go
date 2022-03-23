package v2

import (
	"github.com/marvin-hansen/go-typedb/proto/common"
	"strings"
)

// Thing
// 605

func convertStringToByte(str string) []byte {
	//trimming the spaces, this method will remote the '\t', '\n', '\v', '\f', '\r', ' ' chars
	return []byte(strings.TrimSpace(str))
}

func convertThingToProto(iid string) *common.Thing {
	return &common.Thing{Iid: convertStringToByte(iid)}
}

// getTypeTx coverts a ThingReq into a Transaction_Req
func getThingTx(req *common.Thing_Req, iid string) *common.Transaction_Req {
	req.Iid = convertStringToByte(iid)
	r := &common.Transaction_Req_ThingReq{ThingReq: req}
	return &common.Transaction_Req{Req: r}
}
