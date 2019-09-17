package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(compareVersion("2.03", "2.2"))
}

func compareVersion(version1, version2 string) int {

	for {
		v1, v2 := 0, 0
		index1 := strings.Index(version1, ".")
		if index1 == -1 {
			v1, _ = strconv.Atoi(version1)
			version1 = ""
		} else {
			v1, _ = strconv.Atoi(version1[:index1])
			version1 = version1[index1+1:]
		}
		index2 := strings.Index(version2, ".")
		if index2 == -1 {
			v2, _ = strconv.Atoi(version2)
			version2 = ""
		} else {
			v2, _ = strconv.Atoi(version2[:index2])
			version2 = version2[index2+1:]
		}
		if v1 == v2 && version1 == version2 {
			return 0
		}
		if v1 < v2 {
			return -1
		} else if v1 > v2 {
			return 1
		}
	}
}
