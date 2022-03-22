package v2

import "github.com/marvin-hansen/go-typedb/proto/common"

// ConceptManager
// https://github.com/vaticle/typedb-client-python/blob/master/typedb/common/rpc/request_builder.py

// getConceptTx coverts a ConceptManager_Req into a Transaction_Req
func getConceptTx(req *common.ConceptManager_Req) *common.Transaction_Req {
	r := &common.Transaction_Req_ConceptManagerReq{ConceptManagerReq: req}
	return &common.Transaction_Req{Req: r}
}

func getConceptPutEntityReq(label string) *common.ConceptManager_Req {
	req := &common.ConceptManager_Req_PutEntityTypeReq{}
	req.PutEntityTypeReq = &common.ConceptManager_PutEntityType_Req{
		Label: label,
	}
	return &common.ConceptManager_Req{Req: req}
}

func getConceptPutRelationReq(label string) *common.ConceptManager_Req {
	req := &common.ConceptManager_Req_PutRelationTypeReq{}
	req.PutRelationTypeReq = &common.ConceptManager_PutRelationType_Req{
		Label: label,
	}
	return &common.ConceptManager_Req{Req: req}
}

func getConceptPutAttributeReq(label string, valueType common.AttributeType_ValueType) *common.ConceptManager_Req {
	req := &common.ConceptManager_Req_PutAttributeTypeReq{}
	req.PutAttributeTypeReq = &common.ConceptManager_PutAttributeType_Req{
		Label:     label,
		ValueType: valueType,
	}
	return &common.ConceptManager_Req{Req: req}
}

func getConceptGetThingTypeReq(label string) *common.ConceptManager_Req {
	req := &common.ConceptManager_Req_GetThingTypeReq{}
	req.GetThingTypeReq = &common.ConceptManager_GetThingType_Req{
		Label: label,
	}
	return &common.ConceptManager_Req{Req: req}
}

func getConceptGetThingReq(iid string) *common.ConceptManager_Req {
	req := &common.ConceptManager_Req_GetThingReq{}
	req.GetThingReq = &common.ConceptManager_GetThing_Req{
		Iid: []byte(iid),
	}
	return &common.ConceptManager_Req{Req: req}
}
