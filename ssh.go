package devOps

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"net"
)

type sshClient struct {
	Host       string
	User       string
	Password   string
	Port       int
	sshSession *ssh.Session
}

func NewSSHClient() *sshClient {
	return new(sshClient)
}

func (s *sshClient) SSHConnect() error {
	// get auth method
	auth := make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(s.Password))
	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	clientConfig := &ssh.ClientConfig{
		User: s.User,
		Auth: auth,
		// Timeout:             30 * time.Second,
		HostKeyCallback: hostKeyCallbk,
	}
	// connet to ssh
	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	client, err := ssh.Dial("tcp", addr, clientConfig)
	if err != nil {
		errorLog(err.Error())
		return err
	}
	// create session
	session, err := client.NewSession()
	if err != nil {
		errorLog(err.Error())
		return err
	}
	s.sshSession = session
	return nil
}

func (s *sshClient) StdOut(stdOut io.Writer) {
	s.sshSession.Stdout = stdOut
}

func (s *sshClient) StdErr(stdErr io.Writer) {
	s.sshSession.Stderr = stdErr
}

func (s *sshClient) DoCmd(cmdString string) error {
	infoLog("命令内容：")
	info2Log(cmdString)
	err := s.sshSession.Run(cmdString)
	if err != nil {
		errorLog(err.Error())
		return err
	}
	return nil
}

func (s *sshClient) Close() {
	err := s.sshSession.Close()
	if err != nil {
		panic(err)
	}
}
