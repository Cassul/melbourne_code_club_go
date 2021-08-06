package search_stuff

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomeFunc(t *testing.T) {
	testCases := []struct {
		name string
		err  error
	}{
		{
			"some test case",
			nil,
		},
	}

	ctx := context.Background()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := someFunc(ctx)

			assert.Equal(t, tc.err, result)
		})
	}
}
