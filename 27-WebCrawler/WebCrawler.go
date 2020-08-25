package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

var finishChan chan int

func init() {
	finishChan = make(chan int)
}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func httpGet(url string) (data string, err error) {
	resp, e := http.Get(url)
	defer resp.Body.Close()
	if e != nil {
		err = e
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("error")
		return
	}

	// 解决读取的网页内容中文乱码问题
	encoding := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body, encoding.NewDecoder())
	bytes, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	data = string(bytes)
	return
}

func crawlProdFromJd(pageNo int) {
	url := "https://search.jd.com/Search?keyword=iphone%20xs%20max&enc=utf-8&qrst=1&rt=1&stop=1&vt=2&wq=iphone%20xs&click=0&page="
	url += strconv.Itoa(pageNo + (pageNo - 1))
	data, err := httpGet(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	priceCompile := regexp.MustCompile(`<div class="p-price">(?s:(.*?))</div>`)
	submatch := priceCompile.FindAllStringSubmatch(data, -1)
	for _, match := range submatch {
		fmt.Println(match)
	}
	finishChan <- pageNo
}

// 京东商品爬虫
func main() {
	startPage := 1
	endPage := 1
	for i := startPage; i <= endPage; i++ {
		go crawlProdFromJd(i)
	}
	for i := startPage; i <= endPage; i++ {
		pageNo := <-finishChan
		fmt.Println("第", pageNo, "页数据处理完毕")
	}
}
