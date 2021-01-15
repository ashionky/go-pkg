/**
 * @Author pibing
 * @create 2021/1/13 11:26 AM
 */

package es

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/olivere/elastic"
	"go-pkg/pkg/util"
	"strconv"
)

//标准知识库的索引分布数
const EsStdIdxCount = 10
const EsMsgIdxName = "message_"
const ALIASE = "_aliase" //别名
const TYPE = "_doc"

//添加/更新内容到es
func AddMessage(message map[string]interface{}, account string) error {
	idx := util.GetHashCode(account, EsStdIdxCount)
	id, _ := message["_id"].(string)
	delete(message, "_id")
	esClient := getClient()
	if esClient != nil {
		_, err := esClient.Index().Index(EsMsgIdxName + strconv.Itoa(idx) + ALIASE).Type(TYPE).Routing(account).Id(id).BodyJson(message).Do(ctx)
		if err != nil {
			fmt.Printf("add/update doc sync es err: %v ", err)
			return err
		}
		return nil
	} else {
		fmt.Println("es client context deadlineexceeded")
		return errors.New("es client context deadlineexceeded")
	}
}


/*
  account   //类似路由
  eq       //完全匹配
  msgContent //模糊匹配
  begin end  //范围匹配
  return []string  ***id数组
*/
func SearchMessage(account string, msgContent string) (ids []string, err error) {
	idx := util.GetHashCode(account, EsStdIdxCount)
	qesMatch := elastic.NewMatchPhraseQuery("name", msgContent).Analyzer("standard") //模糊匹配内容
	//eqMatch := elastic.NewMatchQuery("name", eq)   //完全相等匹配
	//sresMatch := elastic.NewRangeQuery("date").Gte(begin).Lte(end)    //范围查询

	query := elastic.NewBoolQuery().Must(qesMatch)
	searchResult, err := getClient().Search().TrackTotalHits(false).
		Index(EsMsgIdxName + strconv.Itoa(idx) + ALIASE).
		Type(TYPE).Routing(account).
		Query(query).
		StoredFields("_id").     //搜索需要的字段
		From(0).Size(1000).  //搜索尺码记录数
		Do(ctx)
	if err != nil {
		fmt.Printf("--search  doc from es err: %v \n", err)
		return
	}
	//if searchResult.TotalHits() == 0 {
	//	return
	//}
	for _, item := range searchResult.Hits.Hits {
		ids = append(ids, item.Id)
	}
	return
}

/**
  全量数据查询
  account   //类似路由
  msgContent //模糊匹配
  begin end  //范围匹配
 return []map[string]interface{}  ** 直接返回数据（所有字段）
*/
func SearchMessageAll(account string, msgContent string) (data []map[string]interface{}, err error) {
	idx := util.GetHashCode(account, EsStdIdxCount)
	qesMatch := elastic.NewMatchPhraseQuery("name", msgContent).Analyzer("standard") //模糊匹配内容
	//sresMatch := elastic.NewRangeQuery("date").Gte(begin).Lte(end)
	query := elastic.NewBoolQuery().Must(qesMatch)

	searchResult, err := getClient().Search().TrackTotalHits(false).
		Index(EsMsgIdxName + strconv.Itoa(idx) + ALIASE).
		Type(TYPE).Routing(account).
		Query(query).
		From(0).Size(1000).
		Do(ctx)
	if err != nil {
		fmt.Printf("-----search  doc from es err : %v ", err)
		return
	}
	//if searchResult.TotalHits() == 0 {
	//	return
	//}
	for _, hit := range searchResult.Hits.Hits {
		var t map[string]interface{}
		err := json.Unmarshal(*hit.Source, &t) //map
		if err != nil {
			fmt.Printf("search  doc es Unmarshal err: %v ", err)
		}
		data = append(data, t)
	}
	return data, nil
}
