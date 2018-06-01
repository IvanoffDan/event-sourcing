package controllers

import (
	"errors"
	"testing"
	"time"

	"github.com/IvanoffDan/event-sourcing/controllers/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type RouterContextResponse struct {
	method     string
	statusCode int
	response   gin.H
}

func TestRouterContext(t *testing.T) {
	assert := assert.New(t)

	tests := map[string]struct {
		cpuPercent func(time.Duration, bool) ([]float64, error)
		response   RouterContextResponse
	}{
		"successful": {
			cpuPercent: func(time.Duration, bool) ([]float64, error) {
				return []float64{100}, nil
			},
			response: RouterContextResponse{
				method:     "JSON",
				statusCode: 200,
				response: gin.H{
					"cpus": []float64{100},
				},
			},
		},
		"with error": {
			cpuPercent: func(time.Duration, bool) ([]float64, error) {
				return []float64{}, errors.New("TestError")
			},
			response: RouterContextResponse{
				method:     "AbortWithStatusJSON",
				statusCode: 500,
				response: gin.H{
					"error": "TestError",
				},
			},
		},
	}

	for name, test := range tests {
		t.Logf("Running test case: %s", name)
		cpuPercent = test.cpuPercent
		routerContext := &mocks.RouterContext{}

		routerContext.
			On(test.response.method, test.response.statusCode, test.response.response).
			Return().
			Once()

		healthCheck(routerContext)

		assert.Equal(routerContext.AssertCalled(t, test.response.method, test.response.statusCode, test.response.response), true)
		routerContext.AssertExpectations(t)
	}
}
