
package engine
// url  string
//解析器 方法
//
type Request struct{
	Url string
	ParseFunc func (b []byte)
}
type RequestResult struct{
	Items []interface{} ///存储解析器解析出来的数据
	result []Request
}