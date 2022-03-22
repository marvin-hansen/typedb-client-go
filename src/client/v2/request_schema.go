package v2

import pb "github.com/marvin-hansen/go-typedb/proto/core"

// Schema

func getDBSchemaReq(dbName string) *pb.CoreDatabase_Schema_Req {
	return &pb.CoreDatabase_Schema_Req{Name: dbName}
}
