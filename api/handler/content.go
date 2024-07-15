package handler

import (
	"api/api/auth"
	pb "api/genproto/content"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateStory godoc
// @Summary create story
// @Description create new story
// @Tags stories
// @Param info body content.CreateStoriesRequest true "info"
// @Success 200 {object} content.RegisterResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/auth/register [post]
func (h *Handler) CreateStory(c *gin.Context) {
	h.Log.Info("CreateStory started")
	req := pb.CreateStoriesRequest{}
	if err := c.BindJSON(&req); err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	accessToken := c.GetHeader("Authorization")
	id, err := auth.GetUserIdFromAccessToken(accessToken)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req.UserId = id
	res, err := h.Content.CreateStories(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
}
