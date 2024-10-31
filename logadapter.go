package logadapter

import (
	"fmt"
	"log/slog"
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

func (s *logAdapter) With(args ...interface{}) logger.Logger {
	return &logAdapter{s.log.With().Fields(args).Logger()}
}

func (s *logAdapter) Debug(v ...interface{}) {
	msg, args := getArgs(v...)
	withArgs(s.log.Debug(), args).Msg(msg)
}

func (s *logAdapter) Debugf(format string, v ...interface{}) {
	s.log.Debug().Msg(fmt.Sprintf(format, v...))
}

func (s *logAdapter) Info(v ...interface{}) {
	msg, args := getArgs(v...)
	withArgs(s.log.Info(), args).Msg(msg)
}

func (s *logAdapter) Infof(format string, v ...interface{}) {
	s.log.Info().Msg(fmt.Sprintf(format, v...))
}

func (s *logAdapter) Warn(v ...interface{}) {
	msg, args := getArgs(v...)
	withArgs(s.log.Warn(), args).Msg(msg)
}

func (s *logAdapter) Warnf(format string, v ...interface{}) {
	s.log.Warn().Msg(fmt.Sprintf(format, v...))
}

func (s *logAdapter) Error(v ...interface{}) {
	msg, args := getArgs(v...)
	withArgs(s.log.Error(), args).Msg(msg)
}

func (s *logAdapter) Errorf(format string, v ...interface{}) {
	s.log.Error().Msg(fmt.Sprintf(format, v...))
}

func (s *logAdapter) Fatal(v ...interface{}) {
	msg, args := getArgs(v...)
	withArgs(s.log.Fatal(), args).Msg(msg)
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

// getArgs returns the message and arguments from the variadic arguments.
func getArgs(v ...interface{}) (string, []slog.Attr) {
	if len(v) == 0 {
		return "", nil
	}
	if len(v) == 1 {
		return fmt.Sprint(v[0]), nil
	}
	// validate that the first argument is a string
	msg, ok := v[0].(string)
	if !ok {
		return fmt.Sprint(v...), nil
	}
	args := make([]slog.Attr, 0, len(v)-1)
	// validate that the rest of the arguments are Attr
	for i := 1; i < len(v); i++ {
		arg, ok := v[i].(slog.Attr)
		if !ok {
			return fmt.Sprint(v...), nil
		}
		args = append(args, arg)
	}
	return msg, args
}

func withArgs(logger *zerolog.Event, args []slog.Attr) *zerolog.Event {
	for _, args := range args {
		var value interface{}
		switch t := args.Value.Kind(); t {
		case slog.KindAny:
			value = args.Value.Any()
		case slog.KindBool:
			value = args.Value.Bool()
		case slog.KindDuration:
			value = args.Value.Duration()
		case slog.KindFloat64:
			value = args.Value.Float64()
		case slog.KindInt64:
			value = args.Value.Int64()
		case slog.KindString:
			value = args.Value.String()
		case slog.KindTime:
			value = args.Value.Time()
		case slog.KindUint64:
			value = args.Value.Uint64()
		case slog.KindGroup:
			value = args.Value.Group()
		case slog.KindLogValuer:
			value = args.Value.LogValuer()
		}
		logger = logger.Interface(args.Key, value)
	}
	return logger
}
