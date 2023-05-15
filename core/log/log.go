package log

import (
	"go-gin/config"

	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/gookit/slog/rotatefile"
)

type Log struct {}

var log = &slog.Logger{}

func Init () {
	log = newLogger()
	return
}

func Debug (message ...interface{}) {
	log.Debug(message)
	return
}

func Info (message ...interface{}) {
	log.Info(message)
	return
}

func Error (message ...interface{}) {
	log.Error(message)
	return
}

func Debugf (format string, message ...interface{}) {
	log.Debugf(format, message)
	return
}

func Infof (format string, message ...interface{}) {
	log.Infof(format, message)
	return
}

func Errorf (format string, message ...interface{}) {
	log.Errorf(format, message)
	return
}

// Example Initialize the system log
// 初始化系统日志
func InitSystemLog () *rotatefile.Writer {

	system := config.Log["system"].(string)
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

func newLogger () *slog.Logger {

	template := "[{{datetime}}] [{{level}}] [{{caller}}] {{message}}\n"

	path := config.Log["path"].(string)
	if path == "" {
		path = "storage/log/zero.log"
	}

	return slog.New().Configure(
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
}
