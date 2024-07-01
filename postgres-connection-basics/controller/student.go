package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/serdarfirlayis/echo-framework-tutorial.git/postgres-connection-basics/model"
	"github.com/serdarfirlayis/echo-framework-tutorial.git/postgres-connection-basics/storage"
	"net/http"
)

func GetStudents(c echo.Context) error {
	students, _ := GetRepoStudents()
	return c.JSON(http.StatusOK, students)
}

func GetRepoStudents() ([]model.Students, error) {
	db := storage.GetDBInstance()
	students := []model.Students{}

	if err := db.Find(&students).Error; err != nil {
		return nil, err
	}

	return students, nil
}
