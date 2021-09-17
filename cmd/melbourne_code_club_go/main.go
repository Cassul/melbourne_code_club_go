package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/manifoldco/promptui"

	"github.com/zendesk/melbourne_code_club_go/internal/search_stuff"
	"github.com/zendesk/melbourne_code_club_go/internal/types"
)

func main() {
	ctx := context.Background()

	query := promptUser()
	index := loadAndIndexData(ctx)
	searchData(index, query)
}

func promptUser() types.Query {
	datasetPrompt := promptui.Select{
		Label: "Select Data Type",
		Items: []string{"tickets", "organizations", "users"},
	}

	_, dataset, err := datasetPrompt.Run()

	acceptedFields := types.DataTypes[dataset]

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
		panic(fmt.Sprintf("Prompt failed, err - %v", err))
	}

	var value interface{}
	json.Unmarshal([]byte(inputValue), &value)

	return types.Query{Dataset: dataset, Field: field, Value: value}
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
	records := make(chan types.Record, 1)
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		fmt.Println("Going to load users")
		users := search_stuff.LoadUsers(ctx)

		for _, u := range users {
			fmt.Println("Loaded a user")
			records <- types.Record(u)
		}
		fmt.Println("Loaded all users")

		wg.Done()
	}()

	go func() {
		fmt.Println("Going to load organizations")
		organizations := search_stuff.LoadOrganizations(ctx)

		for _, u := range organizations {
			fmt.Println("Loaded an organization")
			records <- types.Record(u)
		}
		fmt.Println("Loaded all organizations")

		wg.Done()
	}()

	go func() {
		fmt.Println("Going to load tickets")
		tickets := search_stuff.LoadTickets(ctx)

		for _, u := range tickets {
			fmt.Println("Loaded a ticket")
			records <- types.Record(u)
		}
		fmt.Println("Loaded all tickets")

		wg.Done()
	}()

	index := map[types.Query][]types.Record{}
	go func() {
		wg.Wait()
		close(records)
	}()

	fmt.Println("Going to start looping through records")
	for record := range records {
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
	fmt.Println(index[query], query)
	result := index[query]
	fmt.Println("Result: ", result)
}
