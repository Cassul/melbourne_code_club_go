package search

import (
	"fmt"

	"github.com/zendesk/melbourne_code_club_go/internal/types"
)

func SearchData(index types.Index, query types.Query) string {
	results := index[query]

	var resultSum string
	for _, result := range results {
		resultSum = resultSum + result.Print(index) + "\n"
	}

	resultSum = resultSum + fmt.Sprintln("Number of results ", len(results))

	return resultSum
}
