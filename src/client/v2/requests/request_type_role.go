// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package requests

import (
	"github.com/marvin-hansen/typedb-client-go/common"
	. "github.com/marvin-hansen/typedb-client-go/src/type/common_type"
)

// RoleType
// 410

func GetRoleType(label *Label, encoding common.Type_Encoding) *common.Type {
	return &common.Type{
		Label:    label.GetName(),
		Scope:    label.GetScope(),
		Encoding: encoding,
	}
}

func GetRoleTypeRelationTypeReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_RoleTypeGetRelationTypesReq{}}
	return GetTypeTx(req, label)
}

func GetRoleTypePlayersReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_RoleTypeGetPlayersReq{}}
	return GetTypeTx(req, label)
}
