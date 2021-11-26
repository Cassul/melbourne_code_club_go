package types

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
	Print(Index) string
	PrintBasicInfo() string
	KeysForIndex() []Query
}

type Index map[Query][]Record

func findOne(index Index, query Query) Record {
	if index[query] != nil {
		return index[query][0]
	}
	return nil
}
