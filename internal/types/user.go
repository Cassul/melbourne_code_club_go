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

type User struct {
	Id             float64  `json:"_id"`
	Url            string   `json:"url"`
	ExternalId     string   `json:"external_id"`
	Name           string   `json:"name"`
	Alias          string   `json:"alias"`
	CreatedAt      string   `json:"created_at"`
	Active         bool     `json:"active"`
	Verified       bool     `json:"verified"`
	Shared         bool     `json:"shared"`
	Locale         string   `json:"locale"`
	Timezone       string   `json:"timezone"`
	LastLoginAt    string   `json:"last_login_at"`
	Email          string   `json:"email"`
	Phone          string   `json:"phone"`
	Signature      string   `json:"signature"`
	OrganizationId float64  `json:"organization_id"`
	Tags           []string `json:"tags"`
	Suspended      bool     `json:"suspended"`
	Role           string   `json:"role"`
}

var UserFields []string = []string{"_id", "url", "external_id", "name", "alias", "created_at", "active", "verified", "shared", "locale", "timezone", "last_login_at", "email", "phone", "signature", "organization_id", "tags", "suspended", "role"}

func (u User) KeysForIndex() []Query {
	query := []Query{
		{Dataset: "users", Field: "_id", Value: u.Id},
		{Dataset: "users", Field: "url", Value: u.Url},
		{Dataset: "users", Field: "external_id", Value: u.ExternalId},
		{Dataset: "users", Field: "created_at", Value: u.CreatedAt},
		{Dataset: "users", Field: "name", Value: u.Name},
		{Dataset: "users", Field: "alias", Value: u.Alias},
		{Dataset: "users", Field: "active", Value: u.Active},
		{Dataset: "users", Field: "verified", Value: u.Verified},
		{Dataset: "users", Field: "shared", Value: u.Shared},
		{Dataset: "users", Field: "locale", Value: u.Locale},
		{Dataset: "users", Field: "timezone", Value: u.Timezone},
		{Dataset: "users", Field: "last_login_at", Value: u.LastLoginAt},
		{Dataset: "users", Field: "email", Value: u.Email},
		{Dataset: "users", Field: "phone", Value: u.Phone},
		{Dataset: "users", Field: "signature", Value: u.Signature},
		{Dataset: "users", Field: "suspended", Value: u.Suspended},
		{Dataset: "users", Field: "role", Value: u.Role},
	}

	for _, tag := range u.Tags {
		query = append(query, Query{Dataset: "users", Field: "tags", Value: tag})
	}

	return query
}

func (u User) Print(index Index) string {
	organization := findOne(index, Query{Dataset: "organizations", Field: "_id", Value: u.OrganizationId})

	return fmt.Sprintf("## User.\n%s\n%s", u.PrintBasicInfo(), u.PrintAssociatedRecords(organization))
}

func (u User) PrintAssociatedRecords(organization Record) string {
	//organization
	if organization != nil {
		return fmt.Sprintf("### Organization.\n%s", organization.PrintBasicInfo())
	}

	return ""
}

func (u User) PrintBasicInfo() string {
	var buf bytes.Buffer

	templateBody :=
		`           _id: {{.Id}}
	           url: {{.Url}}
	   external_id: {{.ExternalId}}
	    created_at: {{.CreatedAt}}
	          type: {{.Name}}
	       subject: {{.Alias}}
	    desciption: {{.Active}}
	      priority: {{.Verified}}
	        status: {{.Shared}}
	 has_incidents: {{.Locale}}
	        due_at: {{.Timezone}}
	         email: {{.Email}}
	 last_login_at: {{.LastLoginAt}}
	         phone: {{.Phone}}
	     signature: {{.Signature}}
	          tags: {{.Tags}}
	     suspended: {{.Suspended}}
	          role: {{.Role}}`
	tmpl, err := template.New("test").Parse(templateBody)

	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(&buf, u)
	if err != nil {
		panic(err)
	}

	return buf.String()
}

func LoadUsers(ctx context.Context) []User {
	jsonFile, err := os.Open("data/users.json")
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users []User

	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		panic(err)
	}

	return users
}
