package main

import (
	"github.com/gin-gonic/gin"
	"healthcheck-service/cmd/api"
	"healthcheck-service/internal/healthcheck"
	"log"
)

func main() {
	healthCheckAPI := api.HealthCheckAPI{HealthCheckService: healthcheck.Service{}}
	r := gin.Default()
	r.POST("/check", healthCheckAPI.CheckSitesHandler)
	log.Fatal(r.Run(":8080"))
}
