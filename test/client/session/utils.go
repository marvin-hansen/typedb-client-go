package session

import (
	"context"
	"encoding/hex"
	v2 "github.com/marvin-hansen/typedb-client-go/src/client/v2"
	"github.com/marvin-hansen/typedb-client-go/test/client/utils"
)

const (
	dbName  = utils.DBName
	verbose = true
)

func byteToString(sessionID []byte) string {
	return hex.EncodeToString(sessionID)
}

func testPrint(msg string) {
	if verbose {
		println(msg)
	}
}

func getClient() (*v2.Client, context.CancelFunc) {
	conf := v2.NewLocalConfig(dbName)
	client, cancel := v2.NewClient(conf)
	return client, cancel
}
