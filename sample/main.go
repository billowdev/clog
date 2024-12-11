package main

import (
	"errors"
	"time"
	"github.com/billowdev/clog"
)

func performDatabaseOperation() error {
	// Simulate a database operation that fails
	return errors.New("connection timeout")
}

func main() {
	// Configure the logger
	clog.SetLogLevel(clog.DebugLevel)
	clog.SetShowFileLine(true)
	clog.SetShowGoroutineID(true)

	// Application startup logging
	clog.Init("Starting application...")
	clog.Config("Loading configuration from config.yaml")

	// General information logging
	clog.Info("Application initialized successfully")
	clog.Success("Server started on port %d", 8080)

	// Debug and trace information
	clog.Debug("Connected to database with timeout: %v", 30*time.Second)
	clog.Trace("Establishing connection pool with size: %d", 10)

	// Performance metrics
	clog.Metric("active_connections", 42, "db=postgres", "host=primary")
	clog.Metric("request_latency_ms", 156.3, "endpoint=/api/users")

	// Warning example
	clog.Warning("High memory usage detected: %d%%", 85)

	// Error handling
	if err := performDatabaseOperation(); err != nil {
		clog.Error("Database operation failed: %v", err)
	}

	// Example of different log levels in a workflow
	clog.Debug("Processing new user registration...")
	clog.Info("New user registered: %s", "john.doe@example.com")
	clog.Success("Welcome email sent successfully")

	// Demonstrating panic (commented out to prevent program termination)
	// clog.Panic("Critical system error: %v", errors.New("unexpected state"))

	// More metrics
	clog.Metric("user_registrations", 1, "source=web")
	clog.Metric("email_delivery_time_ms", 245.8)
}