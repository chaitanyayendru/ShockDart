package pkg

import (
	"github.com/sirupsen/logrus"
)

func Initialize() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
}
