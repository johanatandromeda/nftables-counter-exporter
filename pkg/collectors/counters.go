package collectors

import (
	"fmt"
	"github.com/johanatandromeda/nftables-exporter/pkg/nftables"
	"github.com/prometheus/client_golang/prometheus"
	"log/slog"
	"time"
)

var nftCounterBytes = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "nftables",
	Subsystem: "counter",
	Name:      "bytes_total",
	Help:      "nftables named counter bytes"},
	[]string{"table", "name"})
var nftCounterPackages = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "nftables",
	Subsystem: "counter",
	Name:      "packages_total",
	Help:      "nftables named counter packages"},
	[]string{"table", "name"})

func InitNFtablesCounters() {

	slog.Debug("Init nftables counter metrics")

	prometheus.MustRegister(nftCounterBytes)
	prometheus.MustRegister(nftCounterPackages)

	for {
		counters, err := nftables.GetCounters()
		if err != nil {
			slog.Error(fmt.Sprintf("Could not fetch counters due to %s"), err)
			time.Sleep(15 * time.Second)
			continue
		}
		for _, c := range counters {
			nftCounterBytes.With(prometheus.Labels{"table": c.Table, "name": c.Name}).Set(float64(c.Bytes))
			nftCounterPackages.With(prometheus.Labels{"table": c.Table, "name": c.Name}).Set(float64(c.Packets))
		}
		time.Sleep(3 * time.Second)
	}

}
