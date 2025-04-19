package app

import "github.com/Andrew-71/gisopvk-bot/internal/app/query"

type Application struct {
	Commands Commands
	Queries  Queries
}

// Technically, we have CQS
// ...but our app has no commands :D
type Commands struct {
}

type Queries struct {
	GetReply query.GetReplyHandler
}
