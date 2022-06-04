package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"topsis/handler/constant"
	"topsis/handler/model"
)

func (h *Handler) BulkCreateScoreRating(c *gin.Context) {
	logrus.Info("Start api bulk create score rating...")

	// Parse request
	var scoreRatingRequest []*model.ScoreRatingRequest
	err := c.ShouldBindJSON(&scoreRatingRequest)
	if err != nil {
		logrus.Errorf("Parse request create score rating fail: %v", err)
		c.JSON(http.StatusBadRequest, &model.ErrorSystem{
			Code:    http.StatusBadRequest,
			Message: constant.ParseRequestFail,
		})
		return
	}

	// Validate request
	var listErr []*model.ErrorValidateRequest
	for _, request := range scoreRatingRequest {
		listErr = append(listErr, h.ValidateRequest(request)...)
	}
	if len(listErr) > 0 {
		logrus.Errorf("Request is invalid: %v", listErr)
		c.JSON(http.StatusBadRequest, listErr)
		return
	}

	err = h.scoreRatingDomain.BulkCreateScoreRating(c, scoreRatingRequest)
	if err != nil {
		logrus.Errorf("Bulk create new score rating fail: %v", err)
		c.JSON(http.StatusBadRequest, &model.ErrorSystem{
			Code:    http.StatusBadRequest,
			Message: constant.BulkCreateNewScoreRatingFail,
		})
		return
	}

	logrus.Info("Bulk create new score rating success")
	c.JSON(http.StatusCreated, &model.BulkCreateResponse{
		Type:   "Bulk create",
		Status: model.SUCCESS,
	})
	return
}

func (h *Handler) DeleteScoreRating(c *gin.Context) {
	logrus.Info("Start api delete score rating...")

	// Get request param
	id := c.Param("id")

	err := h.scoreRatingDomain.DeleteScoreRating(c, &model.ScoreRating{
		ID: id,
	})
	if err != nil {
		logrus.Errorf("Delete score rating fail: %v", err)
		c.JSON(http.StatusBadRequest, &model.ErrorSystem{
			Code:    http.StatusBadRequest,
			Message: constant.DeleteScoreRatingFail,
		})
		return
	}

	logrus.Info("Delete score rating success")
	c.JSON(http.StatusOK, &model.DeletedResponse{
		Id:     id,
		Status: model.SUCCESS,
	})
	return
}
