// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package requests

import (
	"github.com/marvin-hansen/typedb-client-go/common"
	. "github.com/marvin-hansen/typedb-client-go/src/type/common_type"
)

// ThingType
// 432

func GetThingType(label *Label, encoding common.Type_Encoding) *common.Type {
	return &common.Type{
		Label:    label.GetName(),
		Scope:    label.GetScope(),
		Encoding: encoding,
	}
}

func GetThingSetAbstractReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_ThingTypeSetAbstractReq{}}
	return GetTypeTx(req, label)
}

func GetThingUnsetAbstractReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_ThingTypeUnsetAbstractReq{}}
	return GetTypeTx(req, label)
}

func GetThingSetSupertypeReq(label *Label, superType *common.Type) *common.Transaction_Req {
	s := &common.Type_SetSupertype_Req{Type: superType}
	req := &common.Type_Req{Req: &common.Type_Req_TypeSetSupertypeReq{TypeSetSupertypeReq: s}}
	return GetTypeTx(req, label)
}

func GetThingGetPlaysReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_ThingTypeGetPlaysReq{}}
	return GetTypeTx(req, label)
}

func GetThingSetPlaysReq(label *Label, roleType *common.Type, overriddenRolType *common.Type) *common.Transaction_Req {
	s := &common.ThingType_SetPlays_Req{Role: roleType}
	if overriddenRolType != nil {
		s.Overridden = &common.ThingType_SetPlays_Req_OverriddenRole{OverriddenRole: overriddenRolType}
	}
	req := &common.Type_Req{Req: &common.Type_Req_ThingTypeSetPlaysReq{ThingTypeSetPlaysReq: s}}
	return GetTypeTx(req, label)
}

func GetThingUnsetPlaysReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_ThingTypeUnsetPlaysReq{}}
	return GetTypeTx(req, label)
}

func GetThingGetOwnsReq(label *Label, valueType *common.AttributeType_ValueType, keysOnly bool) *common.Transaction_Req {
	// Corresponds to #483

	r := &common.ThingType_GetOwns_Req{KeysOnly: keysOnly}

	if valueType != nil {
		r.Filter = &common.ThingType_GetOwns_Req_ValueType{ValueType: *valueType}
	}
	req := &common.Type_Req{Req: &common.Type_Req_ThingTypeGetOwnsReq{ThingTypeGetOwnsReq: r}}
	return GetTypeTx(req, label)
}

func GetThingSetOwnsReq(label *Label, attributeType *common.Type, isKey bool) *common.Transaction_Req {
	s := &common.ThingType_SetOwns_Req{
		AttributeType: attributeType,
		Overridden:    nil,
		IsKey:         isKey,
	}
	req := &common.Type_Req{Req: &common.Type_Req_ThingTypeSetOwnsReq{ThingTypeSetOwnsReq: s}}
	return GetTypeTx(req, label)
}

func GetThingUnsetOwnsReq(label *Label, attributeType *common.Type) *common.Transaction_Req {
	s := &common.ThingType_UnsetOwns_Req{AttributeType: attributeType}
	req := &common.Type_Req{Req: &common.Type_Req_ThingTypeUnsetOwnsReq{ThingTypeUnsetOwnsReq: s}}
	return GetTypeTx(req, label)
}

func GetThingGetInstanceReq(label *Label) *common.Transaction_Req {
	s := &common.ThingType_GetInstances_Req{}
	req := &common.Type_Req{Req: &common.Type_Req_ThingTypeGetInstancesReq{ThingTypeGetInstancesReq: s}}
	return GetTypeTx(req, label)
}
