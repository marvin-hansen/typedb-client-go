// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import (
	"context"
	pb "github.com/marvin-hansen/typedb-client-go/core"
	"google.golang.org/grpc"
	"log"
)

type Client struct {
	client             pb.TypeDBClient
	config             *Config
	ctx                context.Context
	DBManager          *DBManager
	SessionManager     *SessionManager
	TransactionManager *TransactionManager
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
		client:         client,
		ctx:            ctx,
		SessionManager: nil,
	}

	// create session manager
	sessionManager := NewSessionManager(typeDBClient)
	typeDBClient.SessionManager = sessionManager

	// create db manager
	dbManager := NewDBManager(typeDBClient)
	typeDBClient.DBManager = dbManager

	txManager := NewTransactionManager(typeDBClient)
	typeDBClient.TransactionManager = txManager

	//
	return typeDBClient, nil
}
