# nftables-counter-exporter

This project emerged from a need to measure the nftables counters.
I did not find any general exporter that did not invoke the `nft` binary.
I wanted a more efficient exporter for my home firewall.

## Design

The `nftables-counter-exporter` only exports named counters. The motivation
for this decision is that I did not find any logical or consistent
way of labeling the metrics for counters without names.

## Usage

The port and listening address can be set with flags. To see all flags,
run:
```shell
./nftables-counter-exporter -h
```
The default listening address is `:9110`.

The tool only exposes named counters
https://wiki.nftables.org/wiki-nftables/index.php/Counters

## Metrics

There are two metrics exported:

| Metric name                     | Description            |
|---------------------------------|------------------------|
| nftables_counter_packages_total | The number of packages |
| nftables_counter_bytes_total    | The number of bytes    |

## Install

Copy the executable to `/usr/bin/` and add a service like this:

```text
[Unit]
Description=nftables-counter-exporter
After=network.target
StartLimitIntervalSec=0

[Service]
Type=simple
Restart=always
RestartSec=1
ExecStart=/usr/bin/nftables-counter-exporter

[Install]
WantedBy=multi-user.target
```

I would love for the distro maintainers to pick this up and add it to their
repositories (Nudge, Nudge, Wink, Wink, Say no more!).