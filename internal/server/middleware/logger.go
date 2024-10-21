package middleware

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
)

func NewLoggerInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			// Send a token with client requests.
			slog.Info("Request started", "procedure", req.Spec().Procedure)

			// 次のハンドラを呼び出して、レスポンスとエラーを取得
			res, err := next(ctx, req)

			// リクエスト終了時にログを出力
			if err != nil {
				slog.Error("Request failed", "procedure", req.Spec().Procedure, "error", err)
			} else {
				slog.Info("Request succeeded", "procedure", req.Spec().Procedure)
			}

			return res, err
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
