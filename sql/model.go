package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
func main()  {
	//1. 连接数据库
     db, error := sql.Open("mysql","root:root@/minisql?charset=utf8")
     defer db.Close()
     checkError(error)

     //insert
     //编辑 sql语法准备
     stmt, error := db.Prepare("insert into user set userid=?,nickname=?,sex=?,create_time=?")
     checkError(error)
     //执行 sql操作
     res, error := stmt.Exec("2012154","张茨","男","20125145")
     fmt.Println(res.RowsAffected())

     //update
     stm, error := db.Prepare("update user set userid=? where create_time=?")
     checkError(error)
     res, error =stm.Exec("66654","6551414")
     fmt.Println(res.LastInsertId())

     //查询
     rows, err := db.Query("select * from user")
     checkError(err)


     //以json格式取出
     //先查询数据库 一条条信息 放到一个切片里面
     columns, err := rows.Columns() //返回[]string 以每条信息的形式
     checkError(err)
     var count int = len(columns) //计算有多少条信息
	//创建一个map切片 ，用来添加每个消息实体，消息实体也是一个切片
     tabledata := make([]map[string]interface{},0)
     //创建消息实体切片，用于存储一条信息实体
     values := make([]interface{},count)
     //用来存储 地址，便于rows一次性读取
     valuePtrs := make([]interface{},count)
	for rows.Next() {
		for i:=0;i<count ;i++  {
			valuePtrs[i] = &values[i] //把消息实体放进地址切片
		}
		rows.Scan(valuePtrs...) //一次性读取
		entry := make(map[string]interface{})

		for i,col:= range columns{
			var v interface{}
			val :=values[i]
			b,ok := val.([]byte)
			if ok {
				v =string(b)
			}else {
				v =val
			}
			entry[col] = v
		}
		tabledata = append(tabledata,entry)
	}
	//将切片转换成json格式
    jsonData,err :=json.Marshal(tabledata)
    checkError(err)
    var jsonString string= string(jsonData)
     fmt.Println(jsonString)



	/*for rows.Next() {
		//声明字段必须要和数据库字段对应
		var userid  string
		var nickname string
		var sex  string
		var create_time string
		error := rows.Scan(&userid,&nickname,&sex,&create_time)
		checkError(error)
		fmt.Println(userid)
		fmt.Println(nickname)
		fmt.Println(sex)
		//string 转 int64
		var num64 int64
		num64, _= strconv.ParseInt(create_time,10,20)
		time := time2.Unix(num64,0)//1970-01-07 09:38:07 +0800 CST
		fmt.Println(time)
		//按照指定格式打印
		timeString := time.Format("2006-01-02 15:04:05")
		fmt.Println(timeString)
	}*/



     //delete
	stmt, err = db.Prepare("delete from user where userid=?")
	checkError(err)
	res, err = stmt.Exec("2012154")
	checkError(err)

}

func checkError(err error)  {
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
}