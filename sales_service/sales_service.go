package salesservice

import (
	"fmt"
	"net/http"

	"log_eval/logging"

	"go.uber.org/zap"
)

var domain logging.Domain = "sales_service"

type SalesService struct {
	logger *zap.Logger
}

func NewSalesService() *SalesService {
	return &SalesService{
		logger: logging.GetDomainLogger(domain),
	}
}

func (s *SalesService) ProcessOrder(w http.ResponseWriter, r *http.Request) {
	s.logger.Debug("processing order started")
	query := r.URL.Query()
	orderID := query.Get("orderID")

	s.logger.Info("processing order", zap.String("order_id", orderID))

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Order with number %s processed\n", orderID)))

	s.logger.Debug("processing order finished")
}
