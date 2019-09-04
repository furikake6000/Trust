package controllers

import (
	"math"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	robotSuccessRate    float64
	humanSuccessRate    float64
	mutualSelectionMode bool
	isHumansTurn        bool
)

func init() {
	rand.Seed(time.Now().UnixNano())
	reset()
	mutualSelectionMode = false
	isHumansTurn = true
}

func reset() {
	humanSuccessRate = rand.Float64()*0.4 + 0.3
	if humanSuccessRate > 0.5 {
		robotSuccessRate = humanSuccessRate + rand.Float64()*0.2 + 0.1
	} else {
		robotSuccessRate = humanSuccessRate - rand.Float64()*0.2 - 0.1
	}
	isHumansTurn = true
}

func Index(c *gin.Context) {
	var robotsSelection = -1
	if mutualSelectionMode && !isHumansTurn {
		if robotSuccessRate > humanSuccessRate {
			robotsSelection = 1
		} else {
			robotsSelection = 0
		}
	}
	c.HTML(200, "index.html", gin.H{
		"humanSuccessRate":    math.Ceil(humanSuccessRate * 100),
		"robotSuccessRate":    math.Ceil(robotSuccessRate * 100),
		"mutualSelectionMode": mutualSelectionMode,
		"robotsSelection":     robotsSelection,
	})
}

func Reset(c *gin.Context) {
	reset()

	c.Redirect(302, "/")
}

func Switch(c *gin.Context) {
	reset()
	mutualSelectionMode = !mutualSelectionMode

	c.Redirect(302, "/")
}

func Execute(c *gin.Context) {
	c.Request.ParseForm()
	isHumansTurn = !isHumansTurn
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
