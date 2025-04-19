package httpport

// NOTE: DDD compels us to *triplicate* the types!!!

type Query struct {
	UUID string `json:"uuid"`
	Body int    `json:"body"`
}

type Reply struct {
	UUID string `json:"uuid"`
	Body int    `json:"body"`
}
