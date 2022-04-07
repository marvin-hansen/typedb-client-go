package v2

import (
	"context"
	"github.com/marvin-hansen/typedb-client-go/common"
)

func newTypeDBSession(session *common.Session_Open_Res, cancelFunc context.CancelFunc) *TypeDBSession {
	return &TypeDBSession{
		session:    session,
		cancelFunc: cancelFunc,
	}
}

// TypeDBSession wraps teh proto session to store the corresponding cancel function
// required to stop monitoring. This way, we avoid a secondary data structure i.e. CancelFuncMap
type TypeDBSession struct {
	session    *common.Session_Open_Res
	cancelFunc context.CancelFunc
}

func (t TypeDBSession) GetSession() *common.Session_Open_Res {
	return t.session
}

func (t TypeDBSession) GetCancelFunc() context.CancelFunc {
	return t.cancelFunc
}
