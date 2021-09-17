> ./coding_challenge list_fields
Users
- _id
_ name

Organizations
- _id
- url


--------------------------------------------

> ./coding_challenge
1. List fields
2. Perform search
3. Exit
Which would you like?
> 1
WHich dataset (users/organizations/tickets)?
> users


## TODO
x Index the data
  x Define the Query struct
  x Define a record interface
  x Have the records return their own query objects
  x Build the index
  x Update search code to use index
- Try a parallel linear search
- Expand search to other fields
x Expand search to other datasets

- Duplicate results bug

- Looping search UI
- Format the results
- Handle related data

x Print error if dataset is not correct
x Allow for search of tickets by ID

- performance
  - A list gives us O(n). N is the number of items in the list.
  - Index by id {10: {name: "Mads", role: "admin"}}
    - O(1)
- Focus on
  - Design patterns. Reading Head first design patterns book.
    - Factory pattern
  - Polymorphism
  - Object oriented ideas
  - Goroutines. Advaned stuff that is too hard to maintain for normal code.
    - Channels
    - Error groups?
- Create types folder to represent those data structures and use them for parsing
- Go through challenge
x Print error if not enough arguments given
x Maybe: Try indexing using goroutines. Benchmark.
x Print the arguments to the program
- Get help flags working
- Listen for signal interupted and kill and tidily close out goroutines
- Implement the loading using goroutines

## Interface
./search users _id 1

dataset := os.Args[1:] // "users"
field := os.Args[2] // "_id"


## Goal
- Print how many users we have
- Repeat for tickets and organizations

## Tidy up items
- Config folder with paths to json files



## Performance
- CPU bound tasks often don't parallelise that well
- Users don't care about the difference between 3ms and 200ms
- Senior developers will be expected to measure and talk about tradeoffs and scaling
- 




{
  1: 1,
  2: 1,
  3: 2,
  4: 3
}


function fib(n) {
  ....
}

fib[28]
fib(28)


validUserField? = {
  "_id": true,
  "url": true,
  ...
}

function validUserField(fieldName) {
  fields = ["_id", "url", ....]
  return contains(fields, fieldName)
}

fields = {
  users: ["_id", "url", ....],
  organizations: ["_id", .....]
}


abcDEF
abCDef

database = {
  "users": [....],
  "tickets": [....]
}


dataset []types.Ticket
dataset []types.User
dataset []types.Organization





validating the user input
loading the data
performing the search



# Indexing

lookup_in_index(index, query)

Inputs

index = 

map[string]map[string]map[string][]Record
    dataset    field      value

{"user.id.1": [User()]}

"user.name"

"user.user.name.1"

map[[]string]Record

{
  ["user", "name", "Logan Campbell"]:
  [User(...)]
}

map[Query]Record

{
  Query(dataset: "user", field: "name", value: "Logan Campbell"):
  [User(...)]
}


{ Users
    "id": 
      1: [Users(....)],
      ...id: [Users(...)]
    "name":
        "Logan Campbell": [Users(....)], // Name
        "Logan Campbell": [Users(....)], // Fathers name
  Organizations

}

{1: [Organization(...)]}

{1: [Ticket(...)]}

query = {dataType: "users",
fieldName: "_id",
queryValue: 1}

query = {dataType: "users",
fieldName: "name",
queryValue: "Logan Campbell"}


query = {dataType: "organizations",
fieldName: "_id",
queryValue: 1}


Output
[] Ticket
