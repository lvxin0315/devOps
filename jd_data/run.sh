#!/bin/bash

#keywords=("商用餐厨具" "商用冰柜" "商用制冰机")

#keywords=("商用料理机")

keywords=("商用料理机" "厨具套装" "厨柜" "刀具" "菜刀" "锅具" "厨房用品" "炒锅" "厨房" "自营厨具" "锅铲" "厨具架" "厨具收纳")
#keywords=("炒锅" "厨房" "自营厨具" "锅铲" "厨具架" "厨具收纳")
#keywords=("一次性餐具四件套" "一次性筷子" "一次性餐具套装")
#keywords=("一次性碗" "一次性餐盒" "一次性餐具" "一次性勺子")


#proxyUrl="http://221.180.170.104:8080"

#218.60.8.99:3129
#221.180.170.104:8080
#59.44.78.30:54069
proxyUrl="http://127.0.0.1:1087"


# shellcheck disable=SC2068
for k in ${keywords[@]}
do
  echo "$k"
  go run jd_item_download.go \
-httpProxyUrl="$proxyUrl" \
-keyWord="$k" \
-mysqlConn="root:root@tcp(127.0.0.1:3306)/jd_data?charset=utf8&parseTime=True&loc=Local"

done