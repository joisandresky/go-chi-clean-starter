package middleware

import (
	"net/http"

	"go.uber.org/zap"
)

type TestMiddleware interface {
	TestMiddleware(next http.Handler) http.Handler
}

type testMw struct {
	logger *zap.SugaredLogger
}

func NewTestMiddleware(logger *zap.SugaredLogger) TestMiddleware {
	return &testMw{
		logger: logger,
	}
}

func (mw *testMw) TestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mw.logger.Info("test middleware")
		next.ServeHTTP(w, r)
	})
}
