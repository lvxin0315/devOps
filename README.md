### 做这个devOps是为了什么?
首先这个仅仅是个日常运维小工具，主要是方便开发人员用自己擅长的代码去处理常用的命令及操作。

### devOps都用了什么？
* github.com/bndr/gojenkins
* mysqldump
* expect
* RDS_for_docker 用docker还原RDS物理备份 「查看详情」


### devOps有哪些功能?
* jenkins build job
```go
jenkinsConn, err := devOps.CreateJenkinsConn(
    "http://jenkins.local.com/",
    "admin",
    "111111")
if err != nil {
    return
}
//build
err = jenkinsConn.BuildJobWithGitParam(
    "testJenkinsApi",
    "any",
    "v1.0.201909111800")
if err != nil {
    return
}
```

* scp 拉取文件
```go
scpInfo := devOps.NewScpInfo()
scpInfo.Host = "172.16.0.225"
scpInfo.Port = 22
scpInfo.User = "jjc"
scpInfo.Password = "1"
scpInfo.ReadFilePath = "~/abc"
scpInfo.LocalSaveFilePath = fmt.Sprintf("/data/%d", time.Now().UnixNano())
scpInfo.Do()
```

* ssh 执行命令
```go
sshClient := devOps.NewSSHClient()
sshClient.User = "jjc"
sshClient.Password = "1"
sshClient.Host = "172.16.0.225"
sshClient.Port = 22

err := sshClient.SSHConnect()
defer sshClient.Close()
if err != nil {
    panic(err)
}

var stdOut, stdErr bytes.Buffer
//设置输出内容
sshClient.StdOut(&stdOut)
sshClient.StdErr(&stdErr)
//执行cmd
err = sshClient.DoCmd("ls -ll")
if err != nil {
    panic(err)
}
fmt.Println("stdOut:", stdOut.String())

fmt.Println("stdErr:", stdErr.String())

```

* mysqldump备份指定表
```go
dumper := devOps.NewDumper()
dumper.Host = "172.16.0.225"
dumper.User = "root"
dumper.Password = "root"
dumper.MysqlDumpPath = "mysqldump"
dumper.Database = "xxxxdatabas"
dumper.SavePath = fmt.Sprintf("/data/%d", time.Now().UnixNano())
dumper.TableNameList = []string{"a", "company"}
dumper.Do()
```

* docker swarm http api
```go
os.Setenv("DOCKER_HOST", "http://172.16.0.228:2375")
os.Setenv("DOCKER_API_VERSION", "1.39")
swarmCli, err := devOps.NewSwarmCli()
if err != nil{
    t.Error(err)
}
//创建一个network
res, err := swarmCli.CreateSwarmNetwork("test123123")
if err != nil{
    t.Error(err)
}
```