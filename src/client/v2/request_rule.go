// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import "github.com/marvin-hansen/typedb-client-go/common"

// Rule
// 758

// getRuleTx coverts a RuleReq into a Transaction_Req
func getRuleTx(req *common.Rule_Req, label string) *common.Transaction_Req {
	req.Label = label
	r := &common.Transaction_Req_RuleReq{RuleReq: req}
	return &common.Transaction_Req{Req: r}
}

func getRuleSetLabelReq(label string, newLabel string) (*common.Transaction_Req, error) {
	s := &common.Rule_SetLabel_Req{Label: newLabel}
	req := &common.Rule_Req{Req: &common.Rule_Req_RuleSetLabelReq{RuleSetLabelReq: s}}
	return getRuleTx(req, label), nil
}

func getRuleDeleteReq(label string) (*common.Transaction_Req, error) {
	s := &common.Rule_Delete_Req{}
	req := &common.Rule_Req{Req: &common.Rule_Req_RuleDeleteReq{RuleDeleteReq: s}}
	return getRuleTx(req, label), nil
}
