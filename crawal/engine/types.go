package engine

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url     string
	Id      string
	Payload interface{}
}

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}
