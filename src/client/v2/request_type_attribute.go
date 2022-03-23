package v2

import (
	"github.com/marvin-hansen/go-typedb/common"
	. "github.com/marvin-hansen/go-typedb/src/common"
)

// EntityType
// 565

func getAttributeGetOwner(label *Label, onlyKey bool) *common.Transaction_Req {
	s := &common.AttributeType_GetOwners_Req{OnlyKey: onlyKey}
	req := &common.Type_Req{Req: &common.Type_Req_AttributeTypeGetOwnersReq{AttributeTypeGetOwnersReq: s}}
	return getTypeTx(req, label)
}

func getAttributePutReq(label *Label, value *common.Attribute_Value) *common.Transaction_Req {
	s := &common.AttributeType_Put_Req{Value: value}
	req := &common.Type_Req{Req: &common.Type_Req_AttributeTypePutReq{AttributeTypePutReq: s}}
	return getTypeTx(req, label)
}

func getAttributeReq(label *Label, value *common.Attribute_Value) *common.Transaction_Req {
	s := &common.AttributeType_Get_Req{Value: value}
	req := &common.Type_Req{Req: &common.Type_Req_AttributeTypeGetReq{AttributeTypeGetReq: s}}
	return getTypeTx(req, label)
}

func getAttributeGetRegexReq(label *Label) *common.Transaction_Req {
	s := &common.AttributeType_GetRegex_Req{}
	req := &common.Type_Req{Req: &common.Type_Req_AttributeTypeGetRegexReq{AttributeTypeGetRegexReq: s}}
	return getTypeTx(req, label)
}

func getAttributeSetRegexReq(label *Label, regex string) *common.Transaction_Req {
	s := &common.AttributeType_SetRegex_Req{Regex: regex}
	req := &common.Type_Req{Req: &common.Type_Req_AttributeTypeSetRegexReq{AttributeTypeSetRegexReq: s}}
	return getTypeTx(req, label)
}
