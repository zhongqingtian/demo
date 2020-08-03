package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

//Conn接口是与Redis协作的主要接口，可以使用Dial,DialWithTimeout或者NewConn函数来创建连接，当任务完成时，应用程序必须调用Close函数来完成操作。

/*func main()  {
	conn,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect redis error :",err)
		return
	}
	defer conn.Close()
}*/
func main() {
	//连接 redis
	c, err := redis.Dial("tcp", "localhost:6369")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()
	// 存储 数据到redis 有效时间5秒
	_, err = c.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	//获得 redis 数据
	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	time.Sleep(8 * time.Second)
	//过了 8 秒在再获取数据
	username, err = redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
}
