package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/501army/go-tv-show/forms"
	"github.com/gin-gonic/gin"
)

// RequestController is
type RequestController struct{}
type schedule forms.ScheduleForm
type search forms.SearchForm

var tvmaze = "http://api.tvmaze.com"
var client = &http.Client{}

func doRequest(path string) []byte {
	req, _ := http.NewRequest("GET", tvmaze+path, nil)
	res, _ := client.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	return body
}

// Schedule is
func (w *RequestController) Schedule(c *gin.Context) {
	var data []schedule
	body := doRequest("/schedule")
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		fmt.Print(jsonErr)
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "success",
		"data":    data,
	})
}

// Search is
func (w *RequestController) Search(c *gin.Context) {
	var data []search
	param := c.Request.URL.Query()
	query := param.Get("q")
	body := doRequest("/search/shows?q=" + query)
	jsonErr := json.Unmarshal(body, &data)
	if jsonErr != nil {
		fmt.Print(jsonErr)
	}
	c.JSON(200, gin.H{
		"status":  200,
		"message": "success",
		"data":    data,
	})
}
