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

	_ = types.Database{
		Users:         search_stuff.LoadUsers(ctx),
		Organizations: search_stuff.LoadOrganizations(ctx),
		Tickets:       search_stuff.LoadTickets(ctx),
	}

	args := os.Args[1:]
	if len(args) != 3 {
		panic("number of arguments should equal 3")
	}
	dataType := args[0]
	field_name := args[1]
	// query_value := args[2]

	acceptedTypes := []string{"users", "tickets", "organizations"}

	if !util.ContainsString(acceptedTypes, dataType) {
		panic("wrong data type")
	}

	acceptedFields := types.DataTypes[dataType]
	if !util.ContainsString(acceptedFields, field_name) {
		panic("wrong field name")
	}

	fmt.Println("args - ", args)
}
