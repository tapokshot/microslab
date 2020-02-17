package controller

import (
    "context"
    "github.com/gin-gonic/gin"
    userservice "github.com/tapokshot/microslab/internal/app/grpc"
    "google.golang.org/grpc"
    "log"
    "net/http"
)

type Health struct {
}

func (health *Health) Health(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, "OK")
    conn, err := grpc.Dial("localhost:5300", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Не могу подключиться: %v", err)
    }
    defer conn.Close()
    client := userservice.NewSearchUserClient(conn)

    client.Search(context.Background(), &userservice.SearchUserRequest{
        UserId: 1,
        Phone:  "1111",
    })
}
