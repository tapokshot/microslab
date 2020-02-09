package microslab

import (
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "github.com/tapokshot/microslab/internal/app/microslab/controller"
    "github.com/tapokshot/microslab/internal/app/microslab/database"
)

type MicroslabServer struct {
    config *Config
    logger *logrus.Logger
    router *gin.Engine
    store  *database.Store
}

func New(config *Config) *MicroslabServer {
    return &MicroslabServer{
        config: config,
        logger: logrus.New(),
        router: gin.Default(),
    }
}

func (microslab *MicroslabServer) Launch() error {
    if err := microslab.configureLogger(); err != nil {
        return err
    }
    microslab.logger.Info("Launch app")
    //todo format struct output
    microslab.logger.Debug("Config:")
    microslab.logger.Debug(microslab.config)
    if err := microslab.configureDB(); err != nil {
        return err
    }
    microslab.configureRouter()

    return microslab.router.Run(microslab.config.BindAddr)
}

func (microslab *MicroslabServer) configureLogger() error {
    level, err := logrus.ParseLevel(microslab.config.LogLevel)
    if err != nil {
        return err
    }
    microslab.logger.SetLevel(level)
    return nil
}

func (microslab *MicroslabServer) configureDB() error {
    store, err := (&database.Store{}).New(microslab.config.DatabaseUrl)
    if err != nil {
        return err
    }
    microslab.store = store
    return nil
}

func (microslab *MicroslabServer) configureRouter() {
    healthController := &controller.Health{}

    microslab.router.GET("/health", healthController.Health)

    userController := controller.NewUserController(*microslab.store, microslab.logger)

    v1 := microslab.router.Group("/api/v1")

    v1.GET("/user/:id", userController.GetUserById)
    v1.POST("/user/:id", userController.Update)
    v1.POST("/user", userController.Save)

    v1.DELETE("/user", userController.Delete)
}
