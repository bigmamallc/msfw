package log

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

func Default() zerolog.Logger {
	return zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}).With().Timestamp().Logger()
}
