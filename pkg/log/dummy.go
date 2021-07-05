package log

import "github.com/rs/zerolog"

func DummyLog() {
	_ = zerolog.Nop()
}
