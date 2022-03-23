package v2

import (
	common2 "github.com/marvin-hansen/go-typedb/common"
	. "github.com/marvin-hansen/go-typedb/src/common"
)

// EntityType
// 520

func getCreateEntityReq(label *Label) *common2.Transaction_Req {
	s := &common2.EntityType_Create_Req{}
	req := &common2.Type_Req{Req: &common2.Type_Req_EntityTypeCreateReq{EntityTypeCreateReq: s}}
	return getTypeTx(req, label)
}

// # RelationType
// 529

func getRelationTypeCreateReq(label *Label) *common2.Transaction_Req {
	s := &common2.RelationType_Create_Req{}
	req := &common2.Type_Req{Req: &common2.Type_Req_RelationTypeCreateReq{RelationTypeCreateReq: s}}
	return getTypeTx(req, label)
}

func getRelationTypeGetRelatesReq(label *Label, roleLabel *string) *common2.Transaction_Req {
	if roleLabel != nil {
		s := &common2.RelationType_GetRelatesForRoleLabel_Req{Label: label.GetName()}
		req := &common2.Type_Req{Req: &common2.Type_Req_RelationTypeGetRelatesForRoleLabelReq{
			RelationTypeGetRelatesForRoleLabelReq: s,
		}}
		return getTypeTx(req, label)
	} else {
		s := &common2.RelationType_GetRelates_Req{}
		req := &common2.Type_Req{Req: &common2.Type_Req_RelationTypeGetRelatesReq{RelationTypeGetRelatesReq: s}}
		return getTypeTx(req, label)
	}
}

// FIXME: How to cast string to  isRelationType_SetRelates_Req_Overridden`?
func getRelationTypeSetRelatesReq(label *Label, roleLabel *string, overriddenLabel *string) *common2.Transaction_Req {
	s := &common2.RelationType_SetRelates_Req{Label: *roleLabel}
	if overriddenLabel != nil {
		s.Overridden = nil
	}
	req := &common2.Type_Req{Req: &common2.Type_Req_RelationTypeSetRelatesReq{RelationTypeSetRelatesReq: s}}
	return getTypeTx(req, label)

}

func getRelationTypeUnsetRelatesReq(label *Label, roleLabel *string) *common2.Transaction_Req {
	s := &common2.RelationType_UnsetRelates_Req{Label: label.GetName()}
	req := &common2.Type_Req{Req: &common2.Type_Req_RelationTypeUnsetRelatesReq{RelationTypeUnsetRelatesReq: s}}
	return getTypeTx(req, label)
}
