package beegoLog

import (
	"bufio"
	"fmt"
	_ "github.com/influxdata/influxdb1-client"
	influxDBClient "github.com/influxdata/influxdb1-client/v2"
	"io"
	"net"
	"time"
)

type beegoLogTcpServer struct {
	listenTcpPort string     //tcp端口 :7020
	storage       LogStorage //日志存储方式
	debug         bool       //打开debug 屏显
	autoReopen    bool       //自动重启
}

type LogStorage interface {
	Init()
	SaveLog(logStr string) error
}

//初始化Server
func NewBeegoLogTcpServer(listenTcpPort string, storage LogStorage) *beegoLogTcpServer {
	return &beegoLogTcpServer{
		listenTcpPort: listenTcpPort,
		storage:       storage,
	}
}

//开启debug
func (ser *beegoLogTcpServer) Debug() {
	ser.debug = true
}

//自动从新监听tcp
func (ser *beegoLogTcpServer) AutoReopen() {
	ser.autoReopen = true
}

//日志
func (ser *beegoLogTcpServer) log(a ...interface{}) {
	if ser.debug {
		fmt.Println(a)
	}
	return
}

//监听tcp
func (ser *beegoLogTcpServer) listenTcp() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ser.listenTcpPort)
	if err != nil {
		ser.log("ResolveTCPAddr:", err)
		return err
	}
	tcpListener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		ser.log("ListenTCP:", err)
		return err
	}
	defer tcpListener.Close()

	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			ser.log("AcceptTCP:", err)
			continue
		}
		ser.log("client:", tcpConn.RemoteAddr().String())
		//接受到客户端，处理交给pipe
		go ser.pipe(tcpConn)
	}
}

//客户端内容处理
func (ser *beegoLogTcpServer) pipe(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	defer conn.Close()
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				ser.log("ReadString:", err)
			}
			break
		}
		ser.log("get message:", message)
		//处理日志
		err = ser.storage.SaveLog(message)
		if err != nil {
			ser.log("SaveLog:", err)
			break
		}
	}
}

//开启服务
func (ser *beegoLogTcpServer) Run() {
	//初始化storage
	ser.storage.Init()
	for {
		err := ser.listenTcp()
		if err != nil {
			fmt.Println("listenTcp is error: ", err)
		}
		//未开启自动重连
		if !ser.autoReopen {
			break
		}
		//延迟重启
		time.Sleep(5 * time.Second)
	}
}

type InfluxDBLogStorage struct {
	influxDBUrl     string
	influxDBName    string
	influxTableName string
	influxDBClient  influxDBClient.Client
}

func NewInfluxDBLogStorage(influxDBUrl, influxDBName, influxTableName string) *InfluxDBLogStorage {
	return &InfluxDBLogStorage{
		influxDBUrl:     influxDBUrl,
		influxDBName:    influxDBName,
		influxTableName: influxTableName,
	}
}

func (s *InfluxDBLogStorage) Init() {
	c, err := influxDBClient.NewHTTPClient(influxDBClient.HTTPConfig{
		Addr: s.influxDBUrl,
	})
	if err != nil {
		panic(err)
	}
	s.influxDBClient = c
	s.createDatabase()
	go s.healthy()
}

func (s *InfluxDBLogStorage) SaveLog(logStr string) error {
	if logStr == "" {
		return fmt.Errorf("logStr is empty")
	}
	bp, _ := influxDBClient.NewBatchPoints(influxDBClient.BatchPointsConfig{
		Database:  s.influxDBName,
		Precision: "ns",
	})

	fields := map[string]interface{}{
		"log": logStr,
	}
	pt, err := influxDBClient.NewPoint(s.influxTableName, nil, fields, time.Now())
	if err != nil {
		fmt.Println("influxDBClient.NewPoint: ", err.Error())
		return err
	}
	bp.AddPoint(pt)
	// Write the batch
	err = s.influxDBClient.Write(bp)
	if err != nil {
		fmt.Println("influxDBClient.Write: ", err.Error())
		return err
	}
	return nil
}

//创建数据库
func (s *InfluxDBLogStorage) createDatabase() {
	//创建数据库
	_, _ = s.influxDBClient.Query(influxDBClient.NewQuery(fmt.Sprintf("create database %s", s.influxDBName), "", ""))
}

//定时查询内容，判断健康状态
func (s *InfluxDBLogStorage) healthy() {
	for {
		time.Sleep(5 * time.Second)
		q := influxDBClient.NewQuery(fmt.Sprintf("SELECT count(log) FROM %s", s.influxTableName), s.influxDBName, "ns")
		response, err := s.influxDBClient.Query(q)
		if err != nil {
			fmt.Println("influxDBClient.Query:", err)
			continue
		}
		if response.Error() != nil {
			fmt.Println("response.Error():", response.Error())
			continue
		}
		fmt.Println("influxDB response.Results:", response.Results)
	}
}
