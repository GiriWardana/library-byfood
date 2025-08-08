package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name           string
	input          CleanURLRequest
	expectedStatus int
	expectedURL    string
}

func TestCleanURLHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/clean_url", CleanURLHandler)

	tests := []testCase{
		{
			name: "Canonical only",
			input: CleanURLRequest{
				URL:       "https://byfood.com/path/with/query?ref=123/",
				Operation: "canonical",
			},
			expectedStatus: http.StatusOK,
			expectedURL:    "https://byfood.com/path/with/query",
		},
		{
			name: "Redirection only",
			input: CleanURLRequest{
				URL:       "https://BYFOOD.com/Food-EXPeriences",
				Operation: "redirection",
			},
			expectedStatus: http.StatusOK,
			expectedURL:    "https://www.byfood.com/food-experiences",
		},
		{
			name: "All operations",
			input: CleanURLRequest{
				URL:       "https://BYFOOD.com/food-EXPeriences?query=abc/",
				Operation: "all",
			},
			expectedStatus: http.StatusOK,
			expectedURL:    "https://www.byfood.com/food-experiences",
		},
		{
			name: "Invalid URL",
			input: CleanURLRequest{
				URL:       "invalid-url",
				Operation: "canonical",
			},
			expectedStatus: http.StatusBadRequest,
			expectedURL:    "",
		},
		{
			name: "Missing operation",
			input: CleanURLRequest{
				URL:       "https://byfood.com",
				Operation: "",
			},
			expectedStatus: http.StatusBadRequest,
			expectedURL:    "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.input)
			req := httptest.NewRequest("POST", "/clean_url", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()
			r.ServeHTTP(resp, req)

			assert.Equal(t, tc.expectedStatus, resp.Code)

			if resp.Code == http.StatusOK {
				var response CleanURLResponse
				err := json.Unmarshal(resp.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedURL, response.ProcessedURL)
			}
		})
	}
}
