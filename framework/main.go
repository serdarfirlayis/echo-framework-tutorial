package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strconv"
)

func main() {
	// terminal command: export MY_APP_PORT=4000
	port := os.Getenv("MY_APP_PORT")
	if port == "" {
		port = "8080"
	}

	e := echo.New()
	proucts := []map[int]string{{1: "mobile"}, {2: "tv"}, {3: "laptop"}}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/products", func(c echo.Context) error {
		return c.JSON(http.StatusOK, proucts)
	})

	e.GET("/products/:id", func(c echo.Context) error {
		var product map[int]string

		for _, p := range proucts {
			for k := range p {
				pID, err := strconv.Atoi(c.Param("id"))
				if err != nil {
					return err
				}
				if pID == k {
					product = p
				}
			}
		}

		if product == nil {
			return c.JSON(http.StatusNotFound, "Product not found")
		}
		return c.JSON(http.StatusOK, product)
	})

	e.POST("/products", func(c echo.Context) error {
		type body struct {
			Name string `json:"product_name"`
		}
		var reqBody body
		if err := c.Bind(&reqBody); err != nil {
			return err
		}
		product := map[int]string{len(proucts) + 1: reqBody.Name}
		proucts = append(proucts, product)
		return c.JSON(http.StatusOK, product)
	})

	e.Logger.Print(fmt.Sprintf("Listening on port %s", port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", port)))

	/*
		e := echo.New()
		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World!")
		})
		e.Logger.Print("Listening on port 8080")
		e.Logger.Fatal(e.Start(":8080"))
	*/
}
