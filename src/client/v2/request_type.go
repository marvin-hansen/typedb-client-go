package v2

import (
	"github.com/marvin-hansen/go-typedb/proto/common"
	. "github.com/marvin-hansen/go-typedb/src/common"
)

// Type
// https://github.com/vaticle/typedb-client-python/blob/master/typedb/common/rpc/request_builder.py

// getTypeTx coverts a LogicManager_Req into a Transaction_Req
func getTypeTx(req *common.Type_Req, label *Label) *common.Transaction_Req {
	req.Label = label.GetName()
	if label.HasScope() {
		req.Scope = label.GetScope()
	}

	r := &common.Transaction_Req_TypeReq{TypeReq: req}
	return &common.Transaction_Req{Req: r}
}

func getTypeIsAbstractReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_TypeIsAbstractReq{}}
	return getTypeTx(req, label)
}

func getTypeSetLabelReq(label *Label, newLabel string) *common.Transaction_Req {
	req := &common.Type_Req_TypeSetLabelReq{}
	req.TypeSetLabelReq.Label = newLabel
	return getTypeTx(&common.Type_Req{Req: req}, label)
}

func getTypeSupertypeReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_TypeGetSupertypeReq{}}
	return getTypeTx(req, label)
}

func getTypeAllSupertypesReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_TypeGetSupertypesReq{}}
	return getTypeTx(req, label)
}

func getTypeAllSubtypesReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_TypeGetSubtypesReq{}}
	return getTypeTx(req, label)
}

func getTypeDeleteReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_TypeDeleteReq{}}
	return getTypeTx(req, label)
}

// RoleType
// 410

func getRoleType(label *Label, encoding common.Type_Encoding) *common.Type {
	return &common.Type{
		Label:    label.GetName(),
		Scope:    label.GetScope(),
		Encoding: encoding,
	}
}

func getRoleTypeRelationTypeReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_RoleTypeGetRelationTypesReq{}}
	return getTypeTx(req, label)
}

func getRoleTypePlayersReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_RoleTypeGetPlayersReq{}}
	return getTypeTx(req, label)
}
