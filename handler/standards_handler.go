package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"topsis/handler/constant"
	"topsis/handler/model"
)

func (h *Handler) CreateStandard(c *gin.Context) {
	logrus.Info("Start api create standard...")

	// Parse request
	var standardRequest *model.StandardRequest
	err := c.ShouldBindJSON(&standardRequest)
	if err != nil {
		logrus.Errorf("Parse request create standard fail: %v", err)
		c.JSON(http.StatusBadRequest, &model.ErrorSystem{
			Code:    http.StatusBadRequest,
			Message: constant.ParseRequestFail,
		})
		return
	}

	// Validate request
	listErr := h.ValidateRequest(standardRequest)
	if len(listErr) > 0 {
		logrus.Errorf("Request is invalid: %v", listErr)
		c.JSON(http.StatusBadRequest, listErr)
		return
	}

	standardCreated, err := h.standardDomain.CreateStandard(c, standardRequest)
	if err != nil {
		logrus.Errorf("Create new standard fail: %v", err)
		c.JSON(http.StatusBadRequest, &model.ErrorSystem{
			Code:    http.StatusBadRequest,
			Message: constant.CreateNewStandardFail,
		})
		return
	}

	logrus.Info("Create new standard success")
	c.JSON(http.StatusCreated, &model.StandardResponse{
		ID:           standardCreated.ID,
		UserID:       standardCreated.UserId,
		StandardName: standardCreated.StandardName,
		Weight:       standardCreated.Weight,
	})
	return
}

func (h *Handler) GetStandards(c *gin.Context) {
	logrus.Info("Start api get list standard...")

	// Get request param
	userID := c.Param("user_id")

	standards, err := h.standardDomain.GetStandards(c, &model.Standard{
		UserID: userID,
	})
	if err != nil {
		logrus.Errorf("Get list standard fail: %v", err)
		c.JSON(http.StatusBadRequest, &model.ErrorSystem{
			Code:    http.StatusBadRequest,
			Message: constant.GetStandardsFail,
		})
		return
	}

	logrus.Info("Get list standard success")
	var res []*model.StandardResponse
	for _, s := range standards {
		res = append(res, s.ToResponse())
	}
	c.JSON(http.StatusOK, res)
	return
}

func (h *Handler) DeleteStandard(c *gin.Context) {
	logrus.Info("Start api delete standard...")

	// Get request param
	id := c.Param("id")

	err := h.standardDomain.DeleteStandard(c, &model.Standard{
		ID: id,
	})
	if err != nil {
		logrus.Errorf("Delete standard fail: %v", err)
		c.JSON(http.StatusBadRequest, &model.ErrorSystem{
			Code:    http.StatusBadRequest,
			Message: constant.DeleteStandardFail,
		})
		return
	}

	logrus.Info("Delete standard success")
	c.JSON(http.StatusOK, &model.DeletedResponse{
		Id:     id,
		Status: model.SUCCESS,
	})
	return
}
