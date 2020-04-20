package core

import (
	"fmt"
	"github.com/bndr/gojenkins"
	"time"
)

type JenkinsConn struct {
	jenkins *gojenkins.Jenkins
}

//连接jenkins
func CreateJenkinsConn(jenkinsUrl, jenkinsAccount, jenkinsPassword string) (jenkinsConn *JenkinsConn, err error) {
	jenkinsConn = new(JenkinsConn)
	jenkinsConn.jenkins = gojenkins.CreateJenkins(nil, jenkinsUrl, jenkinsAccount, jenkinsPassword)
	//输出jobs信息
	allJobs, err := jenkinsConn.jenkins.GetAllJobNames()
	if err != nil {
		errorLog("CreateJenkins Error!:", err.Error())
		return jenkinsConn, err
	}
	fmt.Println(yellowBg, "jobs信息:", reset)
	for jobNum, job := range allJobs {
		infoLog(fmt.Sprintf("%d. %s", jobNum, job.Name))
	}
	return jenkinsConn, err
}

//build job
func (jenkinsConn *JenkinsConn) BuildJobWithGitParam(jobName, paramName, paramValue string) error {
	jobId, err := jenkinsConn.jenkins.BuildJob(jobName, map[string]string{paramName: paramValue})
	if err != nil {
		errorLog("BuildJobWithGitParam Error!:", err.Error())
		return err
	}
	for {
		task, err := jenkinsConn.jenkins.GetQueueItem(jobId)
		if err != nil {
			errorLog("GetQueueItem Error!:", err.Error())
			break
		}
		time.Sleep(500 * time.Millisecond)
		if task.Raw.Why == "" {
			info2Log("task over")
			break
		} else {
			infoLog(task.Raw.Why)
		}
	}
	return nil
}
