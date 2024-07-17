package handler

import (
	"api/api/auth"
	pb "api/genproto/story"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CreateStory godoc
// @Security ApiKeyAuth
// @Summary create story
// @Description create new story
// @Tags stories
// @Param info body story.CreateStoriesRequest true "info"
// @Success 200 {object} story.CreateStoriesResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories [post]
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
	res, err := h.Story.CreateStories(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("CreateStory ended")
}

// UpdateStory godoc
// @Security ApiKeyAuth
// @Summary Update story
// @Description Update new story
// @Tags stories
// @Param story_id path string true "story_id"
// @Param info body story.UpdateStoriesReq true "info"
// @Success 200 {object} story.UpdateStoriesRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories/{story_id} [put]
func (h *Handler) UpdateStory(c *gin.Context) {
	h.Log.Info("UpdateStory started")
	req := pb.UpdateStoriesReq{}
	if err := c.BindJSON(&req); err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req.Id = c.Param("story_id")
	res, err := h.Story.UpdateStories(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("UpdateStory ended")
}

// DeleteStory godoc
// @Security ApiKeyAuth
// @Summary Delete story
// @Description Delete new story
// @Tags stories
// @Param story_id path string true "story_id"
// @Success 200 {object} string "succesfully deleted"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories/{story_id} [delete]
func (h *Handler) DeleteStory(c *gin.Context) {
	h.Log.Info("DeleteStory started")
	req := pb.StoryId{}
	req.Id = c.Param("story_id")
	_, err := h.Story.DeleteStories(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "succesfully deleted"})
	h.Log.Info("DeleteStory ended")
}

// GetAllStories godoc
// @Security ApiKeyAuth
// @Summary get all story
// @Description get all stories
// @Tags stories
// @Param limit query string false "Number of stories to fetch"
// @Param offset query string false "Number of stories to omit"
// @Success 200 {object} story.GetAllStoriesRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories [get]
func (h *Handler) GetAllStories(c *gin.Context) {
	h.Log.Info("GetAllStories started")
	req := pb.GetAllStoriesReq{}
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			h.Log.Error(err.Error())
			return
		}
		req.Limit = int64(limit)
	} else {
		req.Limit = 10
	}

	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			h.Log.Error(err.Error())
			return
		}
		req.Offset = int64(offset)
	} else {
		req.Offset = 0
	}

	res, err := h.Story.GetAllStories(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("GetAllStories ended")
}

// GetStory godoc
// @Security ApiKeyAuth
// @Summary Get story
// @Description Get story by id
// @Tags stories
// @Param story_id path string true "story_id"
// @Success 200 {object} story.GetStoryRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories/{story_id} [get]
func (h *Handler) GetStory(c *gin.Context) {
	h.Log.Info("GetStory started")
	req := pb.StoryId{}
	req.Id = c.Param("story_id")
	if len(req.Id) <= 0 {

		h.Log.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})

	}
	res, err := h.Story.GetStory(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("GetStory ended")
}

// CommentStory godoc
// @Security ApiKeyAuth
// @Summary comment story
// @Description comment to story
// @Tags stories
// @Param story_id path string true "story_id"
// @Param info body story.CommentStoryReq true "story_id"
// @Success 200 {object} story.CommentStoryRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories/{story_id}/comments [post]
func (h *Handler) CommentStory(c *gin.Context) {
	h.Log.Info("CommentStory started")
	accessToken := c.GetHeader("Authorization")
	id, err := auth.GetUserIdFromAccessToken(accessToken)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
	}
	req := pb.CommentStoryReq{}
	c.BindJSON(&req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
	}
	req.AuthorId = id
	req.StoryId = c.Param("story_id")
	if len(req.StoryId) <= 0 {

		h.Log.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})

	}
	res, err := h.Story.CommentStory(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("CommentStory ended")
}

// GetCommentsOfStory godoc
// @Security ApiKeyAuth
// @Summary comment of story
// @Description get comment of story
// @Tags stories
// @Param story_id path string true "story_id"
// @Param limit query string false "Number of stories to fetch"
// @Param offset query string false "Number of stories to omit"
// @Success 200 {object} story.GetCommentsOfStoryRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories/{story_id}/comments [get]
func (h *Handler) GetCommentsOfStory(c *gin.Context) {
	h.Log.Info("GetCommentsOfStory started")
	req := pb.GetCommentsOfStoryReq{}
	req.StoryId = c.Param("story_id")
	if len(req.StoryId) <= 0 {

		h.Log.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})

	}
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			h.Log.Error(err.Error())
			return
		}
		req.Limit = int64(limit)
	} else {
		req.Limit = 10
	}

	if offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"error": err.Error()})
			h.Log.Error(err.Error())
			return
		}
		req.Offset = int64(offset)
	} else {
		req.Offset = 0
	}

	res, err := h.Story.GetCommentsOfStory(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("GetCommentsOfStory ended")
}

// Like godoc
// @Security ApiKeyAuth
// @Summary comment story
// @Description comment to story
// @Tags stories
// @Param story_id path string true "story_id"
// @Success 200 {object} story.LikeRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/stories/{story_id}/like [post]
func (h *Handler) Like(c *gin.Context) {
	h.Log.Info("Like started")
	accessToken := c.GetHeader("Authorization")
	id, err := auth.GetUserIdFromAccessToken(accessToken)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
	}
	req := pb.LikeReq{}
	req.UserId = id
	req.StoryId = c.Param("story_id")
	if len(req.StoryId) <= 0 {
		h.Log.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})
	}
	res, err := h.Story.Like(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("Like ended")
}
