package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"log"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

func main() {
	fmt.Println("######################################")
	read1("./hosts")
	session, err := connect("jumpwbx", "befYc4xgK80ZsN94p078sitBC0OGlt", "221.228.108.29", 22)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Run("ipconfig | grep IPv4")

}

//文件读2
func read1(path string) {
	fi, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	con := bufio.NewReader(fi)
	if err != nil {
		panic(err)
	}
	for {
		line, err := con.ReadString('\n')
		if err == io.EOF {
			return
		}
		if match, _ := regexp.MatchString("Gate", line); match {
			a := strings.FieldsFunc(line, Split)
			fmt.Println(a[0])

		}
	}

}

//多分隔符分割数据
func Split(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}

//文件读3
func read3(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)

	fmt.Printf("%s", fd)
	return string(fd)
}

//文件读1
func read0(path string) string {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%s\n", err)
		panic(err)
	}
	return string(f)
}

//SSH函数
func connect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)

	clientConfig = &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		Timeout: 5 * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}
