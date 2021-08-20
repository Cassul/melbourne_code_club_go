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
- Allow for search of tickets by ID
x Print error if dataset is not correct

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

