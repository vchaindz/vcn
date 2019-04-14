package logs

import (
	"os"

	"github.com/sirupsen/logrus"
)

// LOG is the global instance of logrus.Logger
var LOG = logrus.New()

func init() {
	ll := os.Getenv("LOG_LEVEL")
	switch ll {
	case "TRACE":
		LOG.SetLevel(logrus.TraceLevel)
	case "DEBUG":
		LOG.SetLevel(logrus.DebugLevel)
	case "INFO":
		LOG.SetLevel(logrus.InfoLevel)
	case "WARN":
		LOG.SetLevel(logrus.WarnLevel)
	case "ERROR":
		LOG.SetLevel(logrus.ErrorLevel)
	case "FATAL":
		LOG.SetLevel(logrus.FatalLevel)
	case "PANIC":
		LOG.SetLevel(logrus.PanicLevel)
	default:
		LOG.SetLevel(logrus.WarnLevel)
	}
}
