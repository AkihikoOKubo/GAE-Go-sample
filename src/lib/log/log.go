package log

import (
	"context"
	log "github.com/yfuruyama/stackdriver-request-context-log"
	"net/http"
)

func SetupContextLogger(r *http.Request) context.Context {
	return context.WithValue(r.Context(), "logger", log.RequestContextLogger(r))
}

func Logger(ctx context.Context) *log.ContextLogger {
	if ctx == nil {
		// ライブラリの仕様でContextLoggerがnilの場合にハンドリングできない
		// Severityを大きくして全てのロギングをスキップさせてエラーを起こさせない
		return &log.ContextLogger{
			Severity: 999,
		}
	}

	i := ctx.Value("logger")

	if i == nil {
		return &log.ContextLogger{
			Severity: 999,
		}
	}

	return i.(*log.ContextLogger)
}

