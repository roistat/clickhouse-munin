package main

import (
    "fmt"
    "os"
    "strings"
)

type Widget struct {

}

type AvailableWidgets struct {

}

func main() {
    widgetName := ""
    action := "data"
    args := os.Args
    if (len(args) > 1) {
        action = args[1]
    }
    nameParts := strings.Split(args[0], "_")
    if (len(nameParts) > 1) {
        widgetName = nameParts[1]
    }
    fmt.Printf("Widget: %s, Action: %s\n", widgetName, action)
}
