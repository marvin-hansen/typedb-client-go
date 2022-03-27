package v2

import (
	"fmt"
	"github.com/marvin-hansen/typedb-client-go/common"
)

//
// Methods with streaming results i.e one initial request -> multiple partial stream results.
//
// From the proto spec in common/query.proto @line 56
//
//   message ResPart {
//    oneof res {
//      Match.ResPart match_res_part = 100;
//      MatchGroup.ResPart match_group_res_part = 101;
//      MatchGroupAggregate.ResPart match_group_aggregate_res_part = 102;
//      Insert.ResPart insert_res_part = 103;
//      Update.ResPart update_res_part = 104;
//      Explain.ResPart explain_res_part = 105;
//    }
//  }

func (c *Client) RunMatchQuery(requestId []byte, query string, metadata map[string]string, options *common.Options) (queryResults []*common.QueryManager_Match_ResPart, err error) {

	if options == nil {
		options = &common.Options{}
	}

	// Create a request and attach meta data & request ID
	r1 := getMatchQueryReq(query, options, requestId, metadata)
	// Stuff req into slice/array
	var req []*common.Transaction_Req
	req[0] = r1

	// Create a Transaction
	tx, txErr := c.client.Transaction(c.ctx)
	if txErr != nil {
		return nil, txErr
	}

	// Send request through
	sendErr := tx.Send(getTransactionClient(req))
	if sendErr != nil {
		return nil, sendErr
	}

	// process server stream

	for {
		// get return value
		recs, recErr := tx.Recv()
		if recErr != nil {
			return nil, fmt.Errorf("could not receive query response: %w", recErr)
		}

		state := recs.GetResPart().GetStreamResPart().GetState()

		// When the server sends a Stream.ResPart with state = CONTINUE it indicates that there are more answers to fetch,
		// so the client should respond with Stream.Req
		if state == CONTINUE {
			// Create a request and attach meta data & request ID
			rc := getTransactionStreamReq()
			// Stuff req into slice/array
			var reqCont []*common.Transaction_Req
			reqCont[0] = rc
			// run query
			_, queryErr := c.runQuery(req)
			if queryErr != nil {
				return nil, fmt.Errorf("could not send query request iterator: %w", queryErr)
			}

		}

		//  if the Stream.ResPart has state = DONE, it indicates that there are no more answers to fetch, so the client doesn't need to respond.
		if state == DONE {
			break
		} else {
			part := recs.GetResPart().GetQueryManagerResPart().GetMatchResPart()
			queryResults = append(queryResults, part)
		}

	}

	return queryResults, nil
}
