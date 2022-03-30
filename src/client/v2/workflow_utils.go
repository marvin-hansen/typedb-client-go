// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import "github.com/segmentio/ksuid"

// CreateNewRequestMetadata returns a new map of empty metadata
func CreateNewRequestMetadata() map[string]string {
	return map[string]string{}
}

// CreateNewRequestID generates a new unique request ID to use in transactions
func CreateNewRequestID() []byte {
	return ksuid.New().Bytes()
}

// CreateNewStringRequestID generates a new unique request ID to use in transactions
func CreateNewStringRequestID() string {
	return ksuid.New().String()
}

func getSchemaQuery() string {
	return `
	match
		$x sub thing;
	get
		$x;
	`
}
