package v2

import (
	common "github.com/marvin-hansen/typedb-client-go/common"
	. "github.com/marvin-hansen/typedb-client-go/src/common"
)

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
