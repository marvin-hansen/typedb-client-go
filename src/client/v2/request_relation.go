package v2

import "github.com/marvin-hansen/typedb-client-go/common"

// Relation
// 675

func getRelationAddPlayerReq(iid string, roleType *common.Type, player *common.Thing) *common.Transaction_Req {
	s := &common.Relation_AddPlayer_Req{
		RoleType: roleType,
		Player:   player,
	}
	req := &common.Thing_Req{Req: &common.Thing_Req_RelationAddPlayerReq{RelationAddPlayerReq: s}}
	return getThingTx(req, iid)
}

func getRelationRemovePlayerReq(iid string, roleType *common.Type, player *common.Thing) *common.Transaction_Req {
	s := &common.Relation_RemovePlayer_Req{
		RoleType: roleType,
		Player:   player,
	}
	req := &common.Thing_Req{Req: &common.Thing_Req_RelationRemovePlayerReq{RelationRemovePlayerReq: s}}
	return getThingTx(req, iid)
}

func getRelationGetAllPlayersReq(iid string, roleTypes *[]*common.Type) *common.Transaction_Req {
	s := &common.Relation_GetPlayers_Req{RoleTypes: *roleTypes}
	req := &common.Thing_Req{Req: &common.Thing_Req_RelationGetPlayersReq{RelationGetPlayersReq: s}}
	return getThingTx(req, iid)
}

func getRelationGetPlayersByRoleTypeReq(iid string) *common.Transaction_Req {
	s := &common.Relation_GetPlayersByRoleType_Req{}
	req := &common.Thing_Req{Req: &common.Thing_Req_RelationGetPlayersByRoleTypeReq{RelationGetPlayersByRoleTypeReq: s}}
	return getThingTx(req, iid)
}

func getRelationGetRelatingReq(iid string) *common.Transaction_Req {
	s := &common.Relation_GetRelating_Req{}
	req := &common.Thing_Req{Req: &common.Thing_Req_RelationGetRelatingReq{RelationGetRelatingReq: s}}
	return getThingTx(req, iid)
}
