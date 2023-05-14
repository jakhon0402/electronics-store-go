package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func NewLogger() *logrus.Logger {
	// Create a new logger instance
	log := logrus.New()

	// Set the log level
	log.SetLevel(logrus.DebugLevel)

	// Set the output to stdout
	log.SetOutput(os.Stdout)

	// Add a hook to log errors to Sentry (optional)
	// client, err := sentry.NewClient(...)
	// if err == nil {
	//     log.AddHook(&sentryhook{client})
	// }

	return log
}
