package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "net/http"
)

func handleErrors(err error, ctx *gin.Context, logger *logrus.Logger) {
    ctx.JSON(http.StatusInternalServerError, err)
    logger.Error(err)
}
