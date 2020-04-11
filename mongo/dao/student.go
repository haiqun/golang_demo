package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang_demo/mongo/db"
	"log"
	"time"
)

// 数据结构体
type Student struct {
	Id int32
	Name string
	Age int
}

var (
	opt = "mongodb://root:root123@localhost:27018"
 	name = "user1"
	maxTime = time.Duration(2)
	num uint64 = 50
	table = "student"
	toDB *mongo.Database
	collection *mongo.Collection
	)

func init()  {
	var err error
	toDB, err = db.ConnectToDB(opt, name,maxTime,num)
	if err!= nil {
		panic("链接数据库有误!")
	}
	collection = toDB.Collection(table)
}

// GetList 获取全量的数据
func GetList()  {
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	var all []*Student
	err = cur.All(context.Background(), &all)
	if err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())

	log.Println("collection.Find curl.All: ", all)
	for _, one := range all {
		log.Println("Id:",one.Id," - name:",one.Name," - age:",one.Age)
	}
}

// AddOne 新增一条数据
func AddOne(s1 *Student)  {
	objId, err := collection.InsertOne(context.TODO(), &s1)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("录入数据成功，objId:",objId)
}

// EditOne 编辑一条数据
func EditOne(student *Student,m bson.M)  {
	update := bson.M{"$set": student}
	updateResult, err := collection.UpdateOne(context.Background(),  m, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.UpdateOne:", updateResult)
}

// 更新数据 - 存在更新，不存在就新增
func Update(student *Student,m bson.M)  {
	update := bson.M{"$set": student}
	updateOpts := options.Update().SetUpsert(true)
	updateResult, err := collection.UpdateOne(context.Background(), m, update, updateOpts)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.UpdateOne:", updateResult)
}

// 删除一条数据
func Del(m bson.M)  {
	deleteResult, err := collection.DeleteOne(context.Background(), m)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.DeleteOne:", deleteResult)
}

// Sectle 模糊查询
// bson.M{"name": primitive.Regex{Pattern: "深入"}}
func Sectle(m bson.M)  {
	cur, err := collection.Find(context.Background(), m)
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.Background()) {
		var s Student
		if err = cur.Decode(&s); err != nil {
			log.Fatal(err)
		}
		log.Println("collection.Find name=primitive.Regex{xx}: ", s)
	}
	cur.Close(context.Background())
}

// 统计总数
func Count()  {
	count, err := collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(count)
	}
	log.Println("collection.CountDocuments:", count)
}

// 搜索
func GetOne(m bson.M)  {
	var one Student
	err := collection.FindOne(context.Background(), m).Decode(&one)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("collection.FindOne: ", one)
}
