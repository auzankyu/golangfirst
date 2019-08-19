package controllers

import (
	"encoding/json"
	"golangapi/models"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func GetApi(c *gin.Context) {
	var obj models.ResponseSource
	endpoint := "https://newsapi.org/v2/sources?apiKey=7f0f2964ca7d479ea1b25ff492de4e71"

	response, err := http.Get(endpoint)
	if err != nil {
		log.Print("gagal get")
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Print("gagal read all")
	}
	json.Unmarshal(responseData, &obj)

	success(c, obj.Sources)
}
