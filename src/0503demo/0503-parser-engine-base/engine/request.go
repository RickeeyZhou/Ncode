package engine

// url string
// 解析器 方法

//func (b []byte) 类型
//CityListParser 变量名
type Request struct {
	Url string  // 每一个请求对应的Url
	ParserFunc func (b []byte) // 每一个页面取到的数据都传给解析器进行解析
}
