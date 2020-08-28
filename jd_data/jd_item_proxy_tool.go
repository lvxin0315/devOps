package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var testHttpProxyUrl = []string{
	"http://123.56.75.226:3128",
	"http://124.93.201.59:42672",
	"http://124.93.201.59:42672",
	"http://175.148.70.85:1133",
	"http://221.180.170.104:8080",
	"http://218.60.8.83:3129",
	"http://119.108.177.21:9000",
	"http://59.44.78.30:54069",
	"http://218.60.8.83:3129",
	"http://119.119.106.181:9000",
	"http://123.56.118.36:8080",
	"http://124.93.201.59:42672",
	"http://119.119.227.167:9000",
	"http://218.60.8.99:3129",
	"http://218.60.8.83:3129",
	"http://119.119.251.26:9000",
	"http://119.119.100.76:9000",
	"http://59.44.78.30:54069",
	"http://119.119.249.32:9000",
	"http://218.60.8.99:3129",
	"http://119.119.249.32:9000",
	"http://218.60.8.99:3129",
	"http://221.180.170.104:8080",
	"http://119.119.232.216:9000",
	"http://221.180.170.104:8080",
	"http://119.119.232.216:9000",
	"http://218.60.8.99:3129",
	"http://123.56.75.226:3128",
	"http://218.60.8.99:3129",
	"http://123.56.75.226:3128",
	"http://221.180.170.104:8080",
	"http://223.100.166.3:369",
	"http://59.44.78.30:54069",
	"http://119.119.246.251:9000",
	"http://221.180.170.104:8080",
	"http://221.180.170.104:8080",
	"http://123.56.75.226:3128",
	"http://123.56.75.226:3128",
	"http://119.108.187.138:9000",
	"http://119.108.187.138:9000",
	"http://119.119.224.205:9000",
	"http://119.119.224.205:9000",
	"http://218.60.8.83:3129",
	"http://218.60.8.83:3129",
	"http://119.119.227.76:9000",
	"http://119.119.227.76:9000",
}

func main() {
	//判断代理OK不OK
	for _, u := range testHttpProxyUrl {
		_, err := proxyGet(u)
		if err != nil {
			fmt.Println("err:", err)
		} else {
			fmt.Println(u)
		}
	}
}

//get 请求
func proxyGet(httpProxyUrl string) ([]byte, error) {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: func(_ *http.Request) (*url.URL, error) {
				return url.Parse(httpProxyUrl)
			},
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequest("GET", "https://www.jd.com", nil)
	if err != nil {
		return nil, err
	}
	// 自定义Header
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	//函数结束后关闭相关链接
	defer resp.Body.Close()
	//状态码
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code not 200")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
