package log

import (
	"io"
	"os"
	"sign-lottery/pkg/constants"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Entry

const LogPath = "./log/lottery.log"

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
	})
	logrus.SetFormatter(&logrus.TextFormatter{})
	if constants.Mode == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetOutput(os.Stdout)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
		fileWriter, err := getWriter()
		if err != nil {
			logrus.Error(err)
		}
		fileAndStdout := io.MultiWriter(os.Stdout, fileWriter)
		logrus.SetOutput(fileAndStdout)
	}
	Log = logrus.WithFields(logrus.Fields{
		"vsersion": constants.Version,
	})
}

func getWriter() (io.Writer, error) {
	return rotatelogs.New(LogPath+".%Y%m%d%H",
		rotatelogs.WithRotationTime(60*time.Second),
		rotatelogs.WithMaxAge(7*24*time.Hour),
	)
}
