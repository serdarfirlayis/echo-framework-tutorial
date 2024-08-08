package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Statik dosya servisi
	e.Static("/static", "static")

	// Login ve Dashboard endpoint'leri
	e.GET("/login", loginPage)
	e.GET("/dashboard", dashboardPage)

	// Post request for login
	e.POST("/login", login)

	e.Logger.Fatal(e.Start(":8080"))
}

func loginPage(c echo.Context) error {
	return c.File("views/login.html")
}

func dashboardPage(c echo.Context) error {
	return c.File("views/dashboard.html")
}

func login(c echo.Context) error {
	// Bu örnekte sadece form verilerini okuyoruz ve redirect yapıyoruz.
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Burada kullanıcı doğrulama işlemi yapılmalıdır (örneğin, Keycloak ile)
	if email == "user@example.com" && password == "password" {
		return c.Redirect(http.StatusSeeOther, "/dashboard")
	}
	return c.Redirect(http.StatusSeeOther, "/login")
}
