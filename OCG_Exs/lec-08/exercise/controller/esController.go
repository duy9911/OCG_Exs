package controller

import (
	"context"
	"exercise/model"
	"fmt"
	"reflect"
	"strconv"

	"github.com/olivere/elastic/v7"
)

func AddToEs() {
	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		// Handle error
		panic(err)
	}

	info, code, err := client.Ping("http://localhost:9200").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	lines, err := model.ReadCsv("./test.csv")
	if err != nil {
		panic(err)
	}

	for i, line := range lines {
		data := model.CsvLine{
			Type:  line[0],
			Title: line[1],
			Body:  line[2],
		}
		put1, err := client.Index().
			Index("train").
			Id(strconv.Itoa(i)).
			BodyJson(data).
			Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}

		fmt.Printf("Indexedline %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	}

	fmt.Println("Done...!")

}

func QueryEs(field string, val string) {
	ctx := context.Background()
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		// Handle error
		panic(err)
	}
	termQuery := elastic.NewTermQuery(field, val)
	searchResult, err := client.Search().
		Index("train").   // search in index "train"
		Query(termQuery). // specify the query
		Do(ctx)           // execute
	if err != nil {
		// Handle error
		panic(err)
	}

	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	fmt.Printf("Query with term {%v: %v} took %d milliseconds\n", field, val, searchResult.TookInMillis)
	var ttyp model.CsvLine
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if t, ok := item.(model.CsvLine); ok {
			fmt.Printf("Filed by %s: %s\n", t.Body, t.Type)
		}
	}
	// TotalHits is another convenience function that works even when something goes wrong.
	fmt.Printf("Found a total of %d field\n", searchResult.TotalHits())
}
