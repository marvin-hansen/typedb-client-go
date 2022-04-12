// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import (
	"context"
	"github.com/marvin-hansen/typedb-client-go/common"
	pb "github.com/marvin-hansen/typedb-client-go/core"
	"google.golang.org/grpc"
	"log"
)

type Client struct {
	client      pb.TypeDBClient
	config      *Config
	ctx         context.Context
	DBManager   *DBManager
	Session     *SessionManager
	Transaction *TransactionManager
	Query       *QueryManager
}

func NewOptions() *common.Options {
	return &common.Options{}
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
		client:  client,
		ctx:     ctx,
		Session: nil,
	}

	// construct session manager
	sessionManager := NewSessionManager(typeDBClient)
	typeDBClient.Session = sessionManager

	// construct db manager
	dbManager := NewDBManager(typeDBClient)
	typeDBClient.DBManager = dbManager

	// create TX manager
	txManager := NewTransactionManager(typeDBClient)
	typeDBClient.Transaction = txManager

	// construct query manager
	qManager := newQueryManager(typeDBClient)
	typeDBClient.Query = qManager

	return typeDBClient, nil
}

func (c Client) Close() {
	c.Session.Shutdown() // Closes all remaining sessions
}
