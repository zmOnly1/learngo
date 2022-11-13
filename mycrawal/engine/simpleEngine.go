package engine

import "log"

type Item struct {
	Url     string
	Id      int
	Payload interface{}
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Request struct {
	Url       string
	ParseFunc func(content []byte, url string) ParseResult
}

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			panic(err)
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			ItemSaver(item)
		}
	}
}

func ItemSaver(item Item) {
	log.Printf("got item: %v", item)
}

func worker(r Request) (ParseResult, error) {
	return ParseResult{}, nil
}
