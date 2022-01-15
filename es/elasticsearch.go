package es

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/olivere/elastic/v7"
)

// Elasticsearch demo
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}

func EsClient() {
	client, err := elastic.NewClient(elastic.SetSniff(false),elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Println("connect to es success")
	p1 := Person{Name: "lmh", Age: 18, Married: false}
	put1, err := client.Index().
		Index("user").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
	/*---------*/
	// create()
	// delete()
	// update()
	// gets()
	// query()
	// list(1, 3)
}

/*-------------------------*/
var client *elastic.Client
var host = "http://127.0.0.1:9200/"

type Employee struct {
	Id        string   `json:"id"`
	FirstName string   `json:"first_name"`
	Sex       string   `json:"sex"`
	LastName  string   `json:"last_name"`
	Age       int      `json:"age"`
	About     string   `json:"about"`
	Interests []string `json:"interests"`
	Ctime     int64    `json:"ctime"`
}

const mapping = `
		{
        	"settings":{
        		"number_of_shards": 1,
        		"number_of_replicas": 0
        	},
        	"mappings":{
				"properties":{
					"id":{
						"type":"text"
					},
					"first_name":{
						"type":"text"
					},
                    "sex":{
						"type":"keyword"
					},
					"last_name":{
						"type":"text"
					},
					"age":{
						"type":"long"
					},
					"about":{
						"type":"text"
					},
					"interests": {
						"type":"text"
					}, 
					"ctime": {
                          "type": "date"
                    }
				}
        	}
        }
	`

//初始化
func init() {
	errorlog := logrus.New()
	var err error
	client, err = elastic.NewClient(elastic.SetSniff(false),elastic.SetErrorLog(errorlog), elastic.SetURL(host))
	if err != nil {
		panic(err)
	}
	info, code, err := client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

}

/*下面是简单的CURD*/
var empIndex = "employer"
var empType = "emploType"

//创建
func create() {

	//使用结构体
	e1 := Employee{"", "Jane", "y", "Smith 1", 32, "I like to collect rock albums", []string{"music"}, time.Now().UnixNano()}
	/*isExist, err := client.IndexExists(empIndex).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(isExist)
	if !isExist {
		CreateResult, err := client.CreateIndex(empIndex).BodyString(mapping).Do(context.TODO())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(
			CreateResult)
	} else {
		CreateResult, err := client.DeleteIndex(empIndex).Do(context.TODO())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(
			CreateResult)
	}*/
	// 当向一个不存在的索引中写入文档时，会自动创建索引。
	put1, err := client.Index().
		Index(empIndex).
		BodyJson(e1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put1.Id, put1.Index, put1.Type)

	//使用字符串
	e2 := `{"first_name":"John","sex":"x","last_name":"Smith kk","age":25,"about":"I love to go rock climbing","interests":["sports","music"],"ctime":2}`
	put2, err := client.Index().
		Index(empIndex).
		BodyJson(e2).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put2.Id, put2.Index, put2.Type)

	e3 := `{"first_name":"Douglas kk","sex":"x","last_name":"Fir","age":5,"about":"I like to build cabinets","interests":["forestry"],"ctime":12}`
	put3, err := client.Index().
		Index(empIndex).
		BodyJson(e3).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put3.Id, put3.Index, put3.Type)

}

//删除
func delete() {

	res, err := client.Delete().Index("megacorp").
		Type("employee").
		Id("1").
		Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n", res.Result)
}

//修改
func update(e *Employee) {
	e.Ctime = 100
	res, err := client.Update().
		Index(empIndex).
		Id("10uUOHcB2NMlvAKAP7aD").
		Doc(e).
		Do(context.Background())
	if err != nil {
		println(err.Error())
	}

	fmt.Printf("update age %s\n", res.Result)

}

//查找
func gets() *Employee {
	//通过id查找
	get1, err := client.Get().Index(empIndex).Id("10uUOHcB2NMlvAKAP7aD").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
		var typ Employee
		err := json.Unmarshal(get1.Source, &typ)
		if err != nil {
			fmt.Println(err)
		}
		typ.Id = get1.Id
		return &typ
	}
	return nil
}

//搜索
func query() {
	var res *elastic.SearchResult
	var err error
	//取所有
	// true 升序
	// matchPhraseQuery := elastic.NewMatchPhraseQuery("about", "rock climbing")
	// query := elastic.NewMatchQuery("about", "kk")
	// query2 := elastic.NewMatchQuery("last_name", "kk").Fuzziness("AUTO").Operator("or")
	// 创建bool查询
	// tempQuery := elastic.NewTermQuery("last_name", "Smith")
	//tempQuery2 := elastic.NewTermQuery("about", "like")
	boolQuery := elastic.NewBoolQuery()
	rangeQuery := elastic.NewRangeQuery("ctime")
	rangeQuery.Gt(0)
	rangeQuery.Lt(1611313273576438041)
	//boolQuery.Must(rangeQuery)
	boolQuery.Must(elastic.NewMatchQuery("sex", "y"))
	boolQuery.Must(elastic.NewMatchQuery("ctime", "12"))

	boolQuery.Should(elastic.NewMatchQuery("age", "35")) // 只要有must 后面should 语句条件全部失效
	boolQuery.Should(elastic.NewMatchQuery("last_name", "kk")).MinimumNumberShouldMatch(1)
	boolQuery.Filter(rangeQuery)
	res, err = client.Search(empIndex).Query(boolQuery).From(0).Size(10).Do(context.Background())
	printEmployee(res, err)
	//字段相等
	/*	q := elastic.NewQueryStringQuery("last_name:Smith")
		res, err = client.Search("megacorp").Type("employee").Query(q).Do(context.Background())
		if err != nil {
			println(err.Error())
		}
		printEmployee(res, err)

		//条件查询
		//年龄大于30岁的
		boolQ := elastic.NewBoolQuery()
		boolQ.Must(elastic.NewMatchQuery("last_name", "smith"))
		boolQ.Filter(elastic.NewRangeQuery("age").Gt(30))
		res, err = client.Search("megacorp").Type("employee").Query(q).Do(context.Background())
		printEmployee(res, err)

		//短语搜索 搜索about字段中有 rock climbing
		matchPhraseQuery := elastic.NewMatchPhraseQuery("about", "rock climbing")
		res, err = client.Search("megacorp").Type("employee").Query(matchPhraseQuery).Do(context.Background())
		printEmployee(res, err)

		//分析 interests
		aggs := elastic.NewTermsAggregation().Field("interests")
		res, err = client.Search("megacorp").Type("employee").Aggregation("all_interests", aggs).Do(context.Background())
		printEmployee(res, err)*/

}

//简单分页
func list(size, page int) {
	if size < 0 || page < 1 {
		fmt.Printf("param error")
		return
	}
	res, err := client.Search(empIndex).
		TypedKeys(true).
		Size(size).
		From((page - 1) * size).
		Do(context.Background())
	printEmployee(res, err)

}

//打印查询到的Employee
func printEmployee(res *elastic.SearchResult, err error) {
	if err != nil {
		print(err.Error())
		return
	}
	// var typ Employee
	/*	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(Employee)
		fmt.Printf("%#v\n", t)
		//fmt.Printf("%v\n", t)
	}*/
	total := res.TotalHits()
	fmt.Println(total)
	for _, hit := range res.Hits.Hits {
		var typ Employee
		err := json.Unmarshal(hit.Source, &typ)
		if err != nil {
			fmt.Println(err)
		}
		typ.Id = hit.Id
		fmt.Printf("%#v\n", typ)
	}
}

func getRecordInfo() {

}
