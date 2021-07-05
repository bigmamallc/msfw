package signal

import (
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"syscall"
)

func AwaitSignalAndExit(log zerolog.Logger) {
	wait := make(chan os.Signal)
	signal.Notify(wait, syscall.SIGINT, syscall.SIGTERM)
	sig := <-wait
	log.Info().Str("signal", sig.String()).Msg("exiting by signal")
}
