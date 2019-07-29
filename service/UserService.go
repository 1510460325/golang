package UserService

var Id = 13

/*
  大写的函数或者全局变量才能被外包引用
*/
func Test(a int, b int) int {
	return a - b + Id
}
