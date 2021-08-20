package search_stuff

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/zendesk/melbourne_code_club_go/internal/types"
)

func LoadUsers(ctx context.Context) []types.User {
	jsonFile, err := os.Open("data/users.json")
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users []types.User

	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		panic(err)
	}

	fmt.Println("number of users - ", len(users))

	return users
}
