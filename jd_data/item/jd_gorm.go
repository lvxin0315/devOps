package item

import (
	"encoding/json"
	"fmt"
	"time"
)

type JdItemModel struct {
	SkuID string `gorm:"primary_key;not null;unique"`
	//品牌
	BrandID     string
	BrandName   string
	PName       string
	SkuName     string
	ProductArea string
	SaleUnit    string
	//分类id集合
	CategoryJson string
	//包装尺寸(mm)
	Length int64
	Width  int64
	Height int64
	//规格
	SalePropJson    string
	SalePropSeqJson string
	//新版规格
	NewColorSizeJson string
	//jd价格
	//"l": "",    //划线价
	//"m": "10998.00",
	//"nup": "",
	//"op": "5149.00",
	//"p": "4999.00",
	PriceL  string
	PriceM  string
	PriceOP string
	PriceP  string
	//库存情况
	StockState     int64
	StockStateName string
	StockAreaJson  string
}

func (m *JdItemModel) TableName() string {
	return "jd_item_" + time.Now().Format("20060102")
}

func SaveJdItemModel(itemOnly *JdItemOnly, itemInfo *JdItemInfo) {
	jdItemModel := &JdItemModel{
		SkuID:            itemOnly.Item.SkuID,
		BrandID:          itemOnly.Item.BrandID,
		BrandName:        itemOnly.Item.BrandName,
		PName:            itemOnly.Item.PName,
		SkuName:          itemOnly.Item.SkuName,
		ProductArea:      itemOnly.Item.ProductArea,
		SaleUnit:         itemOnly.Item.SaleUnit,
		CategoryJson:     toJsonString(itemOnly.Item.Category),
		Length:           itemOnly.Item.Length,
		Width:            itemOnly.Item.Width,
		Height:           itemOnly.Item.Height,
		SalePropJson:     toJsonString(itemOnly.Item.SaleProp),
		SalePropSeqJson:  toJsonString(itemOnly.Item.SalePropSeq),
		NewColorSizeJson: toJsonString(itemOnly.Item.NewColorSize),
		PriceL:           itemInfo.Price.L,
		PriceM:           itemInfo.Price.M,
		PriceOP:          itemInfo.Price.Op,
		PriceP:           itemInfo.Price.P,
		StockState:       itemInfo.Stock.StockState,
		StockStateName:   itemInfo.Stock.StockStateName,
		StockAreaJson:    toJsonString(itemInfo.Stock.Area),
	}
	err := NewDB().Save(jdItemModel).Error
	if err != nil {
		fmt.Println("NewDB().Save(jdItemModel)", err)
	}
}

func toJsonString(data interface{}) string {
	jsonByte, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(jsonByte)
}

func InitTable() {
	if !NewDB().HasTable(JdItemModel{}) {
		NewDB().CreateTable(JdItemModel{})
	}
}
