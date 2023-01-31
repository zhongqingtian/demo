package clickhouse

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"time"
)

var dbch driver.Conn
var ctxch context.Context

type User struct {
	Id    uint8     `ch:"id"` // 默认是名称字段
	Name  string    `ch:"name"`
	Age   int64     `ch:"age"`
	Ids   []uint8   `ch:"ids"`
	Ctime time.Time `ch:"ctime"`
}

type Example struct {
	Id    uint64    `ch:"id"` // 默认是名称字段
	Name  string    `ch:"name"`
	Age   int64     `ch:"age"`
	Ids   []uint8   `ch:"ids"`
	Ctime time.Time `ch:"ctime"`
}

func init() {
	var err error
	dbch, err = clickhouse.Open(&clickhouse.Options{
		Addr: []string{"10.13.128.139:9090"},
		Auth: clickhouse.Auth{
			Database: "temp",
			Username: "root",
			Password: "092j3AnV",
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: 5 * time.Second,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		Debug: true,
		/*Debugf: func(format string, v ...interface{}) {
			fmt.Printf(format, v)
		},*/
		//Protocol:         clickhouse.HTTP,
		/*MaxOpenConns:     5,
		MaxIdleConns:     5,
		ConnMaxLifetime:  time.Duration(10) * time.Minute,
		ConnOpenStrategy: clickhouse.ConnOpenInOrder,
		BlockBufferSize:  10,*/
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	ctxch = clickhouse.Context(context.Background(), clickhouse.WithSettings(clickhouse.Settings{
		"max_block_size": 10,
	}))
	fmt.Println(dbch.Ping(context.Background()))

	/*sql := fmt.Sprintf("SELECT `name`,`age` FROM test")
	fmt.Println("数据库sql:", sql)
	var users []User
	err = dbch.Select(ctxch, &users, sql)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(users)*/

	/*row := dbch.QueryRow(context.Background(), "SELECT `name`,`age` FROM test WHERE id=1")
	var (
		name string
		age  int32
	)
	fmt.Println(row.Scan(&name, &age))
	fmt.Println(name, age)*/
}

func Select(read User) ([]User, error) {
	sql := "SELECT * FROM test"
	var whereStr []string
	var params []interface{}
	if read.Id > 0 {
		whereStr = append(whereStr, "id=?")
		params = append(params, read.Id)
	}
	if read.Name != "" {
		whereStr = append(whereStr, "name=?")
		params = append(params, read.Name)
	}
	if read.Age > 0 {
		whereStr = append(whereStr, "age=?")
		params = append(params, read.Age)
	}
	fmt.Println("数据库sql:", sql)
	var users []User
	err := dbch.Select(ctxch, &users, sql, params...)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(users)
	return users, nil
}

func QueryRow() (User, error) {
	var user User
	err := dbch.QueryRow(ctxch, "SELECT * FROM test where id=1").ScanStruct(&user)
	if err != nil {
		fmt.Println(err.Error())
		return User{}, err
	}
	fmt.Println(user)
	return user, nil
}

// 预处理
func QueryWithParameters() error {
	chCtx := clickhouse.Context(context.Background(), clickhouse.WithParameters(clickhouse.Parameters{
		"id": "3",
		//"name":     "zhangsan",
		//"age":      "99",
		"database": "temp",
		"table":    "test",
	}))

	row := dbch.QueryRow(chCtx, "SELECT {id:UInt8}, {name:String} , {age:Int32} FROM  {database:Identifier}.{table:Identifier}")
	var (
		col1 uint8
		col2 string
		col3 int32
	)
	if err := row.Scan(&col1, &col2, &col3); err != nil {
		return err
	}
	fmt.Printf("row: col1=%d, col2=%s, col3=%d\n", col1, col2, col3)
	return nil
}

func QueryParameter() (uint64, error) {
	var count uint64
	// positional bind
	if err := dbch.QueryRow(ctxch, "SELECT count() FROM test WHERE id > ?", 2).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}

// 异步插入(测试不成功)
func AsyncInsert(users []User) error {
	for _, user := range users {
		err := dbch.AsyncInsert(ctxch, fmt.Sprintf("INSERT INTO test (`id`,`name`,`age`,`ids`,`ctime`)VALUES (%d,'%s',%d,'[1,3,2,8,89]','2021-01-01 16:26:19')", user.Id, user.Name, user.Age), true)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}
	return nil
}

// 异步插入(测试不成功)
func AsyncInsertExample(user Example) error {
	err := dbch.AsyncInsert(ctxch, fmt.Sprintf("INSERT INTO example (`id`,`name`,`age`,`ids`,`ctime`)VALUES (%d,'%s',%d,'[1,3,2,8,89]','2021-01-01 16:26:19')", user.Id, user.Name, user.Age), true)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func BatchInsertExample(examples []Example) error {
	batch, err := dbch.PrepareBatch(ctxch, "INSERT INTO example (`id`,`name`,`age`,`ids`,`ctime`)")
	if err != nil {
		return err
	}

	for _, user := range examples {
		err = batch.Append(user.Id, user.Name, user.Age, user.Ids, user.Ctime)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}

	err = batch.Send()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// 批量插入
func BatchInsert(users []User) error {
	batch, err := dbch.PrepareBatch(ctxch, "INSERT INTO test(`id`,`name`,`age`,`ids`,`ctime`)")
	if err != nil {
		return err
	}

	for _, user := range users {
		err = batch.Append(user.Id, user.Name, user.Age, user.Ids, user.Ctime)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
	}

	err = batch.Send()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// 指定插入列(不支持更新)
func ColumnInsert(users []User) error {
	batch, err := dbch.PrepareBatch(context.Background(), "INSERT INTO test")
	if err != nil {
		return err
	}
	var ( // 全部列要填写
		col1 []uint8
		col2 []string
		col3 []int64
		col4 [][]uint8
		col5 []time.Time
	)
	for _, user := range users {
		col1 = append(col1, user.Id)
		col2 = append(col2, user.Name)
		// col3 = append(col3, []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9})
		col3 = append(col3, user.Age)
		col4 = append(col4, user.Ids)
		col5 = append(col5, user.Ctime)
	}
	if err := batch.Column(0).Append(col1); err != nil {
		fmt.Println(err.Error())
		return err
	}
	if err := batch.Column(1).Append(col2); err != nil {
		fmt.Println(err.Error())
		return err
	}
	if err := batch.Column(2).Append(col3); err != nil {
		fmt.Println(err.Error())
		return err
	}
	if err := batch.Column(3).Append(col4); err != nil {
		fmt.Println(err.Error())
		return err
	}
	if err := batch.Column(4).Append(col5); err != nil {
		fmt.Println(err.Error())
		return err
	}
	if err := batch.Send(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func CreateTable() error {
	dbch.Exec(ctxch, `DROP TABLE IF EXISTS example`)
	const ddl = `
		CREATE TABLE example (
			  id UInt64,
			  name String,
			  ids Array(UInt8),
			  ctime DateTime
		) ENGINE = Memory
	`
	if err := dbch.Exec(ctxch, ddl); err != nil {
		return err
	}
	return nil
}
