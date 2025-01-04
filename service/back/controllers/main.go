package controllers

import (
	"github.com/eavlongs/go_backend_template/repository"
	"github.com/eavlongs/go_backend_template/utils"
	"github.com/gin-gonic/gin"
)

type MainController struct {
	repo *repository.MainRepository
}

func NewMainController(repo *repository.MainRepository) *MainController {
	return &MainController{repo: repo}
}

func (c *MainController) Test(ctx *gin.Context) {
	utils.RespondWithSuccess(ctx, struct {
		Data string `json:"data"`
	}{Data: "Hello World"})
}
