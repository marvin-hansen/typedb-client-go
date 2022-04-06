package requests

import "github.com/segmentio/ksuid"

// CreateNewRequestID generates a new unique request ID to use in transactions
func createNewRequestID() []byte {
	return ksuid.New().Bytes()
}
