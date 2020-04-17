package controllers

import (
	"first_docker_project/domain"
	"net/http"
	"strconv"

	"first_docker_project/interfaces/database"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

type PostController struct {
	db *sqlx.DB
}

func NewPostController(db *sqlx.DB) *PostController {
	return &PostController{db: db}
}

func (c *PostController) Index(ctx echo.Context) error {
	posts, err := database.FindAll(c.db)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, posts)
}

func (c *PostController) Create(ctx echo.Context) error {
	post := domain.Post{}
	if err := ctx.Bind(&post); err != nil {
		return err
	}

	id, err := database.Store(c.db, post)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	post.ID = id
	return ctx.JSON(http.StatusOK, post)
}

func (c *PostController) Show(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	post, err := database.FindByID(c.db, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, post)
}

func (c *PostController) Update(ctx echo.Context) error {
	var post domain.Post
	id := ctx.Param("id")
	identifier, _ := strconv.Atoi(id)
	post.ID = identifier

	if err := ctx.Bind(&post); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := database.Update(c.db, post); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return ctx.JSON(http.StatusOK, post)
}

func (c *PostController) Delete(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := database.Delete(c.db, id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	res := map[string]int{
		"ID": id,
	}

	return ctx.JSON(http.StatusOK, res)
}
