package logging

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	once         sync.Once
	globalLogger *zap.Logger
	globalDomain Domain = "global"
)

// Default returns the global logger instance, which is used for logging that is not specific to any domain.
func Default() *zap.Logger {
	once.Do(func() {
		atomicLevel := zap.NewAtomicLevelAt(zap.InfoLevel)
		globalLogger = newLoggerWithLevel(Domain(globalDomain), atomicLevel)
		entry := &domainLogger{
			logger: globalLogger,
			level:  atomicLevel,
		}
		logMap.Store(globalDomain, entry)
	})
	return globalLogger
}

func newLoggerWithLevel(domain Domain, level zap.AtomicLevel) *zap.Logger {
	stdout := zapcore.AddSync(os.Stdout)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		stdout,
		level,
	)

	return zap.New(core).With(zap.String("domain", domain.String()))
}
