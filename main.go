package main

import (
	"fmt"
	"github.com/roistat/go-clickhouse"
	"os"
	"strings"
)

func main() {
	graphName, action := parseOptions()
	graph, ok := AvailableGraphs[graphName]
	if !ok {
		fmt.Printf("Invalid graph name: %s", graphName)
		os.Exit(1)
	}

	if action == "config" {
		renderGraphConfig(graph)
	} else {
		renderGraphData(graph)
	}
}

func renderGraphConfig(g Graph) {
	fmt.Printf(`graph_title %s
graph_category %s
graph_info %s
graph_vlabel %s
graph_period %s
graph_args %s
`, g.graphTitle, g.graphCategory, g.graphInfo, g.graphLabel, g.graphPeriod, g.graphArgs)
	for _, m := range g.metrics {
		fmt.Printf(`
%s.label %s
%s.type %s
%s.min 0
%s.draw %s
%s.colour %s
`, m.id, m.label, m.id, m.metricType, m.id, m.id, m.drawType, m.id, m.color)
	}
}

func renderGraphData(g Graph) {
	stats := loadClickHouseStats()
	total := 0
	if g.isPercent {
		for _, m := range g.metrics {
			total = total + stats[m.clickHouseEvent]
		}
	}
	for _, m := range g.metrics {
		value := stats[m.clickHouseEvent]
		if g.isPercent {
			calcValue := float64(value) / float64(total) * float64(100)
			formattedValue := fmt.Sprintf("%.2f", calcValue)
			fmt.Printf("%s.value %s\n", m.id, formattedValue)
		} else {
			fmt.Printf("%s.value %d\n", m.id, value)
		}
	}

}

func loadClickHouseStats() map[string]int {
	host := os.Getenv("clickhouse_host")
	if host == "" {
		host = "localhost"
	}
	conn := clickhouse.NewConn(host+":8123", clickhouse.NewHttpTransport())
	query := clickhouse.NewQuery("SELECT event, value FROM system.events")
	iterator := query.Iter(conn)
	var (
		event string
		value int
	)
	result := map[string]int{}
	for iterator.Scan(&event, &value) {
		result[event] = value
	}
	return result
}

func parseOptions() (string, string) {
	graphName := "queries"
	action := "data"
	args := os.Args
	if len(args) > 1 {
		action = args[1]
	}
	nameParts := strings.Split(args[0], "_")
	if len(nameParts) > 1 {
		graphName = nameParts[1]
	}
	return graphName, action
}
