package v2

import (
	common2 "github.com/marvin-hansen/go-typedb/common"
)

// LogicManager
// https://github.com/vaticle/typedb-client-python/blob/master/typedb/common/rpc/request_builder.py

// getLogicTx coverts a LogicManager_Req into a Transaction_Req
func getLogicTx(req *common2.LogicManager_Req) *common2.Transaction_Req {
	r := &common2.Transaction_Req_LogicManagerReq{LogicManagerReq: req}
	return &common2.Transaction_Req{Req: r}
}

func getLogicPutRuleReq(label, when, then string) *common2.LogicManager_Req {
	req := &common2.LogicManager_PutRule_Req{
		Label: label,
		When:  when,
		Then:  then,
	}
	return &common2.LogicManager_Req{Req: &common2.LogicManager_Req_PutRuleReq{PutRuleReq: req}}
}

func getLogicGetRuleReq(label string) *common2.LogicManager_Req {
	req := &common2.LogicManager_GetRule_Req{
		Label: label,
	}
	return &common2.LogicManager_Req{Req: &common2.LogicManager_Req_GetRuleReq{GetRuleReq: req}}
}

func getLogicGetAllRulesReq(label string) *common2.LogicManager_Req {
	req := &common2.LogicManager_GetRules_Req{}
	return &common2.LogicManager_Req{Req: &common2.LogicManager_Req_GetRulesReq{GetRulesReq: req}}
}
