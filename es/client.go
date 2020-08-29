package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
)

type esClient struct {
	host    string
	port    int
	client  *elastic.Client
	sniffer bool
	ctx     context.Context
}

//构造EsClient
func NewEsClient(host string, port int, sniffer bool) *esClient {
	ctx := context.Background()
	esc := &esClient{
		host:    host,
		port:    port,
		sniffer: sniffer,
		ctx:     ctx,
	}
	return esc
}

//连接es server
func (c *esClient) Connect() error {
	client, err := elastic.NewClient(
		elastic.SetURL(fmt.Sprintf("http://%s:%d", c.host, c.port)),
		elastic.SetSniff(c.sniffer))
	if err != nil {
		return err
	}
	c.client = client
	//检查版本
	if _, err := c.checkVersion(); err != nil {
		return err
	}
	return nil
}

//检查version，也算是检查通讯是否OK
func (c *esClient) checkVersion() (string, error) {
	//检查version 必须7.x
	version, err := c.client.ElasticsearchVersion(fmt.Sprintf("http://%s:%d", c.host, c.port))
	if err != nil {
		return version, err
	}
	if strings.Index(version, "7.") != 0 {
		return version, fmt.Errorf("version is not 7.x")
	}
	return version, nil
}

//查询index
func (c *esClient) IndexNames() ([]string, error) {
	return c.client.IndexNames()
}

//判断index是否存在
func (c *esClient) Exists(indexKey string) (bool, error) {
	return c.client.IndexExists(indexKey).Do(c.ctx)
}

//创建index
func (c *esClient) CreateIndexWithMapping(indexKey string, mapping *Mapping) error {
	bs, err := mapping.MappingToJsonString()
	if err != nil {
		return err
	}
	//验证indexKey 是否存在
	exists, err := c.Exists(indexKey)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("indexKey exists")
	}
	_, err = c.client.CreateIndex(indexKey).BodyString(string(bs)).Do(c.ctx)
	if err != nil {
		return err
	}
	return nil
}

func (c *esClient) CreateIndexWithProperties(indexKey string, pt Properties) error {
	mapping := NewMapping(pt)
	return c.CreateIndexWithMapping(indexKey, mapping)
}

//删除index
func (c *esClient) DeleteIndex(indexKey string) error {
	_, err := c.client.DeleteIndex(indexKey).Do(c.ctx)
	return err
}

//获取index数据条数
func (c *esClient) ItemCount(indexKey string) (int64, error) {
	return c.client.Count(indexKey).Do(c.ctx)
}

//添加index数据
func (c *esClient) AddItem(indexKey string, data interface{}) (string, error) {
	//判断index是否存在
	exists, err := c.Exists(indexKey)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", fmt.Errorf("indexKey not exists")
	}
	indexResponse, err := c.client.Index().Index(indexKey).BodyJson(data).Do(c.ctx)
	if err != nil {
		return "", err
	}
	return indexResponse.Id, err
}

//批量插入index数据
func (c *esClient) AddItems(indexKey string, dataList ...interface{}) (int64, error) {
	//判断index是否存在
	exists, err := c.Exists(indexKey)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, fmt.Errorf("indexKey not exists")
	}
	bs := elastic.NewBulkService(c.client).Index(indexKey)
	for _, data := range dataList {
		br := elastic.NewBulkIndexRequest()
		br.Doc(data)
		bs.Add(br)
	}
	r, err := bs.Do(c.ctx)
	if err != nil {
		return 0, err
	}
	//获取数量
	return int64(len(r.Items)), nil
}

//更新index数据
func (c *esClient) UpdateItem(indexKey string, id string, data interface{}) (int64, error) {
	res, err := c.client.Update().Index(indexKey).Id(id).Doc(data).Do(c.ctx)
	if err != nil {
		return 0, err
	}
	return res.Version, nil
}

//复制index及相关数据
func (c *esClient) ReIndex(sourceIndex string, newIndex string) (int64, error) {
	//判断源index是否存在
	exists, err := c.Exists(sourceIndex)
	if err != nil {
		return 0, err
	}
	if !exists {
		return 0, fmt.Errorf("sourceIndex not exists")
	}
	//判断新index是否存在
	exists, err = c.Exists(newIndex)
	if err != nil {
		return 0, err
	}
	if exists {
		return 0, fmt.Errorf("newIndex exists")
	}
	//nrd := elastic.NewReindexDestination()
	//nrd.Index(newIndex)
	//nrd.OpType("create")
	rs := elastic.NewReindexService(c.client).SourceIndex(sourceIndex).DestinationIndex(newIndex)
	r, err := rs.Do(c.ctx)
	if err != nil {
		return 0, err
	}
	fmt.Println(r.Total)
	return 0, nil
}
