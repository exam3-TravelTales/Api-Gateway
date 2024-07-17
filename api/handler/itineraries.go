package handler

import (
	"api/api/auth"
	pb "api/genproto/itineraries"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Itineraries godoc
// @Security ApiKeyAuth
// @Summary create
// @Description create itineraries
// @Tags itineraries
// @Param info body itineraries.ItinerariesReq true "info"
// @Success 200 {object} itineraries.ItinerariesRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/itineraries [post]
func (h *Handler) Itineraries(c *gin.Context) {
	h.Log.Info("Itineraries started")
	accessToken := c.GetHeader("Authorization")
	id, err := auth.GetUserIdFromAccessToken(accessToken)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
	}
	req := pb.ItinerariesReq{}
	err = c.BindJSON(&req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
	}
	req.UserId = id
	res, err := h.Itinerary.Itineraries(c, &req)
	if err != nil {
		h.Log.Error("there")
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("Itineraries ended")
}

// UpdateItineraries godoc
// @Security ApiKeyAuth
// @Summary update
// @Description update itineraries
// @Tags itineraries
// @Param itinerary_id path string true "itinerary_id"
// @Param info body itineraries.UpdateItinerariesReq true "info"
// @Success 200 {object} itineraries.ItinerariesRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/itineraries/{itinerary_id} [put]
func (h *Handler) UpdateItineraries(c *gin.Context) {
	h.Log.Info("UpdateItineraries started")
	req := pb.UpdateItinerariesReq{}
	err := c.BindJSON(&req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
	}
	req.Id = c.Param("itinerary_id")
	if len(req.Id) <= 0 {

		h.Log.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})

	}
	res, err := h.Itinerary.UpdateItineraries(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("UpdateItineraries ended")
}

// DeleteItineraries godoc
// @Security ApiKeyAuth
// @Summary Delete
// @Description Delete itineraries
// @Tags itineraries
// @Param itinerary_id path string true "itinerary_id"
// @Success 200 {object} string "successfully deleted"
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/itineraries/{itinerary_id} [delete]
func (h *Handler) DeleteItineraries(c *gin.Context) {
	h.Log.Info("DeleteItineraries started")
	req := pb.StoryId{}
	req.Id = c.Param("itinerary_id")
	if len(req.Id) <= 0 {

		h.Log.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})

	}
	_, err := h.Itinerary.DeleteItineraries(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "successfully deleted"})
	h.Log.Info("DeleteItineraries ended")
}

// GetItineraries godoc
// @Security ApiKeyAuth
// @Summary Get
// @Description Get itineraries
// @Tags itineraries
// @Param limit query string false "Number of itineraries to fetch"
// @Param offset query string false "Number of itineraries to omit"
// @Success 200 {object} itineraries.GetItinerariesRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/itineraries [get]
func (h *Handler) GetItineraries(c *gin.Context) {
	h.Log.Info("GetItineraries started")
	req := pb.GetItinerariesReq{}
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

	res, err := h.Itinerary.GetItineraries(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("GetItineraries ended")
}

// GetItinerariesById godoc
// @Security ApiKeyAuth
// @Summary Get
// @Description Get itineraries by id
// @Tags itineraries
// @Param itinerary_id path string true "itinerary_id"
// @Success 200 {object} itineraries.GetItinerariesByIdRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/itineraries/{itinerary_id} [get]
func (h *Handler) GetItinerariesById(c *gin.Context) {
	h.Log.Info("GetItinerariesById started")
	req := pb.StoryId{}
	req.Id = c.Param("itinerary_id")
	if len(req.Id) <= 0 {
		h.Log.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})
	}
	res, err := h.Itinerary.GetItinerariesById(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("GetItinerariesById ended")
}

// CommentItineraries godoc
// @Security ApiKeyAuth
// @Summary comment
// @Description comment itineraries
// @Tags itineraries
// @Param itinerary_id path string true "itinerary_id"
// @Param info body itineraries.CommentItinerariesReq true "info"
// @Success 200 {object} itineraries.CommentItinerariesRes
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error"
// @Router /api/v1/itineraries/{itinerary_id}/comments [post]
func (h *Handler) CommentItineraries(c *gin.Context) {
	h.Log.Info("CommentItineraries started")
	accessToken := c.GetHeader("Authorization")
	id, err := auth.GetUserIdFromAccessToken(accessToken)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
	}
	req := pb.CommentItinerariesReq{}
	err = c.BindJSON(&req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
	}
	req.ItineraryId = c.Param("itinerary_id")
	if len(req.ItineraryId) <= 0 {
		h.Log.Error("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})
	}
	req.AuthorId = id
	res, err := h.Itinerary.CommentItineraries(c, &req)
	if err != nil {
		h.Log.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
	}
	c.JSON(200, &res)
	h.Log.Info("CommentItineraries ended")
}
