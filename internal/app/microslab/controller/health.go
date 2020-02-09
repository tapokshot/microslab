package controller

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Health struct {
}

func (health *Health) Health(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, "OK")
}
