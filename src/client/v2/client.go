package v2

import (
	"context"
	pb "github.com/marvin-hansen/go-typedb/proto/core"
	"google.golang.org/grpc"
	"time"
)

type Client struct {
	client pb.TypeDBClient
	config *Config
	ctx    context.Context
	cnc    context.CancelFunc
}

func NewClient(conn *grpc.ClientConn) (*Client, error) {
	// CheckConnection(conn, nil)
	client := pb.NewTypeDBClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	amdbClient := &Client{
		client: client,
		ctx:    ctx,
		cnc:    cancel,
	}
	return amdbClient, nil
}
