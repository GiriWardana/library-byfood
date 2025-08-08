package controllers

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// Request represents the expected JSON request body
type CleanURLRequest struct {
	URL       string `json:"url" binding:"required,url"`
	Operation string `json:"operation" binding:"required,oneof=canonical redirection all"`
}

// Response structure
type CleanURLResponse struct {
	ProcessedURL string `json:"processed_url"`
}

// CleanURLHandler godoc
//
// @Summary      Clean or redirect URL based on operation
// @Description  Accepts a URL and an operation type, processes the URL, and returns the result.
// @Tags         url
// @Accept       json
// @Produce      json
// @Param        request body CleanURLRequest true "URL and operation type"
// @Success      200 {object} CleanURLResponse
// @Failure      400 {object} models.ErrorResponseCleanURL
// @Router       /clean_url [post]
func CleanURLHandler(c *gin.Context) {
	var req CleanURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	parsedURL, err := url.Parse(req.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse URL"})
		return
	}

	// Apply canonical operation
	if req.Operation == "canonical" || req.Operation == "all" {
		parsedURL.RawQuery = ""                                  // Remove query
		parsedURL.Path = strings.TrimSuffix(parsedURL.Path, "/") // Remove trailing slash
	}

	// Apply redirection operation
	if req.Operation == "redirection" || req.Operation == "all" {
		parsedURL.Host = "www.byfood.com" // Ensure domain
		parsedURL.Scheme = strings.ToLower(parsedURL.Scheme)
		parsedURL.Path = strings.ToLower(parsedURL.Path)
	}

	c.JSON(http.StatusOK, CleanURLResponse{
		ProcessedURL: parsedURL.String(),
	})
}
