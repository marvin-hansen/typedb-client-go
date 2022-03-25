// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

func newQueryManager(client *Client, sessionId []byte) *QueryManager {
	return &QueryManager{client: client, sessionId: sessionId}
}

type QueryManager struct {
	client    *Client
	sessionId []byte
}
