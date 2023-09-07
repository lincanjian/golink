package golink

import (
	"errors"
	"strings"
)

func Link_Verify(identifier string, number uint, link string) (arr []string, err error) {

	var (
		firstChar string
		lastChar  string
	)
	firstChar = string(link[0])
	lastChar = string(link[len(link)-1])
	arr = strings.Split(link, ":")

	if firstChar == ":" || lastChar == ":" {
		err = errors.New("首个字符或尾部字符不能为\":\"！")
		return
	}

	if len(arr) < 5 || len(arr) > 5 {
		err = errors.New("参数数量有误！")
		return
	}

	if arr[0] != "email" {
		err = errors.New("标识符有误！")
		return
	}
	return
}
