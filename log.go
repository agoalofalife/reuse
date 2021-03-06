package reuse

import (
	"github.com/op/go-logging"
	"os"
)

var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

const module = `log`

type Log struct {
	Log *logging.Logger
}

func NewLog() *Log {
	log := Log{logging.MustGetLogger(module)}
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backend1Leveled := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backend1Leveled)
	return &log
}

func (log *Log) LoadModule(app Application) bool {
	app.Container.SetInstance(`log`, log)
	return true
}
