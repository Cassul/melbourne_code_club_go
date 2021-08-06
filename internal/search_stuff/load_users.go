package search_stuff

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/zendesk/melbourne_code_club_go/internal/types"
)

func LoadUsers(ctx context.Context, done chan bool) error {
	jsonFile, err := os.Open("data/users.json")
	if err != nil {
		fmt.Printf("error reading file - %v", err)
		return err
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users []types.User

	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		panic(err)
	}

	fmt.Println("number of users - ", len(users))

	done <- true
	return nil
}
