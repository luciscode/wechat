package main

import (
	"fmt"
	// "io/ioutil"
	"bufio"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("hellowrod")
	path, err := os.Getwd()
	fr, err := os.Open(path + "/post/hello-world.md")
	fw, _ := os.Create(path + "/html/hello-world.html")
	defer fr.Close()
	defer fw.Close()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fr.Name())
		rd := bufio.NewReader(fr)
		for {
			line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
			splite := strings.Split(line, " ")
			if splite[0] == "##" {

				line = "<h1>" + strings.TrimPrefix(line, splite[0]) + "</h1>"

			} else if splite[0] == "###" {
				line = "<h2>" + strings.TrimPrefix(line, splite[0]) + "</h2>"

			} else if splite[0] == "title:" {
				line = "<h1>" + strings.TrimPrefix(line, splite[0]) + "</h1>"

			}

			fw.Write([]byte(line))
			if err != nil || io.EOF == err {
				break
			}
		}

	}
}
