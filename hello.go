package main


/*

%+v  {X:3 Y:4}
%v   {3 4}
%q   {'\x03' '\x04'}

%T   打印类型名

*/

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/cmplx"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/zhanxunchang/go-hello/morestrings"
	//"github.com/google/go-cmp/cmp"
	//"golang.org/x/tour/pic"
)



// 类型在变化名后面
// 原因参考 [这篇关于 Go 语法声明的文章](http://blog.go-zh.org/gos-declaration-syntax)
func add(x int, y int) int {
	return x + y
}

// 当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略
func add2(x, y int) int {
	return x + y
}

// 函数可以返回任意数量的返回值
func swap(x, y string) (string, string) {
	return y, x
}

// Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。
// 返回值的名称应当具有一定的意义，它可以作为文档使用。
// 没有参数的 return 语句返回已命名的返回值。也就是 直接 返回。
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 变量
// var 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后。
// 就像在这个例子中看到的一样，var 语句可以出现在包或函数级别。
var c, python, java bool

func variableDemo() {
	var i int
	fmt.Println(i, c, python, java)
}

// 变量的初始化
// 变量声明可以包含初始值，每个变量对应一个。
var i, j int = 1, 2

func variableDemo2() {
	// 如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

// 短变量声明
// *在函数中*，简洁赋值语句 := 可在类型明确的地方代替 var 声明。
// 函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。
func variableDemo3() {
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(k, c, python, java)
}

/*
基本类型
Go 的基本类型有

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名，表示一个 Unicode 码点

float32 float64

complex64 complex128
本例展示了几种类型的变量。 同导入语句一样，变量声明也可以“分组”成一个语法块。

int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。 当你需要一个整数值时应使用 int 类型，除非你有特殊的理由使用固定大小或无符号的整数类型。
*/

var (
	ToBe   bool       = false
	MaxUint64 uint64  = 1<<64 - 1
	MaxInt64 int64    = 1<<63 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	name   string     = "jason"
)

func basicTypeDemo() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxUint64, MaxUint64)
	fmt.Printf("Type: %T Value: %v\n", MaxInt64, MaxInt64)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	fmt.Printf("Type: %T Value: %v\n", name, name)
}

func runeDemo() {
	// 字符串长度
	var str = "hello 你好"
	// 12, golang中string底层是通过byte数组实现的，所以直接求len 实际是在按字节长度计算  所以一个汉字占3个字节算了3个长度
	fmt.Println("len(str):", len(str))
	// 8，实际字符长度
	fmt.Println("len of rune:", len([]rune(str)))
}

/*
零值
没有明确初始值的变量声明会被赋予它们的 零值。

零值是：

数值类型为 0，
布尔类型为 false，
字符串为 ""（空字符串）
*/
func defaultVal() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v [%v] [%q]\n", i, f, b, s, s) // %q会加上双引号
}

/*
类型转换
表达式 T(v) 将值 v 转换为类型 T。

一些关于数值的转换：

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
或者，更加简单的形式：

i := 42
f := float64(i)
u := uint(f)
与 C 不同的是，Go 在不同类型的项之间赋值时需要显式转换。试着移除例子中 float64 或 uint 的转换看看会发生什么。
*/

/*
常量
常量的声明与变量类似，只不过是使用 const 关键字。

常量可以是字符、字符串、布尔值或数值。

常量不能用 := 语法声明。
*/
const Pi = 3.14
func constDemo() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

const (
	FlagMask1 = 1 << iota
	FlagMask2
	FlagMask3
	FlagMask4
)

func constDemo2()  {
	// FlagMask: 1 2 4 8
	fmt.Println("FlagMask:", FlagMask1, FlagMask2, FlagMask3, FlagMask4)
}

//---------------------------------
/*
Go 只有一种循环结构：`for` 循环
Go 的 for 语句后面的三个构成部分外没有小括号， 大括号 `{ }` 则是必须的。

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}


初始化语句和后置语句是可选的。
func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}


for 是 Go 中的 “while”
此时你可以去掉分号，因为 C 的 while 在 Go 中叫做 for。
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}



无限循环
如果省略循环条件，该循环就不会结束，因此无限循环可以写得很紧凑。
func main() {
	for {
	}
}
*/

/*

if
Go 的 if 语句与 for 循环类似，表达式外无需小括号 ( ) ，而大括号 { } 则是必须的。

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}


if 的简短语句
同 for 一样， if 语句可以在条件表达式前执行一个简单的语句。
该语句声明的变量作用域仅在 if 之内。

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

*/

func Sqrt(x float64) float64 {
	z := x / 2
	for i := 1; i < 10; i++ {
		t := z
		z -= (z*z - x) / (2 * z)
		fmt.Printf("%v %v\n", i, z)
		if math.Abs(t - z) < 0.1 {
			return z
		}
	}
	return z
}


/*
Go 的 switch 语句类似于 C、C++、Java、JavaScript 和 PHP 中的，不过 Go 只运行选定的 case，而非之后所有的 case。 实际上，Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。 除非以 fallthrough 语句结束，否则分支会自动终止。 Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数。
*/
func switchDemo() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Windows.")
		// fallthrough 此处加上fallthrough就不会break了
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

/*
switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。

例如，

switch i {
case 0:
case f():
}
在 i==0 时 f 不会被调用。
*/
func switchDemo2() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	fmt.Printf("%T %v\n", time.Saturday, time.Saturday)
}

/*
没有条件的 switch
没有条件的 switch 同 switch true 一样。

这种形式能将一长串 if-then-else 写得更加清晰。
*/
func switchDemo3() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}


/*
defer
defer 语句会将函数推迟到外层函数返回之后执行。

推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。
*/
func invokeByDeferParam() (string) {
	fmt.Println("secondly")
	return "finally"
}

func deferDemo() {
	fmt.Println("firstly")

	defer fmt.Println(invokeByDeferParam())

	fmt.Println("thirthly")
}


/*
defer 栈
推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
*/
func deferDemo2() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}


// ------------------------------
/*
指针
Go 拥有指针。指针保存了值的内存地址。

类型 *T 是指向 T 类型值的指针。其零值为 nil。

var p *int
*/

func pointerDemo() {
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}


/*
结构体
一个结构体（struct）就是一组字段（field）。
*/
type Vertex struct {
	X int
	Y int
}

func structDemo() {
	v := Vertex{1, 2}
	v.X = 3
	fmt.Println("struct:", v)
	
	fmt.Printf("struct: %q, %v\n", v, v)
	fmt.Printf("struct: %T, v.X=%v, v.Y= %v\n", v, v.X, v.Y)
	
	// 结构体的指针
	p := &v
	p.X = 1e9 // 隐式间接引用：p.X是(*p).X的简写，等价的
	(*p).Y = 10
	fmt.Println("struct:", v)
	
	
	/*
	结构体文法
	结构体文法通过直接列出字段的值来新分配一个结构体。

	使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）

	特殊的前缀 & 返回一个指向结构体的指针。
	*/
	var (
		v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
		v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
		v3 = Vertex{}      // X:0 Y:0
		p1  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
	)
	fmt.Println(v1, p1, v2, v3)

}


/*
数组
类型 [n]T 表示拥有 n 个 T 类型的值的数组。

表达式

var a [10]int
会将变量 a 声明为拥有 10 个整数的数组。

数组的长度是其类型的一部分，因此数组不能改变大小。这看起来是个限制，不过没关系，Go 提供了更加便利的方式来使用数组。
*/
func arrDemo() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}


/*
切片（Slices）
每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。在实践中，切片比数组更常用。

类型 []T 表示一个元素类型为 T 的切片。
切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：
a[low : high]
它会选择一个半开区间，包括第一个元素，但排除最后一个元素。
以下表达式创建了一个切片，它包含 a 中下标从 1 到 3 的元素：
a[1:4]

*/
func slicesDemo() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}


/*
切片就像数组的引用
切片并不存储任何数据，它只是描述了底层数组中的一段。

更改切片的元素会修改其底层数组中对应的元素。

与它共享底层数组的切片都会观测到这些修改。
*/
func slicesDemo2() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

/*
切片的长度与容量
切片拥有 长度 和 容量。

切片的长度就是它所包含的元素个数。

切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。

切片 s 的长度和容量可通过表达式 len(s) 和 cap(s) 来获取。

你可以通过重新切片来扩展一个切片，给它提供足够的容量。试着修改示例程序中的切片操作，向外扩展它的容量，看看会发生什么。
*/
func slicesDemo3() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)
	

	// nil切片
	var nilS []int
	fmt.Println(nilS, len(nilS), cap(nilS)) // 输出“[] 0 0“， nil打印出来是[]
	if nilS == nil {
		fmt.Println("nil!")
	}
}


func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}


/*
用 make 创建切片
切片可以用内建函数 make 来创建，这也是你创建动态数组的方式。

make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：

a := make([]int, 5)  // len(a)=5
要指定它的容量，需向 make 传入第三个参数：

b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
*/

/*
切片的切片
切片可包含任何类型，甚至包括其它的切片。
*/
func slicesDemo4() {
	// 创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("|%s|\n", strings.Join(board[i], " "))
	}
}

/*
向切片追加元素
*/
func slicesDemo5() {
	var s []int
	printSlice(s)

	// 添加一个元素
	s = append(s, 0)
	printSlice(s)
	
	// 这个切片会按需增长
	s = append(s, 1)
	printSlice(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	printSlice(s)
}



/*
Range
for 循环的 range 形式可遍历切片或映射。

当使用 for 循环遍历切片时，每次迭代都会返回两个值。第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。
*/
func rangeDemo() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	
/*	
可以将下标或值赋予 _ 来忽略它。
for i, _ := range pow
for _, value := range pow
若你只需要索引，忽略第二个变量即可。
for i := range pow
*/
	pow = make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}


type Location struct {
	Lat, Long float64
}

/*
映射(map)
映射将键映射到值。

映射的零值为 nil 。nil 映射既没有键，也不能添加键。

make 函数会返回给定类型的映射，并将其初始化备用。
*/
var m map[string]Location

func mapDemo() {
	m = make(map[string]Location)
	m["Bell Labs"] = Location{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}

/*
映射的文法
*/
func mapDemo2() {
	// 映射的文法与结构体相似，不过必须有键名。
	var m = map[string]Location{
		"Bell Labs": Location{
			40.68433, -74.399671110004,
		},
		"Google": Location{
			37.42202, -122.08408111555,
		},
	}
	fmt.Println(m)
	
	// 若顶级类型只是一个类型名，你可以在文法的元素中省略它。
	m = map[string]Location{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08409},
	}
	fmt.Println(m)
}


/*
删除元素：
delete(m, key)

通过双赋值检测某个键是否存在：
elem, ok = m[key]

若 key 在 m 中，ok 为 true ；否则，ok 为 false。
若 key 不在m中，那么 elem 是该映射元素类型的零值。
同样的，当从映射中读取某个不存在的键时，结果是映射的元素类型的零值。
*/
func mapDemo3() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}


func WordCount(s string) map[string]int {
	m := make(map[string]int)
	var arr []string = strings.Fields(s)
	for _, val := range arr {
		m[val]++
	}
	
	return m
}

/*

函数值
函数也是值。它们可以像其它值一样传递。

函数值可以用作函数的参数或返回值。
*/
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func funcDemo() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}



/*
函数的闭包
Go 函数可以是一个闭包。闭包是一个函数值，它引用了其函数体之外的变量。该函数可以访问并赋予其引用的变量的值，换句话说，该函数被这些变量“绑定”在一起。

例如，函数 adder 返回一个闭包。每个闭包都被绑定在其各自的 sum 变量上。
*/
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func funcDemo2() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}


// ------------------------------------------


type VertexF struct {
	X, Y float64
}

/*
方法
Go 没有类。不过你可以为结构体类型定义方法。
方法就是一类带特殊的 接收者 参数的函数。
方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
*/

// 方法(值接收者)
func (v VertexF) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 函数
func Abs(v VertexF) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func methodDemo() {
	v := VertexF{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))
}



/*
指针与函数

指针参数是址传递（跟C一样）
*/

// 值传递
func Scale(v VertexF, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 址传递，函数的指针参数
func ScaleP(v *VertexF, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}



// 指针接收者
func (v *VertexF) ScaleP(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func pointerFuncDemo() {
	v := VertexF{3, 4}
	
	Scale(v, 10)
	fmt.Println(Abs(v))
	
	ScaleP(&v, 10)
	fmt.Println(Abs(v))
	
	p := &v
	p.ScaleP(10)
	fmt.Println(Abs(v))
	
	// 而以指针为接收者的方法被调用时，接收者既能为值又能为指针
	// 由于 ScaleP 方法有一个指针接收者，为方便起见，Go 会将语句 v.Scale(5) 解释为 (&v).Scale(5)。
	v.ScaleP(10)
	fmt.Println(Abs(v))
}


/*
接口
接口类型 是由一组方法签名定义的集合。

接口类型的变量可以保存任何实现了这些方法的值。

接口隐式实现，没有“implements”关键字。
*/
type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func interfaceDemo() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := VertexF{3, 4}

	a = f  // a MyFloat implements Abser
	a = v // a VertexF implements Abser

	fmt.Println(a.Abs())
}


/*
底层值为 nil 的接口值
即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。

在一些语言中，这会触发一个空指针异常，但在 Go 中通常会写一些方法来优雅地处理它（如本例中的 M 方法）。

注意: 保存了 nil 具体值的接口其自身并不为 nil。
*/

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}


type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func interfaceDemo2() {
	var i I

	var t *T
	i = t // 保存了 nil 具体值的接口其自身并不为 nil
	describe(i) // (<nil>, *main.T)
	i.M() // 不会报空指针
	
	
//  nil 接口值
//  nil 接口值既不保存值也不保存具体类型。
//  为 nil 接口调用方法会产生运行时错误，因为接口的元组内并未包含能够指明该调用哪个 具体 方法的类型。
//	var np I
//	np.M() // 空指针

	i = &T{"hello"}
	describe(i)
	i.M()
}


func describe2(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}


/*

空接口
指定了零个方法的接口值被称为 *空接口：*

interface{}
空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）【类似java中的Object】

空接口被用来处理未知类型的值。
*/
func interfaceDemo3() {
	var i interface{}
	describe2(i)

	i = 42
	describe2(i)

	i = "hello"
	describe2(i)
}

// 简单的计数器服务。
type Counter struct {
	n int
}


// 对对象类型定义的接口
func (ctr Counter) countNotEffect(val int) {
	ctr.n += val
}

// 对指针类型定义的接口才能使对成员变量的修改生效
func (ctr *Counter) count(val int) {
	ctr.n += val
}

func interfaceDemo4()  {
	counter1 := Counter{0}
	counter1.countNotEffect(1)
	fmt.Printf("counter1 = %d\n", counter1.n) // 0
	counter1.countNotEffect(2)
	fmt.Printf("counter1 = %d\n", counter1.n) // 0

	counter2 := Counter{0}
	(&counter2).count(1)
	fmt.Printf("counter2 = %d\n", counter2.n) // 1
	counter2.count(5) // 自动转成指针调用
	fmt.Printf("counter2 = %d\n", counter2.n) // 3

	counter3 := new(Counter)
	counter3.count(2)
	fmt.Printf("counter = %d\n", (*counter3).n) // 2
	counter3.count(3)
	fmt.Printf("counter = %d\n", counter3.n) // 5
}



/*
类型断言
类型断言 提供了访问接口值底层具体值的方式。

t := i.(T)
该语句断言接口值 i 保存了具体类型 T，并将其底层类型为 T 的值赋予变量 t。

若 i 并未保存 T 类型的值，该语句就会触发一个恐慌。

为了 判断 一个接口值是否保存了一个特定的类型，类型断言可返回两个值：其底层值以及一个报告断言是否成功的布尔值。

t, ok := i.(T)
若 i 保存了一个 T，那么 t 将会是其底层值，而 ok 为 true。

否则，ok 将为 false 而 t 将为 T 类型的零值，程序并不会产生恐慌。

请注意这种语法和读取一个映射时的相同之处。
*/
func typeAssertionDemo() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

//	f = i.(float64) // 报错(panic)
//	fmt.Println(f)
}


/*
类型选择
类型选择 是一种按顺序从几个类型断言中选择分支的结构。

类型选择与一般的 switch 语句相似，不过类型选择中的 case 为类型（而非值）， 它们针对给定接口值所存储的值的类型进行比较。

switch v := i.(type) {
case T:
    // v 的类型为 T
case S:
    // v 的类型为 S
default:
    // 没有匹配，v 与 i 的类型相同
}
类型选择中的声明与类型断言 i.(T) 的语法相同，只是具体类型 T 被替换成了关键字 type。

此选择语句判断接口值 i 保存的值类型是 T 还是 S。在 T 或 S 的情况下，变量 v 会分别按 T 或 S 类型保存 i 拥有的值。在默认（即没有匹配）的情况下，变量 v 与 i 的接口类型和值相同。
*/
func typeJudge(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func typeSwitchDemo() {
	typeJudge(21)
	typeJudge("hello")
	typeJudge(true)
}


/*
Stringer(类似java的toString)
fmt 包中定义的 Stringer 是最普遍的接口之一。

type Stringer interface {
    String() string
}
Stringer 是一个可以用字符串描述自己的类型。fmt 包（还有很多包）都通过此接口来打印值。
%v输出的也是此值。
*/
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func stringerDemo() {
	a := Person{"Arthur Dent", 42}
	fmt.Println(a)
	fmt.Printf("Person a: %v\n", a)

	fmt.Println("a.String():", a.String())
}



type IPAddr [4]byte

// 给 IPAddr 添加一个 "String() string" 方法
func (ip IPAddr)String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func stringerDemo2() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}


/*
错误
Go 程序使用 error 值来表示错误状态。

与 fmt.Stringer 类似，error 类型是一个内建接口：

type error interface {
    Error() string
}
（与 fmt.Stringer 类似，fmt 包在打印值时也会满足 error。）

通常函数会返回一个 error 值，调用的它的代码应当判断这个错误是否等于 nil 来进行错误处理。

i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
error 为 nil 时表示成功；非 nil 的 error 表示失败。
*/
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"error message",
	}
}

func errorDemo() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

/*
练习：错误
从之前的练习中复制 Sqrt 函数，修改它使其返回 error 值。

Sqrt 接受到一个负数时，应当返回一个非 nil 的错误值。复数同样也不被支持。

创建一个新的类型

type ErrNegativeSqrt float64
并为其实现

func (e ErrNegativeSqrt) Error() string
方法使其拥有 error 值，通过 ErrNegativeSqrt(-2).Error() 调用该方法应返回 "cannot Sqrt negative number: -2"。

注意: 在 Error 方法内调用 fmt.Sprint(e) 会让程序陷入死循环。可以通过先转换 e 来避免这个问题：fmt.Sprint(float64(e))。这是为什么呢？

修改 Sqrt 函数，使其接受一个负数时，返回 ErrNegativeSqrt 值。
*/
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt)Error() string {
	if e < 0 {
		return fmt.Sprintf("should not be negative: %v", float64(e))
	} else {
		return "ok"
	}
}

func Sqrt2(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func errorDemo2() {
	f, e := Sqrt2(2)
	if e != nil {
		fmt.Println("error:", e)
	} else {
		fmt.Println("ok:", f)
	}
	
	f, e = Sqrt2(-2)
	if e != nil {
		fmt.Println("error:", e)
	} else {
		fmt.Println("ok:", f)
	}
	
	fmt.Println(ErrNegativeSqrt(3).Error())
	fmt.Println(ErrNegativeSqrt(-3).Error())
}


// ---------------------------

/*
Reader
io 包指定了 io.Reader 接口，它表示从数据流的末尾进行读取。

Go 标准库包含了该接口的许多实现，包括文件、网络连接、压缩和加密等等。

io.Reader 接口有一个 Read 方法：

func (T) Read(b []byte) (n int, err error)
Read 用数据填充给定的字节切片并返回填充的字节数和错误值。在遇到数据流的结尾时，它会返回一个 io.EOF 错误。

示例代码创建了一个 strings.Reader 并以每次 8 字节的速度读取它的输出。
*/
func readerDemo() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}



type MyReader struct{}

// 给 MyReader 添加一个 Read([]byte) (int, error) 方法
func (m MyReader)Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func readerDemo2() {
	
}



// ---------------------------------------------------------------
/*
Go 程(Goroutines)
Go 程（goroutine）是由 Go 运行时管理的轻量级线程。

go f(x, y, z)
会启动一个新的 Go 程并执行

f(x, y, z)
f, x, y 和 z 的求值发生在当前的 Go 程中，而 f 的执行发生在新的 Go 程中。

Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。sync 包提供了这种能力，不过在 Go 中并不经常用到，因为还有其它的办法（见下一页）。
*/
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func routineDemo() {
	go say("world")
	say("hello")
}



/*
信道
信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。

ch <- v    // 将 v 发送至信道 ch。
v := <-ch  // 从 ch 接收值并赋予 v。
（“箭头”就是数据流的方向。）

和映射与切片一样，信道在使用前必须创建：

ch := make(chan int)
默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。

以下示例对切片中的数进行求和，将任务分配给两个 Go 程。一旦两个 Go 程完成了它们的计算，它就能算出最终的结果。
*/
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func channelDemo() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int) // 创建信道
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从 c 中接收

	fmt.Println(x, y, x+y)
}


/*
带缓冲的信道
信道可以是 带缓冲的。将缓冲长度作为第二个参数提供给 make 来初始化一个带缓冲的信道：

ch := make(chan int, 100)
仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。

修改示例填满缓冲区，然后看看会发生什么。
*/
func channelDemo2() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}


type Vector []float64

// 将此操作应用至 v[i], v[i+1] ... 直到 v[n-1]
func (v Vector) DoSome(i, n int, c chan int) {
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(1000000000 * rand.Intn(2))) // 1s
	for ; i < n; i++ {
		v[i] = 2
	}
	c <- 1    // 发信号表示这一块计算完成。
}


func (v Vector) DoAll() {
	NCPU := runtime.NumCPU()  // CPU核心数
	fmt.Println("NCPU:", NCPU)
	runtime.GOMAXPROCS(NCPU)

	c := make(chan int, NCPU)  // 缓冲区是可选的，但明显用上更好
	for i := 0; i < NCPU; i++ {
		//time.Sleep(1000000)
		go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, c)
	}
	// 排空信道，等待每个任务完成
	for i := 0; i < NCPU; i++ {
		<-c
		fmt.Println("done tasks: ", i + 1, "v: ", v)
	}
	// 一切完成。
	fmt.Println("all done  : ", NCPU, "v: ", v)
}

func channelDemo3()  {
	v := make(Vector, 17)
	v.DoAll()
}


/*
range 和 close
一般是二者配合使用
发送者可通过 close 关闭一个信道来表示没有需要发送的值了。接收者可以通过为接收表达式分配第二个参数来测试信道是否被关闭：若没有值可以接收且信道已被关闭，那么在执行完

v, ok := <-ch
之后 ok 会被设置为 false。

循环 for i := range c 会不断从信道接收值，直到它被关闭。

*注意：* 只有发送者才能关闭信道，而接收者不能。向一个已经关闭的信道发送数据会引发程序恐慌（panic）。

*还要注意：* 信道与文件不同，通常情况下无需关闭它们。只有在必须告诉接收者不再有需要发送的值时才有必要关闭，例如终止一个 range 循环。
*/
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func channelDemo4() {
	c := make(chan int, 6)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}


/*
select 语句
select 语句使一个 Go 程可以等待多个通信操作。

select 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。
*/
func fibonacciSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func selectDemo() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciSelect(c, quit)
}


/*
默认选择(Default Selection)
当 select 中的其它分支都没有准备好时，default 分支就会执行。
为了在尝试发送或者接收时不发生阻塞，可使用 default 分支：

select {
case i := <-c:
    // 使用 i
default:
    // 从 c 中接收会阻塞时执行
}
*/
func defaultSelectionDemo() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}



/*
sync.Mutex
我们已经看到信道非常适合在各个 Go 程间进行通信。

但是如果我们并不需要通信呢？比如说，若我们只是想保证每次只有一个 Go 程能够访问一个共享的变量，从而避免冲突？

这里涉及的概念叫做 *互斥（mutual*exclusion）* ，我们通常使用 *互斥锁（Mutex）* 这一数据结构来提供这种机制。

Go 标准库中提供了 sync.Mutex 互斥锁类型及其两个方法：

Lock
Unlock
我们可以通过在代码前调用 Lock 方法，在代码后调用 Unlock 方法来保证一段代码的互斥执行。参见 Inc 方法。

我们也可以用 defer 语句来保证互斥锁一定会被解锁。参见 Value 方法。
*/
// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func multexDemo() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}


// --------------------------------
/*
练习：Web 爬虫
在这个练习中，我们将会使用 Go 的并发特性来并行化一个 Web 爬虫。

修改 Crawl 函数来并行地抓取 URL，并且保证不重复。

提示：你可以用一个 map 来缓存已经获取的 URL，但是要注意 map 本身并不是并发安全的！
*/

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: 并行的抓取 URL。
	// TODO: 不重复抓取页面。
        // 下面并没有实现上面两种情况：
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}


func crawlDemo() {
	Crawl("https://golang.org/", 4, fetcher)
}


// 捕获异常
func safelyDo() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
			log.Println("continue do after recover")
		}
	}()
	func() {
		panic("throw a panic")
	}()
}

// 抛出异常
func panicDo() {
	var user = os.Getenv("USER")
	if user == "" {
		panic("no $USER")
	}
}


// --------------------------------

func main() {

	fmt.Println("Hello, Go.")
	fmt.Println("The time is", time.Now())

	constDemo2()
	
	rand.Seed(100)
	fmt.Println("Get a random: ", rand.Intn(10))
	
	
	fmt.Println("math.Sqrt(7): ", math.Sqrt(7), " math.Pi: ", math.Pi)
	
	
	fmt.Println(morestrings.ReverseRunes(".3,2,1"))
	//fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	
	
	fmt.Println("1+3=", add2(1, 3))
	
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	
	fmt.Println(split(17))
	
	variableDemo()
	variableDemo2()
	variableDemo3()
	
	basicTypeDemo()

	runeDemo()

	defaultVal()
	
	switchDemo()
	switchDemo2()
	switchDemo3()

	deferDemo()
	deferDemo2()
	
	pointerDemo()
	
	structDemo()
	
	arrDemo()
	
	slicesDemo()
	slicesDemo2()
	slicesDemo3()
	
	rangeDemo()
	
	mapDemo()
	mapDemo2()
	
	fmt.Println(WordCount("Hello Go I like Go I am learning Go"))
	
	funcDemo()
	
	pointerFuncDemo()
	
	interfaceDemo3()
	interfaceDemo4()
	
	typeAssertionDemo()
	typeSwitchDemo()
	
	stringerDemo()
	
	errorDemo()
	errorDemo2()
	
	readerDemo()
	
	routineDemo()
	
	channelDemo()
	channelDemo3()
	
	selectDemo()
	defaultSelectionDemo()
	
	multexDemo()

	safelyDo()
	panicDo()
}