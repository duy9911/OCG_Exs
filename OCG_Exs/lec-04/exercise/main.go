package main

import (
	c "lec-04/exercise/crawlers"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	c.CrawlerFilm(&wg)
	// c.CrawlerProduct(&wg)
}
