package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang_demo/mongo/dao"
	//"math/rand"
)

func main() {
	// table 统计
	dao.Count()
	// 更新数据
	//st := dao.Student{
	//	Name: "kk3",
	//	Age:  24,
	//	Id: rand.Int31n(9999999),
	//}
	// 录入一条数据 ，objId 查询信息
	//dao.AddOne(&st)
	//dao.EditOne(&st,bson.M{"Id": "aq"})
	// 不存在就更新
	//dao.Update(&st,bson.M{"Id": "aq"})
	// 删除数据
	dao.Del(bson.M{"id:": 2019727887})
	// 获取全部的集合
	//dao.GetList()
	// 模糊搜索  bson.M{"name": primitive.Regex{Pattern: "深入"}}
	dao.Sectle(bson.M{"name": primitive.Regex{Pattern: "kk"}})
	// 准确搜索
	dao.GetOne(bson.M{"id":2019727887})

}
