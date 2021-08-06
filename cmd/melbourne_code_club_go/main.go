package main

import (
	"context"
	"fmt"
	"time"

	"github.com/zendesk/melbourne_code_club_go/internal/search_stuff"
)

func main() {
	ctx := context.Background()

	start := time.Now()

	done := make(chan bool, 3)

	go search_stuff.LoadUsers(ctx, done)
	go search_stuff.LoadOrganizations(ctx, done)
	go search_stuff.LoadTickets(ctx, done)

	<-done
	<-done
	<-done
	duration := time.Since(start)
	fmt.Println("goroutine time - ", duration)

	sequentialStart := time.Now()
	search_stuff.LoadUsers(ctx, done)
	search_stuff.LoadOrganizations(ctx, done)
	search_stuff.LoadTickets(ctx, done)
	sequentialDuration := time.Since(sequentialStart)
	fmt.Println("sequential time - ", sequentialDuration)

}
