package types

import "fmt"

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

var DataTypes map[string][]string = map[string][]string{
	"users":         UserFields,
	"organizations": OrganizationFields,
	"tickets":       TicketFields,
}

type Database struct {
	Users         []User
	Tickets       []Ticket
	Organizations []Organization
}

type Query struct {
	Dataset string
	Field   string
	Value   interface{}
}

type Record interface {
	Print(Index)
	PrintBasicInfo()
	KeysForIndex() []Query
}

type Index map[Query][]Record

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
	t.PrintAssociatedRecords(submitter, assignee, organization)
}

func findOne(index Index, query Query) Record {
	if index[query] != nil {
		return index[query][0]
	}
	return nil
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

func (t Ticket) PrintAssociatedRecords(submitter Record, assignee Record, organization Record) {
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

func (u User) Print(index Index) {
	organization := findOne(index, Query{Dataset: "organizations", Field: "_id", Value: u.OrganizationId})

	fmt.Println("## User.")
	u.PrintBasicInfo()
	u.PrintAssociatedRecords(organization)
}

func (u User) PrintAssociatedRecords(organization Record) {
	//organization
	if organization != nil {
		fmt.Println("### Organization.")
		organization.PrintBasicInfo()
		fmt.Println("")
	}
}

func (u User) PrintBasicInfo() {
	fmt.Println("           _id: ", u.Id)
	fmt.Println("           url: ", u.Url)
	fmt.Println("   external_id: ", u.ExternalId)
	fmt.Println("    created_at: ", u.CreatedAt)
	fmt.Println("          type: ", u.Name)
	fmt.Println("       subject: ", u.Alias)
	fmt.Println("    desciption: ", u.Active)
	fmt.Println("      priority: ", u.Verified)
	fmt.Println("        status: ", u.Shared)
	fmt.Println(" has_incidents: ", u.Locale)
	fmt.Println("        due_at: ", u.Timezone)
	fmt.Println("         email: ", u.Email)
	fmt.Println(" last_login_at: ", u.LastLoginAt)
	fmt.Println("         phone: ", u.Phone)
	fmt.Println("     signature: ", u.Signature)
	fmt.Println("          tags: ", u.Tags)
	fmt.Println("     suspended: ", u.Suspended)
	fmt.Println("          role: ", u.Role)
	fmt.Println("")
}

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
	o.PrintBasicInfo()
}

func (o Organization) PrintBasicInfo() {
	fmt.Println("            _id: ", o.Id)
	fmt.Println("    external_id: ", o.ExternalId)
	fmt.Println("   domain_names: ", o.DomainNames)
	fmt.Println("           name: ", o.Name)
	fmt.Println("     created_at: ", o.CreatedAt)
	fmt.Println(" shared_tickets: ", o.SharedTickets)
	fmt.Println("           tags: ", o.Tags)
	fmt.Println("        details: ", o.Details)
}
