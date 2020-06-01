### 将阿里云的RDS备份，放到mysql:5.7 的docker容器中
1. Dockerfile 中使用的是mysql:5.7的镜像，所以在启动时需要；
2. hins_data_qp.xb是阿里云RDS备份文件，请自行下载并改名，或者修改dockerfile；
3. setup1.sh 里面是修改阿里云源的，如果网络OK也可以不修改，请自行修改dockerfile；

### 运行
```shell
docker build -t demo:0.6 .

docker run -p 3308:3306 -e MYSQL_ROOT_PASSWORD=123456 -d demo:0.6
```