package main

import (
	"context"

	"github.com/zendesk/melbourne_code_club_go/internal/search_stuff"
)

func main() {
	ctx := context.Background()

	_ = search_stuff.SomeFunc(ctx)
}
