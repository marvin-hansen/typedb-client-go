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

func getThingSetSupertypeReq(label *Label) *common.Transaction_Req {
	// FIXME
	return nil
}

func getThingGetPlaysReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_ThingTypeGetPlaysReq{}}
	return getTypeTx(req, label)
}

func getThingSetPlaysReq(label *Label, roleType *common.Type, overriddenRolType *common.Type) *common.Transaction_Req {
	spr := &common.Type_Req_ThingTypeSetPlaysReq{}
	spr.ThingTypeSetPlaysReq.Role = roleType

	// FIXME
	if overriddenRolType != nil {
		spr.ThingTypeSetPlaysReq.Overridden = nil
	}
	req := &common.Type_Req{Req: spr}

	return getTypeTx(req, label)
}

func getThingUnsetPlaysReq(label *Label) *common.Transaction_Req {
	req := &common.Type_Req{Req: &common.Type_Req_ThingTypeUnsetPlaysReq{}}
	return getTypeTx(req, label)
}
