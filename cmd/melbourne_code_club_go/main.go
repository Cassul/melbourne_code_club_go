package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/manifoldco/promptui"

	"github.com/zendesk/melbourne_code_club_go/internal/search_stuff"
	"github.com/zendesk/melbourne_code_club_go/internal/types"
)

func main() {
	ctx := context.Background()

	dataSetPrompt := promptui.Select{
		Label: "Select Data Type",
		Items: []string{"tickets", "organizations", "users"},
	}

	_, dataSet, err := dataSetPrompt.Run()

	acceptedFields := types.DataTypes[dataSet]

	fieldPrompt := promptui.Select{
		Label: "Select Field",
		Items: acceptedFields,
	}
	_, field, err := fieldPrompt.Run()

	inputValuePrompt := promptui.Prompt{
		Label:    "What are you searching for, dear User?",
		Validate: validateSearchQuery,
	}

	inputValue, err := inputValuePrompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	var value interface{}
	json.Unmarshal([]byte(inputValue), &value)

	query := types.Query{Dataset: dataSet, Field: field, Value: value}

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

func validateSearchQuery(searchQuery string) error {
	if !json.Valid([]byte(searchQuery)) {
		return fmt.Errorf("Invalid search query, must be json")
	}
	return nil
}

func loadAndIndexData(ctx context.Context) map[types.Query][]types.Record {
	var records []types.Record

	users := search_stuff.LoadUsers(ctx)
	organizations := search_stuff.LoadOrganizations(ctx)
	tickets := search_stuff.LoadTickets(ctx)

	for _, u := range users {
		records = append(records, (types.Record(u)))
	}

	for _, u := range organizations {
		records = append(records, (types.Record(u)))
	}

	for _, u := range tickets {
		records = append(records, (types.Record(u)))
	}

	index := map[types.Query][]types.Record{}

	for _, record := range records {
		for _, query := range record.KeysForIndex() {
			if len(index[query]) > 0 {
				index[query] = append(index[query], record)
			} else {
				index[query] = []types.Record{record}
			}
		}
	}

	return index
}

func searchData(index map[types.Query][]types.Record, query types.Query) {
	result := index[query]
	fmt.Println("Result: ", result)
}
