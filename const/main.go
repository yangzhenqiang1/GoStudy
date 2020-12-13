package main

func main() {
	/*相对于变量，常量是恒定不变的值，多用于定义程序运行期间不会改变的那些值。
	常量的声明和变量声明非常类似，只是把var换成了const，常量在定义的时候必须赋值。*/
	const pi = 3.1415
	const e = 2.7182

	iota是go语言的常量计数器，只能在常量的表达式中使用。

// iota在const关键字出现时将被重置为0。
// const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。

// 举个例子：

const (
		n1 = iota //0
		n2        //1
		n3        //2
		n4        //3
	)


	/*
	几个常见的iota示例:
使用_跳过某些值

const (
		n1 = iota //0
		n2        //1
		_
		n4        //3
	)
iota声明中间插队

const (
		n1 = iota //0
		n2 = 100  //100
		n3 = iota //2
		n4        //3
	)
	const n5 = iota //0
定义数量级 （这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，也就是十进制的1024。同理2<<2表示将2的二进制表示向左移2位，也就是由10变成了1000，也就是十进制的8。）

const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)
多个iota定义在一行

const (
		a, b = iota + 1, iota + 2 //1,2
		c, d                      //2,3
		e, f                      //3,4
	)
	*/
}
