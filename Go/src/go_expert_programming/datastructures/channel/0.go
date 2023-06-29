package main

func main() {
	//在 Go 语言中，对通道（channel）的操作可能会导致 panic 的几种情况包括：
	//操作	    nil channel	       closed channel	        not nil, not closed channel
	//close	     panic	           panic	                正常关闭
	//读 <- ch	 阻塞	           读到对应类型的零值	        阻塞或正常读取数据。缓冲型 channel 为空或非缓冲型 channel 没有等待发送者时会阻塞
	//写 ch <-	 阻塞	           panic	                阻塞或正常写入数据。非缓冲型 channel 没有等待接收者或缓冲型 channel buf 满时会被阻塞
	//
	//总结一下，发生 panic 的情况有三种：向一个关闭的 channel 进行写操作；关闭一个 nil 的 channel；重复关闭一个 channel。
	//这些情况下，Go 语言的运行时会引发 panic，以提醒开发者出现了不合理的通道操作。因此，在使用通道时，要特别注意处理好通道的关闭、缓冲情况以及协调发送方和接收方的操作，以避免发生 panic。
}
