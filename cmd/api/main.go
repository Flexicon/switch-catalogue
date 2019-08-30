package main

import (
	"github.com/flexicon/switch-catalogue/pkg/http/handler"
	"github.com/flexicon/switch-catalogue/pkg/http/router"
)

func main() {
	r := router.New()
	base := r.Group("")
	api := r.Group("/api")

	h := handler.NewHandler()
	h.RegisterBase(base)
	h.RegisterApi(api)

	r.Logger.Fatal(r.Start(":3000"))
}
