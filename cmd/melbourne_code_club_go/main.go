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
	db := loadData(ctx)
	searchData(db, args)
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

func loadData(ctx context.Context) types.Database {
	return types.Database{
		Users:         search_stuff.LoadUsers(ctx),
		Organizations: search_stuff.LoadOrganizations(ctx),
		Tickets:       search_stuff.LoadTickets(ctx),
	}
}

func searchData(db types.Database, args []string) {
	dataType := args[0]
	// fieldName := args[1]
	query_value := args[2]

	switch dataType {
	case "tickets":
		fmt.Println("results - ", search(db.Tickets, query_value))
	default:
		panic("Somehow we got here with the wrong datatype")
	}
}
