package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "github.com/tapokshot/microslab/internal/app/microslab/database"
    "github.com/tapokshot/microslab/internal/app/microslab/model"
    "github.com/tapokshot/microslab/internal/app/microslab/service"
    "net/http"
    "strconv"
)

type User struct {
    userService service.User
    logger      *logrus.Logger
}

func NewUserController(store database.Store, logger *logrus.Logger) *User {
    return &User{
        userService: service.NewUserService(store),
        logger:      logger,
    }
}

func (u *User) GetUserById(ctx *gin.Context) {
    userId, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        handleErrors(err, ctx, u.logger)
        return
    }
    result, err := u.userService.GetUserById(userId)
    if err != nil {
        handleErrors(err, ctx, u.logger)
        return
    }
    ctx.JSON(http.StatusOK, result)
}

func (u *User) Save(ctx *gin.Context) {
    var addUser model.AddUserRequest
    if err := ctx.ShouldBindJSON(&addUser); err != nil {
        handleErrors(err, ctx, u.logger)
        return
    }
    err := u.userService.SaveUser(&addUser)
    if err != nil {
        handleErrors(err, ctx, u.logger)
    }
}

func (u *User) Update(ctx *gin.Context) {
}

func (u *User) Delete(ctx *gin.Context) {
}
