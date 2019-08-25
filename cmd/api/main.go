package main

import (
	"github.com/flexicon/switch-catalogue/pkg/http/middleware"
	"github.com/flexicon/switch-catalogue/pkg/http/rpc"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	middleware.SetupGlobalMiddleware(e)
	rpc.SetupRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
