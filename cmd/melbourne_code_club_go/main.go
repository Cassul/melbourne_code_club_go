package main

import (
	"context"

	"github.com/spf13/cobra"
)

func main() {
	ctx := context.Background()
	Start(ctx)
}

// Start starts everything
func Start(ctx context.Context) {

	if err := rootCommand(ctx).Execute(); err != nil {
		if ctx.Done() == nil {
			panic(err)
		}
	}

}

func rootCommand(ctx context.Context) *cobra.Command {
	cmd := cobra.Command{
		Use:   "search_app",
		Short: "Search service",
	}

	err := cmd.Execute()
	if err != nil {
		if ctx.Done() == nil {
			panic(err)
		}
	}
	return &cmd
}
