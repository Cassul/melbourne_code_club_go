package search_stuff

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/zendesk/melbourne_code_club_go/internal/types"
)

func LoadTickets(ctx context.Context, done chan bool) error {
	jsonFile, err := os.Open("data/tickets.json")
	if err != nil {
		fmt.Printf("error reading file - %v", err)
		return err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var tickets []types.Ticket

	err = json.Unmarshal(byteValue, &tickets)
	if err != nil {
		panic(err)
	}

	fmt.Println("number of tickets - ", len(tickets))
	done <- true
	return nil
}
