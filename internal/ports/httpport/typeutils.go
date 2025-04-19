package httpport

import (
	"github.com/Andrew-71/gisopvk-bot/internal/app/query"
	"github.com/google/uuid"
)

func (q Query) FromApi() query.Query {
	return query.Query{
		UUID: q.Uuid.String(),
		Body: q.Body,
	}
}

func mapReplyToApi(r query.Reply) Reply {
	return Reply{
		Uuid: uuid.MustParse(r.UUID),
		Body: r.Body,
	}
}
