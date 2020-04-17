package main

import (
	"first_docker_project/infrastructure"
	"first_docker_project/interfaces/controllers"
	"github.com/labstack/echo"
)

func main() {
	sqlHandler := infrastructure.NewSQLHandler()
	defer sqlHandler.Conn.Close()

	postController := controllers.NewPostController(sqlHandler.Conn)
	e := echo.New()
	e.GET("/posts", postController.Index)
	e.POST("/posts", postController.Create)
	e.GET("/posts/:id", postController.Show)
	e.PUT("/posts/:id", postController.Update)
	e.DELETE("/posts/:id", postController.Delete)
	e.Logger.Fatal(e.Start(":8080"))
}
