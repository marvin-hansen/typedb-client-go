package v2

import (
	"github.com/marvin-hansen/typedb-client-go/common"
)

// NewSession creates a new session for the given client & DB
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

// SessionManager Encapsulates one single session.
type SessionManager struct {
	client  *Client
	session *common.Session_Open_Res
}

// GetSessionId returns the current session ID
func (s SessionManager) GetSessionId() []byte {
	return s.session.GetSessionId()
}

// Reset -  No idea what the reset function actually rests, but it's in the specs....
func (s SessionManager) Reset() {
	s.session.Reset()
}

// Close - Ends the current session.
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
