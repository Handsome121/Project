package main

//用 chan struct{} 来传递信号, chan bool 表达的不够清楚

//当你在结构中看到 chan bool 的定义时，有时不容易理解如何使用该值，例如：

type Service1 struct {
	deleteCh chan bool // what does this bool mean?
}
//但是我们可以将其改为明确的 chan struct {} 来使其更清楚：我们不在乎值（它始终是 struct {}），我们关心可能发生的事件，例如：

type Service2 struct {
	deleteCh chan struct{} // ok, if event than delete something.
}