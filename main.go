package main

import (
    "fmt"
    "os"
    "strings"
    "github.com/roistat/go-clickhouse"
)

func main() {
    widgetName, action := parseOptions()
    widget, ok := AvailableWidgets[widgetName]
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
    fmt.Printf(`graph_title %s
graph_category %s
graph_info %s
graph_vlabel %s
graph_period %s
graph_args %s
`, w.graph_title, w.graph_category, w.graph_info, w.graph_vlabel, w.graph_period, w.graph_args)
    for _, m := range w.metrics {
        fmt.Printf(`
%s.label %s
%s.type %s
%s.min 0
%s.draw %s
%s.colour %s
`, m.id, m.label, m.id, m.metric_type, m.id, m.id, m.draw_type, m.id, m.color)
    }
}

func renderWidgetData(w Widget) {
    stats := loadClickHouseStats()
    total := 0
    if w.is_percent {
        for _, m := range w.metrics {
            total = total + stats[m.clickhouseEvent]
        }
    }
    for _, m := range w.metrics {
        value := stats[m.clickhouseEvent]
        if w.is_percent {
            calcValue := float64(value) / float64(total) * float64(100)
            formattedValue := fmt.Sprintf("%.2f", calcValue)
            fmt.Printf("%s.value %s\n", m.id, formattedValue)
        } else {
            fmt.Printf("%s.value %d\n", m.id, value)
        }
    }

}

func loadClickHouseStats() map[string]int {
    conn := clickhouse.NewConn("localhost:8123", clickhouse.NewHttpTransport())
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