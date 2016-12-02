# clickhouse-munin
# [![Go Report](https://goreportcard.com/badge/github.com/roistat/go-clickhouse)](https://goreportcard.com/report/github.com/roistat/go-clickhouse) ![](https://img.shields.io/github/license/roistat/clickhouse-munin.svg)

Munin plugin for ClickHouse

## Graphs examples

#### Queries

```bash
cp bin/clickhouse_ /etc/munin/plugins/clickhouse_queries
```

#### Cache hit/miss

```bash
cp bin/clickhouse_ /etc/munin/plugins/clickhouse_cache
```

#### Files operations

```bash
cp bin/clickhouse_ /etc/munin/plugins/clickhouse_cache
```
#### Zookeeper operations

```bash
cp bin/clickhouse_ /etc/munin/plugins/clickhouse_zookeeper
```

## Build instructions

```bash
glide install
go build -o bin/clickhouse_
```