package mongdb

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/tag"
	"time"
)

var mgoCli *mongo.Client

func initEngine() {
	var err error
	// clientOptions := options.Client().ApplyURI("mongodb://admin:123@127.0.0.1:27011/?replicaSet=rs0")
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s/?replicaSet=%s", "localhost:37017", "rs"))

	// 连接到MongoDB
	mgoCli, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logrus.Fatal(err)
	}
	// 检查连接
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		logrus.Fatal(err)
	}
}
func GetMgoCli() *mongo.Client {
	if mgoCli == nil {
		initEngine()
	}
	return mgoCli
}

func Connect() {
	var (
		client     = GetMgoCli()
		db         *mongo.Database
		collection *mongo.Collection
	)
	//2.选择数据库 my_db
	db = client.Database("my_db")

	//选择表 my_collection
	collection = db.Collection("my_collection")
	collection = collection
}

type TimePorint struct {
	StartTime int64 `bson:"startTime"` //开始时间
	EndTime   int64 `bson:"endTime"`   //结束时间
}
type LogRecord struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	JobName string             `bson:"jobName"` //任务名
	Command string             `bson:"command"` //shell命令
	Age     int64              `bson:"age"`
	Err     string             `bson:"err"`     //脚本错误
	Content string             `bson:"content"` //脚本输出
	Tp      TimePorint         `bson:"tp"`      //执行时间
	Keys    []string           `bson:"keys"`
	Tags    []*tag.Tag         `bson:"tags"`
}

func Insert() {
	var (
		client     = GetMgoCli()
		err        error
		collection *mongo.Collection
		//lr         *LogRecord
		iResult *mongo.InsertOneResult
		id      primitive.ObjectID
	)
	lr := &LogRecord{
		JobName: "test",
		Command: "ding ding dang",
		Age:     12,
		Err:     "",
		Content: "这个是个小测试",
		Tp:      TimePorint{StartTime: time.Now().UnixNano()},
	}
	//2.选择数据库 my_db里的某个表
	collection = client.Database("my_db").Collection("my_collection")
	//插入某一条数据
	if iResult, err = collection.InsertOne(context.TODO(), lr); err != nil {
		fmt.Print(err)
		return
	}
	//_id:默认生成一个全局唯一ID
	id = iResult.InsertedID.(primitive.ObjectID)
	fmt.Println("自增ID", id.Hex())
}

func Minsert() error {
	var (
		client     = GetMgoCli()
		err        error
		collection *mongo.Collection
		result     *mongo.InsertManyResult
		id         primitive.ObjectID
	)
	collection = client.Database("my_db").Collection("test")

	//批量插入
	result, err = collection.InsertMany(context.TODO(), []interface{}{
		LogRecord{
			JobName: "work11",
			Command: "eo",
			Err:     "",
			Age:     0,
			Content: "8",
			Tp: TimePorint{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
			Tags: []*tag.Tag{{
				Name:  "app",
				Value: "高",
			}, {Name: "team", Value: "欧洲"}},
			Keys: []string{"12", "123", "1234"},
		},
		LogRecord{
			JobName: "work12",
			Command: "eco 2",
			Age:     0,
			Err:     "",
			Content: "46",
			Tp: TimePorint{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
			Tags: []*tag.Tag{{
				Name:  "app",
				Value: "低",
			}, {Name: "team", Value: "美州"}},
			Keys: []string{"12", "123", "12564"},
		},
		LogRecord{
			JobName: "work13",
			Command: "echo 95",
			Age:     11000,
			Err:     "",
			Content: "5",
			Tp: TimePorint{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
			Tags: []*tag.Tag{{
				Name:  "app",
				Value: "中",
			}, {Name: "team", Value: "澳洲"}, {Name: "machine", Value: "phy"}},
			Keys: []string{"1", "2", "123", "5687"},
		},
		LogRecord{
			JobName: "kpi",
			Command: "echo 47",
			Age:     12222,
			Err:     "",
			Content: "9",
			Tp: TimePorint{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
			Tags: []*tag.Tag{{
				Name:  "app",
				Value: "高",
			}, {Name: "team", Value: "欧洲"}, {Name: "machine", Value: "huawei"}},
			Keys: []string{"12", "123", "1234", "123456"},
		},
	})
	if err != nil {
		return err
	}
	if result == nil {
		return fmt.Errorf("result nil")
	}
	for _, v := range result.InsertedIDs {
		id = v.(primitive.ObjectID)
		fmt.Println("自增ID", id.Hex())
	}
	return nil
}

//查询实体
type FindByJobName struct {
	JobName string `bson:"jobName"` //任务名
}

func Search() {
	var (
		client     = GetMgoCli()
		err        error
		collection *mongo.Collection
		cursor     *mongo.Cursor
	)
	//2.选择数据库 my_db里的某个表
	collection = client.Database("my_db").Collection("test")

	//如果直接使用 LogRecord{JobName: "job10"}是查不到数据的，因为其他字段有初始值0或者“”
	cond := bson.D{}
	options := &options.FindOptions{Sort: bson.M{"age": -1}}
	//按照jobName字段进行过滤jobName="job10",翻页参数0-2

	if cursor, err = collection.Find(context.TODO(), cond, options); err != nil {
		fmt.Println(err)
		return
	}
	//延迟关闭游标
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			logrus.Error(err)
		}
	}()

	//遍历游标获取结果数据
	for cursor.Next(context.TODO()) {
		var lr LogRecord
		//反序列化Bson到对象
		if cursor.Decode(&lr) != nil {
			fmt.Print(err)
			return
		}
		//打印结果数据
		fmt.Printf("%v \n", lr)
	}

	//这里的结果遍历可以使用另外一种更方便的方式：
	/*var results []LogRecord
	if err = cursor.All(context.TODO(), &results); err != nil {
		logrus.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}*/
}

// 删除
func DropColl() {
	client := GetMgoCli()
	//2.选择数据库 my_db里的某个表
	collection := client.Database("my_db").Collection("test")
	//插入某一条数据
	if err := collection.Drop(context.Background()); err != nil {
		fmt.Print(err)
		return
	}
}

func Find() {
	var (
		client     = GetMgoCli()
		err        error
		collection *mongo.Collection
		cursor     *mongo.Cursor
	)
	//2.选择数据库 my_db里的某个表
	collection = client.Database("my_db").Collection("test")
	//如果直接使用 LogRecord{JobName: "job10"}是查不到数据的，因为其他字段有初始值0或者“”
	// filter := bson.M{"age": bson.M{"$lt": 32}, "command": bson.M{"$eq": "eo"}}
	//按照jobName字段进行过滤jobName="job10",翻页参数0-2

	// 查询条件
	/*	filter := bson.D{
		{"$and", bson.A{
			bson.M{"_id": id},
			bson.M{"deleted": false}, // 排除已删除的业务
		}},
	}*/
	/*id1, err := primitive.ObjectIDFromHex("601a09c366dab06f29ab73c0")
	fmt.Println(err)
	id2, err := primitive.ObjectIDFromHex("601a09c366dab06f29ab73c2")
	fmt.Println(err)*/
	// filter := bson.M{"_id": bson.M{"$in": []primitive.ObjectID{id1, id2}}}
	/*ids := make([]primitive.ObjectID, 0)
	ids = append(ids, id1)
	ids = append(ids, id2)*/
	//	filter := bson.D{{"_id", bson.M{"$in": ids}}}
	filter := bson.M{"age": 0}
	//options := &options.FindOptions{Sort: bson.M{"age":1}}
	if cursor, err = collection.Find(context.TODO(), filter, nil); err != nil {
		fmt.Println(err)
		return
	}
	//延迟关闭游标
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			logrus.Error(err)
		}
	}()

	//遍历游标获取结果数据
	for cursor.Next(context.TODO()) {
		var lr LogRecord
		//反序列化Bson到对象
		if cursor.Decode(&lr) != nil {
			fmt.Print(err)
			return
		}
		//打印结果数据
		fmt.Println(lr)
	}

}

func Delete(id string) error {
	client := GetMgoCli()
	coll := client.Database("my_db").Collection("test")
	objId, err := primitive.ObjectIDFromHex(id)
	fmt.Println(err)
	filter := bson.D{{"_id", objId}}
	/*filter := bson.D{{"_id", bson.M{"$ne": ""}}} // 全部删除
	res, err := coll.DeleteMany(context.Background(), filter)*/
	res, err := coll.DeleteOne(context.Background(), filter)
	fmt.Println(res)
	fmt.Println(err)
	return err
}

func FindByTag() {
	var (
		client     = GetMgoCli()
		err        error
		collection *mongo.Collection
		cursor     *mongo.Cursor
	)
	//2.选择数据库 my_db里的某个表
	collection = client.Database("my_db").Collection("test")
	//如果直接使用 LogRecord{JobName: "job10"}是查不到数据的，因为其他字段有初始值0或者“”
	// filter := bson.M{"age": bson.M{"$lt": 32}, "command": bson.M{"$eq": "eo"}}
	//按照jobName字段进行过滤jobName="job10",翻页参数0-2
	filter := bson.M{"tags.name": "team"}
	if cursor, err = collection.Find(context.TODO(), filter); err != nil {
		fmt.Println(err)
		return
	}
	//延迟关闭游标
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			logrus.Error(err)
		}
	}()

	//遍历游标获取结果数据
	for cursor.Next(context.TODO()) {
		var lr LogRecord
		//反序列化Bson到对象
		if cursor.Decode(&lr) != nil {
			fmt.Print(err)
			return
		}
		//打印结果数据
		fmt.Println(lr)
	}

}

func GetRecord() primitive.ObjectID {
	var (
		client     = GetMgoCli()
		collection *mongo.Collection
	)
	//2.选择数据库 my_db里的某个表
	collection = client.Database("my_db").Collection("test")
	//如果直接使用 LogRecord{JobName: "job10"}是查不到数据的，因为其他字段有初始值0或者“”
	// filter := bson.M{"age": bson.M{"$lt": 32}, "command": bson.M{"$eq": "eo"}}
	//按照jobName字段进行过滤jobName="job10",翻页参数0-2
	filter := bson.M{"age": 16, "jobName": "job1111"}
	var lr *LogRecord
	res, err := collection.Find(context.TODO(), filter)
	for res.Next(context.Background()) {
		lr = new(LogRecord)
		res.Decode(lr)
		return lr.Id
	}
	fmt.Println(err)
	res.Decode(lr)

	fmt.Printf("%#v", lr)
	return [12]byte{}
}

type RegEx struct {
	Pattern string
	Options string
}

func FuzzyFind() {
	var (
		client     = GetMgoCli()
		collection *mongo.Collection
	)
	//2.选择数据库 my_db里的某个表
	collection = client.Database("my_db").Collection("test")
	//如果直接使用 LogRecord{JobName: "job10"}是查不到数据的，因为其他字段有初始值0或者“”
	// filter := bson.M{"age": bson.M{"$lt": 32}, "command": bson.M{"$eq": "eo"}}
	//按照jobName字段进行过滤jobName="job10",翻页参数0-2
	//	filter := bson.M{"age": 16, "jobName": "job1111"}
	filter := bson.M{}
	/*m := make([]bson.M, 0)
	m = append(m, bson.M{
		"key": bson.M{"$regex": "job", "$options": "$i"},
	})*/

	// filter := bson.M{"jobName": bson.M{"$regex":"\\job"}}
	filter["jobName"] = bson.M{"$regex": "2", "$options": "$i"}
	var lr *LogRecord
	res, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println(err)
		return
	}
	for res.Next(context.Background()) {
		lr = new(LogRecord)
		res.Decode(lr)
		fmt.Printf("%#v\n", lr)
	}

}

func Session() {
	var (
		client     = GetMgoCli()
		collection *mongo.Collection
	)
	//2.选择数据库 my_db里的某个表
	collection = client.Database("my_db").Collection("test")

	sessionCtx, err := client.StartSession()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sessionCtx.EndSession(context.Background())
	id1, err := primitive.ObjectIDFromHex("601f87c651266e19194dd610")
	id2, err := primitive.ObjectIDFromHex("601f87c651266e19194dd611")
	//	id3, err := primitive.ObjectIDFromHex("601f87c651266e19194dd612")
	_, err = collection.InsertOne(context.Background(), LogRecord{
		Id:      id1,
		JobName: "jj2",
		Command: "test",
		Age:     100,
	})
	if err != nil {
		fmt.Println(err)
	}
	/*sessionCtx.StartTransaction()
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx := context.Background()
	_, err = collection.InsertOne(ctx, LogRecord{
		Id:      id2,
		JobName: "jj3",
		Command: "test",
		Age:     100,
	})
	if err != nil {
		sessionCtx.AbortTransaction(ctx)
	}
	_, err = collection.InsertOne(ctx, LogRecord{
		Id:      id1,
		JobName: "jj4",
		Command: "test",
		Age:     100,
	})
	if err != nil {
		sessionCtx.AbortTransaction(ctx)
	} else {
		sessionCtx.CommitTransaction(ctx)
	}*/
	callBack := func(sessionCtx mongo.SessionContext) (interface{}, error) {
		// ctx := context.Background()
		_, err = collection.InsertOne(sessionCtx, LogRecord{
			Id:      id2,
			JobName: "jj3",
			Command: "test",
			Age:     100,
		})
		if err != nil {
			sessionCtx.AbortTransaction(sessionCtx)
		}
		_, err = collection.InsertOne(sessionCtx, LogRecord{
			Id:      id1,
			JobName: "jj4",
			Command: "test",
			Age:     100,
		})
		if err != nil {
			sessionCtx.AbortTransaction(sessionCtx)
		} else {
			sessionCtx.CommitTransaction(sessionCtx)
		}
		return nil, nil
	}
	// 提交事务
	_, err = sessionCtx.WithTransaction(context.Background(), callBack)
	fmt.Println(err)

}

func UseSession() {
	var (
		client     = GetMgoCli()
		collection *mongo.Collection
	)
	//2.选择数据库 my_db里的某个表
	collection = client.Database("my_db").Collection("test")

	ctx := context.Background()
	id1, err := primitive.ObjectIDFromHex("601f87c651266e19194dd610")
	id2, err := primitive.ObjectIDFromHex("601f87c651267e19194dd611")
	//	id3, err := primitive.ObjectIDFromHex("601f87c651266e19194dd612")
	_, err = collection.InsertOne(context.Background(), LogRecord{
		Id:      id1,
		JobName: "jj2",
		Command: "test",
		Age:     100,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	//第二个事务：执行失败，事务没提交，因最后插入了一条重复id "111",
	err = client.UseSession(ctx, func(sessionContext mongo.SessionContext) error { // 这种要显式手动确认事务
		err := sessionContext.StartTransaction()
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer sessionContext.EndSession(context.Background())
		//在事务内写一条id为“222”的记录
		_, err = collection.InsertOne(sessionContext, LogRecord{
			Id:      id2,
			JobName: "jj3",
			Command: "test",
			Age:     100,
		})
		if err != nil {
			fmt.Println(err)
			return err
		}
		//写重复id
		/*_, err = collection.InsertOne(sessionContext, LogRecord{
			Id:      id1,
			JobName: "jj5",
			Command: "test",
			Age:     100,
		})*/
		if err != nil { // 要显式确认
			sessionContext.AbortTransaction(sessionContext)
			return err
		} else {
			sessionContext.CommitTransaction(sessionContext)
		}
		return nil
	})
	fmt.Println(err)
}

func AllSession() {
	var (
		client     = GetMgoCli()
		collection *mongo.Collection
	)
	//2.选择数据库 my_db里的某个表
	collection = client.Database("my_db").Collection("test")

	sessionCtx, err := client.StartSession()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sessionCtx.EndSession(context.Background())
	id1, err := primitive.ObjectIDFromHex("601f87c651266e19194dd610")
	id2, err := primitive.ObjectIDFromHex("601f87c651266e19103dd611")
	//	id3, err := primitive.ObjectIDFromHex("601f87c651266e19194dd612")
	_, err = collection.InsertOne(context.Background(), LogRecord{
		Id:      id1,
		JobName: "jj2",
		Command: "test",
		Age:     100,
	})
	if err != nil {
		fmt.Println(err)
	}
	callBack := func(sessionCtx mongo.SessionContext) (interface{}, error) {
		// ctx := context.Background()
		_, err = collection.InsertOne(sessionCtx, LogRecord{
			Id:      id2,
			JobName: "jj3",
			Command: "test",
			Age:     100,
		})
		if err != nil {
			return nil, err
		}
		_, err = collection.InsertOne(sessionCtx, LogRecord{
			Id:      id1,
			JobName: "jj4",
			Command: "test",
			Age:     100,
		})
		/*if err != nil { // 可以显式去commit 事务，或者通过err 系统自动判断
			sessionCtx.AbortTransaction(sessionCtx)
		} else {
			sessionCtx.CommitTransaction(sessionCtx)
		}*/
		return nil, err
	}
	// 提交事务
	_, err = sessionCtx.WithTransaction(context.Background(), callBack)
	fmt.Println(err)

}

func MoreSession() {
	var (
		client = GetMgoCli()
	)
	//2.选择数据库 my_db里的某个表
	collection1 := client.Database("my_db").Collection("test")
	collection2 := client.Database("my_db").Collection("my_collection")

	sessionCtx, err := client.StartSession()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sessionCtx.EndSession(context.Background())
	id1, err := primitive.ObjectIDFromHex("601f87c651266e19194dd610")
	id2, err := primitive.ObjectIDFromHex("601f87c651266e19103dd611")
	//	id3, err := primitive.ObjectIDFromHex("601f87c651266e19194dd612")
	_, err = collection1.InsertOne(context.Background(), LogRecord{
		Id:      id1,
		JobName: "jj2",
		Command: "test",
		Age:     100,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	callBack := func(sessionCtx mongo.SessionContext) (interface{}, error) {
		// ctx := context.Background()
		_, err = collection2.InsertOne(sessionCtx, LogRecord{
			Id:      id2,
			JobName: "jj3",
			Command: "test",
			Age:     100,
		})
		if err != nil {
			return nil, err
		}
		_, err = collection1.InsertOne(sessionCtx, LogRecord{
			Id:      id1,
			JobName: "jj4",
			Command: "test",
			Age:     100,
		})
		/*if err != nil { // 可以显式去commit 事务，或者通过err 系统自动判断
			sessionCtx.AbortTransaction(sessionCtx)
		} else {
			sessionCtx.CommitTransaction(sessionCtx)
		}*/
		return nil, err
	}
	// 提交事务
	_, err = sessionCtx.WithTransaction(context.Background(), callBack)
	fmt.Println(err)

}

func FindJob() {
	var (
		client     = GetMgoCli()
		err        error
		collection *mongo.Collection
		cursor     *mongo.Cursor
	)
	//2.选择数据库 my_db里的某个表
	collection = client.Database("my_db").Collection("test")
	//如果直接使用 LogRecord{JobName: "job10"}是查不到数据的，因为其他字段有初始值0或者“”
	// filter := bson.M{"age": bson.M{"$lt": 32}, "command": bson.M{"$eq": "eo"}}
	//按照jobName字段进行过滤jobName="job10",翻页参数0-2
	// 查询条件
	/*	filter := bson.D{
		{"$and", bson.A{
			bson.M{"_id": bson.M{"$in": ids}},
			bson.M{"deleted": false}, // 排除已删除的业务
		}},
	}*/
	id1, err := primitive.ObjectIDFromHex("60370b35c5e1b14b8e6e6c3a")
	fmt.Println(err)
	id2, err := primitive.ObjectIDFromHex("60370b35c5e1b14b8e6e6c39")
	fmt.Println(err)
	/*filter := bson.D{
		{"$and", bson.A{
			bson.M{"_id": bson.M{"$in": ids}},
			bson.M{"deleted": false}, // 排除已删除的业务
		}},
	}*/
	filter := bson.M{"_id": bson.M{"$in": []primitive.ObjectID{id1, id2}}}
	/*ids := make([]primitive.ObjectID, 0)
	ids = append(ids, id1)
	ids = append(ids, id2)*/
	//	filter := bson.D{{"_id", bson.M{"$in": ids}}}
	//options := &options.FindOptions{Sort: bson.M{"age":1}}
	if cursor, err = collection.Find(context.TODO(), filter, nil); err != nil {
		fmt.Println(err)
		return
	}
	//延迟关闭游标
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			logrus.Error(err)
		}
	}()

	//遍历游标获取结果数据
	for cursor.Next(context.TODO()) {
		var lr LogRecord
		//反序列化Bson到对象
		if cursor.Decode(&lr) != nil {
			fmt.Print(err)
			return
		}
		//打印结果数据
		fmt.Println(lr)
	}

}
func FindSlice() {
	var (
		client     = GetMgoCli()
		err        error
		collection *mongo.Collection
		cursor     *mongo.Cursor
	)
	//2.选择数据库 my_db里的某个表
	collection = client.Database("my_db").Collection("test")

	es := make([]bson.E, 0)

	/*es = append(es, bson.E{Key: "$or", Value: bson.A{
		bson.M{"keys": bson.M{"$all": []string{"12"}}}, // $all 数组包括一个值$all
		bson.M{"keys": bson.M{"$all": []string{"1"}}},  // $all 数组包括一个值$all
		// bson.M{"jobName": "work13"},
	}})*/
	es = append(es, bson.E{Key: "$and", Value: bson.A{
		bson.M{"jobName": "work13"},
	}})
	filter := bson.D{
		{"$or", bson.A{
			bson.M{"keys": bson.M{"$all": []string{"12"}}}, // $all 数组包括一个值$all
			bson.M{"keys": bson.M{"$all": []string{"1"}}},  // $all 数组包括一个值$all
			// bson.M{"jobName": "work13"},
		}}, {
			"$and", bson.A{
				bson.M{"jobName": "work13"},
			}},
	}
	filter := bson.D{}
	filter = append(filter, es...)
	//	filter := bson.D{{"_id", bson.M{"$in": ids}}}
	//options := &options.FindOptions{Sort: bson.M{"age":1}}
	if cursor, err = collection.Find(context.TODO(), filter, nil); err != nil {
		fmt.Println(err)
		return
	}
	//延迟关闭游标
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			logrus.Error(err)
		}
	}()

	//遍历游标获取结果数据
	for cursor.Next(context.TODO()) {
		var lr LogRecord
		//反序列化Bson到对象
		if cursor.Decode(&lr) != nil {
			fmt.Print(err)
			return
		}
		//打印结果数据
		fmt.Println(lr)
	}
}
