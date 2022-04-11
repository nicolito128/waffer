package logger

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var defaultFlags = log.Ldate | log.Ltime

type InfoLogger interface {
	Infoln(msg string)
	Infof(msg string, args ...any)
}

type ErrorLogger interface {
	Fail(msg string)
	Failf(msg string, args ...any)
}

// SystemLogger is a structure for handle log outputs.
type SystemLogger struct {
	// Info manager
	Info *log.Logger

	// Error manager
	Errors *log.Logger

	// Define prefix for outputs
	Flags int

	// Bot session access
	session *discordgo.Session
}

// New() create and return a new SystemLogger struct.
func New(s *discordgo.Session) *SystemLogger {
	logger := &SystemLogger{}
	logger.Info = log.New(os.Stdout, "INFO: ", defaultFlags)
	logger.Errors = log.New(os.Stdout, "ERROR: ", defaultFlags)
	logger.Flags = defaultFlags
	logger.session = s

	return logger
}

// Println() call I.Output to print an INFO message.
func (sl *SystemLogger) Println(message string) {
	sl.Info.Println(message)
}

// Printf() call I.Output to print an INFO message with format.
func (sl *SystemLogger) Printf(message string, args ...any) {
	sl.Info.Printf(message, args...)
}

// Fatal() call I.Output to print an ERROR message followed by a call to os.Exit(1).
func (sl *SystemLogger) Fatal(message string) {
	sl.Errors.Fatal(message)
}

// Fatalf() call I.Output to print an ERROR message with format followed by a call to os.Exit(1).
func (sl *SystemLogger) Fatalf(message string, args ...any) {
	sl.Errors.Fatalf(message, args...)
}
