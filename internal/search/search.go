package search

import (
	"fmt"

	"github.com/zendesk/melbourne_code_club_go/internal/types"
)

func SearchData(index types.Index, query types.Query) {
	results := index[query]
	// fmt.Println("Result: ", result)

	for _, result := range results {
		result.Print(index)
	}

	fmt.Println("Number of results", len(results))
}
