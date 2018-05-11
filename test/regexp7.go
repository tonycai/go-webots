package main

import (
	"fmt"
	"regexp"
)

// 示例：第一匹配和最长匹配
func main() {
	b := []byte("abc12345def123")
	pat := "[0-9]+"
	reg1 := regexp.MustCompile(pat)      // 第一匹配
	reg2 := regexp.MustCompilePOSIX(pat) // 最长匹配
	fmt.Printf("%s\n", reg1.Find(b))     // abc1
	fmt.Printf("%s\n", reg2.Find(b))     // abc1def1

	b = []byte("abc1def1")
	pat = `(abc|abc1def)*1`
	reg1 = regexp.MustCompile(pat)      // 第一匹配
	reg2 = regexp.MustCompilePOSIX(pat) // 最长匹配
	fmt.Printf("%s\n", reg1.Find(b))    // abc1
	fmt.Printf("%s\n", reg2.Find(b))    // abc1def1
}
