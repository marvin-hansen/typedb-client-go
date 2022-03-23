package v2

import (
	common2 "github.com/marvin-hansen/go-typedb/common"
	"strings"
)

// Thing
// 605

func convertStringToByte(str string) []byte {
	//trimming the spaces, this method will remote the '\t', '\n', '\v', '\f', '\r', ' ' chars
	return []byte(strings.TrimSpace(str))
}

func convertThingToProto(iid string) *common2.Thing {
	return &common2.Thing{Iid: convertStringToByte(iid)}
}

// getTypeTx coverts a ThingReq into a Transaction_Req
func getThingTx(req *common2.Thing_Req, iid string) *common2.Transaction_Req {
	req.Iid = convertStringToByte(iid)
	r := &common2.Transaction_Req_ThingReq{ThingReq: req}
	return &common2.Transaction_Req{Req: r}
}
