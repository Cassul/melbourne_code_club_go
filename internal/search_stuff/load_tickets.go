package search_stuff

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/zendesk/melbourne_code_club_go/internal/types"
)

func LoadTickets(ctx context.Context) []types.Ticket {
	jsonFile, err := os.Open("data/tickets.json")
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var tickets []types.Ticket

	err = json.Unmarshal(byteValue, &tickets)
	if err != nil {
		panic(err)
	}

	return tickets
}
