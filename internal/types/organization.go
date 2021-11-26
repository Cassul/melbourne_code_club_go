package types

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

type Organization struct {
	Id            float64  `json:"_id"`
	Url           string   `json:"url"`
	ExternalId    string   `json:"external_id"`
	DomainNames   []string `json:"domain_names"`
	Name          string   `json:"name"`
	CreatedAt     string   `json:"created_at"`
	SharedTickets bool     `json:"shared_tickets"`
	Tags          []string `json:"tags"`
	Details       string   `json:"details"`
}

var OrganizationFields []string = []string{"_id", "url", "external_id", "domain_names", "name", "created_at", "shared_tickets", "tags", "details"}

func (o Organization) KeysForIndex() []Query {
	query := []Query{
		{Dataset: "organizations", Field: "_id", Value: o.Id},
		{Dataset: "organizations", Field: "url", Value: o.Url},
		{Dataset: "organizations", Field: "external_id", Value: o.ExternalId},
		{Dataset: "organizations", Field: "name", Value: o.Name},
		{Dataset: "organizations", Field: "created_at", Value: o.CreatedAt},
		{Dataset: "organizations", Field: "shared_tickets", Value: o.SharedTickets},
		{Dataset: "organizations", Field: "details", Value: o.Details},
	}

	for _, tag := range o.Tags {
		query = append(query, Query{Dataset: "organizations", Field: "tags", Value: tag})
	}

	for _, tag := range o.DomainNames {
		query = append(query, Query{Dataset: "organizations", Field: "domain_names", Value: tag})
	}

	return query
}

func (o Organization) Print(index Index) {
	fmt.Println("## Organization.")
	fmt.Println(o.PrintBasicInfo())
}

func (o Organization) PrintBasicInfo() string {

	var buf bytes.Buffer

	templateBody :=
		`		      _id: {{.Id}}
	      external_id: {{.ExternalId}}
  	   domain_names: {{.DomainNames}}
	             name: {{.Name}}
	       created_at: {{.CreatedAt}}
	   shared_tickets: {{.SharedTickets}}
	             tags: {{.Tags}}
	          details: {{.Details}}`

	tmpl, err := template.New("test").Parse(templateBody)

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(&buf, o)
	if err != nil {
		panic(err)
	}

	return buf.String()
}

func LoadOrganizations(ctx context.Context) []Organization {
	jsonFile, err := os.Open("data/organizations.json")
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var organizations []Organization

	err = json.Unmarshal(byteValue, &organizations)
	if err != nil {
		panic(err)
	}

	return organizations
}
