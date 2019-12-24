package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("######################################")
	read1("./hosts")

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
