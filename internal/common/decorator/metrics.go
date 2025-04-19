package decorator

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type MetricsClient interface {
	Inc(key string, value float64)
}

type queryMetricsDecorator[C any, R any] struct {
	base   QueryHandler[C, R]
	client MetricsClient
}

func (d queryMetricsDecorator[C, R]) Handle(ctx context.Context, query C) (result R, err error) {
	start := time.Now()

	actionName := strings.ToLower(generateActionName(query))

	defer func() {
		end := time.Since(start)

		d.client.Inc(fmt.Sprintf("queries.%s.duration", actionName), end.Seconds())

		if err == nil {
			d.client.Inc(fmt.Sprintf("queries.%s.success", actionName), 1)
		} else {
			d.client.Inc(fmt.Sprintf("queries.%s.failure", actionName), 1)
		}
	}()

	return d.base.Handle(ctx, query)
}
