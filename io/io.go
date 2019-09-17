package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Data struct {
	msg string
}

// 不用指针类型，为了在输出是不会修改源对象的属性，更安全
func (data Data) String() string {
	return data.msg
}

func toString() {
	fmt.Println(Data{msg: "hello"})
}

func readFromReader(reader io.Reader, num int) ([]byte, error) {
	buffer := make([]byte, num)
	i, err := reader.Read(buffer)
	return buffer[:i], err
}

func readFromStdin() {
	//data := make([]byte, 500)
	//for ; ; {
	//	i, _ := os.Stdin.Read(data)
	//	if i == 0 {
	//		break
	//	}
	//	fmt.Print(string(data[:i]))
	//}

	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, _ := reader.ReadLine()
		if line != nil {
			fmt.Println(string(line))
		}
	}
}

func readFromFile() {
	// filepath.abs 获取一个相对路劲的绝对路径
	dir, _ := filepath.Abs("./")
	file, err := os.Open(dir + "/io/io.go")
	if err != nil {
		fmt.Printf("读取有误,错误 : %v\n", err)
		return
	}
	data, _ := readFromReader(file, 50000)
	fmt.Println(string(data))
}

func buffedIO() {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = fmt.Fprint(writer, "hello")
	_, _ = fmt.Fprint(writer, " world!")
	_ = writer.Flush()
}

func writeToFile() {
	path, _ := filepath.Abs(".")
	file, _ := os.OpenFile(path+"/io/test", os.O_APPEND, os.ModeAppend)
	if file == nil {
		file, _ = os.Create(path + "/io/test")
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, _ = writer.WriteString("1212")
	_, _ = writer.WriteString("asdf")
	_ = writer.Flush()
}

func main() {
	//toString()
	readFromStdin()
	//readFromFile()
	buffedIO()
	writeToFile()
}
