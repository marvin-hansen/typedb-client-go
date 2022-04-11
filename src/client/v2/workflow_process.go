package v2

import "github.com/marvin-hansen/typedb-client-go/common"

func (c *Client) isStreamDone(reqResult *common.Transaction_ResPart, stream bool) (loopBreak bool) {
	if stream {
		state := reqResult.GetStreamResPart().GetState()
		if state == DONE {
			return true
		}
		if state == CONTINUE {
			return false
		}
	}
	return false // no stream
}
