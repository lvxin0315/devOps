package es

import (
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"reflect"
)

type QueryData interface {
	SetScore(float64)
	SetEsId(string)
}

type queryResult struct {
	Total    int64
	MaxScore float64
	DataList []QueryData
}

func (c *esClient) Index(indexKey string) *esClient {
	c.query = c.client.Search().Index(indexKey).Pretty(true)
	return c
}

//执行查询索引
func (c *esClient) Query(query elastic.Query) *esClient {
	c.query = c.query.Query(query)
	return c
}

func (c *esClient) SearchService() *elastic.SearchService {
	return c.query
}

//分数排序
func (c *esClient) DefaultSort() *esClient {
	//排序规则
	scoreSort := elastic.NewScoreSort()
	scoreSort.Order(false)
	c.query.SortBy(scoreSort)
	return c
}

//分页
func (c *esClient) Page(pageNo int, limit int) *esClient {
	c.query.From((pageNo - 1) * limit).Size(limit)
	return c
}

//查询
func (c *esClient) Select(typ reflect.Type) (*queryResult, error) {
	searchResult, err := c.query.Do(c.ctx)
	if err != nil {
		return nil, err
	}
	qr := new(queryResult)
	qr.Total = searchResult.TotalHits()
	qr.MaxScore = *searchResult.Hits.MaxScore
	//结果处理
	for _, h := range searchResult.Hits.Hits {
		v := reflect.New(typ).Elem()
		if h.Source == nil {
			qr.DataList = append(qr.DataList, v.Interface().(QueryData))
			continue
		}
		if err := json.Unmarshal(h.Source, v.Addr().Interface()); err == nil {
			//填充分数
			v.Interface().(QueryData).SetScore(*h.Score)
			v.Interface().(QueryData).SetEsId(h.Id)
			qr.DataList = append(qr.DataList, v.Interface().(QueryData))
		}
	}
	return qr, nil
}
