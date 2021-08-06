package search_stuff

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/zendesk/melbourne_code_club_go/internal/search_stuff/types"
)

func SomeFunc(ctx context.Context) error {
	fmt.Printf("Hello World! Context - %+v", ctx)
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

	fmt.Printf("number of users - %d potato\n", len(users))
	return nil
}
