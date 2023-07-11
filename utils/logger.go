package utils

import (
	log "github.com/sirupsen/logrus"
)

func InitLogger() {
	log.SetReportCaller(true)
	log.SetFormatter(
		&log.TextFormatter{
			ForceColors:               true,
			ForceQuote:                true,
			EnvironmentOverrideColors: true,
			FullTimestamp:             true,
			TimestampFormat:           "2006-01-02 15:04:05",
			DisableLevelTruncation:    true,
		},
	)
}
