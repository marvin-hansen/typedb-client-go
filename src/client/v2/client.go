package v2

import (
	"context"
	pb "github.com/marvin-hansen/go-typedb/proto/core"
	"google.golang.org/grpc"
	"log"
	"sync"
)

type Client struct {
	client pb.TypeDBClient
	config *Config
	ctx    context.Context
	mtx    sync.RWMutex
}

func NewClient(conf *Config) (*Client, context.CancelFunc) {
	dbHost := conf.GetConnectionString()

	conn, cErr := newConnection(dbHost)
	if cErr != nil {
		log.Fatal("While trying to dial gRPC at: " + dbHost)
	}

	dbClient, clErr := newClient(conn)
	if clErr != nil {
		log.Fatal("While creating TypeDB client")
	}

	cancelFn := func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error while closing TypeDB connection:%v", err)
		}
	}

	return dbClient, cancelFn
}

func newConnection(target string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatal("While trying to dial gRPC")
	}
	return conn, err
}

func newClient(conn *grpc.ClientConn) (*Client, error) {
	checkConnection(conn, nil)
	client := pb.NewTypeDBClient(conn)
	ctx := context.Background()
	typeDBClient := &Client{
		client: client,
		ctx:    ctx,
	}
	return typeDBClient, nil
}
