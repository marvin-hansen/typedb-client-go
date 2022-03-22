package v2

import (
	"github.com/marvin-hansen/go-typedb/proto/common"
	pb "github.com/marvin-hansen/go-typedb/proto/core"
	"log"
)

func NewTransactionClient(client *pb.TypeDB_TransactionClient, sessionId []byte) *TransactionClient {
	return &TransactionClient{}
}

type TransactionClient struct {
	client        pb.TypeDB_TransactionClient
	sessionId     []byte
	metadata      map[string]string
	latencyMillis int
}

func (c TransactionClient) OpenTransaction() {
	err := c.client.Send(&common.Transaction_Client{Reqs: nil})
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func (c TransactionClient) CommitTransaction() {

}

func (c TransactionClient) RollbackTransaction() {

}

func (c TransactionClient) CloseTransaction() {

}
