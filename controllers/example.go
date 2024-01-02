package controllers

import (
	"gin-boilerplate/helpers/response"
	"gin-boilerplate/models"
	"gin-boilerplate/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllExample(ctx *gin.Context) {
	example := services.FindAllExample()
	response.ReturnResponse(ctx, http.StatusOK, example)
}

func CreateExample(ctx *gin.Context) {
	var json models.Example
	if err := ctx.ShouldBindJSON(&json); err != nil {
		response.ReturnResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	services.CreateExample(&json)
	response.ReturnResponse(ctx, http.StatusOK, json)
}

func UpdateExample(ctx *gin.Context) {
	var json models.Example
	if err := ctx.ShouldBindJSON(&json); err != nil {
		response.ReturnResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	services.UpdateExample(&json)
	response.ReturnResponse(ctx, http.StatusOK, json)
}

func DeleteExample(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := uuid.Validate(id); err != nil {
		response.ReturnResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	services.DeleteExample(id)
	response.ReturnResponse(ctx, http.StatusOK, nil)
}
