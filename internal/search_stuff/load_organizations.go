package search_stuff

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/zendesk/melbourne_code_club_go/internal/types"
)

func LoadOrganizations(ctx context.Context, done chan bool) error {
	jsonFile, err := os.Open("data/organizations.json")
	if err != nil {
		fmt.Printf("error reading file - %v", err)
		return err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var organizations []types.Organization

	err = json.Unmarshal(byteValue, &organizations)
	if err != nil {
		panic(err)
	}

	fmt.Println("number of organizations - ", len(organizations))
	done <- true
	return nil
}