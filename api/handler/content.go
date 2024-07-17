package handler

import (
	"api/api/auth"
	pb "api/genproto/content"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetDestinations godoc
// @Security ApiKeyAuth
// @Summary get
// @Description get destination
// @Tags destinations
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Param name query string false "name"
// @Success 200 {object} content.GetDestinationsRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/destinations [get]
func (h *Handler) GetDestinations(c *gin.Context) {
	h.Log.Info("GetDestinations started")
	req := pb.GetDestinationsReq{}
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	req.Name = c.Query("name")
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

	res, err := h.Content.GetDestinations(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("GetDestinations ended")
}

// GetDestinationsById godoc
// @Security ApiKeyAuth
// @Summary Get
// @Description Get destination by id
// @Tags destinations
// @Param destination_id path string true "destination_id"
// @Success 200 {object} content.GetDestinationsByIdRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/destinations/{destination_id} [get]
func (h *Handler) GetDestinationsById(c *gin.Context) {
	h.Log.Info("GetDestinationsById started")
	req := pb.GetDestinationsByIdReq{}
	req.Id = c.Param("destination_id")
	if len(req.Id) <= 0 {
		h.Log.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})
	}
	res, err := h.Content.GetDestinationsById(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("GetDestinationsById ended")
}

// SendMessage godoc
// @Security ApiKeyAuth
// @Summary Send Message
// @Description Send Message
// @Tags message
// @Param info body content.SendMessageReq true "info"
// @Success 200 {object} content.SendMessageRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/messages [post]
func (h *Handler) SendMessage(c *gin.Context) {
	h.Log.Info("SendMessage started")
	accessToken := c.GetHeader("Authorization")
	id, err := auth.GetUserIdFromAccessToken(accessToken)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req := pb.SendMessageReq{}
	err = c.BindJSON(&req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req.UserId = id
	res, err := h.Content.SendMessage(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, &res)
	h.Log.Info("SendMessage ended")
}

// GetMessages godoc
// @Security ApiKeyAuth
// @Summary get Message
// @Description get Message
// @Tags message
// @Param limit query string false "Number of messages to fetch"
// @Param offset query string false "Number of messages to omit"
// @Success 200 {object} content.GetMessagesRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/messages [get]
func (h *Handler) GetMessages(c *gin.Context) {
	h.Log.Info("GetMessages started")
	req := pb.GetMessagesReq{}
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

	res, err := h.Content.GetMessages(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}

	c.JSON(200, &res)
	h.Log.Info("GetMessages ended")
}

// CreateTips godoc
// @Security ApiKeyAuth
// @Summary create
// @Description create tips
// @Tags tips
// @Param info body content.CreateTipsReq true "destination_id"
// @Success 200 {object} content.CreateTipsRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/travel-tips [post]
func (h *Handler) CreateTips(c *gin.Context) {
	h.Log.Info("CreateTips started")
	accessToken := c.GetHeader("Authorization")
	id, err := auth.GetUserIdFromAccessToken(accessToken)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	req := pb.CreateTipsReq{}
	err = c.BindJSON(&req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
	}
	req.UserId = id
	res, err := h.Content.CreateTips(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("CreateTips ended")
}

// GetTips godoc
// @Security ApiKeyAuth
// @Summary get
// @Description get tips
// @Tags tips
// @Param limit query string false "Number of messages to fetch"
// @Param offset query string false "Number of messages to omit"
// @Param category query string false "category"
// @Success 200 {object} content.GetTipsRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/travel-tips [get]
func (h *Handler) GetTips(c *gin.Context) {
	h.Log.Info("GetTips started")

	req := pb.GetTipsReq{}
	req.Category = c.Query("category")
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

	res, err := h.Content.GetTips(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("GetTips ended")
}

// GetUserStat godoc
// @Security ApiKeyAuth
// @Summary best
// @Description get user
// @Tags users
// @Param user_id path string true "user_id"
// @Success 200 {object} content.GetUserStatRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/users/{user_id}/statistics [get]
func (h *Handler) GetUserStat(c *gin.Context) {
	h.Log.Info("GetUserStat started")
	req := pb.GetUserStatReq{}
	req.UserId = c.Param("user_id")
	if len(req.UserId) <= 0 {
		h.Log.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})
	}
	res, err := h.Content.GetUserStat(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("GetUserStat ended")
}

// TopDestinations godoc
// @Security ApiKeyAuth
// @Summary top places
// @Description get top places
// @Tags top
// @Success 200 {object} content.Answer
// @Failure 500 {object} string "Server error"
// @Router /api/v1/trending-destinations [get]
func (h *Handler) TopDestinations(c *gin.Context) {
	h.Log.Info("TopDestinations started")
	req := pb.Void{}
	res, err := h.Content.TopDestinations(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	fmt.Println(res)
	c.JSON(200, &res)
	h.Log.Info("TopDestinations ended")
}
