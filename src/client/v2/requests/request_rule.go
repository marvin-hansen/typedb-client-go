// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package requests

import "github.com/marvin-hansen/typedb-client-go/common"

// Rule
// 758

// GetRuleTx coverts a RuleReq into a Transaction_Req
func GetRuleTx(req *common.Rule_Req, label string) *common.Transaction_Req {
	req.Label = label
	r := &common.Transaction_Req_RuleReq{RuleReq: req}
	return &common.Transaction_Req{Req: r}
}

func GetRuleSetLabelReq(label string, newLabel string) (*common.Transaction_Req, error) {
	s := &common.Rule_SetLabel_Req{Label: newLabel}
	req := &common.Rule_Req{Req: &common.Rule_Req_RuleSetLabelReq{RuleSetLabelReq: s}}
	return GetRuleTx(req, label), nil
}

func GetRuleDeleteReq(label string) (*common.Transaction_Req, error) {
	s := &common.Rule_Delete_Req{}
	req := &common.Rule_Req{Req: &common.Rule_Req_RuleDeleteReq{RuleDeleteReq: s}}
	return GetRuleTx(req, label), nil
}
