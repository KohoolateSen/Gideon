package logging

import (
	. "Gideon/internal/config"
	"strings"

	"github.com/sirupsen/logrus"
)

// Global logger setup
func SetupLogger() error {
	log_level, err := logrus.ParseLevel(
		strings.ToLower(AppConfig.Logging.Level),
	)

	if err != nil {
		return err
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
		ForceQuote:    true,
	})

	logrus.SetLevel(log_level)

	return nil
}
