package equilibrium

import (
	"github.com/ChainSafe/chainbridge-utils/core"
	"github.com/ChainSafe/log15"
)

type TransferLogger interface {
	log15.Logger
	ErrorTransfer(msg string, transferContext core.MessageContext, ctx ...interface{})
}

type loggerGelfWrapper struct {
	inner log15.Logger
}

func NewTransferLogger(l log15.Logger) TransferLogger {
	return &loggerGelfWrapper{inner: l}
}

func (l *loggerGelfWrapper) New(ctx ...interface{}) log15.Logger {
	child := l.inner.New(ctx...)
	return &loggerGelfWrapper{child}
}

func (l *loggerGelfWrapper) Trace(msg string, ctx ...interface{}) {
	l.inner.Trace(msg, ctx...)
}

func (l *loggerGelfWrapper) Debug(msg string, ctx ...interface{}) {
	l.inner.Debug(msg, ctx...)
}

func (l *loggerGelfWrapper) Info(msg string, ctx ...interface{}) {
	l.inner.Info(msg, ctx...)
}

func (l *loggerGelfWrapper) Warn(msg string, ctx ...interface{}) {
	l.inner.Warn(msg, ctx...)
}

func (l *loggerGelfWrapper) Error(msg string, ctx ...interface{}) {
	l.inner.Error(msg, ctx...)
	Error(msg, make(core.MessageContext), ctx...)
}

func (l *loggerGelfWrapper) Crit(msg string, ctx ...interface{}) {
	l.inner.Crit(msg, ctx...)
	Crit(msg, make(core.MessageContext), ctx...)
}

func (l *loggerGelfWrapper) GetHandler() log15.Handler {
	return l.inner.GetHandler()
}

func (l *loggerGelfWrapper) SetHandler(h log15.Handler) {
	l.inner.SetHandler(h)
}

func (l *loggerGelfWrapper) ErrorTransfer(msg string, transferContext core.MessageContext, ctx ...interface{}) {
	l.inner.Error(msg, ctx...)
	Error(msg, transferContext, ctx...)
}