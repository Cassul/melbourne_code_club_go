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
- Get help flags working
- Go through challenge
- Create types folder to represent those data structures and use them for parsing
- performance
  - A list gives us O(n). N is the number of items in the list.
  - Index by id {10: {name: "Mads", role: "admin"}}
    - O(1)
- Maybe: Try indexing using goroutines. Benchmark.
- Focus on
  - Design patterns. Reading Head first design patterns book.
    - Factory pattern
  - Polymorphism
  - Object oriented ideas
  - Goroutines. Advaned stuff that is too hard to maintain for normal code.
    - Channels
    - Error groups?


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