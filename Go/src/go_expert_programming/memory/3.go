package main

//逃逸分析:是指由编译器决定内存分配的位置，不需要程序员指定，在函数中申请一个新的对象:
//如果分配在栈中，则函数执行结束后可自动将内存回收

//逃逸策略:
//1、如果函数外部没有引用，则优先放到栈中
//2、如果函数外部存在引用，则必定放到堆中

//逃逸场景:
//1、指针逃逸
//2、栈空间不足逃逸
//3、动态类型逃逸
//4、闭包引用对象逃逸


//小结
//1、栈上分配内存比堆中分配内存有更高的效率
//2、栈上分配的内存不需要GC处理
//3、堆上分配的内存需要GC处理
//4、逃逸分析的目的是决定分配地址是堆还是栈
//5、逃逸分析在编译阶段完成



