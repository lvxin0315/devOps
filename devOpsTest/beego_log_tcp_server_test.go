package devOpsTest

import (
	"github.com/lvxin0315/devOps/beegoLog"
	"testing"
)

const (
	influxDBUrl     = "http://127.0.0.1:8086"
	influxDBName    = "qingGoUCenter"
	influxTableName = "beegoLog"
	listenTcpPort   = ":7020"
)

func TestBeegoLogTcpServer(t *testing.T) {
	logStorage := beegoLog.NewInfluxDBLogStorage(influxDBUrl, influxDBName, influxTableName)
	ser := beegoLog.NewBeegoLogTcpServer(listenTcpPort, logStorage)
	ser.Debug()
	ser.AutoReopen()
	ser.Run()
}
