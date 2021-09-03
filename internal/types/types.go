package types

type User struct {
	Id             int      `json:"_id"`
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
	OrganizationId int      `json:"organization_id"`
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
	SubmitterId    int      `json:"submitter_id"`
	AssigneeId     int      `json:"assignee_id"`
	OrganizationId int      `json:"organization_id"`
	Tags           []string `json:"tags"`
	HasIncidents   bool     `json:"has_incidents"`
	DueAt          string   `json:"due_at"`
	Via            string   `json:"via"`
}

var TicketFields []string = []string{"_id", "url", "external_id", "created_at", "type", "subject", "desciption", "priority", "status", "submitter_id", "assignee_id", "organization_id", "tags", "has_incidents", "due_at", "via"}

type Organization struct {
	Id            int      `json:"_id"`
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
	KeysForIndex() []Query
}

type Index map[Query]Record

func (t Ticket) KeysForIndex() []Query {
	return []Query{
		{Dataset: "tickets", Field: "_id", Value: t.Id},
	}
}

func (u User) KeysForIndex() []Query {
	return []Query{
		{Dataset: "users", Field: "_id", Value: u.Id},
	}
}

func (o Organization) KeysForIndex() []Query {
	return []Query{
		{Dataset: "organizations", Field: "_id", Value: o.Id},
	}
}
