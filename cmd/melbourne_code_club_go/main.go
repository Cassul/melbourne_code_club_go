package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	indexpkg "github.com/zendesk/melbourne_code_club_go/internal/index"
	"github.com/zendesk/melbourne_code_club_go/internal/search"
	"github.com/zendesk/melbourne_code_club_go/internal/types"
	"github.com/zendesk/melbourne_code_club_go/internal/ui"
)

func main() {
	ctx := context.Background()
	indexChannel := make(chan types.Index)
	var index types.Index
	var syncOnce sync.Once
	enableGracefulShutdown()

	// Do this in the background
	go func() {
		indexChannel <- indexpkg.LoadAndIndexData(ctx)
	}()

	// Loop these two
	for {
		query, err := ui.PromptUser()

		if err != nil {
			fmt.Println("Goodbye")
			return
		}

		syncOnce.Do(
			func() {
				index = <-indexChannel
			})
		fmt.Println(search.SearchData(index, query))
	}
}

func enableGracefulShutdown() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println("got a signal - ", sig)
	}()
	fmt.Println("awaiting signal")
}
