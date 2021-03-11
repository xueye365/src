package test

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

/**
文件名必须以_test.go结尾，方法名必须以Test开头
*/
func TestFirst(t *testing.T) {
	// 斐波那契数列
	t.Log("test")
	a := 1
	b := 1
	t.Log(a)
	for i := 1; i < 5; i++ {
		t.Log(b)
		temp := a
		a = b
		b = temp + a
	}
}

func TestExchange(t *testing.T) {
	// 交换两个变量的值
	// 可以在一个赋值语句里面赋值
	a := 1
	b := 2
	//temp:=a
	//a=b
	//b=temp
	a, b = b, a
	t.Log(a, b)
}

// 连续常量可以使用该简洁写法
const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readable    = 1 << iota // 可读
	Writeable               // 可写
	Executeable             // 可执行
)

func TestConst(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday)
	t.Log(Readable, Writeable, Executeable)
	a := 7
	// 判断a是否可读可写可执行
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executeable == Executeable)
}

type MyInt int64

func TestImplicit(t *testing.T) {
	// 哪怕是别名的int64，go里面也不允许隐形转换数据类型，只能显示转换
	var a int64 = 2
	var c MyInt
	c = MyInt(a)
	t.Log(a, c)
	i := math.MaxInt64
	t.Log(i)
}

func TestPoint(t *testing.T) {
	a := 3
	// 取址符
	temp := &a
	t.Log(temp)
	// 按照格式化输出结果
	t.Logf("%T %T", a, temp)
	// 但是go语言不支持指针运算
}

func TestCompareArray(t *testing.T) {
	// 长度不同的数组进行比较会得到一个编译错误
	// 在其他编程语言里它是指针类型进行比较
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	c := [...]int{2, 3, 4, 5}
	// 顺序也必须要相等
	t.Log(a == b)
	t.Log(a == c)

}

func TestBitClear(t *testing.T) {
	// 按位清零 &^
	// 只要是右边的是1，则清零，如果右边是0，则左边是什么就是什么
	t.Log(1 &^ 0)
	t.Log(1 &^ 1)
	t.Log(0 &^ 1)
	t.Log(0 &^ 0)

	a := 7               // 0111
	a = a &^ Executeable // 0100 将第三位给清为0了
	t.Log(a&Readable == Readable, a&Writeable == Writeable, a&Executeable == Executeable)

}

func TestFor(t *testing.T) {
	// for循环和while循环是一样的

	// if 条件语句可以支持变量赋值
	//if v, err := someFun(); err == nil {
	//	t.Log("如果有错误的时候返回")
	//} else {
	//	t.Log("如果没有错误的时候使用返回值", v)
	//}

	// switch 语句也支持赋值语句，
	// 不需要写break，默认有break
	// case可以写多个，命中任意一个都算命中
	// 没有数据类型的限制
	// 可以不设定switch之后的条件表达式，在此种情况下，相当于if else
}

func TestArray(t *testing.T) {
	var arr [3]int
	arr[1] = 2
	t.Log(arr)
	a := [...]int{1, 2, 3, 4, 5}
	for i, v := range a {
		t.Log(i, v)
	}

	// 左侧包含，右侧不包含
	t.Log(a[1:2])
	t.Log(a[1:3])
	t.Log(a[1:len(a)])
	t.Log(a[1:])
	t.Log(a[:2])

}

func TestSlice(t *testing.T) {
	// 切片传的是值
	// 切片是一个数据结构，包含了指向后面数据的指针，即便函数传递的是值，因为整体结构被复制，两个值所指向的是同一个空间，所以会
	// 切片 可变长，共享空间
	// len：可访问的长度，cap:容量
	// 切片本质上是一个共享存储结构
	// make可以初始化数组长度
	s1 := make([]int, 3, 5)
	t.Log(len(s1), cap(s1))
	s1 = append(s1, 3)
	t.Log(len(s1), cap(s1))
	t.Log(s1)

	// 初始化一个数组,并且观察它的增长过程
	s := []int{}
	for i := 0; i < 10; i++ {
		// 每次增长两倍，增长的时候都是一个新数组，新地址，所以需要重新赋值给s
		s = append(s, i)
		t.Log(len(s), cap(s))
	}

	s2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := s2[3:5]
	t.Log(slice2, len(slice2), cap(slice2))
	// 因为切片是指针，当在切片里修改数据的时候，远数组就变化了，将6变化成为11
	slice2 = append(slice2, 11)
	t.Log(slice2, len(slice2), cap(slice2))
	t.Log(s2, len(s2), cap(s2))

}

func TestMap(t *testing.T) {
	// 数组容量不可伸缩，切片容量可伸缩
	// 数组可以进行比较，切片不可以进行比较(只能对nil进行比较)
	m1 := map[int]int{1: 2, 3: 4, 5: 6}
	t.Logf("len: %d", len(m1))
	// map是不能用cap来得到容量的，但是10设置的是容量
	m3 := make(map[int]int, 10)
	t.Logf("len: %d", len(m3))
	// 当数据不存在的时候会默认返回一个0，不会存在空指针的情况，但是需要手动判断数据是不存在还是存的就是0
	m2 := map[int]int{}
	t.Log(m2[3])
	if v, ok := m2[3]; ok {
		t.Log(v)
	} else {
		t.Log("不存在该key")
	}

	for k, v := range m1 {
		t.Log(k, v)
	}

}

func TestMap2(t *testing.T) {
	// value可以是一个函数，通过这个可以创建一个工厂模式
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}
	t.Log(m[1](2), m[2](3), m[3](4))
}

func TestSet(t *testing.T) {
	// go语言里没有set集合，通常使用map来代替set
	m := map[int]bool{2: true, 3: true, 4: false}
	t.Log(m)
	delete(m, 3)
}

func TestString(t *testing.T) {
	// String是只读的切片，len函数返回它包含的byte数，

	// String类型在初始化的时候会被初始化成空字符串
	// 在其他编程语言里它是指针类型
	var s string
	t.Log(len(s))
	t.Log("*" + s + "*")

	// String的byte数组可以存放任何数据，可以存储任何二进制数据
	s = "\xE4\xB8\xA5"
	t.Log(s)
	// String是一个不可以变的byte切片

	s1 := "中"
	// byte数组的长度
	t.Log(len(s1))
	// 存储字符的长度
	runes := []rune(s1)
	t.Log(len(runes))
	// 4e2d（在编码集中的编码）
	t.Logf("中 Unicode %x", runes[0])
	// e4b8ad 它在byte中的存储
	t.Logf("中 UTF8 %x", s1)

	s3 := "中华人民共和国"
	for _, i := range s3 {
		// [1]代表都是和第一个字符对应，只是不同的对应,%c汉字，%d数字，%x十六进制
		t.Logf("%[1]c %[1]d, %[1]x", i)
	}

	// 切割成一个切片
	//split := strings.Split(s3, ",")

}

func TestConv(t *testing.T) {
	// 数字变成字符串
	s := strconv.Itoa(10)
	t.Log("abc" + s)

	if s2, err := strconv.Atoi("10"); err != nil {
		t.Log("不是数字")
	} else {
		t.Log("是数字", 10+s2)
	}
}

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
}

// 装饰者模式
// 对于传进来的函数进行计时
func timeSpent(innr func(op int) int) func(op int) int {
	return func(op int) int {
		start := time.Now()
		ret := innr(op)
		fmt.Println("time spent", time.Since(start).Seconds())
		return ret
	}
}

// 慢方法
func slowFun(n int) int {
	time.Sleep(time.Second * 1)
	return n
}

// 调用被封装的慢方法
func TestFunc(t *testing.T) {
	spent := timeSpent(slowFun)
	i := spent(10)
	t.Log(i)
}

// 可变参数
func Sum(ops ...int) int {
	ret := 0
	for _, value := range ops {
		ret += value
	}
	return ret
}
