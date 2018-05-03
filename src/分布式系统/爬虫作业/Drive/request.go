package Drive


type Request struct{
	Url string
	Parserfunc func (b []byte)(r RequestResult)//
	//链接解析以后,会产生其他的数据

}

type RequestResult struct{
	Items []interface{} //接受任意类型的数据
	R []Request         //说明解析结果会产生另外的请求
}

func NilParser( b []byte )(r RequestResult){
	return r
}    //再一次解析,有可能产生另外的请求