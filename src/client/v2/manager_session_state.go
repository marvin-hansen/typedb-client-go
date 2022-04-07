package v2

// SessionManagerState separates state from function
type SessionManagerState struct {
	client     *Client
	sessionMap map[string]*TypeDBSession
}

func newSessionManagerState(client *Client) *SessionManagerState {
	return &SessionManagerState{
		client:     client,
		sessionMap: make(map[string]*TypeDBSession),
	}
}
