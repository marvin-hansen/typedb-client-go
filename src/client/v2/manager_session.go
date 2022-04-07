// Copyright (c) 2022. Marvin Hansen | marvin.hansen@gmail.com

package v2

import (
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/marvin-hansen/typedb-client-go/src/client/v2/requests"
)

func NewSessionManager(client *Client) *SessionManager {
	return &SessionManager{
		client:     client,
		sessionMap: map[string]*TypeDBSession{},
	}
}

// SessionManager Encapsulates multiple session.
type SessionManager struct {
	client     *Client
	sessionMap map[string]*TypeDBSession
}

// NewSession creates a new session for the given client & DB
func (s SessionManager) NewSession(dbName string, sessionType common.Session_Type) (sessionID []byte, err error) {

	openReq := requests.GetSessionOpenReq(dbName, sessionType, &common.Options{})
	session, openErr := s.client.client.SessionOpen(s.client.ctx, openReq)
	if openErr != nil {
		return nil, openErr
	}

	// Start monitoring the session
	cancelFunc := s.startMonitorSession(s.client.ctx, session.SessionId)

	// construct session object from proto session & cancelFunc.
	typeDBSession := newTypeDBSession(session, cancelFunc)

	// Add session to the map
	sessionId := string(session.SessionId)
	s.sessionMap[sessionId] = typeDBSession

	return session.SessionId, nil
}

func (s SessionManager) GetSession(sessionID []byte) (session *TypeDBSession, ok bool, err error) {
	sessionId := string(sessionID)
	if s.checkSessionExists(sessionId) {
		session = s.sessionMap[sessionId]
		return session, true, nil
	} else {
		return nil, false, fmt.Errorf("Session does not exists for key: " + sessionId)
	}
}

func (s SessionManager) ResetSession(sessionID []byte) error {
	sessionId := string(sessionID)
	if s.checkSessionExists(sessionId) {
		session := s.sessionMap[sessionId].GetSession()
		session.Reset()
		return nil
	}
	return fmt.Errorf("Session does not exist for key: " + sessionId)
}

func (s SessionManager) CloseSession(sessionID []byte) (err error) {
	sessionId := string(sessionID)
	if s.checkSessionExists(sessionId) {
		closeReq := requests.GetSessionCloseReq(sessionID)
		_, closeErr := s.client.client.SessionClose(s.client.ctx, closeReq)

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
	if _, ok := s.sessionMap[sessionID]; ok {
		return true
	} else {
		return false
	}
}

func (s SessionManager) deleteSession(sessionID string) {
	if _, ok := s.sessionMap[sessionID]; ok {
		delete(s.sessionMap, sessionID)
	}
}
