// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com
package v2

import (
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/marvin-hansen/typedb-client-go/src/err"
	"strings"
)

// Thing
// 605

func convertStringToByte(str string) []byte {
	//trimming the spaces, this method will remote the '\t', '\n', '\v', '\f', '\r', ' ' chars
	return []byte(strings.TrimSpace(str))
}

func convertThingToProto(iid string) *common.Thing {
	return &common.Thing{Iid: convertStringToByte(iid)}
}

// getThingTx coverts a ThingReq into a Transaction_Req
func getThingTx(req *common.Thing_Req, iid string) *common.Transaction_Req {
	req.Iid = convertStringToByte(iid)
	r := &common.Transaction_Req_ThingReq{ThingReq: req}
	return &common.Transaction_Req{Req: r}
}

func getThingHasReq(iid string, attributeTypes *[]*common.Type, onlyKey bool) (*common.Transaction_Req, error) {

	if attributeTypes != nil && onlyKey {
		return nil, err.TypeDBConceptError(err.GET_HAS_WITH_MULTIPLE_FILTERS)
	}

	s := &common.Thing_GetHas_Req{}
	if onlyKey {
		s.KeysOnly = onlyKey
	}

	if attributeTypes != nil {
		s.AttributeTypes = *attributeTypes
	}

	req := &common.Thing_Req{Req: &common.Thing_Req_ThingGetHasReq{ThingGetHasReq: s}}
	return getThingTx(req, iid), nil
}

func getThingSetHasReq(iid string, attribute *common.Thing) (*common.Transaction_Req, error) {
	s := &common.Thing_SetHas_Req{Attribute: attribute}
	req := &common.Thing_Req{Req: &common.Thing_Req_ThingSetHasReq{ThingSetHasReq: s}}
	return getThingTx(req, iid), nil
}

func getThingUnsetHasReq(iid string, attribute *common.Thing) (*common.Transaction_Req, error) {
	s := &common.Thing_UnsetHas_Req{Attribute: attribute}
	req := &common.Thing_Req{Req: &common.Thing_Req_ThingUnsetHasReq{ThingUnsetHasReq: s}}
	return getThingTx(req, iid), nil
}

func getThingPlayingReq(iid string) (*common.Transaction_Req, error) {
	s := &common.Thing_GetPlaying_Req{}
	req := &common.Thing_Req{Req: &common.Thing_Req_ThingGetPlayingReq{ThingGetPlayingReq: s}}
	return getThingTx(req, iid), nil
}

func getThinRelationsReq(iid string, roleTypes *[]*common.Type) (*common.Transaction_Req, error) {
	s := &common.Thing_GetRelations_Req{RoleTypes: *roleTypes}
	req := &common.Thing_Req{Req: &common.Thing_Req_ThingGetRelationsReq{ThingGetRelationsReq: s}}
	return getThingTx(req, iid), nil
}

func getThingDeleteReq(iid string) (*common.Transaction_Req, error) {
	s := &common.Thing_Delete_Req{}
	req := &common.Thing_Req{Req: &common.Thing_Req_ThingDeleteReq{ThingDeleteReq: s}}
	return getThingTx(req, iid), nil
}
