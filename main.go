package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	infraservice "log_eval/infra_service"
	"log_eval/logging"
	salesservice "log_eval/sales_service"
)

func main() {
	startServer()

	waitForShutDown()
}

func waitForShutDown() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	logging.Default().Info("shutting down gracefully")
}

func startServer() {
	logging.Default().Info("application starting")

	salesService := salesservice.NewSalesService()
	infraService := infraservice.Newinfraservice()

	http.HandleFunc("/set-log-level", logging.HandleLogLevelUpdate)
	http.HandleFunc("/process-order", salesService.ProcessOrder)
	http.HandleFunc("/build-infra", infraService.BuildInfra)

	go func() {
		logging.Default().Info("server running on :8080")
		http.ListenAndServe(":8080", nil)
	}()
}
