package main

import (
	"context"
	"fmt"
	"os"

	"github.com/zendesk/melbourne_code_club_go/internal/search_stuff"
	"github.com/zendesk/melbourne_code_club_go/internal/types"
	"github.com/zendesk/melbourne_code_club_go/internal/util"
)

func main() {
	ctx := context.Background()
	args := os.Args[1:]

	validate(args)
	query := types.Query{Dataset: args[0], Field: args[1], Value: args[2]}
	index := loadAndIndexData(ctx)
	searchData(index, query)
}

func search(tickets []types.Ticket, search_val string) []types.Ticket {
	results := []types.Ticket{}
	for _, ticket := range tickets {
		if ticket.Id == search_val {
			results = append(results, ticket)
		}
	}
	return results
}

func validate(args []string) {
	if len(args) != 3 {
		panic("number of arguments should equal 3")
	}
	dataType := args[0]
	fieldName := args[1]

	acceptedTypes := []string{"users", "tickets", "organizations"}

	if !util.ContainsString(acceptedTypes, dataType) {
		panic("wrong data type")
	}

	acceptedFields := types.DataTypes[dataType]
	if !util.ContainsString(acceptedFields, fieldName) {
		panic("wrong field name")
	}
}

func loadAndIndexData(ctx context.Context) types.Index {
	var records []types.Record

	users := search_stuff.LoadUsers(ctx)
	organizations := search_stuff.LoadOrganizations(ctx)
	tickets := search_stuff.LoadTickets(ctx)

	// Covariance
	// Contravariance

	for _, u := range users {
		records = append(records, (types.Record(u)))
	}

	for _, u := range organizations {
		records = append(records, (types.Record(u)))
	}

	for _, u := range tickets {
		records = append(records, (types.Record(u)))
	}

	index := map[types.Query]types.Record{}

	for _, record := range records {
		for _, query := range record.KeysForIndex() {
			index[query] = record
		}
	}

	return index
}

func searchData(index types.Index, query types.Query) {
	result := index[query]
	fmt.Println("Result: ", result)
}
