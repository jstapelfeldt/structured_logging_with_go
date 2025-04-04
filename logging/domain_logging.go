package logging

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
)

type domainLogger struct {
	logger *zap.Logger
	level  zap.AtomicLevel
}

// key: domain name, value: *loggerEntry (pointers)
var logMap = sync.Map{}

// Domain represents the name or identifier for a specific domain's logger.
type Domain string

func (d Domain) String() string {
	return string(d)
}

func GetDomainLogger(domain Domain) *zap.Logger {
	if domain == "" {
		return globalLogger
	}

	if existingDomainLogger, exists := logMap.Load(domain); exists {
		if dl, ok := existingDomainLogger.(*domainLogger); ok {
			return dl.logger
		}
		globalLogger.Warn(fmt.Sprintf("invalid type in logMap, got %T", existingDomainLogger), zap.String("domain", domain.String()))
	}

	return createNewDomainLogger(domain)
}

func createNewDomainLogger(domain Domain) *zap.Logger {
	atomicLevel := zap.NewAtomicLevelAt(zap.InfoLevel)
	logger := newLoggerWithLevel(domain, atomicLevel)

	entry := &domainLogger{
		logger: logger,
		level:  atomicLevel,
	}

	logMap.Store(domain, entry)
	return logger
}
