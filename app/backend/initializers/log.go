package initializers

import (
	"blood-type-compatibility/helpers"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// Set global log writers
var (
	accessLogWriter io.Writer
	errorLogWriter  io.Writer
)

// Getters for log writers
func AccessLogWriter() io.Writer {
	return accessLogWriter
}
func ErrorLogWriter() io.Writer {
	return errorLogWriter
}

// Implement a hook to write logs of specified LogLevels to specified Writer
type WriterHook struct {
	Writer    io.Writer
	LogLevels []logrus.Level
}

func (hook *WriterHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write([]byte(line))
	return err
}
func (hook *WriterHook) Levels() []logrus.Level {
	return hook.LogLevels
}

// Initialize the logger
func InitLogger() {
	// Create logs directory with current date and time with each run
	dirName := time.Now().Format(helpers.TimeFormat)
	os.MkdirAll(fmt.Sprint("logs/", dirName), os.ModePerm)

	// Create error log file
	errorLogFile := fmt.Sprint("logs/", dirName, "/error.log")
	ef, err := os.OpenFile(errorLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Fatal(err)
	}

	// Create access log file
	accessLogFile := fmt.Sprint("logs/", dirName, "/access.log")
	af, err := os.OpenFile(accessLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Fatal(err)
	}

	// Set global log writers to file and console
	accessLogWriter = io.MultiWriter(af, os.Stdout)
	errorLogWriter = io.MultiWriter(ef, os.Stderr)

	// Set log output to the global log writers via hooks
	logrus.SetOutput(io.Discard)
	logrus.AddHook(&WriterHook{
		Writer:    accessLogWriter,
		LogLevels: []logrus.Level{logrus.TraceLevel, logrus.DebugLevel, logrus.InfoLevel, logrus.WarnLevel},
	})
	logrus.AddHook(&WriterHook{
		Writer:    errorLogWriter,
		LogLevels: []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel},
	})

	// Set log format and level
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: helpers.TimeFormat,
		PrettyPrint:     true,
	})
	logrus.SetLevel(logrus.TraceLevel)
}
