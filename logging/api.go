package logging

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func HandleLogLevelUpdate(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	domain := query.Get("domain")
	level := query.Get("level")

	if domain == "" || level == "" {
		http.Error(w, "Missing domain or level parameter", http.StatusBadRequest)
		return
	}

	if err := updateLogLevel(Domain(domain), level); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Log level for domain %s updated to %s\n", domain, level)))
}

func updateLogLevel(domain Domain, level string) error {
	zapLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		return fmt.Errorf("%w, possible levels are debug, info, warn, error, dpanic, panic, fatal", err)
	}

	if entry, exists := logMap.Load(domain); exists {
		l, ok := entry.(*domainLogger)
		if !ok {
			return fmt.Errorf("invalid type in logMap, got %T", entry)
		}
		l.level.SetLevel(zapLevel)
		l.logger.Info("log level updated", zap.String("new_level", level))
		return nil
	}

	return fmt.Errorf("logger for domain %s not found", domain)
}
