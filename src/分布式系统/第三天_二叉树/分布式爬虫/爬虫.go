package main
// 动态判断编码
// go get golang.org/x/net/html

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"golang.org/x/text/transform"
	"io"
	"golang.org/x/text/encoding"
	"golang.org/x/net/html/charset"
	"bufio"

	"regexp"
)

func main()  {

	resp,err := http.Get("https://www.zhenai.com/zhenghun")
	// 目标网址 http 库中 get方法
	fmt.Printf("测试输出:%v\n测试地址:%T\n",resp,resp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()  //

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:Status Code, ",resp.StatusCode)
		return
	}

	// 动态判断编码
	e := determinEncoding(resp.Body)

	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	all,err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}

	//fmt.Printf("%s\n",all)
	findUrlADD(all)
}

// 将response.Body 作为参数传入到函数中
// 会自动返回编码格式
func determinEncoding(r io.Reader) encoding.Encoding  {

	bytes , err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		panic(err)
	}
	// 查找编码格式
	e,_,_ := charset.DetermineEncoding(bytes,"")
	// 返回编码格式
	return e
}
func findUrlADD( all []byte){
	match:=regexp.MustCompile(`<a href="(.*?)"\s*class="">(.*?)</a>`)
	//s:=match.FindAllString(string(all),-1)
	b:=match.FindAllSubmatch(all,-1)
	//fmt.Println(s)
	//fmt.Printf("%s",b)
	for _,value:=range b{
		fmt.Printf("%s,%s\n",value[1],value[2])
	}
}