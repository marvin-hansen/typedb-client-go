package v2

import (
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/marvin-hansen/typedb-client-go/core"
)

func (c *Client) runSchemaQuery(tx core.TypeDB_TransactionClient, sessionID []byte) (queryResults []*common.QueryManager_ResPart, err error) {

	options := &common.Options{InferOpt: &common.Options_Infer{Infer: true}, ExplainOpt: &common.Options_Explain{Explain: true}}

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

	return nil, nil
}

