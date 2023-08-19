package logger

import (
	"context"
	"io"
	"os"

	"github.com/Vasiliy82/PoolingRouterEmul/internal/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	global       *zap.SugaredLogger
	defaultLevel = zap.NewAtomicLevelAt(zap.ErrorLevel)
)

func init() {
	SetLogger(NewStdOut(defaultLevel))
}

func New(level zapcore.LevelEnabler, w io.Writer, options ...zap.Option) *zap.SugaredLogger {
	if level == nil {
		level = defaultLevel
	}

	cfg := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	enc := zapcore.NewJSONEncoder(cfg)

	return zap.New(zapcore.NewCore(enc, zapcore.AddSync(w), level), options...).Sugar()
}

func NewStdOut(level zapcore.LevelEnabler, options ...zap.Option) *zap.SugaredLogger {
	return New(level, os.Stdout, options...)
}

func SetLogger(l *zap.SugaredLogger) {
	global = l
}

func Logger() *zap.SugaredLogger {
	return global
}

func WithRequestID(ctx context.Context) *zap.SugaredLogger {
	reqID, err := utils.GetRequestID(ctx)
	if err != nil {
		reqID = utils.NewRequestID()
		global.Warnf("no requestID found, use %s instead", reqID)
	}

	return global.With("requestID", reqID)
}

func SetLevel(l zapcore.Level) {
	defaultLevel.SetLevel(l)
}
