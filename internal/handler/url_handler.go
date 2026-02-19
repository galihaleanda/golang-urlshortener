package handler

import (
	"net/http"
	"time"

	"go-url-shortener/internal/service"

	"github.com/gin-gonic/gin"
)

type URLHandler struct {
	service service.URLService
}

func NewURLHandler(s service.URLService) *URLHandler {
	return &URLHandler{s}
}

func (h *URLHandler) Shorten(c *gin.Context) {
	var req struct {
		URL       string  `json:"url"`
		ExpiredAt *string `json:"expired_at"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	url, err := h.service.Shorten(req.URL, req.ExpiredAt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"short_code": url.ShortCode,
		"expires_at": url.ExpiresAt,
	})
}

func (h *URLHandler) Redirect(c *gin.Context) {
	code := c.Param("code")

	url, err := h.service.GetOriginal(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	if url.ExpiresAt != nil && time.Now().After(*url.ExpiresAt) {
		c.JSON(http.StatusGone, gin.H{"error": "link expired"})
		return
	}

	// increment click
	h.service.Increment(code)

	c.Redirect(http.StatusTemporaryRedirect, url.OriginalURL)
}

func (h *URLHandler) Analytics(c *gin.Context) {
	code := c.Param("code")

	url, err := h.service.GetOriginal(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"original_url": url.OriginalURL,
		"click_count":  url.ClickCount,
		"created_at":   url.CreatedAt,
		"expires_at":   url.ExpiresAt,
	})
}
