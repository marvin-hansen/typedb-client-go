// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import (
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/marvin-hansen/typedb-client-go/src/client/v2/requests"
)

func NewSessionManager(client *Client) *SessionManager {
	return &SessionManager{
		state: newSessionManagerState(client),
	}
}

// SessionManager encapsulates multiple session.
type SessionManager struct {
	state *SessionManagerState
}

// NewSession creates a new session for the DB
func (s SessionManager) NewSession(dbName string, sessionType common.Session_Type) (sessionID []byte, err error) {

	openReq := requests.GetSessionOpenReq(dbName, sessionType, &common.Options{})
	session, openErr := s.state.client.client.SessionOpen(s.state.client.ctx, openReq)
	if openErr != nil {
		return nil, openErr
	}

	// Start monitoring the session
	cancelFunc := s.startMonitorSession(s.state.client.ctx, session.SessionId)

	// Construct session object from proto session & cancelFunc.
	typeDBSession := newTypeDBSession(session, cancelFunc)

	// Add session to the map
	sessionId := string(session.SessionId)
	s.state.sessionMap[sessionId] = typeDBSession

	return session.SessionId, nil
}

func (s SessionManager) GetSession(sessionID []byte) (session *TypeDBSession, err error) {
	sessionId := string(sessionID)
	if s.checkSessionExists(sessionId) {
		session = s.state.sessionMap[sessionId]
		return session, nil
	} else {
		return nil, fmt.Errorf("Session does not exists for key: " + sessionId)
	}
}

func (s SessionManager) ResetSession(sessionID []byte) error {
	sessionId := string(sessionID)
	if s.checkSessionExists(sessionId) {
		session := s.state.sessionMap[sessionId].GetSession()
		session.Reset()
		return nil
	}
	return fmt.Errorf("Session does not exist for key: " + sessionId)
}

func (s SessionManager) CloseSession(sessionID []byte) (err error) {
	sessionId := string(sessionID)
	if s.checkSessionExists(sessionId) {
		closeReq := requests.GetSessionCloseReq(sessionID)
		_, closeErr := s.state.client.client.SessionClose(s.state.client.ctx, closeReq)

		// Stop heartbeat regardless of close success.
		// When the server stops receiving heartbeats, it will close the session after 30 sec. of inactivity
		stopErr := s.stopMonitorSession(sessionID)
		if stopErr != nil {
			return stopErr
		}

		// Delete closed session regardless of close success.
		// if client side close fails, the server will close the session after time out.
		s.deleteSession(sessionId)

		// Check for error, just in case
		if closeErr != nil {
			return closeErr
		} else {
			return nil
		}
	} else {
		return fmt.Errorf("Session does not exists for key: " + sessionId)
	}
}

func (s SessionManager) checkSessionExists(sessionID string) (exists bool) {
	if _, ok := s.state.sessionMap[sessionID]; ok {
		return true
	} else {
		return false
	}
}

func (s SessionManager) deleteSession(sessionID string) {
	if _, ok := s.state.sessionMap[sessionID]; ok {
		delete(s.state.sessionMap, sessionID)
	}
}

// Shutdown should be called when closing the client to end all remaining sessions.
func (s SessionManager) Shutdown() {
	if len(s.state.sessionMap) > 0 {
		for _, session := range s.state.sessionMap {
			sessionID := session.GetSession().GetSessionId()
			_ = s.CloseSession(sessionID)
		}
	}
}
