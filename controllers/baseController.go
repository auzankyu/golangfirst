package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func success(c *gin.Context, data interface{}) {
	var result gin.H
	result = gin.H{
		"code": http.StatusOK,
		"message": "success",
		"data": data,
	}
	c.JSON(http.StatusOK, result)
}

func error(c *gin.Context) {
	var result gin.H
	result = gin.H{
		"code": http.StatusNotFound,
		"message": "data not found",
		"data": nil,
	}
	c.JSON(http.StatusNotFound, result)
}