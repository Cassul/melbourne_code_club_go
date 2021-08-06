package main

import (
	"context"

	"github.com/zendesk/melbourne_code_club_go/internal/search_stuff"
)

func main() {
	ctx := context.Background()

	done := make(chan bool, 3)

	go search_stuff.LoadUsers(ctx, done)
	go search_stuff.LoadOrganizations(ctx, done)
	go search_stuff.LoadTickets(ctx, done)

	<-done
	<-done
	<-done
}
