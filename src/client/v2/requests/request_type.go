// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package requests

import (
	"github.com/marvin-hansen/typedb-client-go/common"
	. "github.com/marvin-hansen/typedb-client-go/src/type/common_type"
)

// Type
// https://github.com/vaticle/typedb-client-python/blob/master/typedb/common/rpc/request_builder.py

// GetTypeTx coverts a TypeReq into a Transaction_Req
func GetTypeTx(req *common.Type_Req, label *Label) *common.Transaction_Req {
	req.Label = label.GetName()
	if label.HasScope() {
		req.Scope = label.GetScope()
	}

	r := &common.Transaction_Req_TypeReq{TypeReq: req}
	return &common.Transaction_Req{Req: r}
}

func GetTypeIsAbstractReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_TypeIsAbstractReq{}}
	return GetTypeTx(req, label)
}

func GetTypeSetLabelReq(label *Label, newLabel string) *common.Transaction_Req {
	req := &common.Type_Req_TypeSetLabelReq{}
	req.TypeSetLabelReq.Label = newLabel
	return GetTypeTx(&common.Type_Req{Req: req}, label)
}

func GetTypeSupertypeReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_TypeGetSupertypeReq{}}
	return GetTypeTx(req, label)
}

func GetTypeAllSupertypesReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_TypeGetSupertypesReq{}}
	return GetTypeTx(req, label)
}

func GetTypeAllSubtypesReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_TypeGetSubtypesReq{}}
	return GetTypeTx(req, label)
}

func GetTypeDeleteReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_TypeDeleteReq{}}
	return GetTypeTx(req, label)
}
