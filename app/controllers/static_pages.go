package controllers

import (
	"math"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	robotSuccessRate float64
	humanSuccessRate float64
)

func init() {
	rand.Seed(time.Now().UnixNano())

	robotSuccessRate = rand.Float64()
	humanSuccessRate = rand.Float64()
}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"humanSuccessRate": math.Ceil(humanSuccessRate * 100),
		"robotSuccessRate": math.Ceil(robotSuccessRate * 100),
	})
}

func Execute(c *gin.Context) {
	c.Request.ParseForm()
	var isRobotPlay = c.Request.Form["isRobotPlay"][0] == "Yes"
	var isSuccess bool
	if isRobotPlay {
		isSuccess = (robotSuccessRate > rand.Float64())
	} else {
		isSuccess = (humanSuccessRate > rand.Float64())
	}
	c.HTML(200, "result.html", gin.H{
		"isSuccess": isSuccess,
	})
}
