package logrusInfra

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func Init() {
	Log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	Log.SetLevel(logrus.DebugLevel)
	Log.SetOutput(os.Stdout)
}
