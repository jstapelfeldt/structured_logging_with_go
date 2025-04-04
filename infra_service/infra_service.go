package infraservice

import (
	"fmt"
	"net/http"

	"log_eval/logging"

	"go.uber.org/zap"
)

var domain logging.Domain = "infra_service"

type Infraservice struct {
	logger *zap.Logger
}

func Newinfraservice() *Infraservice {
	return &Infraservice{
		logger: logging.GetDomainLogger(domain),
	}
}

func (s *Infraservice) BuildInfra(w http.ResponseWriter, r *http.Request) {
	s.logger.Debug("building infra started")
	query := r.URL.Query()
	objectType := query.Get("type")

	s.logger.Info("building object", zap.String("object_type", objectType))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("object of type %s was built\n", objectType)))

	s.logger.Debug("building infra finished")
}
