package nftables

import (
	"fmt"
	"github.com/google/nftables"
	"log/slog"
)

type Counter struct {
	Name    string
	Packets uint64
	Bytes   uint64
	Table   string
}

func GetCounters() ([]Counter, error) {

	c, err := nftables.New(nftables.AsLasting())
	if err != nil {
		return nil, err
	}
	defer func(c *nftables.Conn) {
		_ = c.CloseLasting()
	}(c)

	tables, err := c.ListTables()
	if err != nil {
		return nil, err
	}

	counters := make([]Counter, 0)

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
				counters = append(counters, Counter{
					Name:    count.Name,
					Packets: count.Packets,
					Bytes:   count.Bytes,
					Table:   table.Name,
				})
			}
		}
	}
	return counters, nil
}
