package main

// MetricConfig is one line on graph
type MetricConfig struct {
	id               string
	label            string
	metricType       string
	drawType         string
	color            string
	clickHouseMetric string
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
	"memory": {
		graphTitle:    "ClickHouse memory",
		graphCategory: "clickhouse",
		graphInfo:     "Values received from ClickHouse system.events table",
		graphLabel:    "memory",
		graphPeriod:   "minute",
		graphArgs:     "--lower-limit 0",
		metrics: []MetricConfig{
			{
				id:               "memorytracking",
				label:            "MemoryTracking",
				metricType:       "DERIVE",
				drawType:         "AREA",
				color:            "COLOUR0",
				clickHouseMetric: "MemoryTracking",
			},
			{
				id:               "memorytrackingmerges",
				label:            "MemoryTrackingForMerges",
				metricType:       "DERIVE",
				drawType:         "AREA",
				color:            "COLOUR1",
				clickHouseMetric: "MemoryTrackingForMerges",
			},
		},
	},
	"merges": {
		graphTitle:    "ClickHouse merges",
		graphCategory: "clickhouse",
		graphInfo:     "Values received from ClickHouse system.events table",
		graphLabel:    "merges",
		graphPeriod:   "minute",
		graphArgs:     "--lower-limit 0",
		metrics: []MetricConfig{
			{
				id:               "mergedrows",
				label:            "MergedRows",
				metricType:       "DERIVE",
				drawType:         "AREA",
				color:            "COLOUR0",
				clickHouseMetric: "MergedRows",
			},
			{
				id:               "mergestimems",
				label:            "MergesTimeMilliseconds",
				metricType:       "DERIVE",
				drawType:         "STACK",
				color:            "COLOUR1",
				clickHouseMetric: "MergesTimeMilliseconds",
			},
		},
	},
	"queries": {
		graphTitle:    "ClickHouse queries",
		graphCategory: "clickhouse",
		graphInfo:     "Values received from ClickHouse system.events table",
		graphLabel:    "queries / minute",
		graphPeriod:   "minute",
		graphArgs:     "--lower-limit 0",
		metrics: []MetricConfig{
			{
				id:               "select",
				label:            "Selects",
				metricType:       "DERIVE",
				drawType:         "AREA",
				color:            "COLOUR0",
				clickHouseMetric: "SelectQuery",
			},
			{
				id:               "insert",
				label:            "Inserts",
				metricType:       "DERIVE",
				drawType:         "STACK",
				color:            "COLOUR1",
				clickHouseMetric: "InsertQuery",
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
				id:               "file_open",
				label:            "Opens",
				metricType:       "DERIVE",
				drawType:         "AREA",
				color:            "COLOUR0",
				clickHouseMetric: "FileOpen",
			},
			{
				id:               "seek",
				label:            "Seeks",
				metricType:       "DERIVE",
				drawType:         "STACK",
				color:            "COLOUR1",
				clickHouseMetric: "Seek",
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
				id:               "hits",
				label:            "hit ratio",
				metricType:       "GAUGE",
				drawType:         "AREA",
				color:            "COLOUR0",
				clickHouseMetric: "MarkCacheHits",
			},
			{
				id:               "misses",
				label:            "miss ratio",
				metricType:       "GAUGE",
				drawType:         "STACK",
				color:            "COLOUR1",
				clickHouseMetric: "MarkCacheMisses",
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
				id:               "ZooKeeperGetChildren",
				label:            "ZooKeeperGetChildren",
				metricType:       "DERIVE",
				drawType:         "AREA",
				color:            "COLOUR0",
				clickHouseMetric: "ZooKeeperGetChildren",
			},
			{
				id:               "ZooKeeperCreate",
				label:            "ZooKeeperCreate",
				metricType:       "DERIVE",
				drawType:         "STACK",
				color:            "COLOUR1",
				clickHouseMetric: "ZooKeeperCreate",
			},
			{
				id:               "ZooKeeperRemove",
				label:            "ZooKeeperRemove",
				metricType:       "DERIVE",
				drawType:         "STACK",
				color:            "COLOUR2",
				clickHouseMetric: "ZooKeeperRemove",
			},
			{
				id:               "ZooKeeperExists",
				label:            "ZooKeeperExists",
				metricType:       "DERIVE",
				drawType:         "STACK",
				color:            "COLOUR3",
				clickHouseMetric: "ZooKeeperExists",
			},
			{
				id:               "ZooKeeperGet",
				label:            "ZooKeeperGet",
				metricType:       "DERIVE",
				drawType:         "STACK",
				color:            "COLOUR4",
				clickHouseMetric: "ZooKeeperGet",
			},
			{
				id:               "ZooKeeperSet",
				label:            "ZooKeeperSet",
				metricType:       "DERIVE",
				drawType:         "STACK",
				color:            "COLOUR5",
				clickHouseMetric: "ZooKeeperSet",
			},
			{
				id:               "ZooKeeperMulti",
				label:            "ZooKeeperMulti",
				metricType:       "DERIVE",
				drawType:         "STACK",
				color:            "COLOUR6",
				clickHouseMetric: "ZooKeeperMulti",
			},
			{
				id:               "ZooKeeperExceptions",
				label:            "ZooKeeperExceptions",
				metricType:       "DERIVE",
				drawType:         "STACK",
				color:            "COLOUR7",
				clickHouseMetric: "ZooKeeperExceptions",
			},
		},
	},
}
