package main

import (
"os"
	"fmt"
	"reflect"
	"gopkg.in/mcuadros/go-syslog.v2"
	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
)

func main() {
	cnf := NewConfig()

    fmt.Printf("Server started %v - %v\n", cnf.Address, cnf.SocketType)

	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	server := syslog.NewServer()
	server.SetFormat(syslog.Automatic)
	server.SetHandler(handler)
	
	
	if err := server.ListenTCP("0.0.0.0:5000"); err != nil {
			fmt.Fprintf(os.Stderr, "err %v\n", err)
	}
	
	if err := server.ListenUDP("0.0.0.0:5000"); err != nil {
			fmt.Fprintf(os.Stderr, "err %v\n", err)
	}

	server.Boot()

	client := appinsights.NewTelemetryClient("<instrumentation key>")
	client.TrackEvent("custom event")
	client.TrackMetric("custom metric", 123)
	client.TrackTrace("trace message")

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			fmt.Println(logParts)
			fmt.Println(logParts["content"])
			fmt.Println(reflect.TypeOf(logParts))
		}
	}(channel)

	server.Wait()
}
