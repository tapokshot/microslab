package microslab

import (
    "github.com/gorilla/mux"
    "github.com/sirupsen/logrus"
    "io"
    "net/http"
)

type MicroslabServer struct {
    config *Config
    logger *logrus.Logger
    router *mux.Router
}

func New(config *Config) *MicroslabServer {
    return &MicroslabServer{
        config: config,
        logger: logrus.New(),
        router: mux.NewRouter(),
    }
}

func (microslab *MicroslabServer) Launch() error {
    if err := microslab.configureLogger(); err != nil {
        return err
    }

    microslab.configureRouter()

    microslab.logger.Info("Launch app")
    //todo format struct output
    microslab.logger.Debug("Config:")
    microslab.logger.Debug(microslab.config)

    return http.ListenAndServe(microslab.config.BindAddr, microslab.router)
}

func (microslab *MicroslabServer) configureLogger() error {
    level, err := logrus.ParseLevel(microslab.config.LogLevel)
    if err != nil {
        return err
    }
    microslab.logger.SetLevel(level)
    return nil
}

func (microslab *MicroslabServer) configureRouter() {
    microslab.router.HandleFunc("/health", microslab.handleHealth())
}

func (microslab *MicroslabServer) handleHealth() http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {
        io.WriteString(writer, "OK")
    }

}
