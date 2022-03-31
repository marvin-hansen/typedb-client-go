// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com
package requests

import "github.com/marvin-hansen/typedb-client-go/common"

// Relation
// 675

func GetRelationAddPlayerReq(iid string, roleType *common.Type, player *common.Thing) *common.Transaction_Req {
	s := &common.Relation_AddPlayer_Req{
		RoleType: roleType,
		Player:   player,
	}
	req := &common.Thing_Req{Req: &common.Thing_Req_RelationAddPlayerReq{RelationAddPlayerReq: s}}
	return GetThingTx(req, iid)
}

func GetRelationRemovePlayerReq(iid string, roleType *common.Type, player *common.Thing) *common.Transaction_Req {
	s := &common.Relation_RemovePlayer_Req{
		RoleType: roleType,
		Player:   player,
	}
	req := &common.Thing_Req{Req: &common.Thing_Req_RelationRemovePlayerReq{RelationRemovePlayerReq: s}}
	return GetThingTx(req, iid)
}

func GetRelationGetAllPlayersReq(iid string, roleTypes *[]*common.Type) *common.Transaction_Req {
	s := &common.Relation_GetPlayers_Req{RoleTypes: *roleTypes}
	req := &common.Thing_Req{Req: &common.Thing_Req_RelationGetPlayersReq{RelationGetPlayersReq: s}}
	return GetThingTx(req, iid)
}

func GetRelationGetPlayersByRoleTypeReq(iid string) *common.Transaction_Req {
	s := &common.Relation_GetPlayersByRoleType_Req{}
	req := &common.Thing_Req{Req: &common.Thing_Req_RelationGetPlayersByRoleTypeReq{RelationGetPlayersByRoleTypeReq: s}}
	return GetThingTx(req, iid)
}

func GetRelationGetRelatingReq(iid string) *common.Transaction_Req {
	s := &common.Relation_GetRelating_Req{}
	req := &common.Thing_Req{Req: &common.Thing_Req_RelationGetRelatingReq{RelationGetRelatingReq: s}}
	return GetThingTx(req, iid)
}
