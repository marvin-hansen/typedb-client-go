package v2

import "github.com/marvin-hansen/typedb-client-go/common"

func (c *Client) OpenNewDataSession(dbName string) (*common.Session_Open_Res, error) {
	return c.newSession(dbName, common.Session_DATA)
}

func (c *Client) OpenNewSchemaSession(dbName string) (*common.Session_Open_Res, error) {
	return c.newSession(dbName, common.Session_SCHEMA)
}

func (c *Client) newSession(dbName string, sessionType common.Session_Type) (*common.Session_Open_Res, error) {
	openReq := getSessionOpenReq(dbName, sessionType, &common.Options{})
	session, openErr := c.client.SessionOpen(c.ctx, openReq)
	if openErr != nil {
		return nil, openErr
	} else {
		return session, nil
	}
}

func (c *Client) CloseSession(sessionID []byte) error {
	closeReq := getSessionCloseReq(sessionID)
	_, closeErr := c.client.SessionClose(c.ctx, closeReq)
	if closeErr != nil {
		return closeErr
	} else {
		return nil
	}
}
