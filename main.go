package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/get-user", GetUser)
	e.GET("/get-order", GetOrder)
	e.GET("/get-product", GetProduct)

	// Start server
	e.Logger.Fatal(e.Start(":"))
}

// Handler
func GetUser(c echo.Context) error {
	return c.String(http.StatusOK, "Data User Berhasil di Get")
}

func GetOrder(c echo.Context) error {
	return c.String(http.StatusOK, "Data Order Berhasil di Get")
}

func GetProduct(c echo.Context) error {
	return c.String(http.StatusOK, "Data Product Berhasil di Get")
}
