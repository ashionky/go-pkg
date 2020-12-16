package mongodb

import (
	"errors"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"io"
	"mime/multipart"
	"reflect"
)

type B bson.M
type Regex bson.RegEx

var s *mgo.Session
var defaultDb string


func getSession() *mgo.Session {
	//if s == nil {
	//	err := MongoInit()
	//	if err != nil {
	//
	//	}
	//}
	return s.Clone()
}


func MongoInit(url, dbname string,poolsize int) error {
	session, err := mgo.Dial(url)
	if err != nil {
		fmt.Println("mongodb init fail!")
		return err
	}
	session.SetMode(mgo.Strong, false)

	session.SetPoolLimit(poolsize)
	s = session
	defaultDb = dbname
	fmt.Println("mongodb init success!")
	return nil
}

func Insert(c string, docs interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(defaultDb).C(c).Insert(docs)
}

func DbInsert(db, c string, docs interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Insert(docs)
}
func DbInsertSafe(db, c string, docs interface{}, safe *mgo.Safe) error {
	session := getSession()
	session.SetSafe(safe)
	defer session.Close()
	return session.DB(db).C(c).Insert(docs)
}
func DbInsertList(db, c string, docs []interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Insert(docs...)
}

func Find(c string, query B, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(defaultDb).C(c).Find(query).All(result)
}

func DbFind(db, c string, query B, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Find(query).All(result)
}

func DbGetCount(db, c string, query B) (int, error) {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Find(query).Count()
}

func FindPage(c string, query B, skip int, limit int, result interface{}) error {
	session := getSession()
	defer session.Close()

	if limit != 0 {
		return session.DB(defaultDb).C(c).Find(query).Limit(limit).Skip(skip).All(result)
	} else {
		return session.DB(defaultDb).C(c).Find(query).All(result)
	}
}
func DBFindPageSortField(db string, c string, query B, field B, skip int, limit int, result interface{}) error {
	session := getSession()
	defer session.Close()

	if limit != 0 {
		return session.DB(db).C(c).Find(query).Select(field).Limit(limit).Skip(skip).All(result)
	} else {
		return session.DB(db).C(c).Find(query).Select(field).All(result)
	}
}
func DBFindPageSortAndField(db string, c string, query B, field B, skip int, limit int, sort string, result interface{}) error {
	session := getSession()
	defer session.Close()

	if limit != 0 {
		return session.DB(db).C(c).Find(query).Select(field).Sort(sort).Limit(limit).Skip(skip).All(result)
	} else {
		return session.DB(db).C(c).Find(query).Select(field).Sort(sort).All(result)
	}
}
func DbFindPage(dbname, c string, query B, skip int, limit int, result interface{}) error {
	session := getSession()
	defer session.Close()

	if limit != 0 {
		return session.DB(dbname).C(c).Find(query).Limit(limit).Skip(skip).All(result)
	} else {
		return session.DB(dbname).C(c).Find(query).All(result)
	}
}

func FindPageSort(c string, query B, skip int, limit int, sort string, result interface{}) error {
	session := getSession()
	defer session.Close()

	if limit != 0 {
		return session.DB(defaultDb).C(c).Find(query).Limit(limit).Sort(sort).Skip(skip).All(result)
	} else {
		return session.DB(defaultDb).C(c).Find(query).Sort(sort).All(result)
	}
}

func DbFindSort(db string, c string, query B, sort string, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Find(query).Sort(sort).All(result)
}

func DbFindSortDouble(db string, c string, query B, sort1 string, sort2 string, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Find(query).Sort(sort1, sort2).All(result)
}

func DbFindPageSortDouble(db string, c string, query B, skip int, limit int, sort1 string, sort2 string, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Find(query).Sort(sort1, sort2).Limit(limit).Skip(skip).All(result)
}

func DBFindPageSort(db string, c string, query B, skip int, limit int, sort string, result interface{}) error {
	session := getSession()
	defer session.Close()
	if sort == "" {
		return errors.New("sort is not exists")
	}
	if limit != 0 {
		return session.DB(db).C(c).Find(query).Sort(sort).Limit(limit).Skip(skip).All(result)
	} else {
		return session.DB(db).C(c).Find(query).Sort(sort).All(result)
	}
}


func DbFindByFields(db string, c string, query B, field B, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Find(query).Select(field).All(result)
}

func DbFindSortByFields(db string, c string, query B, sort string, field B, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Find(query).Select(field).Sort(sort).All(result)
}

func DbFindOneByFields(db string, c string, query B, field B, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Find(query).Select(field).One(result)
}

func FindByFields(c string, query B, field B, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(defaultDb).C(c).Find(query).Select(field).All(result)
}

func FindOne(c string, query B, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(defaultDb).C(c).Find(query).One(result)
}

func DbFindOne(db string, c string, query B, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Find(query).One(result)
}

func DbSortFindOne(db string, c string, query B, sort string, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Find(query).Sort(sort).Limit(1).One(result)
}

func DbFindById(db string, c string, id interface{}, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).FindId(id).One(result)
}

func DbFindFieldsById(db string, c string, id interface{}, field B, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).FindId(id).Select(field).One(result)
}

func FindById(c string, id interface{}, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(defaultDb).C(c).FindId(id).One(result)
}

func UpdateOne(c string, query B, set B) error {
	session := getSession()
	defer session.Close()
	return session.DB(defaultDb).C(c).Update(query, set)
}

func DbUpdateOne(db, c string, query B, set B) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Update(query, set)
}

func DbUpdateOneSafe(db, c string, query B, set B, safe *mgo.Safe) error {
	session := getSession()
	session.SetSafe(safe)
	defer session.Close()
	return session.DB(db).C(c).Update(query, set)
}

func DBUpdateById(DB string, c string, id interface{}, set B) error {
	session := getSession()
	defer session.Close()
	return session.DB(DB).C(c).UpdateId(id, set)
}

func DBUpdateByIdSafe(DB string, c string, id interface{}, set B, safe *mgo.Safe) error {
	session := getSession()
	session.SetSafe(safe)
	defer session.Close()
	return session.DB(DB).C(c).UpdateId(id, set)
}

func UpdateById(c string, id interface{}, set B) error {
	session := getSession()
	defer session.Close()
	return session.DB(defaultDb).C(c).UpdateId(id, set)
}

func UpdateAll(c string, query B, set B) error {
	session := getSession()
	defer session.Close()
	_, err := session.DB(defaultDb).C(c).UpdateAll(query, set)
	return err
}

func DbUpdateAll(db, c string, query B, set B) error {
	session := getSession()
	defer session.Close()
	_, err := session.DB(db).C(c).UpdateAll(query, set)
	return err
}

func DbUpdateAllSafe(db, c string, query B, set B, safe *mgo.Safe) error {
	session := getSession()
	session.SetSafe(safe)
	defer session.Close()
	_, err := session.DB(db).C(c).UpdateAll(query, set)
	return err
}
func UpsertById(c string, id string, set B) error {
	session := getSession()
	defer session.Close()
	_, err := session.DB(defaultDb).C(c).UpsertId(id, set)
	return err
}

func DbUpsertById(db, c, id string, set B) error {
	session := getSession()
	defer session.Close()
	_, err := session.DB(db).C(c).UpsertId(id, set)
	return err
}

func DeleteOne(c string, query B) error {
	session := getSession()
	defer session.Close()
	return session.DB(defaultDb).C(c).Remove(query)
}

func DBDeleteOne(dbname, c string, query B) error {
	session := getSession()
	defer session.Close()
	return session.DB(dbname).C(c).Remove(query)
}

func DeleteAll(c string, query B) error {
	session := getSession()
	defer session.Close()
	_, err := session.DB(defaultDb).C(c).RemoveAll(query)
	return err
}

func DBDeleteAll(dbname, c string, query B) error {
	session := getSession()
	defer session.Close()
	_, err := session.DB(dbname).C(c).RemoveAll(query)
	return err
}

func DBDeleteAllSafe(dbname, c string, query B, safe *mgo.Safe) error {
	session := getSession()
	session.SetSafe(safe)
	defer session.Close()
	_, err := session.DB(dbname).C(c).RemoveAll(query)
	return err
}

func DeleteById(c string, id interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(defaultDb).C(c).RemoveId(id)
}

func DBDeleteById(db, c string, id interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).RemoveId(id)
}

func GetCount(c string, query B) (int, error) {
	session := getSession()
	defer session.Close()
	return session.DB(defaultDb).C(c).Find(query).Count()
}

func GetCountByDb(dbname, c string, query B) (int, error) {
	session := getSession()
	defer session.Close()
	return session.DB(dbname).C(c).Find(query).Count()
}

func GetMaxLimitCountByDb(dbname, c string, query B, limit int) (int, error) {
	session := getSession()
	defer session.Close()
	return session.DB(dbname).C(c).Find(query).Limit(limit).Count()
}

func GetMaxLimitSortCountByDb(dbname, c string, query B, sort string, skip, limit int) (int, error) {
	session := getSession()
	defer session.Close()
	return session.DB(dbname).C(c).Find(query).Sort(sort).Limit(limit).Skip(skip).Count()
}

func AggregateByDb(db string, c string, match B, group B, sort B, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Pipe([]B{{"$match": match}, {"$group": group}, {"$sort": sort}}).All(result)
}

func GetCountAggregateByDb(db string, c string, match, group B) (int, error) {
	result := make(map[string]interface{})
	session := getSession()
	defer session.Close()
	err := session.DB(db).C(c).Pipe([]B{{"$match": match}, {"$group": group}, {"$count": "count"}}).One(result)
	if len(result) > 0 {
		return result["count"].(int), err
	}
	return 0, err
}

func AggregateProjectByDb(db string, c string, match B, group B, project B, sort B, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Pipe([]B{{"$match": match}, {"$group": group}, {"$project": project}, {"$sort": sort}}).All(result)
}

func AggregateLimitByDb(db string, c string, match B, group B, sort B, skip int, limit int, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Pipe([]B{{"$match": match}, {"$group": group}, {"$sort": sort}, {"$skip": skip}, {"$limit": limit}}).All(result)
}

func AggregateLimit(db string, c string, match B, group B, skip int, limit int, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Pipe([]B{{"$match": match}, {"$group": group}, {"$skip": skip}, {"$limit": limit}}).All(result)
}

func AggregateCountByDb(db string, c string, match B, group B, count string, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Pipe([]B{{"$match": match}, {"$group": group}, {"$count": count}}).One(result)
}

func DBFind(dbname, c string, query B, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(dbname).C(c).Find(query).All(result)
}

func ExportExcel(out multipart.File, data map[string]interface{}) error {
	if !contains("fileName", data) {
		return errors.New("文件名不能为空")
	}
	fileName := data["fileName"].(string)

	session := getSession()
	defer session.Close()

	file, err := session.DB("db_file").GridFS("fs").Create(fileName)

	defer file.Close()
	defer out.Close()

	if err != nil {
		fmt.Println(err)
		return err
	}

	if contains("_id", data) {
		file.SetId(data["_id"].(string))
	}
	if contains("contentType", data) {
		file.SetContentType(data["contentType"].(string))
	}
	if contains("metadata", data) {
		file.SetMeta(data["metadata"])
	}

	_, err = io.Copy(file, out)

	if err != nil {
		return errors.New("写入文件流失败")
	}
	return nil
}

func RemoveFile(id interface{}) bool {
	if id == nil {
		return false
	}
	//直接利用名字移除
	session := getSession()
	defer session.Close()

	err := session.DB("db_file").GridFS("fs").RemoveId(id)
	if err != nil {
		return false
	} else {
		return true
	}
}

func contains(obj interface{}, target interface{}) bool {
	if target == nil {
		return false
	}
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}
	return false
}

func DbUpSertOne(db, c string, query B, set B) error {
	session := getSession()
	defer session.Close()
	_, err := session.DB(db).C(c).Upsert(query, set)
	if err != nil {
		return err
	}
	return nil
}

func AggregateCustomer(db string, c string, condition []B, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Pipe(condition).All(result)
}

func DbUpSertOneSafe(db, c string, query B, set B, safe *mgo.Safe) error {
	session := getSession()
	session.SetSafe(safe)
	defer session.Close()
	_, err := session.DB(db).C(c).Upsert(query, set)
	if err != nil {
		return err
	}
	return nil
}
func Aggregate(db string, c string, match B, group B, project B, sort B, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Pipe([]B{{"$match": match}, {"$group": group}, {"$project": project}, {"$sort": sort}}).All(result)
}

func AggregatePipe(db string, c string, pipeline []B, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Pipe(pipeline).All(result)
}

func AggregateByCop(db string, c string, match B, group B, sort B, limit int, skip int, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Pipe([]B{{"$match": match}, {"$group": group}, {"$sort": sort}, {"$skip": skip}, {"$limit": limit}}).All(result)
}
func AggregateByCop2(db string, c string, match B, group B, sort B, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Pipe([]B{{"$match": match}, {"$group": group}, {"$sort": sort}}).All(result)
}

func DbFindSortFields(db string, c string, query B, field B, sort string, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Find(query).Select(field).Sort(sort).All(result)
}

func DBDistinct(db string, c string, field string, query B, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Find(query).Distinct(field, result)
}

func AggregateByOwner(db string, c string, match B, group B, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Pipe([]B{{"$match": match}, {"$group": group}}).All(result)
}

func AggregetaProjectLimitByDb(db string, c string, match B, project B, sort B, skip int, limit int, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Pipe([]B{{"$match": match}, {"$project": project}, {"$sort": sort}, {"$skip": skip}, {"$limit": limit}}).All(result)
}

func AggregateProjectFindOneDb(db string, c string, match B, project B, result interface{}) error {
	session := getSession()
	defer session.Close()
	return session.DB(db).C(c).Pipe([]B{{"$match": match}, {"$project": project}}).One(result)
}

func DbBulkUpdate(db, c string, pairs []interface{}) error {
	session := getSession()
	defer session.Close()
	bulk := session.DB(db).C(c).Bulk()
	bulk.Update(pairs...)
	_, err := bulk.Run()
	return err
}

func DbBulkUpdateAll(db, c string, pairs []interface{}) error {
	session := getSession()
	defer session.Close()
	bulk := session.DB(db).C(c).Bulk()
	bulk.UpdateAll(pairs...)
	_, err := bulk.Run()
	return err
}
