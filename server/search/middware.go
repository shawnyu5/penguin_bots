package search

import "log"

type loggingMiddleware struct {
	logger log.Logger
	next   SearchService
}

func (lm *loggingMiddleware) SearchByRegex(p *Product) ([]Product, error) {
	lm.logger.Printf("Searching for product: %s", p.Title)
	return lm.next.SearchByRegex(p)
}
