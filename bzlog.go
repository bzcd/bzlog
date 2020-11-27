package bzlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type bzlog struct {
	l *zap.Logger
	f []zap.Field
}

func New(opt *Options) Logger {
	l := zap.New(
		zapcore.NewCore(opt.getEncoder(), opt.getWriteSyncer(), opt.getLevel()),
		opt.getAddCaller()...,
	)

	return &bzlog{
		l: l,
	}
}

func Attach(l *zap.Logger) Logger {
	return &bzlog{l: l}
}

func (g *bzlog) AddCallerSkip(n int) Logger {
	l := g.l.WithOptions(zap.AddCallerSkip(n))
	return &bzlog{l: l, f: g.f}
}

func (g *bzlog) WithFields(key string, val interface{}, kvs ...interface{}) Logger {
	f := append(g.f, zap.Any(key, val))

	if len(kvs) > 0 {
		for i := 0; i < len(kvs)-1; i += 2 {
			key, ok := kvs[i].(string)
			if ok {
				f = append(f, zap.Any(key, kvs[i+1]))
			}
		}
	}

	return &bzlog{l: g.l, f: f}
}

func (g *bzlog) WithMap(fields map[string]interface{}) Logger {
	if fields == nil {
		return g
	}

	f := g.f[:]

	for key, field := range fields {
		f = append(f, zap.Any(key, field))
	}

	return &bzlog{l: g.l, f: f}
}

func (g *bzlog) Debug(msg string) { g.l.Debug(msg, g.f...) }
func (g *bzlog) Info(msg string)  { g.l.Info(msg, g.f...) }
func (g *bzlog) Warn(msg string)  { g.l.Warn(msg, g.f...) }
func (g *bzlog) Error(msg string) { g.l.Error(msg, g.f...) }
func (g *bzlog) Fatal(msg string) { g.l.Fatal(msg, g.f...) }
func (g *bzlog) Panic(msg string) { g.l.Panic(msg, g.f...) }
