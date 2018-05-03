package fetcher

import (
	"net/http"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"io"
	"golang.org/x/text/encoding"
	"bufio"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/net/html/charset"
)

func Fetcher(url string)([]byte,error){
	resp,err := http.Get(url)
	if err!=nil{
		return nil,err
	}
	defer resp.Body.Close() ///一旦链接打开,也需要关上
	
	if resp.StatusCode !=http.StatusOK{
		// fmt.Println("Error: status code,",resp.statuscode
		return nil, fmt.Errorf("Error:Status code:%d",resp.StatusCode)
	}

	// 动态判断编码

	e:= determinEncoding(resp.Body)
	utf8Reader:=transform.NewReader(resp.Body,e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}
func determinEncoding(r io.Reader) encoding.Encoding{
	bytes,err:=bufio.NewReader(r).Peek(1024)
	if err!= nil {
		return unicode.UTF8
	}
	e,_,_:= charset.DetermineEncoding(bytes,"")
	return e

}