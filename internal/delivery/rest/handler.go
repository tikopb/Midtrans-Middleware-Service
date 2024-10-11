package rest

import (
	"github.com/tikopb/Midtrans-Middleware-Service/internal/delivery/service"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type handler struct {
	service service.Repository
}

type handlerRespont struct {
	Status  int
	Message string
	Data    interface{} `json:"data"`
}

func NewHandler(service service.Repository) *handler {
	return &handler{
		service: service,
	}
}

func handleError(c echo.Context, statusCode int, err error, data interface{}) error {
	var response handlerRespont

	if strings.Contains(err.Error(), "data not found") {
		statusCode = http.StatusNotFound
	}

	if statusCode != http.StatusOK && statusCode != http.StatusCreated {
		response = handlerRespont{
			Status:  statusCode,
			Message: "Internal Error: " + err.Error(),
			Data:    data,
		}
	} else {
		response = handlerRespont{
			Status:  statusCode,
			Message: "PROCESS SUCCESS: " + err.Error(),
			Data:    data,
		}
	}

	return c.JSON(statusCode, response)
}

func WriteLogErorr(msg string, err error) {
	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	logrus.SetFormatter(formatter)
	logrus.WithFields(logrus.Fields{
		"err": err,
	}).Error(msg, err.Error())
}

func WriteLogInfo(msg string) {
	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}
	logrus.SetFormatter(formatter)
	logrus.Info(msg)
}
