package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/lvxin0315/devOps/jd_data/item"
	"github.com/robertkrimen/otto"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var jdItemUrl = `https://item.m.jd.com/product/%s.html`

var jdSearchUrlTpl = `https://search.jd.com/Search?keyword=%s&page=%d&enc=utf-8`

var keyWord = "商用制冰机"

var itemUrlList []string
var itemIdList []string

//http cookie
var acCookieList []*http.Cookie

//列表page最大值
var maxPage = 30

//item html 保存目录
var itemTmp = "item_tmp"

var debug = false

var httpProxyUrl = "http://58.218.200.229:4736"

var mysqlConn = "root:root@tcp(127.0.0.1:3306)/jd_data?charset=utf8&parseTime=True&loc=Local"

//添加延迟，防止频率高
var sleepTime = 5

//协程
var wg = sync.WaitGroup{}

//最大并发
var maxPool = 4

func main() {
	//init flag
	initFlag()
	//init mysql
	item.InitMysqlDB(mysqlConn)
	//初始化table
	item.InitTable()
	//先打卡jd
	curlGet("https://www.jd.com")

	for i := 1; i <= maxPage; i++ {
		fmt.Println("page:", i)
		mainPage(i)
	}
	//去重itemUrlList生成itemIdList
	getItemIdList()
	fmt.Println(len(itemIdList))
	nowPool := 0
	//解析具体的详情页面
	for _, itemId := range itemIdList {
		//地址变成手机版
		iUrl := fmt.Sprintf(jdItemUrl, itemId)
		//保存html
		fmt.Println("item_url:", iUrl)
		wg.Add(1)
		nowPool++
		go func(iUrl string) {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println(err)
					wg.Done()
				}
			}()
			downloadPage(iUrl)
			wg.Done()
		}(iUrl)
		if nowPool >= maxPool {
			wg.Wait()
			nowPool = 0
			//添加延迟，防止频率高
			time.Sleep(time.Duration(sleepTime) * time.Second)
		}
	}
}

func mainPage(page int) {
	searchUrl := fmt.Sprintf(jdSearchUrlTpl, url.QueryEscape(keyWord), page)
	html, err := curlGet(searchUrl)
	if err != nil {
		panic(err)
	}
	//err = ioutil.WriteFile("1.html", html, 0755)
	//if err != nil {
	//	panic(err)
	//}
	//解析列表页面，生成对应的商品详情url
	err = parseListHtml(bytes.NewBuffer(html))
	if err != nil {
		panic(err)
	}
}

//get 请求
func curlGet(getUrl string) ([]byte, error) {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy:           getProxy,
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
	return body, nil
}

//解析列表页面
func parseListHtml(r io.Reader) error {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		fmt.Println("NewDocumentFromReader", err)
		return err
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
		if strings.Index(href, "item.jd") < 0 {
			return
		}
		itemUrlList = append(itemUrlList, href)
	})
	return nil
}

//下载并保存页面
func downloadPage(itemUrl string) {
	b, err := curlGet(itemUrl)
	if err != nil {
		fmt.Println("curlGet:", err)
		return
	}
	u, err := url.Parse(itemUrl)
	if err != nil {
		fmt.Println("downloadPage:", err)
		return
	}

	//html 文件保存
	saveHTMLFilePath := itemTmp + u.RequestURI()

	writeFile(saveHTMLFilePath, b)

	//分离js内容
	scriptCode := `var window = {};
var _itemOnly = {};
	`
	firstItemInfoScript := true
	firstItemOnlyScript := true
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(b))
	if err != nil {
		fmt.Println("downloadPage NewDocumentFromReader:", err)
		return
	}
	doc.Find("script").Each(func(i int, selection *goquery.Selection) {
		if strings.Index(selection.Text(), "window._itemInfo") >= 0 && firstItemInfoScript {
			scriptCode += fmt.Sprintf("%s\r", selection.Text())
			firstItemInfoScript = false
		}

		if strings.Index(selection.Text(), "window._itemOnly") >= 0 && firstItemOnlyScript {
			scriptCode += fmt.Sprintf("%s\r", selection.Text())
			firstItemOnlyScript = false
		}
	})
	//js文件保存地址
	saveJSFilePath := strings.ReplaceAll(saveHTMLFilePath, ".html", ".js")
	writeFile(saveJSFilePath, []byte(scriptCode))
	//通过js文件内容分离对象到json
	jsVm := otto.New()
	_, err = jsVm.Run(scriptCode)
	if err != nil {
		fmt.Println("otto vmValue:", err)
		return
	}
	//_itemOnly 取值
	_itemOnlyValue, err := jsVm.Run(`JSON.stringify(window._itemOnly);`)
	if err != nil {
		fmt.Println("window._itemOnly:", err)
		return
	}
	_itemOnlyValueString, err := _itemOnlyValue.ToString()
	if err != nil {
		fmt.Println("_itemOnlyValueString:", err)
		return
	}
	_itemOnlyValueStringJsonFilePath := strings.ReplaceAll(saveHTMLFilePath, ".html", "_itemOnly.json")
	writeFile(_itemOnlyValueStringJsonFilePath, []byte(_itemOnlyValueString))

	//_itemInfo 取值
	_itemInfoValue, err := jsVm.Run(`JSON.stringify(window._itemInfo);`)
	if err != nil {
		fmt.Println("window._itemInfo:", err)
		return
	}
	_itemInfoValueString, err := _itemInfoValue.ToString()
	if err != nil {
		fmt.Println("_itemInfoValueString:", err)
		return
	}
	_itemInfoValueStringJsonFilePath := strings.ReplaceAll(saveHTMLFilePath, ".html", "_itemInfo.json")
	if debug {

	}
	writeFile(_itemInfoValueStringJsonFilePath, []byte(_itemInfoValueString))

	//解析到struct
	itemOnlyValue := new(item.JdItemOnly)
	err = json.Unmarshal([]byte(_itemOnlyValueString), itemOnlyValue)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(_itemOnlyValueString), itemOnlyValue):", err)
		return
	}
	itemInfoValue := new(item.JdItemInfo)
	err = json.Unmarshal([]byte(_itemInfoValueString), itemInfoValue)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(_itemInfoValueString), itemInfoValue):", err)
		return
	}
	//入库
	item.SaveJdItemModel(itemOnlyValue, itemInfoValue)

}

func writeFile(filename string, data []byte) error {
	if !debug {
		return nil
	}
	err := ioutil.WriteFile(filename, data, 0755)
	if err != nil {
		fmt.Println(filename, ":", err)
	}
	return err
}

func getItemIdList() {
	var idList []string
	for _, u := range itemUrlList {
		u = strings.ReplaceAll(u, "#comment", "")
		//根据「/」分隔
		uList := strings.Split(u, "/")
		idList = append(idList, strings.ReplaceAll(uList[len(uList)-1], ".html", ""))
	}
	itemIdList = removeRepeatedElement(idList)
}

func getProxy(_ *http.Request) (*url.URL, error) {
	return url.Parse(httpProxyUrl)
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

//解析flag参数
func initFlag() {
	flag.StringVar(&keyWord, "keyWord", "商用餐厨具", "关键词")
	flag.StringVar(&httpProxyUrl, "httpProxyUrl", "", "http代理地址 eg. http://58.218.200.229:4736")
	flag.StringVar(&mysqlConn, "mysqlConn", "root:root@tcp(127.0.0.1:3306)/jd_data?charset=utf8&parseTime=True&loc=Local", "mysql连接地址 eg. root:root@tcp(127.0.0.1:3306)/jd_data?charset=utf8&parseTime=True&loc=Local")
	flag.IntVar(&maxPage, "maxPage", 30, "最大页码")
	flag.IntVar(&maxPool, "maxPool", 4, "最大并发数")
	flag.Parse()
}
