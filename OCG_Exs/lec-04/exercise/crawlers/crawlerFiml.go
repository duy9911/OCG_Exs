package crawlers

import (
	"fmt"
	"lec-04/exercise/database"
	"log"
	"strings"
	"sync"

	"github.com/gocolly/colly"
)

type Film struct {
	Id     int64 `gorm:"primaryKey"`
	Title  string
	Link   string
	Rating string
}

//create chanel d to communication between  2 threads
//using waitgroup to make sure both goroutine will be finish
func CrawlerFilm(wg *sync.WaitGroup) {
	url := "https://www.imdb.com/chart/top/?ref_=nv_mv_250"
	d := make(chan Film)
	wg.Add(2)
	go crawl(wg, d, url) // thread 1
	go toDB(wg, d)       // thread 2
	wg.Wait()            // wait two threads
	log.Println("Successfully write to DB")
}

// crawler sender
func crawl(wg *sync.WaitGroup, fiml chan Film, url string) {

	defer wg.Done()
	defer close(fiml)

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("OOp~ Sthg went wrong bro", err)
	})

	c.OnHTML(".lister-list tr", func(e *colly.HTMLElement) {

		data := Film{}
		title := e.ChildText(".titleColumn a")
		data.Title = strings.Join(strings.Fields(title), " ") //remove whitespace between each words
		data.Link = "https://www.imdb.com" + e.ChildAttr("a", "href")
		data.Rating = e.ChildText(".ratingColumn strong")
		fmt.Printf("Title: %v", data.Title)
		fiml <- data
	})
	c.Visit(url)
}

//receiver values from crawler sender
func toDB(wg *sync.WaitGroup, d chan Film) {
	defer wg.Done()

	db := database.ConnectDb()

	if (!db.Migrator().HasTable(&Film{})) {
		db.Migrator().CreateTable(&Film{})
	}

	for {
		v, ok := <-d
		if !ok {
			return
		}
		user := Film{
			Title:  v.Title,
			Link:   v.Link,
			Rating: v.Rating,
		}
		db.Debug().Create(&user)
	}
}
