package main

import (
	"fmt"
	"github.com/flexicon/switch-catalogue/pkg/http/handler"
	"github.com/flexicon/switch-catalogue/pkg/http/router"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	r := router.New()
	base := r.Group("")
	api := r.Group("/api")

	h := handler.NewHandler()
	h.RegisterBase(base)
	h.RegisterApi(api)

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	address := fmt.Sprintf(":%s", port)
	r.Logger.Fatal(r.Start(address))
}
