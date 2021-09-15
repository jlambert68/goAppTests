package server

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

// Initiate logger if not already done
func (server *Server) checkIfLoggerIsInitiated() {
	if server.loggerIsInitiated == true {
		return
	}

	server.loggerIsInitiated = true

	server.InitLogger("")
}

func (server *Server) InitLogger(filename string) {
	server.logger = logrus.StandardLogger()

	switch LoggingLevel {

	case logrus.DebugLevel:
		log.Println("'common_config.LoggingLevel': ", LoggingLevel)

	case logrus.InfoLevel:
		log.Println("'common_config.LoggingLevel': ", LoggingLevel)

	case logrus.WarnLevel:
		log.Println("'common_config.LoggingLevel': ", LoggingLevel)

	default:
		log.Println("Not correct value for debugging-level, this was used: ", LoggingLevel)
		os.Exit(0)

	}

	logrus.SetLevel(LoggingLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
		DisableSorting:  true,
	})

	//If no file then set standard out

	if filename == "" {
		server.logger.Out = os.Stdout

	} else {
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
		if err == nil {
			server.logger.Out = file
		} else {
			log.Println("Failed to log to file, using default stderr")
		}
	}

	// Should only be done from init functions
	//grpclog.SetLoggerV2(grpclog.NewLoggerV2(logger.Out, logger.Out, logger.Out))

}
