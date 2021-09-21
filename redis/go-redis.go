package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"net/url"
)

//设置全局变量
var db = &sql.DB{}
var client2 *redis.Client

type Demo struct {
	Id   uint16
	User string
	Pwd  string
}

func init() {
	// 连mysql
	db, _ = sql.Open("mysql", "root:root@tcp(localhost:3306)/test?charset=utf8")
	// 连接redis
	client2 = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//测试ping 通，不通就报错
	pong, err := client2.Ping().Result()
	fmt.Println("redis link status", pong, err)
}

//从redis 查询 数据
func Detail(w http.ResponseWriter, r *http.Request) {
	var params url.Values = r.URL.Query()
	var sid string = params.Get("sid") //获取resquest 传来的id
	fmt.Println("sid=" + sid)
	var dataMap []Demo //创建一个 数据切片 存储数据
	redisData, _ := client2.HGet("demo", sid).Result()
	fmt.Println(redisData)
	//  w.Write([]byte(redisData))//把查询到的数据存在body返回
	w.Write([]byte(redisData))

	//json.Unmarshal([]byte(redisData), &dataMap) //json转go对象再使用
	if redisData == "" { //判断是否为空，如果空，说明redis中没有数据，则从数据库中查询
		fmt.Println("query....")
		rows, err := db.Query("SELECT * FROM demo where id=?", sid)
		fmt.Println("query:", err)
		for rows.Next() { //获得数据库记录，再存储在redis里面
			var id uint16
			var user string
			var pwd string
			rows.Scan(&id, &user, &pwd)     //scan函数把值取出
			dataMap = append(dataMap, Demo{ //存储在切片中
				Id:   id,
				User: user,
				Pwd:  pwd,
			})
		}
		tjson, _ := json.Marshal(dataMap)                    //先转化为json再存入redis
		bor, err := client2.HSet("demo", sid, tjson).Result() //存储在redis
		fmt.Println(bor, err)
		w.Write(tjson)
	}
	/*	t, _ := template.ParseFiles("./detail.html")
		t.Execute(w, map[string]interface{}{
			"DataMap": dataMap,
		})*/
}

func List(w http.ResponseWriter, r *http.Request) {
	var dataMap []Demo                                        //创建存储数据记录的切片
	rows, _ := db.Query("SELECT `id`,`user`,`pwd` FROM demo") //查询所有
	//创建 map 切片
	cacheMap := make(map[string]interface{})
	//迭代获得 数据
	for rows.Next() { //读取数据库记录
		var id uint16
		var user string
		var pwd string
		rows.Scan(&id, &user, &pwd)
		demo := Demo{
			Id:   id,
			User: user,
			Pwd:  pwd,
		}
		//	dataMap = demo
		dataMap = append(dataMap, demo)
		tjson, _ := json.Marshal(demo)
		cacheMap[fmt.Sprintf("%v", id)] = tjson
	}
	wjson, _ := json.Marshal(dataMap)
	w.Write(wjson)
	//批量生成 hset
	client2.HMSet("demo", cacheMap)
	//tjson,_ := json.Marshal(dataMap)
	//bcm := client2.HSet("demo",tjson)
	// -1 表示永远有效
	//client2.Set("demo",tjson,-1)
	//
	//client2.HSet("demo","all",tjson)

}

func ExampleNewclient2() {
	//建立连接
	//测试 ping通网络没有
	pong, err := client2.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

func Exampleclient2() {
	//fmt.Println(client2.HSet("key2","lj","kdkdk"))
	//存储值再redis中
	err := client2.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	// 读取 redis的值
	val, err := client2.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client2.HGet("key2", "lj").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

//存储数据库
func main() {
	//Exampleclient2()

	//把所有mysql数据都加载到redis
	http.HandleFunc("/list", List)

	//查询redis ,如果 redis存在数据就返回数据，不存在就从mysql查询，然后再存在redis
	http.HandleFunc("/detail", Detail)
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println(err)
	}
}
