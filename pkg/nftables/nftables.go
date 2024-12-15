package nftables

import (
	"fmt"
	"github.com/google/nftables"
	"github.com/google/nftables/expr"
	"log/slog"
)

func GetCounter() (map[string]uint64, error) {

	c, err := nftables.New()
	if err != nil {
		return nil, err
	}

	chains, err := c.ListChains()
	if err != nil {
		return nil, err
	}

	for _, chain := range chains {
		chainName := chain.Name
		slog.Debug(fmt.Sprintf("Processing chain %s", chainName))
		rules, err := c.GetRules(chain.Table, chain)
		if err != nil {
			return nil, err
		}
		for _, rule := range rules {
			for _, ex := range rule.Exprs {
				count, ok := ex.(*expr.Counter)
				if ok {
					slog.Debug(fmt.Sprintf("Got count with packages %v and bytes %v", count.Packets, count.Bytes))
				}
			}
		}
	}
	return nil, nil
}
