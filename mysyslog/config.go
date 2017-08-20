package main

import (
	"bytes"
	"github.com/spf13/viper"
)

// Config holds all configuration for our program
type Config struct {
	Address        string
	SocketType     string
	InstrumentationKey string
}

// NewConfig creates a Config instance
func NewConfig() *Config {
	viper.AutomaticEnv()
	
	viper.SetEnvPrefix("")
	
	viper.SetDefault("address", "0.0.0.0")
	viper.SetDefault("port", "5000")
	
	viper.BindEnv("address")
	viper.BindEnv("port")
	viper.BindEnv("instrumentationkey")
	
	address := viper.GetString("address")
	port := viper.GetString("port")
	instrumentationkey := viper.GetString("instrumentationkey")
	
	var buffer bytes.Buffer
	buffer.WriteString(address)
	buffer.WriteString(":")
	buffer.WriteString(port)
	
	cnf := Config{
		Address: buffer.String(),
		SocketType: "UDP",
		InstrumentationKey: instrumentationkey,
	}
	return &cnf
}
