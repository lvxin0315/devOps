package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/djimenez/iconv-go"
	"github.com/fedesog/webdriver"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lvxin0315/devOps/tm_data/item"
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

//http cookie
var acCookieList []*http.Cookie

var itemUrlList = []string{
	"https://detail.m.tmall.com/item.htm?spm=a220m.1000858.1000725.1.696e39e5oGxM5x&id=622278076807&skuId=4400155540185&user_id=2207314869202&cat_id=2&is_b=1&rn=543ef0a7f74ad927ef2f3e0fc857899f",
}

type LData struct {
	Title  string `json:"title"`
	Sales  string `json:"sales"`
	Seller string `json:"seller"`
	Price  string `json:"price"`
	Info   string `json:"info"`
}

func (t *LData) TableName() string {
	return "tm_data"
}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/jd_data?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//db.CreateTable(LData{})
	mainPage()
	//去重
	itemUrlList = removeRepeatedElement(itemUrlList)
	for _, u := range itemUrlList {
		fmt.Println(u)
		if strings.Index(u, "https") < 0 {
			u = "https:" + u
		}
		data := download(u)
		if data != nil {
			db.New().Save(data)
		}
	}
}

func mainPage() {
	html, err := curlGet("https://list.tmall.com/search_product.htm?q=%C9%CC%D3%C3%B3%F8%BE%DF&type=p&spm=a220m.1000858.a2227oh.d100&from=.list.pc_1_searchbutton")
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(html))
	if err != nil {
		panic(err)
	}
	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		target, ok := selection.Attr("target")
		if !ok || target != "_blank" {
			return
		}
		href, ok := selection.Attr("href")
		if !ok || href == "" {
			return
		}
		//连接包括item
		if strings.Index(href, "detail.tmall.com") < 0 {
			return
		}
		//转成手机端
		itemUrlList = append(itemUrlList, strings.ReplaceAll(href, "detail.tmall.com", "detail.m.tmall.com"))
	})
}

func download(itemUrl string) *LData {
	//htmlByte, err := curlGet(itemUrl)
	//if err != nil {
	//	panic(err)
	//}
	htmlByte := webdriverGet(itemUrl)
	b := bytes.NewBuffer(htmlByte)
	//临时保存到js
	ioutil.WriteFile("a.html", htmlByte, 0755)
	//分离js内容
	scriptCode := ""
	doc, err := goquery.NewDocumentFromReader(b)
	if err != nil {
		fmt.Println("downloadPage NewDocumentFromReader:", err)
		return nil
	}

	doc.Find("script").Each(func(i int, selection *goquery.Selection) {
		if strings.Index(selection.Text(), "var _DATA_Detail") >= 0 {
			scriptCode += fmt.Sprintf("%s\r", selection.Text())

		}
		if strings.Index(selection.Text(), "var _DATA_Mdskip") >= 0 {
			scriptCode += fmt.Sprintf("%s\r", selection.Text())
		}
		//scriptCode += fmt.Sprintf("%s\r", selection.Text())
	})
	//临时保存到js
	ioutil.WriteFile("1.js", []byte(scriptCode), 0755)
	jsVm := otto.New()
	_, err = jsVm.Run(scriptCode)
	if err != nil {
		fmt.Println("otto vmValue:", err)
		return nil
	}
	//_itemDetail 取值
	_itemDetailValue, err := jsVm.Run(`JSON.stringify(_DATA_Detail);`)
	if err != nil {
		fmt.Println("_DATA_Detail:", err)
		return nil
	}
	_itemDetailValueString, err := _itemDetailValue.ToString()
	if err != nil {
		fmt.Println("_itemDetailValueString:", err)
		return nil
	}
	//解析到结构体
	itemDetailValue := new(item.TmItem)
	err = json.Unmarshal([]byte(_itemDetailValueString), itemDetailValue)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(_itemOnlyValueString), itemOnlyValue):", err)
		return nil
	}
	popsB, err := json.Marshal(itemDetailValue.Props.GroupProps)
	return &LData{
		Title:  itemDetailValue.Item.Title,
		Sales:  doc.Find(".sales").Text(),
		Seller: itemDetailValue.Seller.ShopName,
		Price:  itemDetailValue.Mock.Price.Price.PriceText,
		Info:   string(popsB),
	}
}

//get 请求
func curlGet(getUrl string) ([]byte, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	req, _ := http.NewRequest("GET", getUrl, nil)
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	for _, c := range acCookieList {
		req.AddCookie(c)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return nil, err
	}
	//函数结束后关闭相关链接
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//保留cookie，下次用
	acCookieList = resp.Cookies()
	//gb2312 -> utf-8
	return gb2312ToUtf8(body), nil
}

func gb2312ToUtf8(input []byte) []byte {
	out := make([]byte, len(input))
	iconv.Convert(input, out, "gb2312", "utf-8")
	return out
}

func webdriverGet(itemUrl string) []byte {
	chromeDriver := webdriver.NewChromeDriver("/Users/lvxin/IdeaProjects/SeleniumDemo/src/test/resources/chromedriver")
	err := chromeDriver.Start()
	if err != nil {
		log.Println(err)
	}
	desired := webdriver.Capabilities{"Platform": "Linux"}
	required := webdriver.Capabilities{}
	session, err := chromeDriver.NewSession(desired, required)
	if err != nil {
		log.Println(err)
	}
	err = session.Url(itemUrl)
	if err != nil {
		log.Println(err)
	}
	time.Sleep(3 * time.Second)
	html, _ := session.Source()
	session.Delete()
	chromeDriver.Stop()
	return []byte(html)
}

//去重
func removeRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
