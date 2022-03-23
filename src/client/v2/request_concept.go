package v2

import (
	common2 "github.com/marvin-hansen/go-typedb/common"
)

// ConceptManager
// https://github.com/vaticle/typedb-client-python/blob/master/typedb/common/rpc/request_builder.py

// getConceptTx coverts a ConceptManager_Req into a Transaction_Req
func getConceptTx(req *common2.ConceptManager_Req) *common2.Transaction_Req {
	r := &common2.Transaction_Req_ConceptManagerReq{ConceptManagerReq: req}
	return &common2.Transaction_Req{Req: r}
}

func getConceptPutEntityReq(label string) *common2.ConceptManager_Req {
	req := &common2.ConceptManager_Req_PutEntityTypeReq{}
	req.PutEntityTypeReq = &common2.ConceptManager_PutEntityType_Req{
		Label: label,
	}
	return &common2.ConceptManager_Req{Req: req}
}

func getConceptPutRelationReq(label string) *common2.ConceptManager_Req {
	req := &common2.ConceptManager_Req_PutRelationTypeReq{}
	req.PutRelationTypeReq = &common2.ConceptManager_PutRelationType_Req{
		Label: label,
	}
	return &common2.ConceptManager_Req{Req: req}
}

func getConceptPutAttributeReq(label string, valueType common2.AttributeType_ValueType) *common2.ConceptManager_Req {
	req := &common2.ConceptManager_Req_PutAttributeTypeReq{}
	req.PutAttributeTypeReq = &common2.ConceptManager_PutAttributeType_Req{
		Label:     label,
		ValueType: valueType,
	}
	return &common2.ConceptManager_Req{Req: req}
}

func getConceptGetThingTypeReq(label string) *common2.ConceptManager_Req {
	req := &common2.ConceptManager_Req_GetThingTypeReq{}
	req.GetThingTypeReq = &common2.ConceptManager_GetThingType_Req{
		Label: label,
	}
	return &common2.ConceptManager_Req{Req: req}
}

func getConceptGetThingReq(iid string) *common2.ConceptManager_Req {
	req := &common2.ConceptManager_Req_GetThingReq{}
	req.GetThingReq = &common2.ConceptManager_GetThing_Req{
		Iid: []byte(iid),
	}
	return &common2.ConceptManager_Req{Req: req}
}
