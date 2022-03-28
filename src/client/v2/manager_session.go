package v2

import "C"
import "github.com/marvin-hansen/typedb-client-go/common"

func NewSession(client *Client, dbName string, sessionType common.Session_Type) (*SessionManager, error) {

	openReq := getSessionOpenReq(dbName, sessionType, &common.Options{})
	session, openErr := client.client.SessionOpen(client.ctx, openReq)
	if openErr != nil {
		return nil, openErr
	}

	return &SessionManager{
		client:  client,
		session: session,
	}, nil
}

type SessionManager struct {
	client  *Client
	session *common.Session_Open_Res
}

func (s SessionManager) GetSessionId() []byte {
	return s.session.GetSessionId()
}

func (s SessionManager) Reset() {
	s.session.Reset() // No idea what the reset function actually rests...
}

func (s SessionManager) Close() error {
	sessionID := s.GetSessionId()
	closeReq := getSessionCloseReq(sessionID)
	_, closeErr := s.client.client.SessionClose(s.client.ctx, closeReq)
	if closeErr != nil {
		return closeErr
	} else {
		return nil
	}
}
