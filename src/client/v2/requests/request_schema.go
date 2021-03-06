// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package requests

import (
	pb "github.com/marvin-hansen/typedb-client-go/core"
)

// Schema

func GetDBSchemaReq(dbName string) *pb.CoreDatabase_Schema_Req {
	return &pb.CoreDatabase_Schema_Req{Name: dbName}
}
