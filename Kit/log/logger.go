package log

import (
	"io"
	"io/ioutil"
	"log"
	"strings"
)

type Logger interface {
	Infof(string, ...interface{})
	Errorf(string, ...interface{})
}

type Prefixer interface {
	Prefix(string) error
	PrefixLogType(string, string) error
}

type Enabler interface {
	Enable(bool)
}

type Closer interface {
	Close() error
}

type LoggerReader struct {
	Logger
}

type baseLogger struct {
	logger *log.Logger
	w      io.Writer

	enabled  bool
	prefixes map[string]string
}

var baseLoggerObj = initBaseLogger(ioutil.Discard, "logger")

func initBaseLogger(w io.Writer, prefix string) *baseLogger {
	if w == nil {
		w = ioutil.Discard
	}

	b := &baseLogger{
		w:       w,
		logger:  log.New(w, "", log.LstdFlags),
		enabled: true,
		prefixes: map[string]string{
			"info":  "INFO <" + prefix + ">",
			"error": "ERROR <" + prefix + ">",
		},
	}

	return b
}

func (s *baseLogger) Infof(fmt string, args ...interface{}) {
	if s.enabled {
		s.logger.Printf(s.prefixes["info"]+" : "+fmt, args...)
	}
}

func (s *baseLogger) Errorf(fmt string, args ...interface{}) {
	if s.enabled {
		s.logger.Printf(s.prefixes["error"]+" : "+fmt, args...)
	}
}

func (s *baseLogger) Prefix(prefix string) error {
	s.prefixes["info"] = "INFO <" + prefix + ">"
	s.prefixes["error"] = "ERROR <" + prefix + ">"

	return nil
}

func (s *baseLogger) PrefixLogType(logtype string, prefix string) error {
	s.prefixes[logtype] = strings.Split(s.prefixes[logtype], "<")[0] + "<" + prefix + ">"
	return nil
}

func (s *baseLogger) Enabler(f bool) {
	s.enabled = f
}
