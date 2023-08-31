package util

import (
	"context"
	"fmt"
	"github.com/atlas-api-helper/util/constants"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCast64(t *testing.T) {
	i := 42
	result := Cast64(&i)
	assert.NotNil(t, result)
	assert.Equal(t, int64(42), *result)
}

func TestFatalLog(t *testing.T) {
	Fatalf(getContextWithTraceId(), "Test log %s", "message")
}

func getContextWithTraceId() context.Context {
	traceID := fmt.Sprintf("TraceID-%d", time.Now().UnixNano())
	ctx := context.Background()
	ctx = context.WithValue(ctx, constants.TraceID, traceID)
	return ctx
}
