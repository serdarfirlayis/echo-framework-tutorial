package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/serdarfirlayis/echo-framework-tutorial.git/postgres-connection-basics/controller"
	"github.com/serdarfirlayis/echo-framework-tutorial.git/postgres-connection-basics/storage"
	"log"
	"net/http"
)

func main() {
	db := storage.NewDB()
	if db == nil {
		log.Panic("Failed to connect to database")
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/students", controller.GetStudents)
	/*
		e.POST("/students", controller.CreateStudent)
		e.GET("/students", controller.GetStudents)
		e.GET("/students/:id", controller.GetStudent)
		e.PUT("/students/:id", controller.UpdateStudent)
		e.DELETE("/students/:id", controller.DeleteStudent)
	*/

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello")
}
