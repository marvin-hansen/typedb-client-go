package utils

import (
	"context"
	v2 "github.com/marvin-hansen/typedb-client-go/src/client/v2"
)

const (
	DbName  = DBName
	verbose = true
)

func TestPrint(msg string) {
	if verbose {
		println(msg)
	}
}

func GetClient() (*v2.Client, context.CancelFunc) {
	conf := v2.NewLocalConfig(DbName)
	client, cancel := v2.NewClient(conf)
	return client, cancel
}
