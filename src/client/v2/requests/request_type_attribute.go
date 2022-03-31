// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package requests

import (
	"github.com/marvin-hansen/typedb-client-go/common"
	. "github.com/marvin-hansen/typedb-client-go/src/type/common_type"
)

// EntityType
// 565

func GetAttributeGetOwner(label *Label, onlyKey bool) *common.Transaction_Req {
	s := &common.AttributeType_GetOwners_Req{OnlyKey: onlyKey}
	req := &common.Type_Req{Req: &common.Type_Req_AttributeTypeGetOwnersReq{AttributeTypeGetOwnersReq: s}}
	return GetTypeTx(req, label)
}

func GetAttributePutReq(label *Label, value *common.Attribute_Value) *common.Transaction_Req {
	s := &common.AttributeType_Put_Req{Value: value}
	req := &common.Type_Req{Req: &common.Type_Req_AttributeTypePutReq{AttributeTypePutReq: s}}
	return GetTypeTx(req, label)
}

func GetAttributeReq(label *Label, value *common.Attribute_Value) *common.Transaction_Req {
	s := &common.AttributeType_Get_Req{Value: value}
	req := &common.Type_Req{Req: &common.Type_Req_AttributeTypeGetReq{AttributeTypeGetReq: s}}
	return GetTypeTx(req, label)
}

func GetAttributeGetRegexReq(label *Label) *common.Transaction_Req {
	s := &common.AttributeType_GetRegex_Req{}
	req := &common.Type_Req{Req: &common.Type_Req_AttributeTypeGetRegexReq{AttributeTypeGetRegexReq: s}}
	return GetTypeTx(req, label)
}

func GetAttributeSetRegexReq(label *Label, regex string) *common.Transaction_Req {
	s := &common.AttributeType_SetRegex_Req{Regex: regex}
	req := &common.Type_Req{Req: &common.Type_Req_AttributeTypeSetRegexReq{AttributeTypeSetRegexReq: s}}
	return GetTypeTx(req, label)
}
