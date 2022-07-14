package bleve_t

import (
	"fmt"
	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/analysis/analyzer/keyword"
	"os"
	"time"
)

// var indexMapping mapping.IndexMapping
var exampleIndex bleve.Index
var err error

func init() {
	Client()
}
func Client() {
	// var index bleve.Index

	// index, err := bleve.New("example.bleve", mapping)
	file, _ := os.Stat("example.bleve")
	/*if err != nil {
		fmt.Println(err)
		return
	}*/
	if file != nil && file.IsDir() {
		exampleIndex, err = bleve.Open("example.bleve")
	} else {
		mapping := bleve.NewIndexMapping()
		mapping.DefaultAnalyzer = keyword.Name
		exampleIndex, err = bleve.New("example.bleve", mapping)
	}
	if err != nil {
		fmt.Println(err)
		return
	}

	/*data := struct {
		Name string
	}{
		Name: "characters that are undesirable for indexing are replaced with whitespace. This allows the original byte offsets in the",
	}

	// index some data
	index.Index("test", data)*/
	// index.Index("test", "characters that are undesirable for indexing are replaced with whitespace. This allows the original byte offsets in the")
	//index.Index("test2", "replaced with whitespace. This allows the original byte offsets in the")
	// search for some text
}

type Data struct {
	Name    string    `json:"name"`
	Age     int64     `json:"age"`
	Ctime   time.Time `json:"ctime"`
	KeyWord string    `json:"key_word"`
}

func Insert() {
	data1 := Data{Name: "zhangsan", Age: 1, Ctime: time.Now(), KeyWord: "newwpp.docer.wps.cn"}
	data2 := Data{Name: "lisi", Age: 2, Ctime: time.Now(), KeyWord: "jimo-mb-api.docer.wps.cn"}
	data3 := Data{Name: "chenli", Age: 3, Ctime: time.Now(), KeyWord: "jimo-beautify.docer.wps.cn"}
	exampleIndex.Index("data1", data1)
	exampleIndex.Index("data2", data2)
	exampleIndex.Index("data3", data3)
	count, _ := exampleIndex.DocCount()
	fmt.Println(count)

	query := bleve.NewMatchAllQuery()
	// query := bleve.NewQueryStringQuery("re")
	search := bleve.NewSearchRequest(query)
	//search.Highlight =  bleve.NewHighlight()
	//search.Highlight.AddField("name")
	searchResults, err := exampleIndex.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	/*for _, hit := range searchResults.Hits {
		for s, i := range hit.Fields {
			fmt.Println(s, i)
		}
	}*/
	fmt.Println(searchResults)
}

func Search()  {
	query := bleve.NewWildcardQuery("*jimo-beautify*")
	// query := bleve.NewQueryStringQuery("re")
	search := bleve.NewSearchRequest(query)
	//search.Highlight =  bleve.NewHighlight()
	//search.Highlight.AddField("name")
	searchResults, err := exampleIndex.Search(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	/*for _, hit := range searchResults.Hits {
		for s, i := range hit.Fields {
			fmt.Println(s, i)
		}
	}*/
	fmt.Println(searchResults)

}