// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com
package v2

import "github.com/marvin-hansen/typedb-client-go/common"

// Relation
// 719

func getAttributeOwnerReq(iid string, ownerType *common.Type) *common.Transaction_Req {
	s := &common.Attribute_GetOwners_Req{}
	if ownerType != nil {
		s.Filter = &common.Attribute_GetOwners_Req_ThingType{
			ThingType: ownerType,
		}
	}
	req := &common.Thing_Req{Req: &common.Thing_Req_AttributeGetOwnersReq{AttributeGetOwnersReq: s}}
	return getThingTx(req, iid)
}
