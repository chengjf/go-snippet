package main

import "fmt"
import "os"
import "bufio"
import "golang.org/x/text/encoding/simplifiedchinese"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
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
