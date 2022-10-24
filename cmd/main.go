package main

import (
	"github.com/juanmabaracat/stock-service/internal/app"
	"github.com/juanmabaracat/stock-service/internal/infrastructure/http"
	"github.com/juanmabaracat/stock-service/internal/infrastructure/storage/memory"
	"github.com/juanmabaracat/stock-service/internal/pkg/uuid"
)

func main() {
	repository := memory.NewRepository()
	uuidProvider := uuid.NewUUIDProvider()
	appServices := app.NewServices(repository, uuidProvider)
	server := http.NewServer(appServices)
	server.ListenAndServe(":8080")
}
