
[key words]
var, func, const


var i int
var i, j int
var k, t int = 10, 11
var k = 10
k := 10

var (
	tobe bool = false
	name string = "wangbo"
)

sum := 10
for i:=1; i < 10; i++ {
	sum += i
}

for sum < 1000 {
	sum += sum
}


2、流程控制
if for switch 都可以先执行一个代码段，一般用于定义自己作用域的变量
switch里面的case居然可以是变量，太爽了; swtich如果没有条件，直接等同于一长串if then else
defer 其实是将函数压入栈中，后进先出，有用的特性

3、数据结构
指针：var p *int = &a
结构体：
type Vertex struct {
	X int
	Y int
}
切片slice，太强大



fmt.print
fmt.Println
fmt.Printf
fmt.Sprint
