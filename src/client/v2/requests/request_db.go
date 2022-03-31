// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package requests

import (
	pb "github.com/marvin-hansen/typedb-client-go/core"
)

// CoreDatabaseManager
// https://github.com/vaticle/typedb-client-python/blob/master/typedb/common/rpc/request_builder.py

func GetAllDBReq() (req *pb.CoreDatabaseManager_All_Req) {
	return &pb.CoreDatabaseManager_All_Req{}
}

func GetCreateDBReq(dbName string) *pb.CoreDatabaseManager_Create_Req {
	return &pb.CoreDatabaseManager_Create_Req{Name: dbName}
}

func GetContainsDBReq(dbName string) *pb.CoreDatabaseManager_Contains_Req {
	return &pb.CoreDatabaseManager_Contains_Req{Name: dbName}
}

func GetDeleteDBReq(dbName string) *pb.CoreDatabase_Delete_Req {
	return &pb.CoreDatabase_Delete_Req{Name: dbName}
}
