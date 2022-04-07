package v2

import (
	"context"
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/src/client/v2/requests"
	"time"
)

const heartBeatInterval = 3

func (s SessionManager) startMonitorSession(ctx context.Context, sessionID []byte) context.CancelFunc {
	return s.runHeartbeat(ctx, sessionID)
}

func (s SessionManager) stopMonitorSession(sessionID []byte) (err error) {
	session, err := s.GetSession(sessionID)
	if err != nil {
		return fmt.Errorf("error. Could not retrieve session: %w", err)
	}
	cancelFunc := session.GetCancelFunc()
	cancelFunc()
	return nil
}

func (s SessionManager) runHeartbeat(ctx context.Context, sessionID []byte) context.CancelFunc {

	//mtd := "runHeartbeat"
	//dbgPrint(mtd, "Create a new context, with its cancellation function from the original context")
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		select {

		case <-ctx.Done():
			//dbgPrint(mtd, "Stopped monitoring session: "+byteToString(sessionID))
			break

		default:
			for {
				err := s.sendPulseRequest(sessionID)
				// If this operation returns an error cancel using this local context created above
				if err != nil {
					//dbgPrint(mtd, "Heartbeat error detected. closing session: "+byteToString(sessionID))
					cancel()
				}
				// sleep a bit before repeat
				time.Sleep(time.Second * heartBeatInterval)
			}
		}
	}()

	// return cancel function for callsite to close at a later stage
	return cancel
}

func (s SessionManager) sendPulseRequest(sessionID []byte) error {

	//mtd := "sendPulse: "
	//dbgPrint(mtd, "Sending pulse request for session: "+byteToString(sessionID))
	req := requests.GetSessionPulseReq(sessionID)
	res, pulseErr := s.state.client.client.SessionPulse(s.state.client.ctx, req)
	if pulseErr != nil {
		//dbgPrint(mtd, "Heartbeat error. Close session: "+byteToString(sessionID))
		return pulseErr
	}
	if res.Alive == false {
		//dbgPrint(mtd, "Server not alive anymore. Close session: "+byteToString(sessionID))
		closeErr := s.CloseSession(sessionID)
		if closeErr != nil {
			return closeErr
		}
	}
	return nil
}
