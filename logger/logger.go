package logger

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"runtime/debug"
	"time"

	"github.com/evalphobia/logrus_sentry"
	raven "github.com/getsentry/raven-go"
	logrus "github.com/sirupsen/logrus"
)

type Clog struct {
	logrus *logrus.Entry
	raven  *raven.Client
}

func (c *Clog) Errorf(r *http.Request, format string, args ...interface{}) {
	req := c.logrus
	req = req.WithField("http_request", r)
	req.Errorf(format, args...)
}

func (c *Clog) Infof(r *http.Request, format string, args ...interface{}) {
	req := c.logrus
	req = req.WithField("http_request", r)
	req.Infof(format, args...)
}

// Recover wraps raven.RecoveryHandler to use JSON output.
// https://godoc.org/github.com/getsentry/raven-go#RecoveryHandler
func (c *Clog) Recover(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rval := recover(); rval != nil {
				debug.PrintStack()
				rvalStr := fmt.Sprint(rval)
				packet := raven.NewPacket(rvalStr, raven.NewException(errors.New(rvalStr), raven.NewStacktrace(2, 3, nil)), raven.NewHttp(r))
				raven.Capture(packet, nil)

				w.Header().Set("Content-Type", "application/json")
				http.Error(w, `{"error": "internal server error"}`, 500)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

// Latency logs request latency.
func (c *Clog) Latency(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		c.logrus.Data = logrus.Fields{"latency": t2.Sub(t1).Seconds()}
	}
	return http.HandlerFunc(fn)
}

// MustNew creates a new logger context, which will panic on error.
func MustNew(hostname, dsn string) *Clog {
	l := logrus.New()
	l.Out = os.Stdout

	client, err := raven.New(dsn)
	if err != nil {
		panic(err)
	}

	hook, err := logrus_sentry.NewAsyncWithClientSentryHook(client, logrus.AllLevels)
	if err != nil {
		panic(err)
	}
	hook.StacktraceConfiguration.Enable = true

	l.Hooks.Add(hook)
	ctx := l.WithFields(logrus.Fields{
		"server_name": hostname,
	})

	return &Clog{
		logrus: ctx,
		raven:  client,
	}
}

