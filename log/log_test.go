package log

import (
	"context"
	"testing"
)

func TestLog(t *testing.T) {
	ctx := WithRequestId(context.Background(), "1234")
	Info(ctx, "hello world")

	newCtx := context.Background()
	Info(newCtx, "hello world")

	modelNameCtx := WithModelName(ctx, "CodeLlama")
	Info(modelNameCtx, "hello world")
}
