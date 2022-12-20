package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"topsis/handler/constant"
	"topsis/handler/model"
)

func (h *Handler) Consult(c *gin.Context) {
	logrus.Info("Start api consult...")

	// Get request param
	userId := c.Param("user_id")

	result, err := h.consultDomain.Consult(c, userId)
	if err != nil {
		logrus.Errorf("Consult result fail: %v", err)
		c.JSON(http.StatusBadRequest, &model.ErrorSystem{
			Code:    http.StatusBadRequest,
			Message: constant.ConsultResultFail,
		})
		return
	}

	logrus.Info("Consult result success")
	res := make([]*model.ConsultResponse, 0)
	for _, value := range result {
		res = append(res, value.ToResponse())
	}
	c.JSON(http.StatusOK, res)
	return
}
