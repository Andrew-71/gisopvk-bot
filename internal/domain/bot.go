package domain

type Bot interface {
	Reply(message Query) (Reply, error)
}
