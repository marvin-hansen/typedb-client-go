package v2

import "github.com/marvin-hansen/go-typedb/proto/common"

// Type
// https://github.com/vaticle/typedb-client-python/blob/master/typedb/common/rpc/request_builder.py

// getTypeTx coverts a LogicManager_Req into a Transaction_Req
func getTypeTx(req *common.Type_Req) *common.Transaction_Req {
	r := &common.Transaction_Req_TypeReq{TypeReq: req}
	return &common.Transaction_Req{Req: r}
}

func getTypeIsAbstractReq(label string) *common.Type_Req {
	req := &common.Type_Req_TypeIsAbstractReq{}
	return &common.Type_Req{Req: req, Label: label}
}

func getTypeSetLabelReq(label, newLabel string) *common.Type_Req {
	req := &common.Type_Req_TypeSetLabelReq{}
	req.TypeSetLabelReq.Label = newLabel
	return &common.Type_Req{Req: req, Label: label}
}

func getTypeSupertypeReq(label string) *common.Type_Req {
	req := &common.Type_Req_TypeGetSupertypeReq{}
	return &common.Type_Req{Req: req, Label: label}
}

func getTypeAllSupertypesReq(label string) *common.Type_Req {
	req := &common.Type_Req_TypeGetSupertypesReq{}
	return &common.Type_Req{Req: req, Label: label}
}

func getTypeAllSubtypesReq(label string) *common.Type_Req {
	req := &common.Type_Req_TypeGetSubtypesReq{}
	return &common.Type_Req{Req: req, Label: label}
}

func getTypeDeleteReq(label string) *common.Type_Req {
	req := &common.Type_Req_TypeDeleteReq{}
	return &common.Type_Req{Req: req, Label: label}
}

// RoleType
// 408
