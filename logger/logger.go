package logger

// Recommend to read this article.
// About log.
// https://blog.lulab.net/programmer/what-should-i-log-with-an-intention-method-and-level/

import (
	log "github.com/sirupsen/logrus"
	"github.com/youtube-dl-server/config"
	"os"
)

func InitLogger(config *config.LoggerConfig) {
	log.SetFormatter(&log.JSONFormatter{})
	f, err := os.OpenFile(config.Path, os.O_APPEND|os.O_WRONLY, 0666)
	if os.IsNotExist(err) {
		f, err = os.Create(config.Path)
	}
	if err != nil {
		log.Panicln(err)
	}
	log.SetOutput(f)
	log.Info("Init Server..")
}
