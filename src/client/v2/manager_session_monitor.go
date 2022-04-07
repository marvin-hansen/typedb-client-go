package v2

import (
	"context"
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/src/client/v2/requests"
)

func (s SessionManager) startMonitorSession(sessionID []byte) {
	// How do I track each goRoutine per session

}

func (s SessionManager) stopMonitorSession(sessionID []byte) {
	// How do I call the correct cancle function to stop the GoRoutine matchin the session?
}

func (s SessionManager) runHeartbeat(ctx context.Context, sessionID []byte) context.CancelFunc {

	// Create a new context, with its cancellation function cfrom the original context
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Stopped monitoring session: ")
		default:
			err := s.sendPulseRequest(sessionID)
			// If this operation returns an error
			// cancel all operations using this local context created above
			if err != nil {
				cancel()
			}
			fmt.Println("done")
		}
	}()

	// return cancel function for callsite to close at a later stage
	return cancel
}

func (s SessionManager) sendPulseRequest(sessionID []byte) error {
	mtd := "sendPulse: "

	req := requests.GetSessionPulseReq(sessionID)
	res, pulseErr := s.client.client.SessionPulse(s.client.ctx, req)
	if pulseErr != nil {
		dbgPrint(mtd, "Heartbeat error. Close session")
		return pulseErr
	}
	if res.Alive == false {
		dbgPrint(mtd, "Server not alive anymore. Close session")
		closeErr := s.CloseSession(sessionID)
		if closeErr != nil {
			return closeErr
		}
	}

	// no error
	return nil
}
