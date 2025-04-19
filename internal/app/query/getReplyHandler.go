package query

import (
	"context"
	"log/slog"

	"git.a71.su/Andrew71/gisopvk-bot/internal/common/decorator"
	"git.a71.su/Andrew71/gisopvk-bot/internal/domain"
)

// NOTE: CQS compels us to duplicate the types

type Query struct {
	UUID string
	Body string
}

func (q Query) ToDomain() domain.Query {
	return domain.Query{UUID: q.UUID, Body: q.Body}
}

type Reply struct {
	UUID string
	Body string
}

func fromDomain(r domain.Reply) Reply {
	return Reply{r.UUID, r.Body}
}

type GetReplyHandler decorator.QueryHandler[Query, Reply]

type getReplyHandler struct {
	bot domain.Bot
}

func NewGetReplyHandler(
	bot domain.Bot,
	logger *slog.Logger,
) GetReplyHandler {
	if bot == nil {
		panic("bot is nil")
	}
	return decorator.ApplyQueryDecorators[Query, Reply](
		getReplyHandler{bot: bot},
		logger,
	)
}

func (h getReplyHandler) Handle(ctx context.Context, query Query) (Reply, error) {
	reply, err := h.bot.Reply(query.ToDomain())
	if err != nil {
		return Reply{}, err
	}
	return fromDomain(reply), nil
}
