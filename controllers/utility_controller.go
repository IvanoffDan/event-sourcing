package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
)

// RouterContext interface isolates *gin.Context for better testing
type RouterContext interface {
	AbortWithStatusJSON(code int, jsonObject interface{})
	JSON(code int, obj interface{})
}

/* Declaring functions here for testing */

var cpuPercent = cpu.Percent

// HealthCheck controller calls HealthCheck service and returns the result
func HealthCheck(c RouterContext) {

	cpuUsage, err := cpuPercent(0, true)

	if err != nil {
		// 500
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		// Have to return explicitly for testing
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cpus": cpuUsage,
	})
}
