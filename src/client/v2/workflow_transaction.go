package v2

import pb "github.com/marvin-hansen/go-typedb/core"

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
}

func (c TransactionClient) CommitTransaction() {

}

func (c TransactionClient) RollbackTransaction() {

}

func (c TransactionClient) CloseTransaction() {

}
