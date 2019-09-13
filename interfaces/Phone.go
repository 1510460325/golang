package main

import (
	"errors"
	"fmt"
)

type Phone interface {
	callNumber(target string) error
	sendMsg(target string) (string, string)
}

type Nokia struct {
	number  string
	version string
}

func (p *Nokia) callNumber(target string) error {
	fmt.Printf("即将打电话给 %s\n", target)
	return errors.New("电话欠费")
}

func (p *Nokia) sendMsg(target string) (string, string) {
	fmt.Printf("即将发送短信给 %s\n", target)
	return "发送成功", target
}

func main() {
	var phone Phone = &Nokia{"001", "001"}
	err := phone.callNumber("002")
	fmt.Println(err)
}
