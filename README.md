# 🚀 clog - Colored Logger for Golang
A powerful, colorful, and feature-rich logging package that provides enhanced logging capabilities with emoji indicators, color-coded output, and comprehensive debugging features.

### ✨ Features

- Color-coded output with emoji indicators
- UTC timestamps with millisecond precision
- File and line number tracking
- Goroutine ID tracking
- Performance metrics logging
- Multiple log levels with filtering
- Panic handling with stack traces
- Configurable output formatting

### 📥 Installation
```bash
go get github.com/billowdev/clog
```

### 🔰 Quick Start
```go
package main

import "github.com/yourusername/clog"

func main() {
    // Configure logging
    clog.SetLogLevel(clog.DebugLevel)
    
    clog.Info("Starting application...")
    clog.Success("Connected to database")
    clog.Error("Failed to connect: %v", err)
}
```

🎯 Log Levels
- PanicLevel  - System is unusable (💥)
- ErrorLevel  - Error events (❌)
- WarningLevel - Warning messages (⚠️)
- InfoLevel   - General information (ℹ️)
- DebugLevel  - Debug information (🔍)
- TraceLevel  - Detailed debugging (📍)

### ⚙️ Configuration
```go
// Set minimum log level
clog.SetLogLevel(clog.DebugLevel)

// Enable/disable file and line info
clog.SetShowFileLine(true)

// Enable/disable goroutine ID
clog.SetShowGoroutineID(true)
```


### 📝 Example Output

- 🚀 INIT     [2024-12-11 15:04:05.123 UTC] [main.go:25] (goroutine 1) Starting service
- ⚙️ CONFIG   [2024-12-11 15:04:05.124 UTC] [main.go:26] Loading configuration
- ℹ️ INFO     [2024-12-11 15:04:05.125 UTC] [main.go:27] Processing request
- ✅ SUCCESS  [2024-12-11 15:04:05.126 UTC] [main.go:28] Request completed
- ❌ ERROR    [2024-12-11 15:04:05.127 UTC] [main.go:29] Connection failed


### 🎨 Log Types

- ℹ️  INFO     - General information
- ✅ SUCCESS  - Successful operations
- 🚀 INIT     - Initialization events
- ⚙️  CONFIG   - Configuration events
- ⚠️  WARNING  - Warning messages
- ❌ ERROR    - Error messages
- 🔍 DEBUG    - Debug information
- 📍 TRACE    - Trace-level debugging
- 💥 PANIC    - Critical errors
- 📊 METRIC   - Performance metrics

### 📋 Usage Examples
Basic Logging:
```go
clog.Info("Processing request from %s", clientIP)
clog.Success("Request completed successfully")
clog.Warning("High memory usage: %d%%", memoryUsage)
clog.Error("Failed to connect: %v", err)
```

Metrics & Debug:
```go
// Performance metrics
clog.Metric("response_time_ms", 150, "endpoint=/api/users", "method=GET")

// Debug info
clog.Debug("Request payload: %+v", payload)
clog.Trace("SQL query: %s", query)
```

### 📄 License
MIT License
🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### 💡 Support
For support, please open an issue in the GitHub repository.
