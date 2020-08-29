package es

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func Test_esClient_Connect(t *testing.T) {
	esClient := NewEsClient("172.16.0.203", 9200, false)
	err := esClient.Connect()
	assert.Equal(t, err, nil)
}

func Test_esClient_IndexNames(t *testing.T) {
	esClient := NewEsClient("172.16.0.203", 9200, false)
	err := esClient.Connect()
	assert.Equal(t, err, nil)
	indexNames, err := esClient.IndexNames()
	fmt.Println(indexNames)
	assert.Equal(t, err, nil)
}

func Test_esClient_CreateIndexWithMapping(t *testing.T) {
	esClient := NewEsClient("172.16.0.203", 9200, false)
	err := esClient.Connect()
	assert.Equal(t, err, nil)
	pt := NewBaseIndexProperties()
	pt.SetProperties("name", &PropertiesType{
		Type: "text",
	})
	pt.SetProperties("des", &PropertiesType{
		Type: "text",
	})
	pt.SetProperties("content", &PropertiesType{
		Type: "text",
	})
	mapping := NewMapping(pt)
	err = esClient.CreateIndexWithMapping("test_index1", mapping)
	assert.Equal(t, err, nil, "创建error")
	//删除index
	err = esClient.DeleteIndex("test_index1")
	assert.Equal(t, err, nil, "删除error")
}

func Test_esClient_AddItem(t *testing.T) {
	esClient := NewEsClient("172.16.0.203", 9200, false)
	err := esClient.Connect()
	assert.Equal(t, err, nil)
	pt := NewBaseIndexProperties()
	pt.SetProperties("name", &PropertiesType{
		Type: "text",
	})
	pt.SetProperties("content", &PropertiesType{
		Type: "text",
	})
	err = esClient.CreateIndexWithProperties("test_index2", pt)
	assert.Equal(t, err, nil, "创建error")
	type data struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}
	id, err := esClient.AddItem("test_index2", &data{
		Name:    "test1",
		Content: "test111",
	})
	fmt.Println(id)
	assert.Equal(t, err, nil, "AddItem error")
	//删除index
	err = esClient.DeleteIndex("test_index2")
	assert.Equal(t, err, nil, "删除error")
}

func Test_esClient_AddItems(t *testing.T) {
	esClient := NewEsClient("172.16.0.203", 9200, false)
	err := esClient.Connect()
	assert.Equal(t, err, nil)
	pt := NewBaseIndexProperties()
	pt.SetProperties("name", &PropertiesType{
		Type: "text",
	})
	pt.SetProperties("content", &PropertiesType{
		Type: "text",
	})
	err = esClient.CreateIndexWithProperties("test_index3", pt)
	assert.Equal(t, err, nil, "创建error")
	type data struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}
	num, err := esClient.AddItems("test_index3", &data{
		Name:    "test1",
		Content: "test111",
	}, &data{
		Name:    "test2",
		Content: "test222",
	}, &data{
		Name:    "test3",
		Content: "test333",
	}, &data{
		Name:    "test4",
		Content: "test444",
	})
	fmt.Println("num:", num)
	assert.Equal(t, err, nil, "AddItems error")
	//删除index
	err = esClient.DeleteIndex("test_index3")
	assert.Equal(t, err, nil, "删除error")
}

func Test_esClient_Update(t *testing.T) {
	esClient := NewEsClient("172.16.0.203", 9200, false)
	err := esClient.Connect()
	assert.Equal(t, err, nil)

	pt := NewBaseIndexProperties()
	pt.SetProperties("name", &PropertiesType{
		Type: "text",
	})
	pt.SetProperties("content", &PropertiesType{
		Type: "text",
	})
	err = esClient.CreateIndexWithProperties("test_index4", pt)
	assert.Equal(t, err, nil, "创建error")
	type data struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}
	id, err := esClient.AddItem("test_index4", &data{
		Name:    "test1",
		Content: "test111",
	})
	fmt.Println(id)
	assert.Equal(t, err, nil, "AddItem error")

	version, err := esClient.UpdateItem("test_index4", id, &data{
		Name:    "test11",
		Content: "test1111",
	})
	fmt.Println("更新后version：", version)
	assert.Equal(t, err, nil, "UpdateItem error")

	//删除index
	err = esClient.DeleteIndex("test_index4")
	assert.Equal(t, err, nil, "删除error")
}

func Test_esClient_ReIndex(t *testing.T) {
	esClient := NewEsClient("172.16.0.203", 9200, false)
	err := esClient.Connect()
	assert.Equal(t, err, nil)
	pt := NewBaseIndexProperties()
	pt.SetProperties("name", &PropertiesType{
		Type: "text",
	})
	pt.SetProperties("content", &PropertiesType{
		Type: "text",
	})
	err = esClient.CreateIndexWithProperties("test_index5", pt)
	assert.Equal(t, err, nil, "创建error")
	type data struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}
	num, err := esClient.AddItems("test_index5", &data{
		Name:    "test1",
		Content: "test111",
	}, &data{
		Name:    "test2",
		Content: "test222",
	}, &data{
		Name:    "test3",
		Content: "test333",
	}, &data{
		Name:    "test4",
		Content: "test444",
	})
	fmt.Println("num:", num)
	assert.Equal(t, err, nil, "AddItems error")
	time.Sleep(3 * time.Second)
	//复制index
	esClient.ReIndex("test_index5", "test_index6")
	time.Sleep(3 * time.Second)
	//删除index
	err = esClient.DeleteIndex("test_index5")
	err = esClient.DeleteIndex("test_index6")
	assert.Equal(t, err, nil, "删除error")
}

type Goods struct {
	EsId      string
	EsScore   float64
	Title     string `json:"title"`      //商品名称
	BrandName string `json:"brand_name"` //品牌
}

func (g *Goods) SetScore(f float64) {
	g.EsScore = f
}

func (g *Goods) SetEsId(id string) {
	g.EsId = id
}

func Test_esClient_Select(t *testing.T) {
	fmt.Println("Test_esClient_Select start")
	esClient := NewEsClient("172.16.0.203", 9200, false)
	err := esClient.Connect()
	assert.Equal(t, err, nil)
	query := elastic.NewRawStringQuery(`{
		"bool" : {
			"must" : [
				{
					"match" : { "brand_name" : "唐宗筷" }
				}
			],
			"should" : [
				{
					"match" : { "title" : "唐宗筷" },
					"match" : { "title" : "一次性竹牙签" }
				}
			],
			"minimum_should_match" : 1,
			"boost" : 1.0
		}
	}`)
	qr, err := esClient.Index("goods-item").Query(query).DefaultSort().Select(reflect.TypeOf(&Goods{}))
	assert.Equal(t, err, nil, "Select error")
	fmt.Println(qr.Total)
	fmt.Println(qr.MaxScore)
	for _, d := range qr.DataList {
		fmt.Println("EsId:", d.(*Goods).EsId)
		fmt.Println("Title:", d.(*Goods).Title)
		fmt.Println("BrandName:", d.(*Goods).BrandName)
		fmt.Println("EsScore:", d.(*Goods).EsScore)
	}
}

func Test_esClient_Select1(t *testing.T) {
	fmt.Println("Test_esClient_Select1 start")
	esClient := NewEsClient("172.16.0.203", 9200, false)
	err := esClient.Connect()
	assert.Equal(t, err, nil)
	query := elastic.NewBoolQuery()
	query.Must(elastic.NewMatchQuery("brand_name", "唐宗筷"))
	query.Should(elastic.NewMatchQuery("title", "唐宗筷"), elastic.NewMatchQuery("title", "一次性竹牙签"))
	query.MinimumShouldMatch("1")
	query.Boost(1.0)
	qr, err := esClient.Index("goods-item").Query(query).DefaultSort().Select(reflect.TypeOf(&Goods{}))
	assert.Equal(t, err, nil, "Select error")
	fmt.Println(qr.Total)
	fmt.Println(qr.MaxScore)
	for _, d := range qr.DataList {
		fmt.Println("EsId:", d.(*Goods).EsId)
		fmt.Println("Title:", d.(*Goods).Title)
		fmt.Println("BrandName:", d.(*Goods).BrandName)
		fmt.Println("EsScore:", d.(*Goods).EsScore)
	}
}
