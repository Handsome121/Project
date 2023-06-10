package main

import "time"

//context.Context是一个接口，该接口定义了四个需要实现的方法,具体签名如下:
type Context interface {
	Deadline() (deadline time.Time, ok bool) //Deadline方法需要返回当前Context被取消的时间，也就是完成工作的截止时间
	Done() <-chan struct{} // Done方法需要返回一个Channel，这个Channel会在当前工作完成或者上下文被取消之后关闭，
	// 多次调用Done方法会返回同一个Channel；
	Err() error //Err方法会返回当前Context结束的原因，它只会在Done返回的Channel被关闭时才会返回非空的值；
	//如果当前Context被取消就会返回Canceled错误；
	//如果当前Context超时就会返回DeadlineExceeded错误；
	Value(key interface{}) interface{} //Value方法会从Context中返回键对应的值，对于同一个上下文来说，
	// 多次调用Value 并传入相同的Key会返回相同的结果，
	// 该方法仅用于传递跨API和进程间跟请求域的数据；
}




