package v2

import (
	common2 "github.com/marvin-hansen/go-typedb/common"
	. "github.com/marvin-hansen/go-typedb/src/common"
)

// EntityType
// 565

func getAttributeGetOwner(label *Label, onlyKey bool) *common2.Transaction_Req {
	s := &common2.AttributeType_GetOwners_Req{OnlyKey: onlyKey}
	req := &common2.Type_Req{Req: &common2.Type_Req_AttributeTypeGetOwnersReq{AttributeTypeGetOwnersReq: s}}
	return getTypeTx(req, label)
}

func getAttributePutReq(label *Label, value *common2.Attribute_Value) *common2.Transaction_Req {
	s := &common2.AttributeType_Put_Req{Value: value}
	req := &common2.Type_Req{Req: &common2.Type_Req_AttributeTypePutReq{AttributeTypePutReq: s}}
	return getTypeTx(req, label)
}

func getAttributeReq(label *Label, value *common2.Attribute_Value) *common2.Transaction_Req {
	s := &common2.AttributeType_Get_Req{Value: value}
	req := &common2.Type_Req{Req: &common2.Type_Req_AttributeTypeGetReq{AttributeTypeGetReq: s}}
	return getTypeTx(req, label)
}

func getAttributeGetRegexReq(label *Label) *common2.Transaction_Req {
	s := &common2.AttributeType_GetRegex_Req{}
	req := &common2.Type_Req{Req: &common2.Type_Req_AttributeTypeGetRegexReq{AttributeTypeGetRegexReq: s}}
	return getTypeTx(req, label)
}

func getAttributeSetRegexReq(label *Label, regex string) *common2.Transaction_Req {
	s := &common2.AttributeType_SetRegex_Req{Regex: regex}
	req := &common2.Type_Req{Req: &common2.Type_Req_AttributeTypeSetRegexReq{AttributeTypeSetRegexReq: s}}
	return getTypeTx(req, label)
}
