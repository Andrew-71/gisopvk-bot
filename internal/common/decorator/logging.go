package decorator

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
)

func generateActionName(handler any) string {
	return strings.Split(fmt.Sprintf("%T", handler), ".")[1]
}

type queryLoggingDecorator[C any, R any] struct {
	base   QueryHandler[C, R]
	logger *slog.Logger
}

func (d queryLoggingDecorator[C, R]) Handle(ctx context.Context, cmd C) (result R, err error) {
	logger := d.logger.With(
		slog.String("query", generateActionName(cmd)),
		slog.String("query_body", fmt.Sprintf("%v", cmd)),
	)

	logger.Debug("Executing query")
	defer func() {
		if err == nil {
			logger.Info("Query executed successfully")
		} else {
			logger.Error("Failed to execute query", "error", err.Error())
		}
	}()

	return d.base.Handle(ctx, cmd)
}
