# Go语言基础知识点
## 函数
### 函数签名（函数类型）
相当于函数指针、java的函数式接口
~~~ golang
// 代表两个参数一个返回值的函数签名
type Op func(int, int) int

func do(op Op, a int, b int) {
    return Op(a, b)
}
~~~
### 闭包
定义：闭包由函数及其相关引用环境组合而成的实体，通过匿名函数中引用外部函数的局部变量或者包全局变量构成，
* 多次调用返回闭包的函数（如test()）,返回的多个闭包的外部变量是多个副本
* 多次调用一个闭包（如f()），多次调用共享一个闭包数据，会有影响
~~~ golang
func test() func() {
	a := 12
	f := func() {
		a++
		fmt.Println(a)
	}
	return f
}
func main() {
	f := test()
	f() // 13
	f() // 14
}
~~~
对象是附有行为的数据，闭包是附有数据的行为