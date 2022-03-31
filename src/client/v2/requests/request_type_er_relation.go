// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package requests

import (
	"github.com/marvin-hansen/typedb-client-go/common"
	. "github.com/marvin-hansen/typedb-client-go/src/type/common_type"
)

// EntityType
// 520

func GetCreateEntityReq(label *Label) *common.Transaction_Req {
	s := &common.EntityType_Create_Req{}
	req := &common.Type_Req{Req: &common.Type_Req_EntityTypeCreateReq{EntityTypeCreateReq: s}}
	return GetTypeTx(req, label)
}

// # RelationType
// 529

func GetRelationTypeCreateReq(label *Label) *common.Transaction_Req {
	s := &common.RelationType_Create_Req{}
	req := &common.Type_Req{Req: &common.Type_Req_RelationTypeCreateReq{RelationTypeCreateReq: s}}
	return GetTypeTx(req, label)
}

func GetRelationTypeGetRelatesReq(label *Label, roleLabel *string) *common.Transaction_Req {
	if roleLabel != nil {
		s := &common.RelationType_GetRelatesForRoleLabel_Req{Label: label.GetName()}
		req := &common.Type_Req{Req: &common.Type_Req_RelationTypeGetRelatesForRoleLabelReq{
			RelationTypeGetRelatesForRoleLabelReq: s,
		}}
		return GetTypeTx(req, label)
	} else {
		s := &common.RelationType_GetRelates_Req{}
		req := &common.Type_Req{Req: &common.Type_Req_RelationTypeGetRelatesReq{RelationTypeGetRelatesReq: s}}
		return GetTypeTx(req, label)
	}
}

func GetRelationTypeSetRelatesReq(label *Label, roleLabel *string, overriddenLabel *string) *common.Transaction_Req {
	s := &common.RelationType_SetRelates_Req{}
	if overriddenLabel != nil {
		s.Overridden = &common.RelationType_SetRelates_Req_OverriddenLabel{OverriddenLabel: *overriddenLabel}
	}
	req := &common.Type_Req{Req: &common.Type_Req_RelationTypeSetRelatesReq{RelationTypeSetRelatesReq: s}}
	return GetTypeTx(req, label)
}

func GetRelationTypeUnsetRelatesReq(label *Label, roleLabel *string) *common.Transaction_Req {
	s := &common.RelationType_UnsetRelates_Req{Label: *roleLabel}
	req := &common.Type_Req{Req: &common.Type_Req_RelationTypeUnsetRelatesReq{RelationTypeUnsetRelatesReq: s}}
	return GetTypeTx(req, label)
}
