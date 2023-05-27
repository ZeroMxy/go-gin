package log

import (
	"go-gin/config"
	"sync"

	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/gookit/slog/rotatefile"
)

type Log struct {}

var (
	log = &slog.Logger {}
	once sync.Once
)

func Debug (message ...interface {}) {

	once.Do(func () {
		new()
	})

	log.Debug(message)
	return
}

func Info (message ...interface {}) {

	once.Do(func () {
		new()
	})

	log.Info(message)
	return
}

func Error (message ...interface {}) {

	once.Do(func () {
		new()
	})

	log.Error(message)
	return
}

func Debugf (format string, message ...interface {}) {

	once.Do(func () {
		new()
	})

	log.Debugf(format, message)
	return
}

func Infof (format string, message ...interface {}) {

	once.Do(func () {
		new()
	})

	log.Infof(format, message)
	return
}

func Errorf (format string, message ...interface {}) {

	once.Do(func () {
		new()
	})
	
	log.Errorf(format, message)
	return
}

// Example Initialize the system log
// 初始化系统日志
func InitSystemLog () *rotatefile.Writer {

	system := config.Log["system"]
	if system == "" {
		system = "storage/log/system.log"
	}

	rotateFileConfig := rotatefile.NewConfigWith()
	rotateFileConfig.Filepath = system
	rotateFileConfig.RotateTime = rotatefile.EveryDay

	systemLog, err := rotateFileConfig.Create()
	if err != nil {
		Error(err)
		return nil
	}

	return systemLog
}

func new () {

	template := "[{{datetime}}] [{{level}}] [{{caller}}] {{message}}\n"

	path := config.Log["path"]
	if path == "" {
		path = "storage/log/zero.log"
	}

	log = slog.New().Configure(
		func(slogger *slog.Logger) {
			// reset
			slogger.ResetProcessors()
			slogger.ResetHandlers()
			slogger.CallerFlag = 6
			slogger.CallerSkip = 7
			fileHandler := handler.MustFileHandler(
				path,
				func(handlerConfig *handler.Config) {
					handlerConfig.RotateTime = rotatefile.EveryDay
					handlerConfig.Levels = slog.AllLevels
				},
			)
			fileHandler.SetFormatter(slog.NewTextFormatter(template))
			slogger.PushHandler(fileHandler)
		},
	)
	return
}
