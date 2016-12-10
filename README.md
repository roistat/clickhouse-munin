# clickhouse-munin
# [![Go Report](https://goreportcard.com/badge/github.com/roistat/clickhouse-munin)](https://goreportcard.com/report/github.com/roistat/clickhouse-munin) ![](https://img.shields.io/github/license/roistat/clickhouse-munin.svg)

Munin plugin for ClickHouse

## Graphs examples

#### Queries

```bash
cp bin/clickhouse_ /etc/munin/plugins/clickhouse_queries
```

![](https://raw.githubusercontent.com/roistat/clickhouse-munin/master/examples/queries.png)

#### Cache hit/miss

```bash
cp bin/clickhouse_ /etc/munin/plugins/clickhouse_cache
```

![](https://raw.githubusercontent.com/roistat/clickhouse-munin/master/examples/cache.png)

#### Files operations

```bash
cp bin/clickhouse_ /etc/munin/plugins/clickhouse_files
```

![](https://raw.githubusercontent.com/roistat/clickhouse-munin/master/examples/files.png)

#### Zookeeper operations

```bash
cp bin/clickhouse_ /etc/munin/plugins/clickhouse_zookeeper
```

![](https://raw.githubusercontent.com/roistat/clickhouse-munin/master/examples/zookeeper.png)

#### Make your own graph

You could add you own graphs in config.go (AvailableGraphs variable)

## Build instructions

```bash
glide install
go build -o bin/clickhouse_
```
