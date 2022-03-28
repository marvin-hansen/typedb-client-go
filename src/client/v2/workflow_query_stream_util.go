package v2

import (
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/segmentio/ksuid"
)

const (
	CONTINUE = common.Transaction_Stream_CONTINUE
	DONE     = common.Transaction_Stream_DONE
	TX_READ  = common.Transaction_READ
	TX_WRITE = common.Transaction_WRITE
)

//runStreamQuery util used by all other streaming return value query methods
func (c *Client) runStreamQuery(sessionID []byte, transactionType common.Transaction_Type, req *common.Transaction_Req, options *common.Options) (queryResults []*common.QueryManager_ResPart, err error) {

	// Create a Transaction
	tx, txErr := c.client.Transaction(c.ctx)
	if txErr != nil {
		return nil, fmt.Errorf("could not create transaction: %w", txErr)
	}

	transactionId := ksuid.New().Bytes()

	// Create open transaction request
	netMillisecondLatency := int32(150)
	openReq := getTransactionOpenReq(sessionID, transactionId, transactionType, options, netMillisecondLatency)
	// Stuff req into slice/array
	reqArray := []*common.Transaction_Req{openReq, req}

	// Send request through
	sendErr := tx.Send(getTransactionClient(reqArray))
	if sendErr != nil {
		return nil, fmt.Errorf("could not send transaction to server: %w", sendErr)
	}

	for {
		// Get return value
		recs, recErr := tx.Recv()
		if recErr != nil {
			return nil, fmt.Errorf("could not receive query response: %w", recErr)
		}

		// Extract state of current partial result
		state := recs.GetResPart().GetStreamResPart().GetState()

		// When the server sends a Stream.ResPart with state = CONTINUE
		// it indicates that there are more answers to fetch,
		// so the client should respond with Stream.Req
		if state == CONTINUE {
			// Create a request and attach meta data & request ID
			reqCont := getTransactionStreamReq()
			// run query
			_, queryErr := c.runQuery(sessionID, reqCont, options)
			if queryErr != nil {
				return nil, fmt.Errorf("could not send query request iterator: %w", queryErr)
			}
		}

		// If the Stream.ResPart has state = DONE,
		// it indicates that there are no more answers to fetch,
		// so the client doesn't need to respond.
		if state == DONE {
			break

		} else {
			part := recs.GetResPart().GetQueryManagerResPart()
			queryResults = append(queryResults, part)
		}
	}

	// Close transaction after the last part returned
	closErr := tx.CloseSend()
	if closErr != nil {
		return nil, fmt.Errorf("could not close query transaction: %w", closErr)
	}

	return queryResults, nil
}
