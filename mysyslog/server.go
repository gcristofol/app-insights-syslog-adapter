package main

import (
	"os"
	"fmt"
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
	
	
	if err := server.ListenTCP(cnf.Address); err != nil {
			fmt.Fprintf(os.Stderr, "err %v\n", err)
	}
	
	if err := server.ListenUDP(cnf.Address); err != nil {
			fmt.Fprintf(os.Stderr, "err %v\n", err)
	}

	server.Boot()

	client := appinsights.NewTelemetryClient(cnf.InstrumentationKey)
	fmt.Printf("App Insights client using key [%v]\n",cnf.InstrumentationKey)
	//client.TrackMetric("custom metric", 123)
	//client.TrackTEvent("trace message")

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			fmt.Println(logParts)
			//var content string
			content := logParts["content"].(string)
			client.TrackTrace(content)
		}
	}(channel)

	server.Wait()
}
