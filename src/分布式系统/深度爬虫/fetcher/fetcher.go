package fetcher

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
)

//
//func Fetcherdata(url string ) ([]byte,error) {
//
//	resp,err := http.Get(url)
//	// 目标网址 http 库中 get方法
//	fmt.Printf("测试输出:%v\n测试地址:%T\n",resp,resp)
//	if err != nil {
//		panic(err)
//	}
//	defer resp.Body.Close()  //
//
//	if resp.StatusCode != http.StatusOK {
//		//fmt.Println("Error:Status Code, ",resp.StatusCode)
//		return nil,fmt.Errorf("Error:status Code:%d",resp.StatusCode)
//	}
//
//	// 动态判断编码
//	e := determinEncoding(resp.Body)
//
//	utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())
//	all,err := ioutil.ReadAll(utf8Reader)
//	return all,err
//
//
//	//fmt.Printf("%s\n",all)
////	findUrlADD(all)
//}
//
//// 将response.Body 作为参数传入到函数中
//// 会自动返回编码格式
//func determinEncoding(r io.Reader) encoding.Encoding  {
//
//	bytes , err := bufio.NewReader(r).Peek(1024)
//	if err != nil {
//	//	panic(err)
//	return unicode.UTF8
//	}
//	// 查找编码格式
//	e,_,_ := charset.DetermineEncoding(bytes,"")
//	// 返回编码格式
//	return e
//}
////func findUrlADD( all []byte){
////	match:=regexp.MustCompile(`<a href="(.*?)"\s*class="">(.*?)</a>`)
////	//s:=match.FindAllString(string(all),-1)
////	b:=match.FindAllSubmatch(all,-1)
////	//fmt.Println(s)
////	//fmt.Printf("%s",b)
////	for _,value:=range b{
////		fmt.Printf("%s,%s\n",value[1],value[2])
////	}
////}
func FetchData(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {

		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//fmt.Println("Error:Status Code, ",resp.StatusCode)
		return nil, fmt.Errorf("Error:Status Code:%d", resp.StatusCode)
	}

	// 动态判断编码
	e := determinEncoding(resp.Body)

	utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())
	// ([]byte, error)
	return ioutil.ReadAll(utf8Reader)

}

// 将response.Body 作为参数传入到函数中
// 会自动返回编码格式
func determinEncoding(r io.Reader) encoding.Encoding {

	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		//panic(err)
		return unicode.UTF8
	}
	// 查找编码格式
	e, _, _ := charset.DetermineEncoding(bytes, "")
	// 返回编码格式
	return e
}