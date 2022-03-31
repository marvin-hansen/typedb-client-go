// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package requests

import (
	"github.com/marvin-hansen/typedb-client-go/common"
)

// LogicManager
// https://github.com/vaticle/typedb-client-python/blob/master/typedb/common/rpc/request_builder.py

// GetLogicTx coverts a LogicManager_Req into a Transaction_Req
func GetLogicTx(req *common.LogicManager_Req) *common.Transaction_Req {
	r := &common.Transaction_Req_LogicManagerReq{LogicManagerReq: req}
	return &common.Transaction_Req{Req: r}
}

func GetLogicPutRuleReq(label, when, then string) *common.Transaction_Req {
	req := &common.LogicManager_PutRule_Req{
		Label: label,
		When:  when,
		Then:  then,
	}
	return GetLogicTx(&common.LogicManager_Req{Req: &common.LogicManager_Req_PutRuleReq{PutRuleReq: req}})
}

func GetLogicGetRuleReq(label string) *common.Transaction_Req {
	req := &common.LogicManager_GetRule_Req{
		Label: label,
	}
	return GetLogicTx(&common.LogicManager_Req{Req: &common.LogicManager_Req_GetRuleReq{GetRuleReq: req}})
}

func GetLogicGetAllRulesReq(label string) *common.Transaction_Req {
	req := &common.LogicManager_GetRules_Req{}
	return GetLogicTx(&common.LogicManager_Req{Req: &common.LogicManager_Req_GetRulesReq{GetRulesReq: req}})
}
