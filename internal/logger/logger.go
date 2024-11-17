package logger

import (
	"github.com/lcok/vault-sync/internal/config"
	"os"

	"github.com/sirupsen/logrus"
)

// Logger Simple logger using logrus
type Logger struct {
	*logrus.Logger
}

func NewLogger(c *config.EnvConfig) (*Logger, error) {
	level, err := logrus.ParseLevel(c.LogLevel)
	if err != nil {
		return nil, err
	}

	log := logrus.New()
	log.SetLevel(level)
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return &Logger{Logger: log}, nil
}

func (l *Logger) Infof(format string, args ...any) {
	l.Logger.Infof(format, args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	l.Logger.Errorf(format, args...)
}

func (l *Logger) Debugf(format string, args ...any) {
	l.Logger.Debugf(format, args...)
}

func (l *Logger) Warnf(format string, args ...any) {
	l.Logger.Warnf(format, args...)
}
