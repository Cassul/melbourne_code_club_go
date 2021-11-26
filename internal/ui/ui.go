package ui

import (
	"encoding/json"

	"github.com/manifoldco/promptui"
	"github.com/zendesk/melbourne_code_club_go/internal/types"
	"github.com/zendesk/melbourne_code_club_go/internal/validation"
)

func PromptUser() (types.Query, error) {
	datasetPrompt := promptui.Select{
		Label: "Select Data Type",
		Items: []string{"tickets", "organizations", "users"},
	}

	_, dataset, err := datasetPrompt.Run()

	if err != nil {
		return types.Query{}, err
	}

	acceptedFields := types.DataTypes[dataset]

	fieldPrompt := promptui.Select{
		Label: "Select Field",
		Items: acceptedFields,
	}
	_, field, err := fieldPrompt.Run()

	if err != nil {
		return types.Query{}, err
	}

	inputValuePrompt := promptui.Prompt{
		Label:    "What are you searching for, dear User?",
		Validate: validation.SearchQuery,
	}

	inputValue, err := inputValuePrompt.Run()

	if err != nil {
		return types.Query{}, err
	}

	var value interface{}
	json.Unmarshal([]byte(inputValue), &value)

	return types.Query{Dataset: dataset, Field: field, Value: value}, nil
}
