package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// Logger instance
var log *logrus.Logger

// CustomFormatter defines the format of the logs to match the Spring Boot style
type CustomFormatter struct {
	DisableColors bool // Flag to disable colors
}

// Format method for CustomFormatter to produce the desired log format
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Get the current timestamp with precise time and timezone
	timestamp := entry.Time.UTC().Format("2006-01-02T15:04:05.000Z07:00")

	// Get the PID of the current process
	pid := os.Getpid()

	// Get the log level and align it to a fixed width
	level := fmt.Sprintf("%-5s", entry.Level.String())

	// Get the calling function information (filename and line number)
	var file string
	var line int
	if entry.HasCaller() {
		// Use filepath.Base to extract only the filename
		file = filepath.Base(entry.Caller.File)
		line = entry.Caller.Line
	}

	// Thread-like simulation, use main as default
	thread := "[main]"

	// Package and class name simulation
	packageName := "[amndns]"

	// Create the log message without colors if DisableColors is set to true
	message := entry.Message
	if f.DisableColors {
		message = removeColorCodes(message) // Helper function to remove color codes
	}

	// Create the formatted log line including filename and line number
	logMessage := fmt.Sprintf("%s %s %d --- %s %s %s:%d : %s\n",
		timestamp, level, pid, packageName, thread, file, line, message)

	return []byte(logMessage), nil
}

// Helper function to remove ANSI color codes
func removeColorCodes(s string) string {
	// Strip any ANSI escape codes for colors
	return s
}

// InitLogger initializes the logrus logger with configurable log directory and file
func InitLogger() *logrus.Logger {
	// Create a new Logrus instance
	log = logrus.New()

	// Set log level (you can dynamically change this based on your needs)
	log.SetLevel(logrus.InfoLevel)

	// Load configuration
	viper.SetConfigName("config")    // Config file name without extension
	viper.SetConfigType("yaml")      // Config file type
	viper.AddConfigPath("./config/") // Path to the config folder
	err := viper.ReadInConfig()      // Read the config file
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Get the log directory and file name from the config file
	logDir := viper.GetString("logging.file.path")
	logFile := viper.GetString("logging.file.name")
	if logDir == "" {
		logDir = "." // Default to current directory if not set
	}
	if logFile == "" {
		logFile = "amndns.log" // Default log file name if not set
	}

	// Ensure the directory exists
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	// Set log file path
	logFilePath := filepath.Join(logDir, logFile)
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		// Create a file-specific formatter without colors
		log.Out = file
		log.SetFormatter(&CustomFormatter{DisableColors: true})
	} else {
		// Fallback to stdout if file creation fails
		log.Out = os.Stdout
		log.SetFormatter(&CustomFormatter{DisableColors: false}) // Use colors for console
		log.Warn("Failed to log to file, using default stderr")
	}

	// Enable reporting of the file name and line number
	log.SetReportCaller(true)

	return log
}

// GetLogger returns the logger instance
func GetLogger() *logrus.Logger {
	if log == nil {
		return InitLogger()
	}
	return log
}
