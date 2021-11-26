package types

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Ticket struct {
	Id             string   `json:"_id"`
	Url            string   `json:"url"`
	ExternalId     string   `json:"external_id"`
	CreatedAt      string   `json:"created_at"`
	Type           string   `json:"type"`
	Subject        string   `json:"subject"`
	Description    string   `json:"desciption"`
	Priority       string   `json:"priority"`
	Status         string   `json:"status"`
	SubmitterId    float64  `json:"submitter_id"`
	AssigneeId     float64  `json:"assignee_id"`
	OrganizationId float64  `json:"organization_id"`
	Tags           []string `json:"tags"`
	HasIncidents   bool     `json:"has_incidents"`
	DueAt          string   `json:"due_at"`
	Via            string   `json:"via"`
}

var TicketFields []string = []string{"_id", "url", "external_id", "created_at", "type", "subject", "desciption", "priority", "status", "submitter_id", "assignee_id", "organization_id", "tags", "has_incidents", "due_at", "via"}

func (t Ticket) KeysForIndex() []Query {
	query := []Query{
		{Dataset: "tickets", Field: "_id", Value: t.Id},
		{Dataset: "tickets", Field: "url", Value: t.Url},
		{Dataset: "tickets", Field: "external_id", Value: t.ExternalId},
		{Dataset: "tickets", Field: "created_at", Value: t.CreatedAt},
		{Dataset: "tickets", Field: "type", Value: t.Type},
		{Dataset: "tickets", Field: "subject", Value: t.Subject},
		{Dataset: "tickets", Field: "desciption", Value: t.Description},
		{Dataset: "tickets", Field: "priority", Value: t.Priority},
		{Dataset: "tickets", Field: "status", Value: t.Status},
		{Dataset: "tickets", Field: "submitter_id", Value: t.SubmitterId},
		{Dataset: "tickets", Field: "assignee_id", Value: t.AssigneeId},
		{Dataset: "tickets", Field: "organization_id", Value: t.OrganizationId},
		{Dataset: "tickets", Field: "has_incidents", Value: t.HasIncidents},
		{Dataset: "tickets", Field: "due_at", Value: t.DueAt},
		{Dataset: "tickets", Field: "via", Value: t.Via},
	}

	for _, tag := range t.Tags {
		query = append(query, Query{Dataset: "tickets", Field: "tags", Value: tag})
	}

	return query
}

func (t Ticket) Print(index Index) {
	// TODO: Potentially a bug. What if the associated doesn't exist?
	submitter := findOne(index, Query{Dataset: "users", Field: "_id", Value: t.SubmitterId})
	assignee := findOne(index, Query{Dataset: "users", Field: "_id", Value: t.AssigneeId})
	organization := findOne(index, Query{Dataset: "organizations", Field: "_id", Value: t.OrganizationId})

	fmt.Println("## Ticket.")
	t.PrintBasicInfo()
	t.printAssociatedRecords(submitter, assignee, organization)
}

func (t Ticket) PrintBasicInfo() {
	fmt.Println("          _id: ", t.Id)
	fmt.Println("          url: ", t.Url)
	fmt.Println("  external_id: ", t.ExternalId)
	fmt.Println("   created_at: ", t.CreatedAt)
	fmt.Println("         type: ", t.Type)
	fmt.Println("      subject: ", t.Subject)
	fmt.Println("   desciption: ", t.Description)
	fmt.Println("     priority: ", t.Priority)
	fmt.Println("       status: ", t.Status)
	fmt.Println("has_incidents: ", t.HasIncidents)
	fmt.Println("       due_at: ", t.DueAt)
	fmt.Println("          via: ", t.Via)
	fmt.Println("")
}

func (t Ticket) printAssociatedRecords(submitter Record, assignee Record, organization Record) {
	//sumitter
	fmt.Println("### Submitter.")
	submitter.PrintBasicInfo()

	//assignee
	if assignee != nil {
		fmt.Println("### Assignee.")
		assignee.PrintBasicInfo()
		fmt.Println("")
	}

	//organization
	if organization != nil {
		fmt.Println("### Organization.")
		organization.PrintBasicInfo()
		fmt.Println("")
	}
}

func LoadTickets(ctx context.Context) []Ticket {
	jsonFile, err := os.Open("data/tickets.json")
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var tickets []Ticket

	err = json.Unmarshal(byteValue, &tickets)
	if err != nil {
		panic(err)
	}

	return tickets
}
