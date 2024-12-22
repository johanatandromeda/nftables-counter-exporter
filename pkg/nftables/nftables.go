package nftables

import (
	"fmt"
	"github.com/google/nftables"
	"log/slog"
)

func GetCounter() (map[string]uint64, error) {

	c, err := nftables.New()
	if err != nil {
		return nil, err
	}

	tables, err := c.ListTables()
	if err != nil {
		return nil, err
	}
	for _, table := range tables {
		slog.Debug(fmt.Sprintf("Processing table %s", table.Name))
		objs, err := c.GetObjects(table)
		if err != nil {
			slog.Debug(fmt.Sprintf("Error getting objects for table %s due to %s", table.Name, err))
			continue
		}
		for _, obj := range objs {
			count, ok := obj.(*nftables.CounterObj)
			if ok {
				slog.Debug(fmt.Sprintf("Got count %s with packages %v and bytes %v", count.Name, count.Packets, count.Bytes))
			}
		}
	}
	return nil, nil
}
