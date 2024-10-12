package setup

import (
	"MoneyBackward/global"
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

const (
	red    = 31
	yellow = 33
	blue   = 36
	gray   = 37
)

type logFormatter struct{}

func InitDefaultLogger() {
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(global.Conf.Logger.ShowLine)
	logrus.SetFormatter(&logFormatter{})
	level, err := logrus.ParseLevel(global.Conf.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
}

func InitLogger() *logrus.Logger {
	logSetting := logrus.New()
	logSetting.SetOutput(os.Stdout)
	logSetting.SetReportCaller(global.Conf.Logger.ShowLine)
	logSetting.SetFormatter(&logFormatter{})
	level, err := logrus.ParseLevel(global.Conf.Logger.Level)
	// infoLevel by default
	if err != nil {
		level = logrus.InfoLevel
	}
	logSetting.SetLevel(level)
	return logSetting
}

func (f *logFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// set color
	var color int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		color = gray
	case logrus.WarnLevel:
		color = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		color = red
	default:
		color = blue
	}
	// set buffer
	var buffer *bytes.Buffer
	if entry.Buffer != nil {
		buffer = entry.Buffer
	} else {
		buffer = &bytes.Buffer{}
	}
	log := global.Conf.Logger
	// set timestamp
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	// set caller info
	if entry.HasCaller() {
		funcVal := entry.Caller.Function
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		_, _ = fmt.Fprintf(buffer, "%s[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n", log.Prefix, timestamp, color, entry.Level, fileVal, funcVal, entry.Message)
	} else {
		_, _ = fmt.Fprintf(buffer, "%s[%s] \x1b[%dm[%s]\x1b[0m %s\n", log.Prefix, timestamp, color, entry.Level, entry.Message)
	}
	return buffer.Bytes(), nil
}
