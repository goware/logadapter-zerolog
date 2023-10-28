package logadapter

import (
	"errors"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestLogadapter(t *testing.T) {
	log := LogAdapter(zerolog.New(os.Stdout))

	log.Info("hi")
	log.With("err", errors.New("uh oh"), "ps", "test", "num", 123).Warn("warnnn..")
	log.Info("yes")
}
