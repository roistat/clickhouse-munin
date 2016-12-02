package main

// MetricConfig is one line on graph
type MetricConfig struct {
	id              string
	label           string
	metricType      string
	drawType        string
	color           string
	clickHouseEvent string
}

// Graph itself
type Graph struct {
	graphTitle    string
	graphCategory string
	graphInfo     string
	graphLabel    string
	graphPeriod   string
	graphArgs     string
	isPercent     bool
	metrics       []MetricConfig
}

// AvailableGraphs is list of available graphs
var AvailableGraphs = map[string]Graph{
	"queries": {
		graphTitle:    "ClickHouse queries",
		graphCategory: "clickhouse",
		graphInfo:     "Values received from ClickHouse system.events table",
		graphLabel:    "queries / minute",
		graphPeriod:   "minute",
		graphArgs:     "--lower-limit 0",
		metrics: []MetricConfig{
			{
				id:              "select",
				label:           "Selects",
				metricType:      "DERIVE",
				drawType:        "AREA",
				color:           "COLOUR0",
				clickHouseEvent: "SelectQuery",
			},
			{
				id:              "insert",
				label:           "Inserts",
				metricType:      "DERIVE",
				drawType:        "STACK",
				color:           "COLOUR1",
				clickHouseEvent: "InsertQuery",
			},
		},
	},
	"files": {
		graphTitle:    "ClickHouse files",
		graphCategory: "clickhouse",
		graphInfo:     "Values received from ClickHouse system.events table",
		graphLabel:    "operations / second",
		graphPeriod:   "second",
		graphArgs:     "--lower-limit 0",
		metrics: []MetricConfig{
			{
				id:              "file_open",
				label:           "Opens",
				metricType:      "DERIVE",
				drawType:        "AREA",
				color:           "COLOUR0",
				clickHouseEvent: "FileOpen",
			},
			{
				id:              "seek",
				label:           "Seeks",
				metricType:      "DERIVE",
				drawType:        "STACK",
				color:           "COLOUR1",
				clickHouseEvent: "Seek",
			},
		},
	},
	"cache": {
		graphTitle:    "ClickHouse cache",
		graphCategory: "clickhouse",
		graphInfo:     "Values received from ClickHouse system.events table",
		graphLabel:    "per second",
		graphPeriod:   "second",
		graphArgs:     "-u 100 -l 0 -r --base 1000",
		isPercent:     true,
		metrics: []MetricConfig{
			{
				id:              "hits",
				label:           "hit ratio",
				metricType:      "GAUGE",
				drawType:        "AREA",
				color:           "COLOUR0",
				clickHouseEvent: "MarkCacheHits",
			},
			{
				id:              "misses",
				label:           "miss ratio",
				metricType:      "GAUGE",
				drawType:        "STACK",
				color:           "COLOUR1",
				clickHouseEvent: "MarkCacheMisses",
			},
		},
	},
	"zookeeper": {
		graphTitle:    "ClickHouse Zookeeper transactions",
		graphCategory: "clickhouse",
		graphInfo:     "Values received from ClickHouse system.events table",
		graphLabel:    "operations / minute",
		graphPeriod:   "minute",
		graphArgs:     "--lower-limit 0",
		metrics: []MetricConfig{
			{
				id:              "ZooKeeperGetChildren",
				label:           "ZooKeeperGetChildren",
				metricType:      "DERIVE",
				drawType:        "AREA",
				color:           "COLOUR0",
				clickHouseEvent: "ZooKeeperGetChildren",
			},
			{
				id:              "ZooKeeperCreate",
				label:           "ZooKeeperCreate",
				metricType:      "DERIVE",
				drawType:        "STACK",
				color:           "COLOUR1",
				clickHouseEvent: "ZooKeeperCreate",
			},
			{
				id:              "ZooKeeperRemove",
				label:           "ZooKeeperRemove",
				metricType:      "DERIVE",
				drawType:        "STACK",
				color:           "COLOUR2",
				clickHouseEvent: "ZooKeeperRemove",
			},
			{
				id:              "ZooKeeperExists",
				label:           "ZooKeeperExists",
				metricType:      "DERIVE",
				drawType:        "STACK",
				color:           "COLOUR3",
				clickHouseEvent: "ZooKeeperExists",
			},
			{
				id:              "ZooKeeperGet",
				label:           "ZooKeeperGet",
				metricType:      "DERIVE",
				drawType:        "STACK",
				color:           "COLOUR4",
				clickHouseEvent: "ZooKeeperGet",
			},
			{
				id:              "ZooKeeperSet",
				label:           "ZooKeeperSet",
				metricType:      "DERIVE",
				drawType:        "STACK",
				color:           "COLOUR5",
				clickHouseEvent: "ZooKeeperSet",
			},
			{
				id:              "ZooKeeperMulti",
				label:           "ZooKeeperMulti",
				metricType:      "DERIVE",
				drawType:        "STACK",
				color:           "COLOUR6",
				clickHouseEvent: "ZooKeeperMulti",
			},
			{
				id:              "ZooKeeperExceptions",
				label:           "ZooKeeperExceptions",
				metricType:      "DERIVE",
				drawType:        "STACK",
				color:           "COLOUR7",
				clickHouseEvent: "ZooKeeperExceptions",
			},
		},
	},
}
