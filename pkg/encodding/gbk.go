package main

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func test() {
	file, err := os.Open("./aa.csv")
	check(err)
	// bufferedReader := bufio.NewReader(file)
	// str, err := bufferedReader.ReadString('\n')
	// check(err)
	// fmt.Println(str)
	gbk := simplifiedchinese.GBK
	decoder := gbk.NewDecoder()
	gbkReader := decoder.Reader(file)
	bufferedReader2 := bufio.NewReader(gbkReader)
	s, err := bufferedReader2.ReadString('\n')
	check(err)
	fmt.Println(s)
}

// 将一个io.Reader使用GBK解码，转为bufio.Reader
func getGBKBufReader(reader io.Reader) (result *bufio.Reader) {
	gbk := simplifiedchinese.GBK
	gbkReader := gbk.NewDecoder().Reader(reader)
	result = bufio.NewReader(gbkReader)
	return
}

// 将一个io.Reader使用GBK解码，转为io.Reader
func getGBKReader(reader io.Reader) (result io.Reader) {
	gbk := simplifiedchinese.GBK
	result = gbk.NewDecoder().Reader(reader)
	return
}

func main() {
	file, err := os.Open("./aa.csv")
	check(err)
	bufferedReader := bufio.NewReader(file)
	gbkReader := getGBKBufReader(bufferedReader)
	s, err := gbkReader.ReadString('\n')
	check(err)
	fmt.Println(s)

	file, err = os.Open("./aa.csv")
	check(err)
	l := getGBKReader(file)
	buf := make([]byte, 1024)
	o, p:= l.Read(buf)
	fmt.Printf("%v", buf)
	fmt.Printf("%v %v %v \n", o, p, string(buf))

}
