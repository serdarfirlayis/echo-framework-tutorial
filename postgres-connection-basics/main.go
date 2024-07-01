package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/serdarfirlayis/echo-framework-tutorial.git/postgres-connection-basics/controller"
	"github.com/serdarfirlayis/echo-framework-tutorial.git/postgres-connection-basics/storage"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()
	storage.NewDB()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/students", controller.GetStudents)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}
