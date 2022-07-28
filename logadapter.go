package logadapterzerolog

import (
	"fmt"
	"os"

	"github.com/goware/logger"
	"github.com/rs/zerolog"
)

func LogAdapter(log zerolog.Logger) logger.Logger {
	return &logAdapter{log}
}

type logAdapter struct {
	log zerolog.Logger
}

var _ logger.Logger = &logAdapter{}

func (s *logAdapter) Debug(v ...interface{}) {
	s.log.Debug().Msg(fmt.Sprint(v...))
}

func (s *logAdapter) Debugf(format string, v ...interface{}) {
	s.log.Debug().Msg(fmt.Sprintf(format, v...))
}

func (s *logAdapter) Info(v ...interface{}) {
	s.log.Info().Msg(fmt.Sprint(v...))
}

func (s *logAdapter) Infof(format string, v ...interface{}) {
	s.log.Info().Msg(fmt.Sprintf(format, v...))
}

func (s *logAdapter) Warn(v ...interface{}) {
	s.log.Warn().Msg(fmt.Sprint(v...))
}

func (s *logAdapter) Warnf(format string, v ...interface{}) {
	s.log.Warn().Msg(fmt.Sprintf(format, v...))
}

func (s *logAdapter) Error(v ...interface{}) {
	s.log.Error().Msg(fmt.Sprint(v...))
}

func (s *logAdapter) Errorf(format string, v ...interface{}) {
	s.log.Error().Msg(fmt.Sprintf(format, v...))
}

func (s *logAdapter) Fatal(v ...interface{}) {
	s.log.Fatal().Msg(fmt.Sprint(v...))
	os.Exit(1)
}

func (s *logAdapter) Fatalf(format string, v ...interface{}) {
	s.log.Fatal().Msg(fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (s *logAdapter) Print(v ...interface{}) {
	s.log.Info().Msg(fmt.Sprint(v...))
}

func (s *logAdapter) Println(v ...interface{}) {
	s.log.Info().Msg(fmt.Sprintln(v...))
}

func (s *logAdapter) Printf(format string, v ...interface{}) {
	s.log.Info().Msg(fmt.Sprintf(format, v...))
}
