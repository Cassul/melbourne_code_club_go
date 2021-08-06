package search_stuff

import (
	"context"
	"fmt"
)

func someFunc(ctx context.Context) error {
	fmt.Printf("Hello World! Context - %+v", ctx)
	return nil
}
