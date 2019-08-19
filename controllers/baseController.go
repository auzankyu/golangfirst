package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func success(c *gin.Context, data interface{}) {
	var result gin.H
	result = gin.H{
		"responseCode":    http.StatusOK,
		"responseMessage": "success",
		"data":            data,
	}
	c.JSON(http.StatusOK, result)
}

func error(c *gin.Context, data interface{}) {
	var result gin.H
	result = gin.H{
		"responseCode":    http.StatusNotFound,
		"responseMessage": "error",
		"data":            data,
	}
	c.JSON(http.StatusNotFound, result)
}
