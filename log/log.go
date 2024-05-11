package log

import (
	"context"
	"go.uber.org/zap/zapcore"

	log "go.uber.org/zap"
)

type requestKey struct{}
type modelNameKey struct{}

func init() {
	cfg := log.NewProductionConfig()
	cfg.Level.SetLevel(zapcore.InfoLevel)
	cfg.OutputPaths = []string{"stderr"}
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05 GMT")
	logger, _ := cfg.Build()
	log.ReplaceGlobals(logger)
}

func WithRequestId(ctx context.Context, requestId string) context.Context {
	return context.WithValue(ctx, requestKey{}, requestId)
}

func WithModelName(ctx context.Context, modelName string) context.Context {
	return context.WithValue(ctx, modelNameKey{}, modelName)
}

func Info(ctx context.Context, msg string, fields ...log.Field) {
	if requestId, err := ctx.Value(requestKey{}).(string); err != false {
		fields = append(fields, log.String("requestId", requestId))
	}

	if modelName, err := ctx.Value(modelNameKey{}).(string); err != false {
		fields = append(fields, log.String("modelName", modelName))
	}
	log.L().Info(msg, fields...)

}
