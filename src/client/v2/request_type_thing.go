package v2

import (
	"github.com/marvin-hansen/go-typedb/proto/common"
	. "github.com/marvin-hansen/go-typedb/src/common"
)

// ThingType
// 432

func getThingType(label *Label, encoding common.Type_Encoding) *common.Type {
	return &common.Type{
		Label:    label.GetName(),
		Scope:    label.GetScope(),
		Encoding: encoding,
	}
}

func getThingSetAbstractReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_ThingTypeSetAbstractReq{}}
	return getTypeTx(req, label)
}

func getThingUnsetAbstractReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_ThingTypeUnsetAbstractReq{}}
	return getTypeTx(req, label)
}

func getThingSetSupertypeReq(label *Label, superType common.Type) *common.Transaction_Req {
	ssr := &common.Type_SetSupertype_Req{Type: &superType}
	req := &common.Type_Req{Req: &common.Type_Req_TypeSetSupertypeReq{TypeSetSupertypeReq: ssr}}
	return getTypeTx(req, label)
}

func getThingGetPlaysReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_ThingTypeGetPlaysReq{}}
	return getTypeTx(req, label)
}

// FIXME: ThingType_SetPlays_Req Doesn't have a overriddenRolType field
func getThingSetPlaysReq(label *Label, roleType *common.Type, overriddenRolType *common.Type) *common.Transaction_Req {
	spr := &common.ThingType_SetPlays_Req{}
	spr.Role = roleType
	req := &common.Type_Req{Req: &common.Type_Req_ThingTypeSetPlaysReq{ThingTypeSetPlaysReq: spr}}
	return getTypeTx(req, label)
}

func getThingUnsetPlaysReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_ThingTypeUnsetPlaysReq{}}
	return getTypeTx(req, label)
}

// FIXME: ThingType_GetOwns_Req Doesn't have a valueType field
func getThingGetOwnsReq(label *Label, valueType *common.AttributeType_ValueType, keysOnly bool) *common.Transaction_Req {
	// Corresponds to #483
	gor := &common.ThingType_GetOwns_Req{KeysOnly: keysOnly}
	req := &common.Type_Req{Req: &common.Type_Req_ThingTypeGetOwnsReq{
		ThingTypeGetOwnsReq: gor,
	}}

	return getTypeTx(req, label)
}

func getThingSetOwnsReq(label *Label, attributeType *common.Type, isKey bool) *common.Transaction_Req {
	sor := &common.ThingType_SetOwns_Req{
		AttributeType: attributeType,
		Overridden:    nil,
		IsKey:         isKey,
	}
	req := &common.Type_Req{Req: &common.Type_Req_ThingTypeSetOwnsReq{ThingTypeSetOwnsReq: sor}}
	return getTypeTx(req, label)
}

func getThingUnsetOwnsReq(label *Label, attributeType *common.Type) *common.Transaction_Req {
	uor := &common.ThingType_UnsetOwns_Req{AttributeType: attributeType}
	req := &common.Type_Req{Req: &common.Type_Req_ThingTypeUnsetOwnsReq{ThingTypeUnsetOwnsReq: uor}}
	return getTypeTx(req, label)
}

func getThingGetInstanceReq(label *Label) *common.Transaction_Req {
	r := &common.ThingType_GetInstances_Req{}
	req := &common.Type_Req{Req: &common.Type_Req_ThingTypeGetInstancesReq{ThingTypeGetInstancesReq: r}}
	return getTypeTx(req, label)
}
