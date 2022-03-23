package v2

import (
	common2 "github.com/marvin-hansen/go-typedb/common"
	. "github.com/marvin-hansen/go-typedb/src/common"
)

// ThingType
// 432

func getThingType(label *Label, encoding common2.Type_Encoding) *common2.Type {
	return &common2.Type{
		Label:    label.GetName(),
		Scope:    label.GetScope(),
		Encoding: encoding,
	}
}

func getThingSetAbstractReq(label *Label) *common2.Transaction_Req {
	req := &common2.Type_Req{Req: &common2.Type_Req_ThingTypeSetAbstractReq{}}
	return getTypeTx(req, label)
}

func getThingUnsetAbstractReq(label *Label) *common2.Transaction_Req {
	req := &common2.Type_Req{Req: &common2.Type_Req_ThingTypeUnsetAbstractReq{}}
	return getTypeTx(req, label)
}

func getThingSetSupertypeReq(label *Label, superType *common2.Type) *common2.Transaction_Req {
	s := &common2.Type_SetSupertype_Req{Type: superType}
	req := &common2.Type_Req{Req: &common2.Type_Req_TypeSetSupertypeReq{TypeSetSupertypeReq: s}}
	return getTypeTx(req, label)
}

func getThingGetPlaysReq(label *Label) *common2.Transaction_Req {
	req := &common2.Type_Req{Req: &common2.Type_Req_ThingTypeGetPlaysReq{}}
	return getTypeTx(req, label)
}

// FIXME: ThingType_SetPlays_Req Doesn't have a overriddenRolType field
func getThingSetPlaysReq(label *Label, roleType *common2.Type, overriddenRolType *common2.Type) *common2.Transaction_Req {
	s := &common2.ThingType_SetPlays_Req{}
	s.Role = roleType
	req := &common2.Type_Req{Req: &common2.Type_Req_ThingTypeSetPlaysReq{ThingTypeSetPlaysReq: s}}
	return getTypeTx(req, label)
}

func getThingUnsetPlaysReq(label *Label) *common2.Transaction_Req {
	req := &common2.Type_Req{Req: &common2.Type_Req_ThingTypeUnsetPlaysReq{}}
	return getTypeTx(req, label)
}

// FIXME: ThingType_GetOwns_Req Doesn't have a valueType field
func getThingGetOwnsReq(label *Label, valueType *common2.AttributeType_ValueType, keysOnly bool) *common2.Transaction_Req {
	// Corresponds to #483
	gor := &common2.ThingType_GetOwns_Req{KeysOnly: keysOnly}
	req := &common2.Type_Req{Req: &common2.Type_Req_ThingTypeGetOwnsReq{
		ThingTypeGetOwnsReq: gor,
	}}

	return getTypeTx(req, label)
}

func getThingSetOwnsReq(label *Label, attributeType *common2.Type, isKey bool) *common2.Transaction_Req {
	s := &common2.ThingType_SetOwns_Req{
		AttributeType: attributeType,
		Overridden:    nil,
		IsKey:         isKey,
	}
	req := &common2.Type_Req{Req: &common2.Type_Req_ThingTypeSetOwnsReq{ThingTypeSetOwnsReq: s}}
	return getTypeTx(req, label)
}

func getThingUnsetOwnsReq(label *Label, attributeType *common2.Type) *common2.Transaction_Req {
	s := &common2.ThingType_UnsetOwns_Req{AttributeType: attributeType}
	req := &common2.Type_Req{Req: &common2.Type_Req_ThingTypeUnsetOwnsReq{ThingTypeUnsetOwnsReq: s}}
	return getTypeTx(req, label)
}

func getThingGetInstanceReq(label *Label) *common2.Transaction_Req {
	s := &common2.ThingType_GetInstances_Req{}
	req := &common2.Type_Req{Req: &common2.Type_Req_ThingTypeGetInstancesReq{ThingTypeGetInstancesReq: s}}
	return getTypeTx(req, label)
}
