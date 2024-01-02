package controllers

import (
	"gin-boilerplate/helpers/response"
	"gin-boilerplate/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllExample(ctx *gin.Context) {
	example := services.FindAllExample()
	response.ReturnResponse(ctx, http.StatusOK, example)
}
