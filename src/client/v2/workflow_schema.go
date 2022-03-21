package v2

import (
	"context"
	pb "github.com/marvin-hansen/go-typedb/proto/core"
	"google.golang.org/grpc"
)

func DatabaseSchema(ctx context.Context, in *pb.CoreDatabase_Schema_Req, opts ...grpc.CallOption) (*pb.CoreDatabase_Schema_Res, error) {
	return nil, nil
}
