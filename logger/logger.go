package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"raddit/config"
)

var logger *zap.Logger
var slogger *zap.SugaredLogger

func Init(cfg *config.LogConfig) {
	logWriter := getWriter(cfg)
	encoder := getEncoder()

	core := zapcore.NewCore(encoder, logWriter, zapcore.DebugLevel)

	logger = zap.New(core)
	slogger = logger.Sugar()
}

func getWriter(cfg *config.LogConfig) zapcore.WriteSyncer {
	//file, _ := os.Create(filename)
	ljLogger := &lumberjack.Logger{
		Filename:  cfg.Filename,
		MaxSize:   cfg.MaxSize,
		MaxAge:    cfg.MaxAge,
		LocalTime: true,
	}
	return zapcore.AddSync(ljLogger)
}
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	return zapcore.NewConsoleEncoder(encoderConfig)
}
