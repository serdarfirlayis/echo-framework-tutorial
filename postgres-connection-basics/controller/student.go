package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/serdarfirlayis/echo-framework-tutorial.git/postgres-hello-world/model"
	"net/http"
)

// GetStudents
func GetStudents(c echo.Context) error {
	students, _ := GetRepoStudents()
	return c.JSON(http.StatusOK, students)
}

func GetRepoStudents() ([]model.Students, error) {
	students := []model.Students{}
	return students, nil
}
