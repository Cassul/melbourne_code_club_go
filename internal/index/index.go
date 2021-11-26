package index

import (
	"context"
	"sync"

	"github.com/zendesk/melbourne_code_club_go/internal/types"
)

func LoadAndIndexData(ctx context.Context) types.Index {
	records := make(chan types.Record, 1)
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		users := types.LoadUsers(ctx)

		for _, u := range users {
			records <- types.Record(u)
		}

		wg.Done()
	}()

	go func() {
		organizations := types.LoadOrganizations(ctx)

		for _, u := range organizations {
			records <- types.Record(u)
		}

		wg.Done()
	}()

	go func() {
		tickets := types.LoadTickets(ctx)

		for _, u := range tickets {
			records <- types.Record(u)
		}

		wg.Done()
	}()

	index := types.Index{}
	go func() {
		wg.Wait()
		close(records)
	}()

	for record := range records {
		for _, query := range record.KeysForIndex() {
			if len(index[query]) > 0 {
				index[query] = append(index[query], record)
			} else {
				index[query] = []types.Record{record}
			}
		}
	}

	return index
}
