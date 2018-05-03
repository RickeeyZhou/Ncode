package engine

// url string
// 解析器 方法

//func (b []byte) 类型
//CityListParser 变量名
type Request struct {
	Url string  // 每一个请求对应的Url
	ParserFunc func (b []byte) (r RequestResult) // 每一个页面取到的数据都传给解析器进行解析
}

// 每一个请求返回回来的结构体对象
type RequestResult struct {
	Items []interface{} //存储解析器解析出来的数据
	R []Request //数据请求数组
}

// 空的解析器
func NilParser(b []byte) (r RequestResult) {
	return r
}