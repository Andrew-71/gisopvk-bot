package app

import "git.a71.su/Andrew71/gisopvk-bot/internal/app/query"

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
