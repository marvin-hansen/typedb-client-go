package query

import (
	"github.com/marvin-hansen/typedb-client-go/common"
	"github.com/marvin-hansen/typedb-client-go/test/client/utils"
)

func printResult(queryResults []*common.QueryManager_Match_ResPart, verbose bool) {
	if verbose {
		utils.TestPrint("* Print Query Results: ")
		println("* Print results: ")
		for _, item := range queryResults {
			println(item.String())
		}
	}
}
