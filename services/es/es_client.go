package esclient

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/httplib"
)

var esUrl string

func init(){
	esUrl = "http://127.0.0.1:9200"
}

// HitData 命中的数据
type HitData struct {
	Total TotalData `json:"total"`
	Hits []HitItemData `json:"hits"`
}

type TotalData struct {
	Value int
	Relation string
}

type HitItemData struct{
	Source json.RawMessage `json:"_source"`
}

type ResSearchData struct {
	Hits HitData `json:"hits"`
}

// EsSearch 搜索功能
func EsSearch(indexName string,query map[string]interface{},from int,size int,sort []map[string]string) (HitData,error) {
	// 拼装数据
	searchQueryData:= map[string]interface{}{
		"query":query,
		"from":from,
		"size":size,
		"sort":sort,
	}

	// 通过Api调用Es
	request := httplib.Post(esUrl + indexName + "_search")
	request.JSONBody(searchQueryData)
	response, err := request.String()
	if err != nil {
		fmt.Println(err)
	}
	// 解析
	var rsd ResSearchData
	err = json.Unmarshal([]byte(response), &rsd)

	if err == nil {
		return rsd.Hits,err
	}

	return rsd.Hits,err
}

// EsAdd 添加
func EsAdd(indexName string,id string,params map[string]interface{}) (err error) {
	request := httplib.Post(esUrl + indexName + "/_doc/" + id)
	jsonBody, err := request.JSONBody(params)
	if err != nil {
		return
	}
	_, err = jsonBody.String()
	return err
}

// EsEdit 修改
func EsEdit(indexName string,id string,params map[string]interface{})  (err error){
	bodyData := map[string]interface{}{
		"doc":params,
	}
	// 请求
	request := httplib.Post(esUrl + indexName + "/_doc/" + id + "/_update")
	request.JSONBody(bodyData)
	_, err = request.String()
	if err != nil {
		return
	}
	return err
}


// EsDelete 删除数据
func EsDelete(indexName string,id string) error {
	request := httplib.Delete(esUrl + indexName + "/_doc/" + id)
	_, err := request.String()
	return err
}


