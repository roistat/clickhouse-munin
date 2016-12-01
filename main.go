package main

import (
    "fmt"
    "os"
    "strings"
)

type MetricConfig struct {

}

type Widget struct {
    graph_title string
    graph_category string
    graph_info string
    graph_vlabel string
    graph_period string
    graph_args string
    metrics []MetricConfig
}

var availableWidgets = map[string]Widget{
    "queries": Widget{
        graph_title: "ClickHouse queries",
        graph_category: "clickhouse",
        graph_info: "Values received from ClickHouse system.events table",
        graph_vlabel: "queries / second",
        graph_period: "second",
        graph_args: "--lower-limit 0",
    },
}

func main() {
    widgetName, action := parseOptions()
    widget, ok := availableWidgets[widgetName]
    if !ok {
        fmt.Printf("Invalid widget name: %s", widgetName)
        os.Exit(1)
    }

    if (action == "config") {
        renderWidgetConfig(widget)
    } else {
        renderWidgetData(widget)
    }
}

func renderWidgetConfig(w Widget) {
    fmt.Printf("RenderWidgetConfig: %o\n", w)
}

func renderWidgetData(w Widget) {
    stats := loadClickHouseStats()
    fmt.Printf("RenderWidgetData: %o %d\n", w, stats["Query"])

}

func loadClickHouseStats() map[string]int {
    return map[string]int{
        "Query": 748356,
        "SelectQuery": 289681,
        "InsertQuery": 1038037,
    }
}

func parseOptions() (string, string) {
    widgetName := "queries"
    action := "data"
    args := os.Args
    if (len(args) > 1) {
        action = args[1]
    }
    nameParts := strings.Split(args[0], "_")
    if (len(nameParts) > 1) {
        widgetName = nameParts[1]
    }
    return widgetName, action
}