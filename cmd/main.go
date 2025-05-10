package main

import (
	"Gideon/internal/bot"
	. "Gideon/internal/config"
	"Gideon/internal/utils/httpproxy"
	"Gideon/internal/utils/logging"

	"github.com/sirupsen/logrus"
)

func main() {
	err := AppConfig.LoadConfig()
	if err != nil {
		logrus.Fatalf("failed to load config ➜ %v", err)
	}

	logrus.Info("config file loaded successfully.")

	err = logging.SetupLogger()
	if err != nil {
		logrus.Warnf("failed to setup custom logger, switching to default ➜ %v", err)
	}

	err = httpproxy.SetupHTTPClient()
	if err != nil {
		logrus.Warnf("failed to setup custom httpclient, switching to default ➜ %v", err)
	}

	tb, err := bot.InitBot()
	if err != nil {
		logrus.Fatalf("failed to initiate telegram bot ➜ %v", err)
	}

	logrus.Info("listening for updates ...")
	tb.StartUpdater()
}
